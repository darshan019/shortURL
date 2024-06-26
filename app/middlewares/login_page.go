package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"short_url/app/auth"
	"short_url/models"
)

type tokens struct {
	Access_token string `json:"access_token"`
}

func Login_Manager() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}
		user := models.User()
		json.NewDecoder(r.Body).Decode(user)
		fmt.Println("user: ", user)

		check, err := auth.Login(user.Username, user.Email, user.Password)

		if err != nil || !check {
			fmt.Println("Error: ", err)
			return
		}

		token_string, err := auth.Create_token(user.Username, user.Email)
		if err != nil {
			fmt.Println(http.StatusInternalServerError)
			fmt.Println("No username found")
		}

		obj := tokens{Access_token: token_string}

		json.NewEncoder(w).Encode(obj)
	}
}
