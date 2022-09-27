package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Chino-LexJs/iso8583-go/db"
	"github.com/Chino-LexJs/iso8583-go/models"
	"gorm.io/datatypes"
)

func PostRequestPayment(w http.ResponseWriter, r *http.Request) {
	var terminal models.RequestPaymentMessage
	json.NewDecoder(r.Body).Decode(&terminal)
	terminalJson, _ := json.Marshal(terminal)
	terminalJsonString := string(terminalJson)
	message_request := models.Message_request{Mti: "0200", Content: datatypes.JSON([]byte(terminalJsonString))}

	createdRequest := db.DB.Create(&message_request)
	err_createdRequest := createdRequest.Error
	if err_createdRequest != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err_createdRequest.Error()))
	}
	var terminal_db models.Terminal
	findedTerminal_db := db.DB.Where("terminal_id = ?", terminal.Device.Serialnr).First(&terminal_db)
	err_findedTerminal_db := findedTerminal_db.Error
	if err_findedTerminal_db != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err_findedTerminal_db.Error()))
	}
	var transaction_key_db models.Transaction_keys
	findedTransaction_db := db.DB.Where("terminal_id = ?", terminal.Device.Serialnr).First(&transaction_key_db)
	err_findedTransaction_db := findedTransaction_db.Error
	if err_findedTransaction_db != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err_findedTransaction_db.Error()))
	}
	date := time.Now()
	requestPaymentResponse := &models.RequestPaymentResponse{
		Servertime: date.String(),
		Rc:         -1,
		Rcmessage:  "bla bla bla",
		Id:         int(message_request.ID),
		Workkey: struct {
			Ksn         string
			Key         string
			Crc32       string
			Check_value string
		}{
			Ksn:         transaction_key_db.Ksn,
			Key:         transaction_key_db.Workkey_key,
			Crc32:       transaction_key_db.Crc32,
			Check_value: transaction_key_db.Check_value,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&requestPaymentResponse)
}
