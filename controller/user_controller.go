package controller

import (
	logger "guavacoder/jiu/logger"
	. "guavacoder/jiu/user"
	"net/http"
)

type UserController struct {
	service UserService
	mux     *http.ServeMux
}

func NewUserController(mux *http.ServeMux, service UserService) UserController {
	return UserController{service: service, mux: mux}
}

func (c UserController) Run() {
	c.getUsers()
}

func (c UserController) getUsers() {
	url := "GET /users"
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		req := func() {
			c.service.GetUsers(w, r)
		}
		logger.Println(url, req)
	})
}
