package dcartapi

import (
	//cm "dcartapi/common"
	"log"
	"net/http"
	"strings"

	px "github.com/Ulbora/GoProxy"
)

//API API
type API struct {
	PrivateKey string
	Proxy      px.Proxy
}

//GetNew GetNew
func (a *API) GetNew() DcartAPI {
	return a
}

//GetOrder GetOrder
func (a *API) GetOrder(invoice string, secureURL string, token string) *Order {
	var rtn Order
	var odrs []Order
	var url = getOrder + invoice
	//log.Println("url: ", url)
	req, fail := GetRequest(url, http.MethodGet, nil)
	if !fail {
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("PrivateKey", a.PrivateKey)
		req.Header.Set("Token", token)
		req.Header.Set("SecureURL", secureURL)
		//code := ProcessServiceCall(req, &odrs)
		suc, code := a.Proxy.Do(req, &odrs)
		//log.Println("code: ", code)
		//log.Println("odrs: ", odrs)
		if suc && len(odrs) > 0 {
			rtn = odrs[0]
		} else {
			log.Println("get order service call success: ", suc)
			log.Println("url: ", url)
			log.Println("code: ", code)
			log.Println("odrs: ", odrs)
		}
	}
	return &rtn
}

//AddShippingAddress AddShippingAddress
func (a *API) AddShippingAddress(s *Shipment, oid string, secureURL string, token string) *[]ShipmentResponse {
	var rtn []ShipmentResponse
	var url = newShippingAddress
	url = strings.Replace(url, "OrderId", oid, 1)
	//log.Println("url: ", url)
	aJSON := GetJSONEncode(s)
	req, fail := GetRequest(url, http.MethodPost, aJSON)
	if !fail {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("PrivateKey", a.PrivateKey)
		req.Header.Set("Token", token)
		req.Header.Set("SecureURL", secureURL)
		//code := ProcessServiceCall(req, &rtn)
		suc, code := a.Proxy.Do(req, &rtn)
		if code != 200 && code != 201 {
			log.Println("AddShippingAddress service call success: ", suc)
			log.Println("url: ", url)
			log.Println("code in add address: ", code)
			log.Println("rtn: ", rtn)
		}
	}
	return &rtn
}
