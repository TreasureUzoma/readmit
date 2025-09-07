package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/treasureuzoma/readmit/readmit/cmd/remote"
	"github.com/treasureuzoma/readmit/readmit/controllers"
	"github.com/treasureuzoma/readmit/readmit/gitreader"
	"github.com/treasureuzoma/readmit/readmit/utils"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [type]",
	Short: "Generate files like README, CONTRIBUTION guide, or commit messages",
	Long: `Supported types:
	- readme 		 Generates README.md
	- contribution 	 Generates CONTRIBUTION.md
	- commit 		 Suggests commit message (printed to console)
	- watchtower Picks up all vulnerbilities found and creates a report.md file
	- other 		 Creates <other>-<uuid>.txt`,
	Example: `	readmit generate readme
				readmit generate contribution
				readmit generate commit`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(`
  o__ __o         o__ __o__/_          o           o__ __o        o          o   __o__  ____o__ __o____ 
 <|     v\       <|    v              <|>         <|     v\      <|\        /|>    |     /   \   /   \  
 / \     <\      < >                  / \         / \     <\     / \\o    o// \   / \         \o/       
 \o/     o/       |                 o/   \o       \o/       \o   \o/ v\  /v \o/   \o/          |        
  |__  _<|        o__/_            <|__ __|>       |         |>   |   <\/>   |     |          < >       
  |       \       |                /       \      / \       //   / \        / \   < >          |        
 <o>       \o    <o>             o/         \o    \o/      /     \o/        \o/    |           o        
  |         v\    |             /v           v\    |      o       |          |     o          <|        
 / \         <\  / \  _\o__/_  />             <\  / \  __/>      / \        / \  __|>_        / \       
                                                                                                        
                                                                                                        
                                                                                                        
`)
		time.Sleep(500 * time.Millisecond)

		fileType := strings.ToLower(args[0])
		validTypes := map[string]bool{"readme": true, "contribution": true, "commit": true, "watchtower": true}
		if !validTypes[fileType] {
			log.Printf("[ERROR] Unsupported type: %s (valid: readme, contribution, commit)", fileType)
		}

		fmt.Printf("✓ Analyzing codebase...\n")
		time.Sleep(1 * time.Second)

		uuid, err := utils.GenerateUUID()
		if err != nil {
			log.Printf("[ERROR] Failed to generate UUID: %v", err)
		}
		fileName := fmt.Sprintf("%s-%s.txt", fileType, uuid)

		var fileBuffer *bytes.Buffer

		if fileType == "commit" {
			diffContent, err := gitreader.GetBestDiff()
			if err != nil {
				log.Printf("[ERROR] %v", err)
			}

			if diffContent == "" {
				fmt.Println("No tags found on the repository. Cannot generate a relevant commit message.")
				return
			}

			var builder strings.Builder
			builder.WriteString("=== GIT DIFF ===\n")
			builder.WriteString(diffContent)
			builder.WriteString("\n\n=== CODEBASE ===\n")

			contentMap := controllers.ReadFiles()
			for filename, fileContent := range contentMap {
				builder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
			}

			fileBuffer = bytes.NewBufferString(builder.String())
		} else {
			var contentBuilder strings.Builder
			contentMap := controllers.ReadFiles()
			for filename, fileContent := range contentMap {
				contentBuilder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
			}
			fileBuffer = bytes.NewBufferString(contentBuilder.String())
		}

		fmt.Printf("✓ Uploading files to remote storage...\n")
		time.Sleep(1 * time.Second)

		signedUrl, err := remote.GetSignedUrl(fileName)
		if err != nil {
			log.Printf("[ERROR] Failed to Upload codebase to AI: %v", err)
		}

		if err := remote.UploadFile(signedUrl, fileBuffer); err != nil {
			log.Printf("[ERROR] Failed to upload file: %v", err)
		}

		fmt.Printf("✓ Generating %s content...\n", fileType)
		time.Sleep(1 * time.Second)

		generatedContent, err := remote.CallGenerateAPI(fileName, fileType)
		if err != nil {
			log.Printf("[ERROR] Generate API failed: %v", err)
			log.Println("Could be your version, please reupdate")
		}

		switch fileType {
		case "readme":
			if err := os.WriteFile("README.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write README.md: %v", err)
			}
			fmt.Printf("✓ Successfully created README.md\n")

		case "contribution":
			if err := os.WriteFile("CONTRIBUTION.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write CONTRIBUTION.md: %v", err)
			}
			fmt.Printf("✓ Successfully created CONTRIBUTION.md\n")

		case "commit":
			fmt.Println(generatedContent)
			fmt.Println("✓ Commit message generated and printed to console.")

		default:
			if err := os.WriteFile(fileName, []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write file locally: %v", err)
			}
			fmt.Printf("✓ Successfully created %s at %s\n", fileType, fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
