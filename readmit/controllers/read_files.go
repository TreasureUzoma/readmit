package controllers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/treasureuzoma/readmit/readmit/utils"
)

const (
	MaxFileSize     = 500 * 1024       // 500KB per file
	MaxTotalSize    = 5 * 1024 * 1024  // 5MB total codebase size
)

// matchIgnore checks if a path matches any of the ignore patterns
func matchIgnore(path string, patterns []string) bool {
	segments := strings.Split(filepath.ToSlash(path), "/")
	for _, pattern := range patterns {
		// Check each segment for a match (e.g. if node_modules is in patterns, match node_modules/sub)
		for _, segment := range segments {
			if ok, _ := filepath.Match(pattern, segment); ok {
				return true
			}
		}
		// Also check full relative path
		if ok, _ := filepath.Match(pattern, filepath.ToSlash(path)); ok {
			return true
		}
	}
	return false
}

// getGitUser fetches Git user.name and user.email from local config
func getGitUser() map[string]string {
	user := map[string]string{
		"name":  "Unknown User",
		"email": "unknown@example.com",
	}

	// Helper to run git config
	runGitConfig := func(key string) string {
		out, err := exec.Command("git", "config", "--get", key).Output()
		if err != nil {
			return ""
		}
		return strings.TrimSpace(string(out))
	}

	name := runGitConfig("user.name")
	if name != "" {
		user["name"] = name
	} else {
		// Fallback to system env
		if sysName := os.Getenv("USERNAME"); sysName != "" {
			user["name"] = sysName
		} else if sysName := os.Getenv("USER"); sysName != "" {
			user["name"] = sysName
		}
	}

	email := runGitConfig("user.email")
	if email != "" {
		user["email"] = email
	}

	return user
}

// ReadFiles reads all files recursively (respecting ignore patterns) and adds Git user info
func ReadFiles() (map[string]string, error) {
	filesData := make(map[string]string)
	var totalSize int64

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

		if info.IsDir() {
			if isIgnored(relPath) || info.Name() == "node_modules" || info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		if isIgnored(relPath) {
			return nil
		}

		// Skip files larger than MaxFileSize
		if info.Size() > MaxFileSize {
			log.Printf("Skipping large file: %s (%d KB)", relPath, info.Size()/1024)
			return nil
		}

		// Check total size limit
		if totalSize+info.Size() > MaxTotalSize {
			return fmt.Errorf("cannot generate: your codebase size exceeds the 5MB limit (current upload: %.2fMB)", float64(totalSize+info.Size())/(1024*1024))
		}

		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return nil
		}

		filesData[relPath] = string(data)
		totalSize += info.Size()
		return nil
	})

	if err != nil {
		return nil, err
	}

	gitUser := getGitUser()
	filesData["__userdata__"] = "name: " + gitUser["name"] + "\nemail: " + gitUser["email"]

	return filesData, nil
}

