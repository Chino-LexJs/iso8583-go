package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Chino-LexJs/iso8583-go/db"
	"github.com/Chino-LexJs/iso8583-go/lib"
	"github.com/Chino-LexJs/iso8583-go/models"
	"github.com/Chino-LexJs/iso8583-go/prosa"
	"github.com/Chino-LexJs/iso8583-go/utils"
)

func ExecutePayment(w http.ResponseWriter, r *http.Request) {
	// find request_message
	var executePayment models.ExecutePaymentMessage
	json.NewDecoder(r.Body).Decode(&executePayment)
	var request_db models.Message_request
	findedRequest_db := db.DB.Where("id= ?", executePayment.Id).First(&request_db)
	err_findedRequest := findedRequest_db.Error
	if err_findedRequest != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err_findedRequest.Error()))
	}
	// build iso8583 0200 message to Prosa
	requestBuilder := lib.GetBuilder("0200")
	director := lib.NewDirector(requestBuilder)
	var requestPaymentMessage models.RequestPaymentMessage
	json.Unmarshal([]byte(request_db.Content), &requestPaymentMessage)
	requestMessage := director.BuildRequestMessage(uint(executePayment.Id), requestPaymentMessage)
	isoMessage := fmt.Sprintf("%s%s%s%s", requestMessage.Header, requestMessage.Mti, requestMessage.Bitmap, requestMessage.DataElements)
	done := make(chan string)
	go prosa.SendMessageProsa(isoMessage, done)
	iso8583Message := <-done
	fmt.Printf("\nRequest Message: %v", iso8583Message)
	// armar executePaymentResponse para enviar a terminal
	requestPaymentResponse := utils.Unpack(iso8583Message)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&requestPaymentResponse)
}
