package remote

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetSignedUrl(fileName string) (string, error) {
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

func UploadFile(url string, fileBuffer *bytes.Buffer) error {
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

func CallGenerateAPI(fileName, mode string) (string, error) {
	body := fmt.Sprintf(`{"fileName":"%s","mode":"%s"}`, fileName, mode)

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
