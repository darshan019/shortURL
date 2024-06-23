package middlewares

import (
	"fmt"
	"net/http"
	"short_url/app/db"

	"github.com/gorilla/mux"
)

func Redirect(router *mux.Router) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		short_url := vars["short_url"]
		var res string

		query := `SELECT long_url FROM urls WHERE short_url = $1`
		err := db.Get_DB().QueryRow(query, short_url).Scan(&res)
		fmt.Println("Got the link")

		if err != nil {
			fmt.Println("Error finding url")
		}
		http.Redirect(w, r, res, http.StatusMovedPermanently)
	}
}
