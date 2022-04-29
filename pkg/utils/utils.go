// Data code gets from User is on JSON. We need to convert it to Golang data using Unmarshalling of JSON

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}) { // Http request body from the User is on var r
	if body, err := ioutil.ReadAll(r.Body); err == nil { // Read the Body of the request using ioutil package
		if err := json.Unmarshal([]byte(body), x); err != nil { // Then Unmarshall it
			return
		}
	}
}
