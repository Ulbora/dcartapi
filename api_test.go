package dcartapi

import (
	"log"
	"os"

	"testing"
)

func TestAPI_GetOrder(t *testing.T) {
	//go test -coverprofile=coverage.out -args privateKey token secureURL
	var api API
	log.Println("args: ", len(os.Args))
	var secureURL string
	var token string
	if len(os.Args) == 5 {
		privateKey := os.Args[2]
		token = os.Args[3]
		secureURL = os.Args[4]
		api.PrivateKey = privateKey
		//api.Token = token
		//api.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		log.Println("token: ", token)
		log.Println("secureURL: ", secureURL)
	}
	var dapi DcartAPI
	dapi = &api
	o := dapi.GetOrder("1041", secureURL, token)
	log.Println("order: ", o)
	if o.OrderID == 0 {
		t.Fail()
	}

}

func TestAPI_AddShippingAddress(t *testing.T) {
	var api API
	log.Println("args: ", len(os.Args))
	var secureURL string
	var token string
	if len(os.Args) == 5 {
		privateKey := os.Args[2]
		token = os.Args[3]
		secureURL = os.Args[4]
		api.PrivateKey = privateKey
		//api.Token = token
		//api.SecureURL = secureURL

		log.Println("privateKey: ", privateKey)
		log.Println("token: ", token)
		log.Println("secureURL: ", secureURL)
	}
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
