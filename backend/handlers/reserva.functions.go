package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/camiloarteaga1/Autos-REST-API.git/server"
	"github.com/camiloarteaga1/Autos-REST-API.git/structures"
)

func ReservationsHandler(w http.ResponseWriter, req *http.Request) {

	var reserva structures.Reserva

	json.NewDecoder(req.Body).Decode(&reserva) //Decodificacion del json enviado desde el cliente al servidor

	//Guardado en server

	createdReserva := server.DB.Create(&reserva) //Se trata de guardar la reserva en DB

	err := createdReserva.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //codigo 400
		w.Write([]byte(err.Error()))

		return
	}

	//luego, se debe cambiar el estado de dicho carro a TRUE

	server.DB.Model(&structures.Carro{}).Where("reservation = ? AND car_id = ?", false, reserva.Carro_ID).Update("reservation", true)

	json.NewEncoder(w).Encode(&reserva) //Se envia la reserva

}

func DeleteReservationsHandler(w http.ResponseWriter, req *http.Request) {

	var reserva structures.Reserva

	var findReserva structures.Reserva

	json.NewDecoder(req.Body).Decode(&reserva)

	reservaID := reserva.ReservaID

	server.DB.Where("reservation = ?", reservaID).First(&findReserva)

	if findReserva.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("Reserve not found"))

		return

	}

	server.DB.Model(&structures.Carro{}).Where("car_id = ?", findReserva.Carro_ID).Update("reservation", false)

	server.DB.Unscoped().Delete(&findReserva) //Eliminación de la reserva
	w.WriteHeader(http.StatusNoContent)       //204

}

func ReportHandler(w http.ResponseWriter, req *http.Request) {

	var reporte structures.Reporte

	var user structures.Usuario //se debe buscar el usuario, sus reservas asociadas y luego
	//iterar cada reserva para buscar la data de los carros
	var precioAcumulado int64

	precioAcumulado = 0

	json.NewDecoder(req.Body).Decode(&user)

	server.DB.First(&user, "mail = ?", user.Email) // email or mail

	if user.ID == 0 {

		w.WriteHeader(http.StatusNotFound) //404
		w.Write([]byte("User not found"))

		return
	}

	server.DB.Model(&user).Association("Reservas").Find(&user.Reservas)

	//se itera sobre la estructura de reservas

	for _, reservation := range user.Reservas {

		var carro structures.Carro

		server.DB.Where("car_id = ?", reservation.Carro_ID).Find(&carro)

		precioAcumulado += int64(carro.Precio)

		reporte.Detalles = append(reporte.Detalles, carro)
	}

	reporte.PrecioTotal = precioAcumulado

	json.NewEncoder(w).Encode(&reporte)

}
