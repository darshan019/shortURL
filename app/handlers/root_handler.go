package handlers

import (
	// "html/template"
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Root_handler(router *mux.Router) {
	router.HandleFunc("/", middlewares.Root_page())
}
