package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/camiloarteaga1/Autos-REST-API.git/server"
	"github.com/camiloarteaga1/Autos-REST-API.git/structures"
	"github.com/gorilla/mux"
)

func GetCarsHandler(w http.ResponseWriter, req *http.Request) {

	var cars []structures.Carro

	server.DB.Find(&cars) //Buscamos toda la tabla Carros

	json.NewEncoder(w).Encode(&cars) //Retorna todas las filas y columnas de la tabla

}

func GetCarsHandlerFiltrado(w http.ResponseWriter, req *http.Request) {

	var findCar_filtrado []structures.Carro

	parametros := mux.Vars(req)

	server.DB.Where(&structures.Carro{Marca: parametros["brand"], Tipo: parametros["motor"], Combustible: parametros["fuel"], Transmision: parametros["transmission"], ModeloNombre: parametros["carType"]}).Find(&findCar_filtrado)

	json.NewEncoder(w).Encode(&findCar_filtrado)

}
