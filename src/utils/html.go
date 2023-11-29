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

func ParseString(templateString string, props map[string]interface{}) string {
	templateHTML, err := template.New("index").Parse(string(templateString))
	if err != nil {
		fmt.Printf("Error parsing template (ParseHTML): %s\n", err)
		return ""
	}

	var resultHTML bytes.Buffer

	err = templateHTML.ExecuteTemplate(&resultHTML, "index", props)
	if err != nil {
		fmt.Printf("Error executing template (ParseHTML): %s\n", err)
		return ""
	}

	resultHTMLBytes := resultHTML.String()

	return resultHTMLBytes
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

func JoinStringWithCallback(strings []string, callback func(string) string) string {
	var result string
	for _, str := range strings {
		result += callback(str)
	}
	return result
}
