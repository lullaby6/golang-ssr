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

func JoinMapListWithCallback(data []map[string]interface{}, callback func(map[string]interface{}) string) string {
	var result string
	for _, item := range data {
		result += callback(item)
	}
	return result
}

func ParseString(templateString string, props map[string]interface{}) string {
	template, err := template.New("index").Parse(string(templateString))
	if err != nil {
		fmt.Printf("Error parsing template (ParseString): %s\n", err)
		return ""
	}

	var templateResult bytes.Buffer
	err = template.ExecuteTemplate(&templateResult, "index", props)
	if err != nil {
		fmt.Printf("Error executing template (ParseString): %s\n", err)
		return ""
	}

	result := templateResult.String()

	return result
}
