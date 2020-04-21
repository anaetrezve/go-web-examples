package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "go:pass@(127.0.0.1:3306)/godb?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}

	_, tableExists := db.Query("SELECT * FROM users;")

	if tableExists != nil {
		query := `
    CREATE TABLE users (
      id INT AUTO_INCREMENT,
      username TEXT NOT NULL,
      password TEXT NOT NULL,
      created_at DATETIME,
      PRIMARY KEY (id)
    );
  `

		_, err2 := db.Exec(query)
		if err2 != nil {
			log.Fatal("Can not create table", err2)
		}
	}

	{ // Creating user
		username := "user"
		password := "pass"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)

		if err != nil {
			log.Println(err)
		}

		// Getting recent user id
		userID, err := result.LastInsertId()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(userID)

	}

	{ // querying single user from database
		var (
			id         int
			username2  string
			password2  string
			createdAt2 time.Time
		)

		query := `SELECT id, username, password, created_at FROM users WHERE id = ?`

		err = db.QueryRow(query, 1).Scan(&id, &username2, &password2, &createdAt2)
		if err != nil {
			log.Println(err)
		}

		fmt.Printf(`
      Data after query
      id: %d
      username: %s
      password: %s
      createdAt: %v
      `, id, username2, password2, createdAt2)
	}
	{ // Query all users from database

		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Println(err)
		}
		var users []user

		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Println(err)
			}
			users = append(users, u)
		}
		rErr := rows.Err()
		if rErr != nil {
			log.Println(rErr)
		}
		fmt.Println("\n\nSlice Of Users Before Remove: ", users)
	}

	{
		/**
		 *  Delete a user from database
		 *
		 */
		_, err3 := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err3 != nil {
			log.Println(err)
		}
	}

}
