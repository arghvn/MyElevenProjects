// for Marshalling and Unmarshalling

package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func parseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReaadALL(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
