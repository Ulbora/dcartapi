package dcartapi

import (
	cm "dcartapi/common"
	"log"
	"net/http"
	"strings"
)

//API API
type API struct {
	PrivateKey string
	Token      string
	SecureURL  string
}

//GetOrder GetOrder
func (a *API) GetOrder(invoice string) *Order {
	var rtn Order
	var odrs []Order
	var url = getOrder + invoice
	log.Println("url: ", url)
	req, fail := cm.GetRequest(url, http.MethodGet, nil)
	if !fail {
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("PrivateKey", a.PrivateKey)
		req.Header.Set("Token", a.Token)
		req.Header.Set("SecureURL", a.SecureURL)
		code := cm.ProcessServiceCall(req, &odrs)
		log.Println("code: ", code)
		log.Println("odrs: ", odrs)
		if len(odrs) > 0 {
			rtn = odrs[0]
		}
	}
	return &rtn
}

//AddShippingAddress AddShippingAddress
func (a *API) AddShippingAddress(s *Shipment, oid string) *[]ShipmentResponse {
	var rtn []ShipmentResponse
	var url = newShippingAddress
	url = strings.Replace(url, "OrderId", oid, 1)
	log.Println("url: ", url)
	aJSON := cm.GetJSONEncode(s)
	req, fail := cm.GetRequest(url, http.MethodPost, aJSON)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("PrivateKey", a.PrivateKey)
		req.Header.Set("Token", a.Token)
		req.Header.Set("SecureURL", a.SecureURL)
		code := cm.ProcessServiceCall(req, &rtn)
		log.Println("code: ", code)
		log.Println("rtn: ", rtn)
	}
	return &rtn
}
