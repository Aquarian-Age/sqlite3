// RecordID PersoIndex PersonID RecordPicture RecordTime RecordType RecordPass Temperature RecordData

package record

// RecordData 字段
type RecordDataStruct struct {
	GauzeResult     int     `json:"GauzeResult"`
	HealthCodeColor string  `json:"HealthCodeColor"`
	ICCardNO        string  `json:"ICCardNO"`
	Latitude        int     `json:"Latitude"`
	Longitude       int     `json:"Longitude"`
	PersonName      string  `json:"PersonName"`
	Price           int     `json:"Price"`
	PriceID         string  `json:"PriceID"`
	QRCode          string  `json:"QRCode"`
	SatetyHatResult int     `json:"SatetyHatResult"`
	Similarity      float32 `json:"Similarity"`
}

// RecordData 字段
type RecordDataStruct1 struct {
	FaceRect struct {
		Height int `json:"Height"`
		Width  int `json:"Width"`
		X      int `json:"X"`
		Y      int `json:"Y"`
	} `json:"FaceRect"`
	GauzeResult     int     `json:"GauzeResult"`
	HealthCodeColor string  `json:"HealthCodeColor"`
	ICCardNO        string  `json:"ICCardNO"`
	Latitude        int     `json:"Latitude"`
	Longitude       int     `json:"Longitude"`
	PersonName      string  `json:"PersonName"`
	Price           int     `json:"Price"`
	PriceID         string  `json:"PriceID"`
	QRCode          string  `json:"QRCode"`
	SatetyHatResult int     `json:"SatetyHatResult"`
	Similarity      float64 `json:"Similarity"`
}

/*
	env CC=arm-linux-gnueabihf-gcc CXX=arm-linux-gnueabihf-g++ \
    CGO_ENABLED=1 GOOS=linux GOARCH=arm GOARM=7 \
    go build -v
*/
