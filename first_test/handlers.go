package main

import (
	"net/http"
)


// home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate
}

// about is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}


