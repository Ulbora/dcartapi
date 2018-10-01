package dcartapi

//API API
type API struct {
	PrivateKey string
	Token      string
	SecureURL  string
}

//GetOrder GetOrder
func (a *API) GetOrder(invoice string) *Order {
	return nil
}

//AddShippingAddress AddShippingAddress
func (a *API) AddShippingAddress(s Shipment) *[]ShipmentResponse {
	return nil
}
