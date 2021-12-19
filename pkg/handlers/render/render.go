package render

import (
	"fmt"
	"net/http"
	"text/template"
)

// render temmplate using html/tmp 
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("erro parsing templates", err)
	}
}
// end render
