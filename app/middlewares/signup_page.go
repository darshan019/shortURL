package middlewares

import (
	"encoding/json"
	"net/http"
	"short_url/app/auth"
	"short_url/models"
)

func Signup_Manager() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		user := models.User()
		json.NewDecoder(r.Body).Decode(user)
		auth.Signup(user.Username, user.Email, user.Password)
	}
}
