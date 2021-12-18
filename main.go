package main

import (
	"fmt"
	"net/http"
)

var portNumber = ":8080"
// home is home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "this is the home page")
}

// about is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is about page and 2 + 2 is %d", sum))
}

// add values twi integer and return sum
func AddValues(x, y int) int {
	return x + y
}

// main is the main application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
