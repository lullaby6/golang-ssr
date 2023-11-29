package components

import (
	"main/src/utils"
)

func Text(text string) string {
	props := map[string]interface{}{
		"Text": text,
	}

	return utils.ParseString(`
		<p>{{.Text}}</p>
	`, props)
}
