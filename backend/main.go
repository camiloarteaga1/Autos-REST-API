package main

import (
	//"encoding/json"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	//Inicializar la base de datos
	//db.DB_Conection()

	// Router setup
	router := NewRouter()

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:     []string{"http://localhost:5432"},
		AllowedCredentials: true,
		AllowedMethods:     []string{"GET", "POST", "DELETE", "PUT"},
	})

	handler := c.Handler(router)

	// Start server
	log.Fatal(http.ListenAndServe(":5432", handler))
	log.Println("Server running on port 5432")
}
