package requestUtil

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func Get(url string, params url.Values) (map[string]interface{}, error) {
	if params != nil {
		url += "?" + params.Encode()
	}
	result := make(map[string]interface{})
	response, err := http.Get(url)
	if err != nil {
		return result, err
	}
	jsonDecoder := json.NewDecoder(response.Body)
	jsonDecoder.Decode(&result)
	return result, nil
}
