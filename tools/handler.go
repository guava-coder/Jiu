package tools

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Response struct {
	StatusCode int
	Body       []byte
}

func WriteJsonResponse(w http.ResponseWriter, res Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(res.Body)
}

func JiuInternalServerError(err error) (statusCode int, response []byte) {
	if err != nil {
		statusCode = http.StatusInternalServerError
		response = []byte(fmt.Sprintf("{Error: %s}", err))
	}
	return
}

func ParseUrlParams(urlStr string) (url.Values, error) {
	value, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	params, err := url.ParseQuery(value.RawQuery)
	return params, err
}

func HandleJsonMarshal(obj interface{}) (statusCode int, response []byte) {
	res, err := json.Marshal(obj)

	if err == nil {
		return http.StatusOK, res
	} else {
		return JiuInternalServerError(err)
	}
}
