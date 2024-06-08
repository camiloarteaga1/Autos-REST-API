package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/camiloarteaga1/Autos-REST-API.git/server"
	"github.com/camiloarteaga1/Autos-REST-API.git/structures"
)

func LoginHandler(w http.ResponseWriter, req *http.Request) {

	var user structures.Usuario

	var findUser structures.Usuario

	json.NewDecoder(req.Body).Decode(&user)

	email := user.Email //Contraseña e Email ingresados por el cliente

	passw := user.Password

	server.DB.Where("email = ? AND password = ?", email, passw).First(&findUser) //de busca la exacta coincidencia dentro del server

	if findUser.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Credenciales invalidas"))

		return
	}

	server.DB.Model(&findUser).Association("Reservas").Find(&findUser.Reservas) //Se asocian las reservas de un usuario especifico

	json.NewEncoder(w).Encode(&findUser) //Se retorna toda la data del usuario logeado

}

func RegisterHandler(w http.ResponseWriter, req *http.Request) {

	var user structures.Usuario

	json.NewDecoder(req.Body).Decode(&user) //Decodificacion del json enviado desde el cliente al servidor

	//Guardado en server

	createdUser := server.DB.Create(&user) //Se trata de guardar el usuario en server

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //codigo 400
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user) //Se envia el usuario creado

}

func UsersHandler(w http.ResponseWriter, req *http.Request) {

	var users []structures.Usuario

	server.DB.Find(&users) //Buscamos toda la tabla Usuarios

	json.NewEncoder(w).Encode(&users) //Retorna todas las filas y columnas de la tabla

}
