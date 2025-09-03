package controllers

import (
	"log"
	"os"
	"path/filepath"
	"readmit/utils"
)

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

		// store file -> contents
		filesData[relPath] = string(data)

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return filesData
}
