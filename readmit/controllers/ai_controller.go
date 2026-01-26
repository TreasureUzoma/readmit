package controllers

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/treasureuzoma/readmit/readmit/cmd/remote"
	"github.com/treasureuzoma/readmit/readmit/gitreader"
	"github.com/treasureuzoma/readmit/readmit/utils"
)

// GenerateAIContent handles the full flow of generating content via Readmit AI
func GenerateAIContent(fileType string, includeFullCodebase bool) (string, error) {
	uuid, err := utils.GenerateUUID()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %w", err)
	}
	fileName := fmt.Sprintf("%s-%s.txt", fileType, uuid)

	var fileBuffer *bytes.Buffer

	if fileType == "commit" {
		diffContent, err := gitreader.GetBestDiff()
		if err != nil {
			return "", fmt.Errorf("failed to get git diff: %w", err)
		}

		if diffContent == "" {
			return "", fmt.Errorf("no changes found to generate a relevant commit message")
		}

		var builder strings.Builder
		builder.WriteString("=== GIT DIFF ===\n")
		builder.WriteString(diffContent)

		if includeFullCodebase {
			builder.WriteString("\n\n=== CODEBASE ===\n")
			contentMap := ReadFiles()
			for filename, fileContent := range contentMap {
				builder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
			}
		}

		fileBuffer = bytes.NewBufferString(builder.String())
	} else {
		var contentBuilder strings.Builder
		contentMap := ReadFiles()
		for filename, fileContent := range contentMap {
			contentBuilder.WriteString(fmt.Sprintf("=== %s ===\n%s\n\n", filename, fileContent))
		}
		fileBuffer = bytes.NewBufferString(contentBuilder.String())
	}

	signedUrl, err := remote.GetSignedUrl(fileName)
	if err != nil {
		return "", fmt.Errorf("failed to get signed URL: %w", err)
	}

	if err := remote.UploadFile(signedUrl, fileBuffer); err != nil {
		return "", fmt.Errorf("failed to upload file: %w", err)
	}

	generatedContent, err := remote.CallGenerateAPI(fileName, fileType)
	if err != nil {
		return "", fmt.Errorf("generate API failed: %w", err)
	}

	return generatedContent, nil
}
