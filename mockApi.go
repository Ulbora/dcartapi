package dcartapi

import (
	px "github.com/Ulbora/GoProxy"
)

//MockAPI MockAPI
type MockAPI struct {
	PrivateKey      string
	Proxy           px.Proxy
	MockOrder       *Order
	MockShipmentRes *[]ShipmentResponse
}

//GetNew GetNew
func (a *MockAPI) GetNew() DcartAPI {
	return a
}

//GetOrder GetOrder
func (a *MockAPI) GetOrder(invoice string, secureURL string, token string) *Order {
	return a.MockOrder
}

//AddShippingAddress AddShippingAddress
func (a *MockAPI) AddShippingAddress(s *Shipment, oid string, secureURL string, token string) *[]ShipmentResponse {
	return a.MockShipmentRes
}
