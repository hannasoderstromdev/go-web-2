package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// RenderTemplate parses a template file and renders it as HTML
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	
	
	}
}