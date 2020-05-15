package dcartapi

import (
	"log"
	"testing"
)

func TestMockAPI_GetOrder(t *testing.T) {
	var api MockAPI
	var od Order
	od.OrderID = 5
	api.MockOrder = &od

	var secureURL string
	var token string

	dapi := api.GetNew()
	o := dapi.GetOrder("1041", secureURL, token)
	log.Println("order: ", o)
	if o.OrderID == 0 {
		t.Fail()
	}
}

func TestMockAPI_AddShippingAddress(t *testing.T) {
	var api MockAPI
	var shs []ShipmentResponse
	var sh ShipmentResponse
	sh.Key = "test"
	shs = append(shs, sh)
	api.MockShipmentRes = &shs

	var secureURL string
	var token string

	var dapi DcartAPI
	dapi = &api
	var s Shipment
	s.ShipmentID = 0
	s.ShipmentFirstName = "FFL"
	s.ShipmentLastName = "Lic # 123456"
	s.ShipmentCompany = "Bobs Arms"
	s.ShipmentAddress = "123 Marietta st"
	s.ShipmentCity = "Atlanta"
	s.ShipmentState = "GA"
	s.ShipmentZipCode = "36847"
	s.ShipmentCountry = "USA"
	s.ShipmentPhone = "800 458 8888"
	s.ShipmentTax = 0
	s.ShipmentWeight = 1
	s.ShipmentTrackingCode = ""
	s.ShipmentUserID = ""
	s.ShipmentNumber = 1
	s.ShipmentAddressTypeID = 0

	res := dapi.AddShippingAddress(&s, "63", secureURL, token)
	log.Println("res: ", res)
	if len(*res) == 0 {
		t.Fail()
	}
}
