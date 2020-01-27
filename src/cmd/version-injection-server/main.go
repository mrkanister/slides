package main

import (
	"fmt"
	"log"
	"net/http"
)

var VersionString string // HL

func main() {
	mux := http.NewServeMux()
	// setup routes ...

	mux.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) { // HL
		fmt.Fprintln(w, VersionString) // HL
	}) // HL

	// define timeouts in production!
	log.Fatal(http.ListenAndServe(":8080", mux))
}
