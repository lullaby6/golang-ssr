package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
)

func GetHTML(path string) ([]byte, error) {
	htmlFileData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file (GetHTML): %s\n", err)
		return nil, err
	}

	resultHTMLBytes := []byte(string(htmlFileData))

	return resultHTMLBytes, nil
}

func ParseHTML(path string, props map[string]interface{}) ([]byte, error) {
	htmlFileData, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading file (ParseHTML): %s\n", err)
		return nil, err
	}

	templateHTML, err := template.New("index").Parse(string(htmlFileData))
	if err != nil {
		fmt.Printf("Error parsing template (ParseHTML): %s\n", err)
		return nil, err
	}

	var resultHTML bytes.Buffer

	err = templateHTML.ExecuteTemplate(&resultHTML, "index", props)
	if err != nil {
		fmt.Printf("Error executing template (ParseHTML): %s\n", err)
		return nil, err
	}

	resultHTMLBytes := resultHTML.Bytes()

	return resultHTMLBytes, nil
}
