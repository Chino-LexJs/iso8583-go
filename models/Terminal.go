package models

import (
	"fmt"

	"github.com/Chino-LexJs/iso8583-go/db"
	"gorm.io/gorm"
)

type Terminal struct {
	gorm.Model

	Terminal_id string `gorm:"primaryKey;uniqueIndex;type:varchar(255);not null" json:"terminal_id"`
	Direction   string `gorm:"type:varchar(255);not null" json:"direction"`
}

type RequestPaymentMessage struct {
	Entry_mode string `json:"entry_mode"`
	Device     struct {
		Serialnr string `json:"serialnr"`
		Version  string `json:"version"`
		Counter  uint   `json:"counter"`
	} `json:"device"`
	Key struct {
		Check_value string `json:"check_value"`
		Crc32       string `json:"crc32"`
		Name        string `json:"name"`
		Rsa         string `json:"rsa"`
	} `json:"key"`
	Localtime string `json:"localtime"`
	Amount    string `json:"amount"`
}

type RequestPaymentResponse struct {
	Servertime string `json:"servertime"`
	Rc         int    `json:"rc"`
	Rcmessage  string `json:"rcmessage"`
	Id         int    `json:"id"`
	Workkey    struct {
		Ksn         string
		Key         string
		Crc32       string
		Check_value string
	} `json:"workkey"`
}

func AddTerminal() {
	terminal := Terminal{Terminal_id: "PB04204S60977", Direction: "CALLE FALSA 123"}
	var terminal_db Terminal
	findedTerminal_db := db.DB.Where("terminal_id = ?", "PB04204S60977").First(&terminal_db)
	err_findedTerminal_db := findedTerminal_db.Error
	if err_findedTerminal_db != nil {
		createdTerminal := db.DB.Create(&terminal)
		err := createdTerminal.Error
		if err != nil {
			fmt.Println("ERROR 02: Error al insertar terminales")
		}
	}
}
