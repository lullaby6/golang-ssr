package utils

import (
	"bytes"
	"fmt"
	"html/template"
)

func JoinStringWithCallback(strings []string, callback func(string) string) string {
	var result string
	for _, str := range strings {
		result += callback(str)
	}
	return result
}

func ParseString(templateString string, props map[string]interface{}) string {
	template, err := template.New("index").Parse(string(templateString))
	if err != nil {
		fmt.Printf("Error parsing template (ParseHTML): %s\n", err)
		return ""
	}

	var resultHTML bytes.Buffer

	err = template.ExecuteTemplate(&resultHTML, "index", props)
	if err != nil {
		fmt.Printf("Error executing template (ParseHTML): %s\n", err)
		return ""
	}

	resultHTMLBytes := resultHTML.String()

	return resultHTMLBytes
}
