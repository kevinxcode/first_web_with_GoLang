package handlers

import (
	"net/http"

	"github.com/kevinxcode/first_web_with_GoLang/pkg/handlers/render"
)

// home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

// about is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}


