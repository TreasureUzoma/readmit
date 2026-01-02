package controllers

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

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

	// Parse .gitignore if it exists
	var gitIgnorePatterns []string
	if data, err := os.ReadFile(".gitignore"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			gitIgnorePatterns = append(gitIgnorePatterns, line)
		}
	}

	// Helper to check if a path is ignored
	isIgnored := func(path string) bool {
		// Checks against default patterns
		if matchIgnore(path, utils.IgnorePatterns) {
			return true
		}
		// Checks against .gitignore patterns
		if matchIgnore(path, gitIgnorePatterns) {
			return true
		}
		return false
	}

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing %s: %v", path, err)
			return nil
		}

		relPath := path
		if rel, err := filepath.Rel(".", path); err == nil {
			relPath = rel
		}

		// Handle explicit ignoring of everything in a folder if "*" is present in gitignore
		// But here we rely on the matchIgnore function to handle glob patterns like "somefolder/*" or just "*"

		if info.IsDir() {
			if isIgnored(relPath) {
				return filepath.SkipDir
			}
			// Special handling for node_modules if not caught by patterns
			if info.Name() == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}

		if isIgnored(relPath) {
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
