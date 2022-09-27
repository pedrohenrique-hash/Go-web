package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./assets")))

	http.ListenAndServe(":3000", mux)
}
