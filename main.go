package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloUser)

	http.ListenAndServe(":8080", nil)
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello, User! Welcome to Track-Em! 🎶 Track your favorite tunes!")
}
