package structures

import "gorm.io/gorm"

type Reserva struct {
	gorm.Model

	ReservaID  uint   `gorm:"not null;unique_index" json:"IDr"`
	Usuario_ID uint   `gorm:"not null;" json:"userID"`
	Carro_ID   uint   `gorm:"not null" json:"carID"`
	Resumen    string `gorm:"not null" json:"summary"`
}
