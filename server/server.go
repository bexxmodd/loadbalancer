package server

import (
	"fmt"
	"html"
	"net/http"
)

func Serve() {
	fmt.Println("Listening at ")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.ListenAndServe(":8096", nil)
}