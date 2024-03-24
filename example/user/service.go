package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type UserService struct {
	repo *UserRepository
}

func NewUserSerivice(r *UserRepository) UserService {
	return UserService{repo: r}
}

type Response struct {
	StatusCode int
	Body       []byte
}

func WriteJsonResponse(w http.ResponseWriter, res Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(res.StatusCode)
	w.Write(res.Body)
}

func assertInternalServerError(err error) (statusCode int, response []byte) {
	if err != nil {
		statusCode = http.StatusInternalServerError
		response = []byte(fmt.Sprintf("{Response: %s}", err))
	}
	return
}

func parseUrlParams(r *http.Request) (url.Values, error) {
	value, err := url.Parse(r.URL.String())
	if err != nil {
		return nil, err
	}
	params, err := url.ParseQuery(value.RawQuery)
	return params, err
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	users := serv.repo.GetUsers()

	res, err := json.Marshal(users)
	var response []byte
	statusCode, response = assertInternalServerError(err)

	if err == nil {
		statusCode = http.StatusOK
		response = res
	}

	WriteJsonResponse(w, Response{StatusCode: statusCode, Body: response})
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	params, _ := parseUrlParams(r)
	users := serv.repo.GetUserByConditions(
		params.Get("name"),
		params.Get("email"),
	)

	res, err := json.Marshal(users)
	var response []byte
	statusCode, response = assertInternalServerError(err)

	if err == nil {
		statusCode = http.StatusOK
		response = res
	}

	WriteJsonResponse(w, Response{StatusCode: statusCode, Body: response})
	return
}
