package user

import (
	"encoding/json"
	"fmt"
	to "guavacoder/jiu/tools"
	"net/http"

	"github.com/google/uuid"
)

type UserService struct {
	repo           *UserRepository
	responseWriter http.ResponseWriter
}

func NewUserSerivice(r *UserRepository) UserService {
	return UserService{
		repo:           r,
		responseWriter: nil,
	}
}

type Response struct {
	StatusCode int
	Body       []byte
}

func (serv UserService) printUserJson(users []User) (statusCode int) {
	statusCode, response := to.MustHandleJsonMarshal(users)

	serv.printJsonResponse(statusCode, response)
	return
}

func (serv UserService) printError(statusCode int, err error) {
	serv.printJsonResponse(statusCode, []byte(fmt.Sprintf("{Error: %s}", err.Error())))
}

func (serv UserService) printJsonResponse(statusCode int, response []byte) {
	w := serv.responseWriter
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	serv.responseWriter = w
	users, err := serv.repo.GetUsers()
	if err == nil {
		statusCode = serv.printUserJson(users)
	} else {
		statusCode = http.StatusNoContent
		serv.printError(statusCode, err)
	}
	return
}

func (serv UserService) GetUserByConditions(w http.ResponseWriter, r *http.Request) (statusCode int) {
	serv.responseWriter = w
	params, err := to.ParseUrlParams(r.URL.String())

	handleConditions := func() {
		users, err := serv.repo.GetUserByConditions(
			params.Get("name"),
			params.Get("email"),
		)
		if err == nil {
			statusCode = serv.printUserJson(users)
		} else {
			statusCode = http.StatusNoContent
			serv.printError(statusCode, err)
		}
	}
	if err == nil {
		handleConditions()
	} else {
		statusCode = http.StatusInternalServerError
		serv.printError(statusCode, err)
	}
	return
}

func (serv UserService) GetUserById(w http.ResponseWriter, r *http.Request) (statusCode int) {
	serv.responseWriter = w
	user, err := serv.repo.GetUserById(r.PathValue("id"))
	if err == nil {
		statusCode = serv.printUserJson([]User{user})
	} else {
		statusCode = http.StatusNotFound
		http.Error(w, err.Error(), statusCode)
	}
	return
}

func (serv UserService) AddUser(w http.ResponseWriter, r *http.Request) (statusCode int) {
	serv.responseWriter = w

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	handleAdd := func() {
		user.Id = uuid.New().String()
		err := serv.repo.AddUser(user)
		if err == nil {
			statusCode = http.StatusOK
			serv.printJsonResponse(statusCode, []byte("User added successfully"))
		} else {
			statusCode = http.StatusBadRequest
			http.Error(w, err.Error(), statusCode)
		}
	}

	if err == nil {
		handleAdd()
	} else {
		statusCode = http.StatusInternalServerError
		http.Error(w, err.Error(), statusCode)
	}
	return
}

func (serv UserService) UpdateUser(w http.ResponseWriter, r *http.Request) (statusCode int) {
	serv.responseWriter = w

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	handleUpdate := func() {
		err = serv.repo.UpdateUser([]User{user})
		if err == nil {
			statusCode = http.StatusOK
			serv.printJsonResponse(statusCode, []byte("User updated successfully"))
		} else {
			statusCode = http.StatusBadRequest
			http.Error(w, err.Error(), statusCode)
		}
	}

	if err == nil {
		handleUpdate()
	} else {
		statusCode = http.StatusInternalServerError
		http.Error(w, err.Error(), statusCode)
	}
	return

}
