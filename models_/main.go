package main

import (
	"html/template"
	"net/http"
	"time"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))

	months_number := [12]int{1, 2, 3,
		4, 5, 6, 7, 8,
		9, 10, 11, 12}

	data_show := time.Now()

	// dia_show:= data_show.Day()

	mes_show := int(data_show.Month())

	// ano_show:= data_show.Year()

	var (
		segunda = "12"
		terca   = "5"
		quarta  = "4"
		quinta  = "6"
		sexta   = "7"
		sabado  = "10"
		domingo = "12"
	)
	var (
		valor_segunda = true
		valor_terca   = true
		valor_quarta  = true
		valor_quinta  = true
		valor_sexta   = true
		valor_sabado  = true
		valor_domingo = true
	)

	if segunda == "12" && months_number[mes_show] == mes_show {
		valor_segunda = false
	}

	if terca == "5" && months_number[mes_show] == mes_show {
		valor_terca = false
	}

	if quarta == "4" && months_number[mes_show] == mes_show {
		valor_quarta = false
	}

	if quinta == "6" && months_number[mes_show] == mes_show {
		valor_quinta = false
	}

	if sexta == "7" && months_number[mes_show] == mes_show {
		valor_sexta = false
	}

	if sabado == "10" && months_number[mes_show] == mes_show {
		valor_sabado = false
	}

	if domingo == "12" && months_number[mes_show] == mes_show {
		valor_domingo = false
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "Lulu Santos",
			Todos: []Todo{
				{Title: "Segunda-Feira", Done: valor_segunda},
				{Title: "Terça-Feira", Done: valor_terca},
				{Title: "Quarta-Feira", Done: valor_quarta},
				{Title: "Quinta-Feira", Done: valor_quinta},
				{Title: "Sexta-Feira", Done: valor_sexta},
				{Title: "Sabado", Done: valor_sabado},
				{Title: "Domingo", Done: valor_domingo},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
