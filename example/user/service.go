package user

import (
	"encoding/json"
	to "guavacoder/jiu/tools"
	"net/http"
)

type UserService struct {
	repo *UserRepository
}

func NewUserSerivice(r *UserRepository) UserService {
	return UserService{repo: r}
}

func handleJsonMarshal(obj interface{}) (statusCode int, response []byte) {
	res, err := json.Marshal(obj)

	if err == nil {
		return http.StatusOK, res
	} else {
		return to.JiuInternalServerError(err)
	}
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	users := serv.repo.GetUsers()

	var response []byte
	statusCode, response = handleJsonMarshal(users)

	to.WriteJsonResponse(w, to.Response{StatusCode: statusCode, Body: response})
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	params, _ := to.ParseUrlParams(r.URL.String())
	users := serv.repo.GetUserByConditions(
		params.Get("name"),
		params.Get("email"),
	)

	var response []byte
	statusCode, response = handleJsonMarshal(users)

	to.WriteJsonResponse(w, to.Response{StatusCode: statusCode, Body: response})
	return
}
