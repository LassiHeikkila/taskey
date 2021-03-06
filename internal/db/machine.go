package db

import (
	"gorm.io/gorm"
)

type Machine struct {
	gorm.Model
	Name           string `gorm:"unique,not null"`
	Description    string
	OS             string
	Arch           string
	OrganizationID uint     `gorm:"not null"`
	Records        []Record `gorm:"foreignKey:MachineID"`
}
