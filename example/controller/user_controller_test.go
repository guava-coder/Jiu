package controller

import (
	"io"
	"net/http"
	"testing"
)

func assertStatusNotOk(res *http.Response, t *testing.T) {
	if res.StatusCode != 200 {
		t.Fatal("status code should be 200")
	}
}

func logResponseBody(res *http.Response, t *testing.T) {
	value, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(value))
}

func TestGetUsers(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:8080/user/all")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assertStatusNotOk(res, t)

	logResponseBody(res, t)
}

func TestGetUserByConditions(t *testing.T) {
	res, err := http.Get("http://127.0.0.1:8080/user?name=John&email=j@j.com")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	assertStatusNotOk(res, t)

	logResponseBody(res, t)
}
