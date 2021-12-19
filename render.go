package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// render temmplate to view
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("erro parsing templates", err)
	}
}
// end render
