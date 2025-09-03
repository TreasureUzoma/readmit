package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"readmit/controllers"
	"strings"

	"crypto/rand"
	"encoding/hex"

	"github.com/spf13/cobra"
)

var fileType string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates Readme or other files",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("[INFO] Starting generate command...")

		if fileType == "" {
			log.Fatal("[ERROR] --type flag is required")
		}

		var contentMap = controllers.ReadFiles()
		var contentBuilder strings.Builder

		for filename, fileContent := range contentMap {
			contentBuilder.WriteString(fmt.Sprintf("=== %s ===\n", filename))
			contentBuilder.WriteString(fileContent)
			contentBuilder.WriteString("\n\n")
		}

		fileBuffer := bytes.NewBufferString(contentBuilder.String())

		log.Println("[INFO] File created in memory")

		var fileName string
		if fileType == "commit" {
			fileName = "temp-commit.txt"
		} else {
			uuid, err := generateUUID()
			if err != nil {
				log.Fatalf("[ERROR] Failed to generate UUID: %v", err)
			}
			fileName = fmt.Sprintf("%s-%s.txt", strings.ToLower(fileType), uuid)
		}

		signedUrl, err := getSignedUrl(fileName)
		if err != nil {
			log.Printf("[ERROR] Failed to get signed URL: %v\n", err)
			return
		}
		log.Printf("[INFO] Signed URL obtained: %s\n", signedUrl)

		if err := uploadFile(signedUrl, fileBuffer); err != nil {
			log.Printf("[ERROR] Failed to upload file: %v\n", err)
			return
		}
		log.Println("[INFO] File uploaded successfully")

		generatedContent, err := callGenerateAPI(fileName, fileType)
		if err != nil {
			log.Printf("[ERROR] Generate API failed: %v\n", err)
			log.Println("[INFO] Attempting to restore original state...")
			return
		}
		log.Println("[INFO] Generate API call successful")

		switch fileType {
		case "readme":
			fileName = "README.md"
			err = os.WriteFile(fileName, []byte(generatedContent), 0644)
			if err != nil {
				log.Printf("[ERROR] Failed to write file locally: %v\n", err)
				return
			}
			log.Printf("[INFO] README.md created successfully")
			log.Printf("[INFO] %s generation complete", fileType)

		case "contribution":
			fileName = "CONTRIBUTION.md"
			err = os.WriteFile(fileName, []byte(generatedContent), 0644)
			if err != nil {
				log.Printf("[ERROR] Failed to write file locally: %v\n", err)
				return
			}
			log.Printf("[INFO] CONTRIBUTION.md created successfully")
			log.Printf("[INFO] %s generation complete", fileType)

		case "commit":
			fmt.Println(generatedContent)
			log.Println("[INFO] Commit message printed to console.")
			log.Printf("[INFO] %s generation complete", fileType)

		default:
			log.Println("[INFO] Unknown file type, writing to default file.")
			err = os.WriteFile(fileName, []byte(generatedContent), 0644)
			if err != nil {
				log.Printf("[ERROR] Failed to write file locally: %v\n", err)
				return
			}
			log.Printf("[INFO] %s created successfully at %s", fileType, fileName)
			log.Printf("[INFO] %s generation complete", fileType)
		}
	},
}

func init() {
	generateCmd.Flags().StringVarP(&fileType, "type", "t", "", "Type of file to generate (readme, contribution, commit, etc.)")
	rootCmd.AddCommand(generateCmd)
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(uuid), nil
}

func getSignedUrl(fileName string) (string, error) {
	body := fmt.Sprintf(`{"path":"%s"}`, fileName)
	resp, err := http.Post("http://localhost:3000/api/upload-url", "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("signed URL API returned %d", resp.StatusCode)
	}

	var data struct {
		SignedUrl string `json:"uploadUrl"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	return data.SignedUrl, nil
}

func uploadFile(url string, fileBuffer *bytes.Buffer) error {
	req, err := http.NewRequest("PUT", url, fileBuffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "text/plain")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed: %s", string(b))
	}
	return nil
}

func callGenerateAPI(fileName, mode string) (string, error) {
	body := fmt.Sprintf(`{"fileName":"%s","mode":"%s"}`, fileName, mode)
	log.Printf("[INFO] Sending request body: %s", body)

	resp, err := http.Post("http://localhost:3000/api/generate", "application/json", strings.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("generate API error: %s", string(b))
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", err
	}

	if content, ok := data[mode].(string); ok {
		return content, nil
	}

	if nestedResult, ok := data[mode].(map[string]interface{}); ok {
		if content, ok := nestedResult["text"].(string); ok {
			return content, nil
		}
	}

	return "", fmt.Errorf("generated content not found or is not a string for mode '%s'", mode)
}
