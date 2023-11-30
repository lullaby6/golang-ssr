package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesFromDirWithSuffix(dir string, suffix string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory (GetFilesFromDir): %s\n", err)
		return nil
	}

	var htmlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			htmlFiles = append(htmlFiles, filepath.Join(dir, file.Name()))
		}
	}

	return htmlFiles
}
