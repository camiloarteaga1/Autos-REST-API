package structures

import "gorm.io/gorm"

type Carro struct {
	gorm.Model

	CarroID       uint   `gorm:"not null;unique_index" json:"IDcar"`
	EstadoReserva bool   `gorm:"default:false" json:"reserved"`
	Precio        uint   `gorm:"not null" json:"price"`
	ModeloYear    uint   `gorm:"not null" json:"modelYear"`
	ModeloNombre  string `gorm:"not null" json:"carType"`
	Marca         string `gorm:"not null" json:"brand"`
	Transmision   string `gorm:"not null" json:"transmission"`
	Combustible   string `gorm:"not null" json:"fuel"`
	Tipo          string `gorm:"not null" json:"motor"`
	Puertas       uint   `gorm:"not null" json:"numDoors"`
	ABS           bool   `gorm:"not null" json:"brakes_abs"`
	Img           string `gorm:"not null" json:"pic"`
}
