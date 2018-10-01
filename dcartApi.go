package dcartapi

//DcartAPI DcartAPI
type DcartAPI interface {
	GetOrder(invoice string) *Order
	AddShippingAddress(s *Shipment, oid string) *[]ShipmentResponse
}

//Order Order
type Order struct {
	InvoiceNumberPrefix string `json:"InvoiceNumberPrefix"`
	InvoiceNumber       int64  `json:"InvoiceNumber"`
	OrderID             int64  `json:"OrderID"`
}

//Shipment Shipment
type Shipment struct {
	ShipmentID            int    `json:"ShipmentID"`
	ShipmentFirstName     string `json:"ShipmentFirstName"`
	ShipmentLastName      string `json:"ShipmentLastName"`
	ShipmentAddress       string `json:"ShipmentAddress"`
	ShipmentCity          string `json:"ShipmentCity"`
	ShipmentCompany       string `json:"ShipmentCompany"`
	ShipmentState         string `json:"ShipmentState"`
	ShipmentZipCode       string `json:"ShipmentZipCode"`
	ShipmentCountry       string `json:"ShipmentCountry"`
	ShipmentPhone         string `json:"ShipmentPhone"`
	ShipmentTax           int    `json:"ShipmentTax"`
	ShipmentWeight        int    `json:"ShipmentWeight"`
	ShipmentTrackingCode  string `json:"ShipmentTrackingCode"`
	ShipmentUserID        string `json:"ShipmentUserID"`
	ShipmentNumber        int    `json:"ShipmentNumber"`
	ShipmentAddressTypeID int    `json:"ShipmentAddressTypeID"`
}

//ShipmentResponse ShipmentResponse
type ShipmentResponse struct {
	Key     string `json:"Key"`
	Value   string `json:"Value"`
	Status  string `json:"Status"`
	Message string `json:"Message"`
}
