package user

import (
	to "guavacoder/jiu/tools"
	"net/http"
)

type UserService struct {
	repo *UserRepository
}

func NewUserSerivice(r *UserRepository) UserService {
	return UserService{repo: r}
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	users := serv.repo.GetUsers()

	statusCode, response := to.HandleJsonMarshal(users)

	to.WriteJsonResponse(w, to.Response{StatusCode: statusCode, Body: response})
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	params, _ := to.ParseUrlParams(r.URL.String())
	users := serv.repo.GetUserByConditions(
		params.Get("name"),
		params.Get("email"),
	)

	statusCode, response := to.HandleJsonMarshal(users)

	to.WriteJsonResponse(w, to.Response{StatusCode: statusCode, Body: response})
	return
}
