package models

import (
	"fmt"
	"log"

	"github.com/Chino-LexJs/iso8583-go/db"
	"gorm.io/gorm"
)

type Terminal struct {
	gorm.Model

	Terminal_id string `gorm:"primaryKey;uniqueIndex;type:varchar(255);not null" json:"terminal_id"`
	Direction   string `gorm:"type:varchar(255);not null" json:"direction"`
}

func AddTerminal() {
	fmt.Println("Agregando terminales...")
	terminal := Terminal{Terminal_id: "PB04204S60977", Direction: "CALLE FALSA 123"}
	createdTerminal := db.DB.Create(&terminal)
	err := createdTerminal.Error
	if err != nil {
		fmt.Println("ERROR 02: Error al insertar terminales")
		log.Println(err)
	}
	fmt.Println("\n\nNo te preocupes capo el codigo sigue normalmente")
}
