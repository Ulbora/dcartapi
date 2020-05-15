package dcartapi

import (
	by "bytes"
	"encoding/json"
	"log"
	"net/http"
)

//GetRequest GetRequest
func GetRequest(url string, method string, obj *[]byte) (*http.Request, bool) {
	//fmt.Print("obj: ")
	//fmt.Println(obj)
	var err = false
	var bf *by.Buffer
	var req *http.Request
	var rErr error
	if obj != nil {
		//fmt.Println("in first: ")
		bf = by.NewBuffer(*obj)
		req, rErr = http.NewRequest(method, url, bf)
		if rErr != nil {
			err = true
			log.Print("request err: ")
			log.Println(rErr)
		}
	} else {
		//fmt.Println("in second: ")
		req, rErr = http.NewRequest(method, url, nil)
		if rErr != nil {
			err = true
			log.Print("request err: ")
			log.Println(rErr)
		}
	}
	//fmt.Println(bf)
	return req, err
}

//GetJSONEncode GetJSONEncode
func GetJSONEncode(obj interface{}) *[]byte {
	//fmt.Print("obj in json: ")
	//fmt.Println(obj)
	aJSON, _ := json.Marshal(obj)
	return &aJSON
}
