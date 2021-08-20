package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	env := os.Getenv(("PORT"))
	fmt.Printf("Starting server on port %s...\n", env)

	http.HandleFunc("/submit", func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			fmt.Fprintf(rw, "Thank you! %s for filling the form", r.Form["name"][0])
			return
		}
		http.Error(rw, "forbidden", 403)
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", fs)

	http.ListenAndServe(":"+env, nil)
}
