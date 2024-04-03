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

	c.addUser()
	c.updateUser()
	c.deleteUser()
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

func (c UserController) addUser() {
	url := fmt.Sprintf("POST %sadd", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.AddUser(w, r) })
	})
}

func (c UserController) updateUser() {
	url := fmt.Sprintf("PUT %supdate", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.UpdateUser(w, r) })
	})
}

func (c UserController) deleteUser() {
	url := fmt.Sprintf("DELETE %sdelete", c.prefix)
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return c.service.DeleteUser(w, r) })
	})
}
