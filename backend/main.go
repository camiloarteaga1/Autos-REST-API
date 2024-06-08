package main

import (
	//"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Start database connection
	//db.DB_Conection()

	// Router setup
	router := mux.NewRouter()

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT"},
	})

	// Routes
	handler := c.Handler(router)

	// Start server
	log.Fatal(http.ListenAndServe(":8080", handler))
	log.Println("Server running on port 8080")
}
