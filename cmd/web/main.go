package main

import (
	"fmt"
	"net/http"

	"github.com/kevinxcode/first_web_with_GoLang/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
