package handlers

import (
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Login_Handler(router *mux.Router) {
	router.HandleFunc("/login", middlewares.Login_Manager())
}
