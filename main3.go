package main

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// )

// var portNumber = ":8080"

// // home is home page handler
// func Home(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "this is the home page")
// }

// func Profile(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "this is the Profile page")
// }

// // about is the about page handler
// func About(w http.ResponseWriter, r *http.Request) {
// 	sum := AddValues(2, 2)
// 	_, _ = fmt.Fprintf(w, fmt.Sprintf("this is about page and 2 + 2 is %d", sum))
// }

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(100.0, 0.0)
// 	if err != nil {
// 		fmt.Fprintf(w, "cannot divide by 0 ")
// 		return
// 	}
// 	fmt.Fprintf(w, fmt.Sprintf("%f divide by %f is %f", 100.0, 0.0, f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}
// 	result := x / y
// 	return result, nil

// }

// // add values twi integer and return sum
// func AddValues(x, y int) int {
// 	return x + y
// }

// // main is the main application
// func main() {
// 	http.HandleFunc("/", Home)
// 	http.HandleFunc("/about", About)
// 	http.HandleFunc("/profile", Profile)
// 	http.HandleFunc("/divide", Divide)

// 	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
// 	_ = http.ListenAndServe(portNumber, nil)
// }
