package tools

import (
	"fmt"
	"net/http"
)

// JiuInternalServerError is a function that handles internal server errors.
// The error should be cause by data format parsing, database connection error,
// or other server side errors.
// If the error is nil, it returns http.StatusOK and an empty byte array.
func CheckInternalServerError(err error, res []byte) (int, []byte) {
	if err != nil {
		return http.StatusInternalServerError, []byte(fmt.Sprintf("{Error: %s}", err.Error()))
	}
	return http.StatusOK, res
}
