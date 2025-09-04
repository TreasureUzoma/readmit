package controllers

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/treasureuzoma/readmit/readmit/utils"
)

// matchIgnore checks if a path matches any of the ignore patterns
func matchIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		if ok, _ := filepath.Match(pattern, filepath.Base(path)); ok {
			return true
		}
		if ok, _ := filepath.Match(pattern, path); ok {
			return true
		}
	}
	return false
}

// getGitUser fetches Git user.name and user.email from local config
func getGitUser() map[string]string {
	user := map[string]string{
		"name":  "",
		"email": "",
	}

	// git config user.name
	nameCmd := exec.Command("git", "config", "user.name")
	nameOut, err := nameCmd.Output()
	if err == nil {
		user["name"] = string(bytes.TrimSpace(nameOut))
	}

	// git config user.email
	emailCmd := exec.Command("git", "config", "user.email")
	emailOut, err := emailCmd.Output()
	if err == nil {
		user["email"] = string(bytes.TrimSpace(emailOut))
	}

	return user
}

// ReadFiles reads all files recursively (respecting ignore patterns) and adds Git user info
func ReadFiles() map[string]string {
	filesData := make(map[string]string)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing %s: %v", path, err)
			return nil
		}

		relPath := path
		if rel, err := filepath.Rel(".", path); err == nil {
			relPath = rel
		}

		if info.IsDir() {
			if matchIgnore(relPath, utils.IgnorePatterns) {
				return filepath.SkipDir
			}
			return nil
		}

		if matchIgnore(relPath, utils.IgnorePatterns) {
			return nil
		}

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return nil
		}

		filesData[relPath] = string(data)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	gitUser := getGitUser()
	filesData["__userdata__"] = "name: " + gitUser["name"] + "\nemail: " + gitUser["email"]

	return filesData
}
