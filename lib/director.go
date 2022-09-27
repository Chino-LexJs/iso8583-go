package lib

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Chino-LexJs/iso8583-go/models"
)

type Director struct {
	builder IBuilder
}

func NewDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) GetBuilder() IBuilder {
	return d.builder
}

func (d *Director) BuildRequestMessage(id_request uint, request_message models.RequestPaymentMessage) Message {
	d.builder.setHeader()
	d.builder.setMti()
	d.builder.setBitmap()
	d.builder.setDataElements(addDataElements(id_request, request_message))
	return d.builder.getMessage()
}
func addDataElements(id_request uint, request_message models.RequestPaymentMessage) string {
	a := strings.Replace(request_message.Amount, ".", "", -1)
	amount, _ := strconv.Atoi(a)
	fmt.Printf("\nAmount: %d\n", amount)
	var dataElements = map[int]string{
		1:   "000000001000018C",
		3:   "000000",
		4:   strings.Replace(fmt.Sprintf("%012d", amount), ".", "", -1),
		7:   fmt.Sprintf("%02d%02d%02d%02d%02d", int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second()),
		11:  fmt.Sprintf("%06d", 4),
		12:  fmt.Sprintf("%02d%02d%02d", time.Now().Hour(), time.Now().Minute(), time.Now().Second()),
		13:  fmt.Sprintf("%02d%02d", int(time.Now().Month()), time.Now().Day()),
		17:  fmt.Sprintf("%02d%02d", int(time.Now().Month()), time.Now().Day()),
		18:  "5399",
		22:  "901", // entry mode viene de RequestPaymentMessage
		25:  "00",
		32:  "1109000000003",
		37:  fmt.Sprintf("%012d", id_request),
		42:  fmt.Sprintf("%015s", request_message.Device.Serialnr),
		43:  "0000000000000000000000000000000000000000",
		48:  "027000000000000000000000000000",
		49:  "484",
		60:  "0160000000000000000",
		61:  "0190000000000000000000",
		63:  "0010", // Tokens ES y EZ
		100: "010",
		120: "02900000000000000000000000000000",
		121: "02000000000000000000000",
		125: "012ADINTR000000",
		126: "03800000000000000000000000000000000000000",
	}
	var fields []string
	keys := make([]int, 0, len(dataElements))
	for k := range dataElements {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		fields = append(fields, dataElements[key])
	}
	isomessage := strings.Join(fields, "")
	return isomessage
}
