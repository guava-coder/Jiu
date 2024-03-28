package controller

import (
	"fmt"
	user "guavacoder/jiu/example/user"
	to "guavacoder/jiu/tools"
	"net/http"
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
	c.getUserById()
}

func (c UserController) getUsers() {
	url := fmt.Sprintf("GET %s%s", c.prefix, "all")
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.GetUsers(w, r) })
	})
}

func (c UserController) getUserByConditions() {
	url := fmt.Sprintf("GET %s", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.GetUserByConditions(w, r) })
	})
}

func (c UserController) getUserById() {
	url := fmt.Sprintf("GET %s{id}", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.GetUserById(w, r) })
	})
}
