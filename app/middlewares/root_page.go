package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"short_url/app/auth"
	"short_url/app/db"
)

func Root_page() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprintf(w, "root page")
		token_string := r.Header.Get("Authorization")
		token_string = token_string[len("Bearer "):]
		if token_string != "e" {
			claims, ok := auth.Verify_token(token_string)
			fmt.Println(token_string, claims)
			if ok != nil {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "Invalid Token")
				fmt.Println(ok)
				return
			}

			type user struct {
				Username string   `json:"username"`
				Email    string   `json:"email"`
				Urls     []string `json:urls`
			}

			var u user
			u.Username = claims["username"].(string)
			u.Email = claims["email"].(string)

			query := `SELECT short_url, long_url FROM urls WHERE user_id = (SELECT id FROM users WHERE username = $1 AND email = $2);`
			rows, _ := db.Get_DB().Query(query, claims["username"], claims["email"])

			defer rows.Close()

			for rows.Next() {
				var short, long string
				rows.Scan(&short, &long)

				url := fmt.Sprintf("short: %s, long: %s", short, long)
				u.Urls = append(u.Urls, url)
			}

			json.NewEncoder(w).Encode(u)
		} else {
			fmt.Fprintf(w, "root page")
		}
	}
}
