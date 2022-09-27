package utils

import (
	"fmt"
	"strconv"
	"time"
)

type ExecutePaymentResponse struct {
	Id            uint   //debe de ser numerico sin leading zeros
	Timestamp     string // mandamos un timestamp de la respuesta "timestamp" YYYY-mm-dd HHmmss
	Rc            int    // request_status
	Rcdatetime    string // fecha y hora que proporciono Prosa en 0210
	Rcmessage     string //  (pueden ser muchos diferentes, va a depender de tabla Prosa y codigo)
	Ticket        uint   // folio? system trace audit uint
	Authorization string // Authorization ID Response
	Keys_expired  bool   // indica si hay que actualizar llaves
}

func Unpack(isomessage string) ExecutePaymentResponse {
	bitmap := fmt.Sprintf(isomessage[16:32])
	secondbitmap := fmt.Sprintf(isomessage[32:48])
	fmt.Println(bitmap)
	fmt.Println(secondbitmap)
	i, err := strconv.ParseUint(bitmap, 16, 64)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fields := fmt.Sprintf("%024b\n", i)
	i2, errr := strconv.ParseUint(secondbitmap, 16, 64)
	if errr != nil {
		fmt.Printf("%s", err)
	}
	fieldsSecond := fmt.Sprintf("%064b\n", i2)
	var fieldsNumber []uint
	for index, v := range fields {
		if uint(v) == 49 {
			fieldsNumber = append(fieldsNumber, uint(index+1))
		}
	}
	for index, v := range fieldsSecond {
		if uint(v) == 49 {
			fieldsNumber = append(fieldsNumber, uint(index+65))
		}
	}
	fieldsMap := make(map[uint]string)
	init := 32
	for _, f := range fieldsNumber {
		if f == 1 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+16])
			init += 16
		}
		if f == 3 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+6])
			init += 6
		}
		if f == 4 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+12])
			init += 12
		}
		if f == 7 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+10])
			init += 10
		}
		if f == 11 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+6])
			init += 6
		}
		if f == 12 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+6])
			init += 6
		}
		if f == 13 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+4])
			init += 4
		}
		if f == 17 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+4])
			init += 4
		}
		if f == 18 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+4])
			init += 4
		}
		if f == 22 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+3])
			init += 3
		}
		if f == 32 || f == 100 {
			long := fmt.Sprintf(isomessage[init : init+2])
			u64, _ := strconv.ParseUint(long, 10, 32)
			intLong := uint(u64) + 2
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+int(intLong)])
			init += int(intLong)
		}
		if f == 35 {
			long := fmt.Sprintf(isomessage[init : init+2])
			u64, _ := strconv.ParseUint(long, 10, 32)
			intLong := uint(u64) + 2 + 16
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+int(intLong)])
			init += int(intLong)
		}
		if f == 37 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+12])
			init += 12
		}
		if f == 38 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+6])
			init += 6
		}
		if f == 39 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+2])
			init += 2
		}
		if f == 41 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+16])
			init += 16
		}
		if f == 48 || f == 61 || f == 120 || f == 121 || f == 125 {
			long := fmt.Sprintf(isomessage[init : init+3])
			u64, _ := strconv.ParseUint(long, 10, 32)
			intLong := uint(u64) + 3
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+int(intLong)])
			init += int(intLong)
		}
		if f == 49 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+3])
			init += 3
		}
		if f == 60 {
			fieldsMap[f] = fmt.Sprintf(isomessage[init : init+19])
			init += 19
		}
	}
	if val, ok := fieldsMap[63]; ok {
		// inicio de llaves
		fmt.Printf("\n61 exists and this is the value: %v\n", val)
		executeReponse := ExecutePaymentResponse{}
		return executeReponse
	}
	executeReponse := ExecutePaymentResponse{
		Id:            parseInt(fieldsMap[37]),
		Timestamp:     time.Now().Format("2006-01-02 15:04:05"),
		Rc:            responseCode(fieldsMap[39]),
		Rcdatetime:    fieldsMap[13],
		Rcmessage:     responseCodeString(fieldsMap[39]),
		Ticket:        parseInt(fieldsMap[11]),
		Authorization: fieldsMap[38],
		Keys_expired:  false,
	}
	// json, _ := json.Marshal(executeReponse)
	// fmt.Printf("\n execute response: %v\n", string(json))
	// fmt.Printf("\n%v", fieldsMap)
	return executeReponse
}

func parseInt(number string) uint {
	u64, _ := strconv.ParseUint(number, 10, 64)
	intnumber := uint(u64)
	return intnumber
}
func responseCode(p39 string) int {
	fmt.Println(p39)
	if p39 == "00" {
		return int(0)
	}
	return 1
}
func responseCodeString(p39 string) string {
	if p39 == "00" {
		return "APROBADA"
	}
	return "DESAPROBADA"
}
