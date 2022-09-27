package lib

type IBuilder interface {
	setHeader()
	setMti()
	setBitmap()
	setDataElements(data string)
	getMessage() Message
}

func GetBuilder(builderType string) IBuilder {
	if builderType == "0200" {
		return newRequestBuilder()
	}
	return nil
}
