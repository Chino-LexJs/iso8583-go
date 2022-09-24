package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Message_request struct {
	gorm.Model

	ID      uint           `gorm:"primaryKey;autoIncrement:true;not null" json:"request_id"`
	Mti     string         `gorm:"type:varchar(50);not null" json:"mti"`
	Content datatypes.JSON `gorm:"not null" json:"content"`
}
