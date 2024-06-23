package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"short_url/app/auth"
	"short_url/models"
)

func Signup_Manager() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Println("error")
			return
		}
		user := models.User()
		json.NewDecoder(r.Body).Decode(user)
		auth.Signup(user.Username, user.Email, user.Password)
	}
}
