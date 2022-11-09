package main

import (
	"database/sql"
	"log"

	"net/http"

	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Clients struct {
	Id       int
	Email    string
	Password string
}

var tmpl = template.Must(template.ParseGlob("template/*.tmpl"))

func dataBaseConnection() (database *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "example"
	dbName := "REGISTRATION"

	database, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@(172.19.0.2:3306)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return database
}

// function Index

func Index(w http.ResponseWriter, r *http.Request) {

	database := dataBaseConnection()

	selectDataBase, err := database.Query("SELECT * FROM clinet ORDER BY id DESC")

	if err != nil {
		panic(err.Error())
	}

	name := Clients{}

	store_values := []Clients{}

	for selectDataBase.Next() {

		var Id int

		var Email, Password string

		err = selectDataBase.Scan(&Id, &Email, &Password)

		if err != nil {
			panic(err.Error())
		}

		name.Id = Id

		name.Email = Email

		name.Password = Password

		store_values = append(store_values, name)

	}

	tmpl.ExecuteTemplate(w, "Index", store_values)

	defer database.Close()
}

// function show

func Show(w http.ResponseWriter, r *http.Request) {

	database := dataBaseConnection()

	numberId := r.URL.Query().Get("id")

	selectDataBase, err := database.Query("SELECT * FROM client WHERE id = ?", numberId)

	if err != nil {
		panic(err.Error())
	}

	name := Clients{}

	for selectDataBase.Next() {
		var Id int

		var Email, Password string

		err = selectDataBase.Scan(&Email, &Password, &Id)

		if err != nil {
			panic(err.Error())
		}

		name.Id = Id
		name.Email = Email
		name.Password = Password

	}

	tmpl.ExecuteTemplate(w, "Show", name)

	defer database.Close()

}

// function New

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	database := dataBaseConnection()

	numberId := r.URL.Query().Get("id")

	selectDataBase, err := database.Query("SELECT * FROM client WHERE id = ?", numberId)

	if err != nil {
		panic(err.Error())
	}

	name := Clients{}

	for selectDataBase.Next() {
		var Id int

		var Email, Password string

		err = selectDataBase.Scan(&Id, &Email, &Password)

		if err != nil {
			panic(err.Error())
		}

		name.Id = Id
		name.Email = Email
		name.Password = Password
	}

	tmpl.ExecuteTemplate(w, "Edit", name)

	defer database.Close()
}

// function Insert

func Insert(w http.ResponseWriter, r *http.Request) {
	database := dataBaseConnection()

	if r.Method == "POST" {
		Email := r.FormValue("email")
		Password := r.FormValue("password")

		insForm, err := database.Prepare("INSERT INTO email (email, password) VALUES(?,?)")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(Email, Password)

		log.Println("INSERT: Email: " + Email + "| Password: " + Password)
	}

	defer database.Close()

	http.Redirect(w, r, "/", 301)
}

// function Update

func Update(w http.ResponseWriter, r *http.Request) {

	database := dataBaseConnection()

	if r.Method == "POST" {

		Email := r.FormValue("email")
		Password := r.FormValue("password")
		Id := r.FormValue("id")

		insForm, err := database.Prepare("UPDATE email SET email = ? WHERE id = ?")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(Email, Password, Id)

		log.Println("UPDATE: Email " + Email + "| Password: " + Password)

	}

	defer database.Close()

	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	database := dataBaseConnection()

	numberId := r.URL.Query().Get("id")

	deleteForm, err := database.Prepare("DELETE FROM client WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}

	deleteForm.Exec(numberId)

	log.Println("DELETE")

	defer database.Close()

	http.Redirect(w, r, "/", 301)
}

func main() {

	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)

	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)

	http.ListenAndServe(":9000", nil)
}
