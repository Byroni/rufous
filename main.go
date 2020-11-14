package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

func main() {
	port := "8081"
	var message strings.Builder
	message.WriteString("CLI running on port ")
	message.WriteString(port)
	fmt.Println(message.String())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	var portString strings.Builder
	portString.WriteString(":")
	portString.WriteString(port)
	log.Fatal(http.ListenAndServe(portString.String(), nil))
}
