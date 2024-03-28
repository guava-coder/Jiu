package user

import (
	"fmt"
	to "guavacoder/jiu/tools"
	"net/http"
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

func (serv UserService) printUserJson(w http.ResponseWriter, users []User) (statusCode int) {
	statusCode, response := to.MustHandleJsonMarshal(users)

	serv.printJsonResponse(w, statusCode, response)
	return
}

func (serv UserService) printError(w http.ResponseWriter, statusCode int, err error) {
	serv.printJsonResponse(w, statusCode, []byte(fmt.Sprintf("{Error: %s}", err.Error())))
}

func (serv UserService) printJsonResponse(w http.ResponseWriter, statusCode int, response []byte) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	users := serv.repo.GetUsers()

	serv.printUserJson(w, users)
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	params, err := to.ParseUrlParams(r.URL.String())
	if err == nil {
		users := serv.repo.GetUserByConditions(
			params.Get("name"),
			params.Get("email"),
		)
		statusCode = serv.printUserJson(w, users)
	} else {
		statusCode = http.StatusInternalServerError
		serv.printError(w, statusCode, err)
	}
	return
}
