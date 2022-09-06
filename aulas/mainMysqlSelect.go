package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:example@(172.19.0.3:3306)/Bank?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	/*
	   	{
	   		query := ` CREATE TABLE transfer_users (
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

	   	{
	   		username := "Dilma Oliveira"
	   		password := "secret"
	   		createdAt := time.Now()

	   		result, err := db.Exec(`INSERT INTO transfer_users (username, password, created_at) VALUES (?,?,?)`, username, password, createdAt)
	   		if err != nil {
	   			log.Fatal(err)
	   		}
	   		access_key, err := result.LastInsertId()
	   		fmt.Println(access_key)
	   	}
	*/
	{
		var (
			access_key int
			username   string
			password   string
			createdAt  time.Time
		)

		query := "SELECT access_key, password, created_at FROM transfer_users  WHERE access_key = ? "
		if err := db.QueryRow(query, 1).Scan(&access_key, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}
		fmt.Println(access_key, username, password, createdAt)
	}
	{
		type user struct {
			access_key int
			username   string
			password   string
			createdAt  time.Time
		}
		rows, err := db.Query(`access_key, password, created_at FROM transfer_users`)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var transfer_user []user

		for rows.Next() {
			var u user
			err := rows.Scan(&u.access_key, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			transfer_user = append(transfer_user, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%#v", transfer_user)
	}
	{
		_, err := db.Exec(`DELETE FROM  transfer_user WHERE access_key = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
