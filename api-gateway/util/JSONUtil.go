package util

import (
	"encoding/json"
	"net/http"
)

func GetJsonIC(response *http.Response, target interface{}) error {
	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}
