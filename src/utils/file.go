package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesFromDirWithSuffix(dir string, suffix string) []string {
	dirFiles, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory (GetFilesFromDir): %s\n", err)
		return nil
	}

	var matchingFiles []string
	for _, file := range dirFiles {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			matchingFiles = append(matchingFiles, filepath.Join(dir, file.Name()))
		}
	}

	return matchingFiles
}
