package controller

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestUserController(t *testing.T) {
	t.Run("test get users", func(t *testing.T) {
		res, err := http.Get("http://127.0.0.1:8080/user/all")
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		assertStatusOk(res, t)

		logResponseBody(res, t)
	})
	t.Run("test get user by conditions", func(t *testing.T) {
		res, err := http.Get("http://127.0.0.1:8080/user/?name=John&email=j@j.com")
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		assertStatusOk(res, t)

		logResponseBody(res, t)
	})
	t.Run("test get user by id", func(t *testing.T) {
		res, err := http.Get("http://127.0.0.1:8080/user/1")
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		assertStatusOk(res, t)

		logResponseBody(res, t)
	})
	t.Run("test add user", func(t *testing.T) {
		data := []byte(`{"Name":"John","Email":"j@j.com"}`)
		res, err := http.Post("http://127.0.0.1:8080/user/add", "application/json", bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
		defer res.Body.Close()

		assertStatusOk(res, t)

		logResponseBody(res, t)
	})
}

func assertStatusOk(res *http.Response, t *testing.T) {
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
