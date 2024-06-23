package middlewares

import (
	"fmt"
	"net/http"
)

func Root_page() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "root page")
	}
}
