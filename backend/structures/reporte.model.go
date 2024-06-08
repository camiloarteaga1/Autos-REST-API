package structures

type Reporte struct { //Este modelo NO se guarda en DB

	PrecioTotal int64   `json:"subTotal"`
	Detalles    []Carro `json:"details"`
}
