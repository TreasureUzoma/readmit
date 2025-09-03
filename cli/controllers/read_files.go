package controllers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"readmit/utils"
)

// matchIgnore checks if a given path matches any ignore pattern (glob-aware).
func matchIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		// First, check basename (like file.ext)
		if ok, _ := filepath.Match(pattern, filepath.Base(path)); ok {
			return true
		}
		// Then, check full relative path
		if ok, _ := filepath.Match(pattern, path); ok {
			return true
		}
	}
	return false
}

// ReadFiles reads files in the current directory, excluding ignored ones.
func ReadFiles() map[string]string {
	filesData := make(map[string]string)

	// walk through directory recursively
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing %s: %v", path, err)
			return nil
		}

		// normalize path (remove leading "./")
		relPath := path
		if rel, err := filepath.Rel(".", path); err == nil {
			relPath = rel
		}

		// skip directories entirely if ignored
		if info.IsDir() {
			if matchIgnore(relPath, utils.IgnorePatterns) {
				return filepath.SkipDir
			}
			return nil
		}

		// skip ignored files
		if matchIgnore(relPath, utils.IgnorePatterns) {
			return nil
		}

		// read file contents
		data, err := os.ReadFile(path)
		if err != nil {
			log.Printf("Error reading file %s: %v", path, err)
			return nil
		}

		// store file -> contents
		filesData[relPath] = string(data)
		fmt.Printf("Read: %s\n", relPath)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return filesData
}
