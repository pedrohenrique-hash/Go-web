package main

import (
	"fmt"

	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:example@(172.19.0.2:3306)/Bank?parseTime=true")
	if err != nil {
		log.Fatal(err)

	}

	/*
	   db,err :sql.Open("mysql","nomeBanco:password@(IP170.00.0.2:3306)/nomeTabela?parseTime=true")
	   if err != nil {
	    log.Fatal(err)
	   }
	   if err:= db.Ping();err != nil {
	    log.Fatal(err)
	   }

	*/
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ // Create a new table
		query := `
            CREATE TABLE transfer_users (
                access_key INT AUTO_INCREMENT,
                username TEXT NOT NULL,
                password TEXT NOT NULL,
                created_at DATETIME,
                PRIMARY KEY (access_key)
            );`

		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{ // Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO transfer_users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		access_key, err := result.LastInsertId()
		fmt.Println(access_key)
	}

}
