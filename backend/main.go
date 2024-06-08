package main

import (
	//"encoding/json"

	"log"
	"net/http"

	handlers "github.com/camiloarteaga1/Autos-REST-API.git/handlers"
	"github.com/camiloarteaga1/Autos-REST-API.git/server"
	"github.com/camiloarteaga1/Autos-REST-API.git/structures"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Start database connection
	server.DB_Conection()

	//Migraciones de las tablas para la DB

	server.DB.AutoMigrate(structures.Usuario{})
	server.DB.AutoMigrate(structures.Carro{})
	server.DB.AutoMigrate(structures.Reserva{})

	//Declaración del enrutador

	router := mux.NewRouter()

	//Funcion para mostrar todos los usuarios registrados

	router.HandleFunc("/users", handlers.UsersHandler).Methods("GET")

	//Funcion para Logeo
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	//Funcion para Registro
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	//Funcion para Buscar un carro en la DB
	router.HandleFunc("/cars", handlers.GetCarsHandler).Methods("GET")

	//Funcion para Buscar un carro en la DB
	router.HandleFunc("/cars/filter", handlers.GetCarsHandlerFiltrado).Queries("marca", "{marca}").Queries("tipo", "{tipo}").Queries("combustible", "{combustible}").Queries("transmision", "{transmision}").Queries("ModeloNombre", "{ModeloNombre}").Methods("GET")

	//Funcion para Reservar un carro
	router.HandleFunc("/reservations", handlers.ReservationsHandler).Methods("POST")

	//Funcion para ELIMINAR una Reserva
	router.HandleFunc("/reservations/delete", handlers.DeleteReservationsHandler).Methods("DELETE")

	//Funcion para ver el reporte de reservas del usuario
	router.HandleFunc("/report", handlers.ReportHandler).Methods("POST")

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
