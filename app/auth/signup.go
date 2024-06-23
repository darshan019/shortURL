package auth

import (
	"fmt"
	"log"
	"short_url/app/db"
	"short_url/app/hashing"
)

func Signup(username, email, password string) {
	hash_password := hashing.Generate_Hash_Password(password)
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := db.Get_DB().Exec(query, username, email, hash_password)
	if err != nil {
		log.Fatal("Error insering user to db ", err)
	} else {
		fmt.Println("user ", username, " inserted")
	}
}
