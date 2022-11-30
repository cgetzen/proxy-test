package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var serveRoot = "/assets"

func main() {
	fmt.Println("Starting...")
	r := mux.NewRouter()
	r.HandleFunc("/index.html", serveFile("index.html"))
	r.HandleFunc("/page1.html", serveFile("page1.html"))
	r.Handle("/redirect", http.RedirectHandler("/page1.html", http.StatusFound))
	log.Fatal(http.ListenAndServe(":8000", r))
}

func serveFile(name string) func(w http.ResponseWriter, r *http.Request) {
	// serveFile loads the file on boot and caches it.
	fmt.Printf("Caching %s\n", name)
	f, err := os.ReadFile(fmt.Sprintf("%s/%s", serveRoot, name))
	if err != nil {
		return func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, name, time.Now(), bytes.NewReader(f))
	}
}
