package dcartapi

import (
	"fmt"

	"testing"
)

//GatewayClusterRouteURL url
type GatewayClusterRouteURL struct {
	RouteID           int64  `json:"routeId"`
	Route             string `json:"route"`
	URLID             int64  `json:"urlId"`
	Name              string `json:"name"`
	URL               string `json:"url"`
	Active            bool   `json:"active"`
	CircuitOpen       bool   `json:"circuitOpen"`
	OpenFailCode      int    `json:"openFailCode"`
	FailoverRouteName string `json:"failoverRouteName"`
}

func Test_getRequestBadURL(t *testing.T) {
	var u = ""
	r, err := GetRequest(u, "bad method", nil)
	if err != true {
		t.Fail()
	} else if r != nil {
		t.Fail()
	}

}

func Test_getRequest(t *testing.T) {
	var u = "http://www.myapigateway.com"
	r, err := GetRequest(u, "GET", nil)
	if err == true {
		t.Fail()
	} else if r == nil {
		t.Fail()
	}

}

func Test_getRequestBadURL2(t *testing.T) {
	var u = ""
	var b = []byte("test")
	r, err := GetRequest(u, "bad method", &b)
	if err != true {
		t.Fail()
	} else if r != nil {
		t.Fail()
	}

}

func Test_getRequest2(t *testing.T) {
	var u = "http://www.myapigateway.com"
	var b = []byte("test")
	r, err := GetRequest(u, "POsT", &b)
	if err == true {
		t.Fail()
	} else if r == nil {
		t.Fail()
	}

}

func Test_getJSONEncode(t *testing.T) {
	var p GatewayClusterRouteURL
	j := GetJSONEncode(&p)
	fmt.Print("j in getJsonEncoding: ")
	fmt.Println(j)
	if j == nil {
		t.Fail()
	}
}
