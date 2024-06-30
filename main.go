package main

import (
	"net/http"
	"short_url/app/db"
	"short_url/app/handlers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
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
	handlers.Delete_handler(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "PUT", "DELETE", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handle := c.Handler(router)

	http.ListenAndServe("localhost:8080", handle)
}
