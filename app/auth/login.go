package auth

import (
	"short_url/app/db"
	"short_url/app/hashing"
)

func Login(username, email, password string) (bool, error) {
	var hash_password string
	query := `SELECT password FROM users WHERE username = $1 AND email = $2`
	err := db.Get_DB().QueryRow(query, username, email).Scan(&hash_password)
	if err != nil {
		return false, err
	}
	check_password := hashing.Check_Password_Hash(hash_password, password)
	return check_password, nil
}
