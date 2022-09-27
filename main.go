package main

import (
	"net/http"

	"github.com/Chino-LexJs/iso8583-go/db"
	"github.com/Chino-LexJs/iso8583-go/models"
	"github.com/Chino-LexJs/iso8583-go/routes"
	"github.com/gorilla/mux"
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
	// creating routes to server
	r := mux.NewRouter()
	// Home router
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/requestPayment", routes.PostRequestPayment).Methods("POST")
	r.HandleFunc("/executePayment", routes.ExecutePayment).Methods("POST")
	// start server
	http.ListenAndServe(":3000", r)
}
