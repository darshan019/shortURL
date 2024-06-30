package handlers

import (
	"short_url/app/auth"
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Delete_handler(router *mux.Router) {
	router.HandleFunc("/url/delete", auth.Protect_routes(middlewares.Delete_url()))
}
