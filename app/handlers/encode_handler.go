package handlers

import (
	"short_url/app/auth"
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Encode_handler(router *mux.Router) {
	router.HandleFunc("/encode", auth.Protect_routes(middlewares.Encode_URL()))
}
