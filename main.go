package main

import (
	"net/http"
	"short_url/app/db"
	"short_url/app/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db.Initialize()
	defer db.Get_DB().Close()
	router := mux.NewRouter()

	handlers.Root_handler(router)
	handlers.Encode_handler(router)
	handlers.Signup_Handler(router)
	handlers.Login_Handler(router)
	handlers.Redirect_Handler(router)

	http.ListenAndServe("localhost:8080", router)
}
