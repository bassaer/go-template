package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	tmpl := `{{range .}}tmplate -> {{.}}
{{end}}`
	s := []string{"aaa", "bbb", "ccc"}
	t := template.New("tmpl")
	template.Must(t.Parse(tmpl))
	var buff bytes.Buffer
	if err := t.Execute(&buff, s); err == nil {
		fmt.Println(buff.String())
		return
	}
	fmt.Println("error")
}
