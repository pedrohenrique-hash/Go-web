package main

import (
	"html/template"
	"net/http"
)

type Registration struct {
	Name     string
	Email    string
	Password string
}

func main() {

	fs := http.FileServer(http.Dir("assets"))

	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("templates/index.html"))

		if r.Method != http.MethodPost {

			tmpl.Execute(w, nil)

			return

		}

		data := Registration{
			Name:     r.FormValue("name"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}

		_ = data

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {

		tmpl := template.Must(template.ParseFiles("templates/file.html"))

		tmpl.Execute(w, nil)
	})

	http.ListenAndServe(":3000", nil)
}
