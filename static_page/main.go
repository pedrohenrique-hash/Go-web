package main

import (
	"net/http"
)

/*
func file(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("assets/file.html"))
	tmplt.Execute(w, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var tmplt = template.Must(template.ParseFiles("assets/index.html"))
	tmplt.Execute(w, nil)
}
*/
func main() {
	/*
		http.HandleFunc("/", index)

		http.HandleFunc("/file", file)
	*/
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./assets")))

	http.ListenAndServe(":3000", mux)
}
