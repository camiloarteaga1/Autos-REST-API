package structures

import "gorm.io/gorm"

type Usuario struct {
	gorm.Model

	Email         string    `gorm:"not null;unique_index" json:"mail"`
	NombreUsuario string    `gorm:"not null;type:varchar(32)" json:"username"`
	Password      string    `gorm:"not null" json:"password"`
	Reservas      []Reserva `json:"reservation"`
}
