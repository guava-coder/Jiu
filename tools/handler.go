package tools

import (
	"encoding/json"
	"net/url"
)

// ParseUrlParams parses the URL parameters from the given URL string.
//
// It takes a URL string as a parameter and returns the parsed URL values and an error.
func ParseUrlParams(urlStr string) (url.Values, error) {
	value, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	params, err := url.ParseQuery(value.RawQuery)
	return params, err
}

// HandleJsonMarshal handles the JSON marshaling of the given object.
//
// It takes an interface{} object as a parameter and returns the HTTP status code and the response byte array.
// The function marshals the object into JSON using the json.Marshal function. If the marshaling is successful,
// it returns the status code http.StatusOK and the marshaled object as a byte array. If there is an error during
// marshaling, it returns http.StatusInternalServerError and an error message as a byte array.
func HandleJsonMarshal(obj interface{}) (int, []byte) {
	res, err := json.Marshal(obj)

	return CheckInternalServerError(err, res)
}
