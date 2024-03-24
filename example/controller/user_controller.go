package controller

import (
	"fmt"
	user "guavacoder/jiu/example/user"
	"net/http"

	lg "guavacoder/jiu/logger"
)

type UserController struct {
	service user.UserService
	mux     *http.ServeMux
	prefix  string
}

func NewUserController(mux *http.ServeMux, service user.UserService) UserController {
	return UserController{
		service: service,
		mux:     mux,
		prefix:  "/user/",
	}
}

func (c UserController) Run() {
	c.getUsers()
	c.getUserByConditions()
}

func (c UserController) getUsers() {
	url := fmt.Sprintf("GET %s%s", c.prefix, "all")
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		lg.PrintlnLatesy(url, func() int { return c.service.GetUsers(w, r) })
	})
}

func (c UserController) getUserByConditions() {
	url := fmt.Sprintf("GET %s", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		lg.PrintlnLatesy(url, func() int { return c.service.GetUserByConditions(w, r) })
	})
}
