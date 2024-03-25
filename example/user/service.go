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

type Response struct {
	StatusCode int
	Body       []byte
}

func (serv UserService) response(w http.ResponseWriter, users []User) {
	statusCode, response := to.HandleJsonMarshal(users)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	users := serv.repo.GetUsers()

	serv.response(w, users)
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	params, _ := to.ParseUrlParams(r.URL.String())
	users := serv.repo.GetUserByConditions(
		params.Get("name"),
		params.Get("email"),
	)

	serv.response(w, users)
	return
}
