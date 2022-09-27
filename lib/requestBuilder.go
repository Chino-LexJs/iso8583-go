package lib

type RequestBuilder struct {
	header       string
	mti          string
	bitmap       string
	dataElements string
}

func newRequestBuilder() *RequestBuilder {
	return &RequestBuilder{}
}
func (b *RequestBuilder) setHeader() {
	b.header = "ISO026000050"
}
func (b *RequestBuilder) setMti() {
	b.mti = "0200"
}
func (b *RequestBuilder) setBitmap() {
	b.bitmap = "B238C4810861801A"
}
func (b *RequestBuilder) setDataElements(data string) {
	b.dataElements = data
}
func (b *RequestBuilder) getMessage() Message {
	return Message{
		Header:       b.header,
		Mti:          b.mti,
		Bitmap:       b.bitmap,
		DataElements: b.dataElements,
	}
}

// fields := []string{
// 	requestMessage.Header,
// 	requestMessage.Mti,
// 	requestMessage.Bitmap,
// 	requestMessage.DataElements,
// }
// isomessage := strings.Join(fields, "")
