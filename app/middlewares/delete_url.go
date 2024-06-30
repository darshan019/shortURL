package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"short_url/app/db"
)

func Delete_url() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			fmt.Fprintf(w, "Use DELETE method")
			return
		}
		type URL struct {
			Short string `json:"short"`
			//Long  string `json:"long"`
		}
		var url URL
		json.NewDecoder(r.Body).Decode(&url)
		fmt.Println("URL is :")
		fmt.Println(url)

		_, err := db.Get_DB().Exec(`DELETE FROM urls WHERE short_url = $1`, url.Short)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Println("Url deleted")
	}
}
