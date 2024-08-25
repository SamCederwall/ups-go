package track

import "net/http"

type TrackResults struct {
	trackResponse trackResponse
}

type trackResponse struct{ shipment []shipment }

type shipment struct {
	inquiryNumber string
	upsPackage    []upsPackage
	userRelation  string
	warnings      []warning
}

type upsPackage struct {
	accessPointInformation  struct{ pickupByDate string }
	activity                []activity
	additionalAttribute     []string
	additionalServices      []string
	alternateTrackingNumber alternateTrackingNumber
	currentStatus           []status
	deliveryDate            []deliveryDate
	deliveryInformation     deliveryInformation
	deliveryTime            deliveryTime
	dimension               dimension
	isSmartPackage          bool
	milestones              []milestone
	packageAddress          []packageAddresses
	paymentInformation      []paymentInformation
	referenceNumber         []referenceNumber
	service                 service
	statusCode              string
	statusDescription       string
	suppressionIndicators   string
	trackingNumber          string
	ucixStatus              string
	weight                  weight
}

type weight struct {
	unitOfMeasurement string
	weight            string
}

type service struct {
	code        string
	description string
	levelCode   string
}

type referenceNumber struct {
	number              string
	referenceNumberType string `json:"type"`
}

type paymentInformation struct {
	amount                 string
	currency               string
	id                     string
	paid                   string
	paymentMethod          string
	paymentInformationType string `json:"type"`
}

type packageAddresses struct {
	packageAddress []packageAddress
	packageCount   int
}

type packageAddress struct {
	address            string
	attentionName      string
	name               string
	packageAddressType string `json:"type"`
}

type milestone struct {
	category       string
	code           string
	current        string
	description    string
	linkedActivity string
	state          string
	subMilestone   string
}

type dimension struct {
	height          string
	length          string
	unitOfDimension string
	width           string
}

type deliveryTime struct {
	endTime          string
	startTime        string
	deliveryTimeType string `json:"type"`
}

type accessPointInformation struct{ pickupByDate string }

type alternateTrackingNumber struct {
	number                      int
	alternateTrackingNumberType string `json: "type"`
}

type deliveryInformation struct {
	deliveryPhoto deliveryPhoto
	location      string
	receivedBy    string
	signature     signature
	pod           pod
}

type deliveryPhoto struct {
	isNonPostalCodeCountry bool
	photo                  string
	photoCaptureInd        string
	photoDispositionCode   string
}

type signature struct{ image string }
type pod struct{ content string }

type deliveryDate struct {
	date             string
	deliveryDateType string `json:"type"`
}

type status struct {
	code        string
	description string
	simplified  string
	statusCode  string
	statusType  string `json:"type"`
}

type activity struct {
	date      string
	gmtDate   string
	gmtOffset string
	location  string
	status    string
	time      string
}

type warning struct {
	code    string
	message string
}

type TrackClient struct {
	accessToken string
	HttpClient  *http.Client
}

func NewTrackClient(accessToken string)
