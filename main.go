package main

import (
	"github.com/Chino-LexJs/iso8583-go/db"
	"github.com/Chino-LexJs/iso8583-go/models"
)

func main() {
	// connecting db
	db.DBconnection()
	// settings models in db
	db.DB.AutoMigrate(models.Terminal{})
	db.DB.AutoMigrate(models.Transaction_keys{})
	db.DB.AutoMigrate(models.Message_request{})
	models.AddTerminal()
	models.AddTransaction()

}
