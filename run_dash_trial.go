package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	//http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	fs := http.FileServer(http.Dir("static"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", indexHandler)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//_, err := fmt.Fprint(w, "Hello, World!")
	err := tpl.Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
