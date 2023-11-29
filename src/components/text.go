package components

import (
	"main/src/utils"
)

func Text(props map[string]interface{}) string {
	return utils.ParseString(`
		<p> {{.Text}} </p>
	`, props)
}
