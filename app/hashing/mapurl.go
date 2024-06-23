package hashing

import (
	"fmt"
	"log"
	"short_url/app/db"
)

func Map_URL(username, email, hash_value, url string) {
	query := `SELECT id FROM users WHERE username = $1 AND email = $2`
	var id int
	err := db.Get_DB().QueryRow(query, username, email).Scan(&id)

	if err != nil {
		log.Fatal("error: ", err)
	}

	insertquery := `INSERT INTO urls (user_id, short_url, long_url) VALUES ($1, $2, $3)`
	_, err = db.Get_DB().Exec(insertquery, id, hash_value, url)

	if err != nil {
		log.Fatal("error: ", err)
	}

	fmt.Println("Saved!!")

}
