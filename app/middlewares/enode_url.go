package middlewares

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"short_url/app/db"
	"short_url/app/hashing"

	"github.com/golang-jwt/jwt/v5"
)

func Encode_URL() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		claims, ok := r.Context().Value("user").(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Could not retrieve user info")
			return
		}

		username := claims["username"].(string)
		email := claims["email"].(string)

		type longURL struct {
			Long_url string `json:"long_url"`
		}

		if r.Method != http.MethodPost {
			fmt.Println("err")
			return
		}

		var temp longURL
		json.NewDecoder(r.Body).Decode(&temp)
		fmt.Println(temp.Long_url)
		short_url := hashing.GenerateHash()

		var unique = false
		for !unique {
			unique = check_db_for_dup(short_url)
			short_url = hashing.GenerateHash()
		}

		hashing.Map_URL(username, email, short_url, temp.Long_url)

		w.Header().Set("Content-Type", "application/json")

		var send_url struct {
			Short_url string `json: "short_url"`
		}
		send_url.Short_url = short_url

		json.NewEncoder(w).Encode(send_url)
	}
}

func check_db_for_dup(short_url string) bool {
	query := `SELECT short_url FROM urls WHERE short_url = $1`
	var res string

	err := db.Get_DB().QueryRow(query, short_url).Scan(&res)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("Error checking for duplicate short_url: %v", err)
	}

	return res != short_url
}
