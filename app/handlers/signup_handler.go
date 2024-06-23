package handlers

import (
	"short_url/app/middlewares"

	"github.com/gorilla/mux"
)

func Signup_Handler(router *mux.Router) {
	router.HandleFunc("/signup", middlewares.Signup_Manager())
}
