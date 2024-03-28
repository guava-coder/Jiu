package tools

import (
	"net/http"
	"testing"
)

type Obj struct {
	Id    int
	Name  string
	Email string
}

func TestHandler(t *testing.T) {
	url := "http://test.com/user?name=John&email=j@j.com"
	t.Run("test parse url params", func(t *testing.T) {
		_, err := ParseUrlParams(url)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("test handle json marshal", func(t *testing.T) {
		statusCode, response := MustHandleJsonMarshal(Obj{
			Id:    1,
			Name:  "John",
			Email: "j@j.com",
		})
		if statusCode != http.StatusOK {
			t.Fatal(string(response))
		}
	})
}
