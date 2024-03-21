package controller

import (
	"fmt"
	user "guavacoder/jiu/user"
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

func (c UserController) Get(suffix string) string {
	return fmt.Sprintf("GET %s%s", c.prefix, suffix)
}

func (c UserController) Run() {
	c.getUsers()
}

func (c UserController) getUsers() {
	url := c.Get("all")
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		lg.PrintlnLatesy(url, func() int { return c.service.GetUsers(w, r) })
	})
}
