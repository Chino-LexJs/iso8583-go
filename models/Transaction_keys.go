package models

import (
	"fmt"

	"github.com/Chino-LexJs/iso8583-go/db"
	"gorm.io/gorm"
)

type Transaction_keys struct {
	gorm.Model

	ID          uint   `gorm:"primaryKey;autoIncrement:true;not null" json:"transaction_keys_id"`
	Timestamp   string `gorm:"type:varchar(255);not null" json:"timestamp"`
	Check_value string `gorm:"type:varchar(255);not null" json:"check_value"`
	Crc32       string `gorm:"type:varchar(255);not null" json:"crc32"`
	Name        string `gorm:"type:varchar(255);not null" json:"name"`
	Rsa         string `gorm:"type:varchar(600);not null" json:"rsa"`
	Ksn         string `gorm:"type:varchar(255);not null" json:"ksn"`
	Workkey_key string `gorm:"type:varchar(255);not null" json:"workkey_key"`
	TerminalID  string `gorm:"type:varchar(255);not null" json:"terminal_id"`
}

func AddTransaction() {
	transaction := Transaction_keys{Timestamp: "Wed Sep 07 2022 12:00:56 GMT-0300", Check_value: "CAA9B0", Crc32: "BF8425A0", Name: "A000BZPY72", Rsa: "8481B7D10576D49CE1AACACF284B13256D8313A104F9C68434E3A931759F659917BD7434198F5A358DCEF0F615FD6D84332710C30ABCD050C5D96752658AB02AFAAF053B3F0B9997DD02D12472B06EC9F0A8F5E740486E875F467572E39C0FC386EDBD882D624273B1AA44945942BAC597CC333CACB6C334743E5495E708A9B10D1A3461ED58F5A48000A1862DAB658B8B4A6F0BB45617E5AAA0538A12F776A9CC752BC8070802ECB5388A3C14D811D318378E13639DB7E96BD3824C127BC5B224A137470DCF7547961EA344B5138898302F04327B3846AA586E036AF8F17FEFF061A431223B15A8E185B85343D1A2387B274E51BD64833D15C7E0F8C6BEE861", Ksn: "00000101403388200001", Workkey_key: "AEA40693C054BE3A65E14D572C58EAB9", TerminalID: "PB04204S60977"}
	createdTransaction := db.DB.Create(&transaction)
	err := createdTransaction.Error
	if err != nil {
		fmt.Println("ERROR 03: Error al insertar transacciones")
	}
}
