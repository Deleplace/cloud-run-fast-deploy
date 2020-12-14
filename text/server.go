package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const color = `red`

func main() {
	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		log.Println("Serving", color)
		fmt.Fprintln(w, color)
	})

	fs := http.StripPrefix("/fir/", http.FileServer(http.Dir("./fir")))
	http.Handle("/fir/", fs)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}
