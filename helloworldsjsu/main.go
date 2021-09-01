package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Hello SJSU Server Listening On Port 8080\n")
	http.HandleFunc("/", HelloSJSUServer)
	http.ListenAndServe(":8080", nil)
}

func HelloSJSUServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World from SJSU\n")
}