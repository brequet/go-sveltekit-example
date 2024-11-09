package main

import (
	"log"
	"net/http"

	"github.com/brequet/go-sveltekit-example/frontend"
)

func main() {
	http.Handle("/api/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}))

	http.Handle("/", http.RedirectHandler("/app", http.StatusMovedPermanently))
	http.Handle("/app/", http.StripPrefix("/app", frontend.Handler()))

	log.Println("Server starting on localhost:8080...")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
