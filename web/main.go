package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>") //Prints to browser
}
func main() {
	http.HandleFunc("/", index) //Router
	fmt.Print("Server starting")
	http.ListenAndServe(":3000", nil) //Listen at port 3000
}
