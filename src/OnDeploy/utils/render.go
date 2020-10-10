package utils

import (
	"bytes"
	"text/template"
)

func RendTmp(temp string, attr interface{}) (string, error) {
	tmp, err := template.New("temp").Parse(temp)
	if err != nil {
		return "New Temp Error", err
	}

	var buf bytes.Buffer
	if err := tmp.Execute(&buf, &attr); err != nil {
		return "模板渲染失败", err
	}
	return buf.String(), nil
}
