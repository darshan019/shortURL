package handlers

import (
	"short_url/app/auth"
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Redirect_Handler(router *mux.Router) {
	router.HandleFunc("/{short_url}", auth.Protect_routes(middlewares.Redirect(router)))
	// router.HandleFunc("/{short_url}", middlewares.Redirect(router))
}
