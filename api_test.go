package dcartapi

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"testing"

	px "github.com/Ulbora/GoProxy"
)

func TestAPI_GetOrder(t *testing.T) {
	//go test -coverprofile=coverage.out -args privateKey token secureURL
	var api API
	var gp px.MockGoProxy
	var res http.Response
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"OrderID":4, "InvoiceNumber":123, "InvoiceNumberPrefix":"1001"}]`))

	gp.MockResp = &res
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	api.Proxy = &gp
	var secureURL string
	var token string

	dapi := api.GetNew()
	o := dapi.GetOrder("1041", secureURL, token)
	log.Println("order: ", o)
	if o.OrderID == 0 {
		t.Fail()
	}
}

func TestAPI_GetOrderFail(t *testing.T) {
	//go test -coverprofile=coverage.out -args privateKey token secureURL
	var api API
	var gp px.MockGoProxy
	var res http.Response
	res.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"OrderID":4, "InvoiceNumber":123, "InvoiceNumberPrefix":"1001"}]`))

	gp.MockResp = &res
	gp.MockDoSuccess1 = false
	gp.MockRespCode = 200
	api.Proxy = &gp
	var secureURL string
	var token string

	dapi := api.GetNew()
	o := dapi.GetOrder("1041", secureURL, token)
	log.Println("order: ", o)
	if o.OrderID != 0 {
		t.Fail()
	}
}

func TestAPI_AddShippingAddress(t *testing.T) {
	var api API
	var gp px.MockGoProxy
	var res1 http.Response
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"Status":"OK", "Value":"123", "Key":"1001", "Message": "ok"}]`))

	gp.MockResp = &res1
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	api.Proxy = &gp
	log.Println("args: ", len(os.Args))
	var secureURL string
	var token string

	dapi := api.GetNew()
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

func TestAPI_AddShippingAddressFail(t *testing.T) {
	var api API
	var gp px.MockGoProxy
	var res1 http.Response
	res1.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"Status":"OK", "Value":"123", "Key":"1001", "Message": "ok"}]`))

	gp.MockResp = &res1
	gp.MockDoSuccess1 = false
	gp.MockRespCode = 400
	api.Proxy = &gp
	log.Println("args: ", len(os.Args))
	var secureURL string
	var token string

	dapi := api.GetNew()
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
