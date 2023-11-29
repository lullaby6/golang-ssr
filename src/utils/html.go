package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetHTMLFilesFromDir(path string) []string {
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Printf("Error reading directory (GetHTMLFilesFromDir): %s\n", err)
		return nil
	}

	var htmlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			htmlFiles = append(htmlFiles, filepath.Join(path, file.Name()))
		}
	}

	return htmlFiles
}
