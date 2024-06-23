package handlers

import (
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Root_handler(router *mux.Router) {
	router.HandleFunc("/", middlewares.Root_page())
}
