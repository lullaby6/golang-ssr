package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

func GetHTML(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file (GetHTML): %s\n", err)
		return nil, err
	}

	result := []byte(string(data))

	return result, nil
}

func ParseHTML(path string, props map[string]interface{}) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file (ParseHTML): %s\n", err)
		return nil, err
	}

	template, err := template.New("index").Parse(string(data))
	if err != nil {
		fmt.Printf("Error parsing template (ParseHTML): %s\n", err)
		return nil, err
	}

	var templateResult bytes.Buffer
	err = template.ExecuteTemplate(&templateResult, "index", props)
	if err != nil {
		fmt.Printf("Error executing template (ParseHTML): %s\n", err)
		return nil, err
	}

	result := templateResult.Bytes()

	return result, nil
}

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
