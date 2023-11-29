package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetHTMLFilesFromDir(dir string) []string {
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory (GetHTMLFilesFromDir): %s\n", err)
		return nil
	}

	var htmlFiles []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".html") {
			htmlFiles = append(htmlFiles, filepath.Join(dir, file.Name()))
		}
	}

	return htmlFiles
}
