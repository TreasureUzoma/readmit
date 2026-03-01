package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/treasureuzoma/readmit/readmit/remote"
	"github.com/treasureuzoma/readmit/readmit/controllers"
	"github.com/treasureuzoma/readmit/readmit/utils"
)

var watchtowerCmd = &cobra.Command{
	Use:   "watchtower",
	Short: "Continuously watch for vulnerabilities",
	Long: `Readmit Watchtower scans your project for vulnerabilities
and prints and creates a REPORT.md.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(Ascii)
		fmt.Println("        👀 Watchtower")

		time.Sleep(500 * time.Millisecond)

		fmt.Println("✓ Analyzing codebase...")
		time.Sleep(1 * time.Second)

		fileType := "report" 

		uuid, err := utils.GenerateUUID()
		if err != nil {
			log.Printf("[ERROR] Failed to generate UUID: %v", err)
		}
		fileName := fmt.Sprintf("%s-%s.txt", fileType, uuid)

		var fileBuffer *bytes.Buffer
		var contentBuilder strings.Builder
		contentMap, err := controllers.ReadFiles()
		if err != nil {
			log.Printf("[ERROR] %v", err)
			return
		}
		for filename, fileContent := range contentMap {
			contentBuilder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
		}
		fileBuffer = bytes.NewBufferString(contentBuilder.String())

		signedUrl, err := remote.GetSignedUrl(fileName)
		if err != nil {
			log.Printf("[ERROR] Failed to get signed URL: %v", err)
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

		if err := os.WriteFile("REPORT.md", []byte(generatedContent), 0644); err != nil {
			log.Printf("[ERROR] Failed to write REPORT.md: %v", err)
		}

		fmt.Println("✓ Successfully created REPORT.md")
	},
}

func init() {
	rootCmd.AddCommand(watchtowerCmd)
}
