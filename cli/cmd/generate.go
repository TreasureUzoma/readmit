package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"readmit/cmd/remote"
	"readmit/controllers"
	"readmit/gitreader"
	"readmit/utils"
	"strings"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate [type]",
    Short: "Generate files like README, CONTRIBUTION guide, or commit messages",
    Long: `Supported types:
     - readme         Generates README.md
     - contribution   Generates CONTRIBUTION.md
     - commit         Suggests commit message (printed to console)
      - other          Creates <other>-<uuid>.txt`,
    Example: `  readmit generate readme
                readmit generate contribution
                readmit generate commit`,
	Args:  cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
        
		fileType := strings.ToLower(args[0])
        validTypes := map[string]bool{"readme": true, "contribution": true, "commit": true}
        if !validTypes[fileType] {
         log.Printf("[ERROR] Unsupported type: %s (valid: readme, contribution, commit)", fileType)
        }

		uuid, err := utils.GenerateUUID()
		if err != nil {
			log.Printf("[ERROR] Failed to generate UUID: %v", err)
		}
		fileName := fmt.Sprintf("%s-%s.txt", fileType, uuid)

		var fileBuffer *bytes.Buffer

		// case commit: Commit generation -> diff + codebase
		if fileType == "commit" {
			diffContent, err := gitreader.GetBestDiff()
			if err != nil {
				log.Printf("[ERROR] %v", err)
			}

			if diffContent == "" {
				log.Println("[INFO] No diffs found (staged, unstaged, or last commit).")
				return
			}

			// Include diff + codebase in buffer
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
			// Case 2: Other file types -> just codebase
			var contentBuilder strings.Builder
			contentMap := controllers.ReadFiles()
			for filename, fileContent := range contentMap {
				contentBuilder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
			}
			fileBuffer = bytes.NewBufferString(contentBuilder.String())
		}

		// Upload file
		signedUrl, err := remote.GetSignedUrl(fileName)
		if err != nil {
			log.Printf("[ERROR] Failed to get signed URL: %v", err)
		}

		if err := remote.UploadFile(signedUrl, fileBuffer); err != nil {
			log.Printf("[ERROR] Failed to upload file: %v", err)
		}

		// Call API
		generatedContent, err := remote.CallGenerateAPI(fileName, fileType)
		if err != nil {
			log.Printf("[ERROR] Generate API failed: %v", err)
		}

		// Handle output
		switch fileType {
		case "readme":
			if err := os.WriteFile("README.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write README.md: %v", err)
			}
			log.Println("[SUCCESS] README.md created successfully")

		case "contribution":
			if err := os.WriteFile("CONTRIBUTION.md", []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write CONTRIBUTION.md: %v", err)
			}
			log.Println("[SUCCESS] CONTRIBUTION.md created successfully")

		case "commit":
			fmt.Println(generatedContent)
			log.Println("[SUCCESS] Commit message printed to console.")

		default:
			if err := os.WriteFile(fileName, []byte(generatedContent), 0644); err != nil {
				log.Printf("[ERROR] Failed to write file locally: %v", err)
			}
			log.Printf("[SUCCESS] %s created at %s", fileType, fileName)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
