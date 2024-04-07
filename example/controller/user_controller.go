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

func (c UserController) handleFunc(url string, serv func(w http.ResponseWriter, r *http.Request) int) {
	c.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		to.PrintlnLatency(url, func() int { return serv(w, r) })
	})
}

func (c UserController) getUsers() {
	c.handleFunc(fmt.Sprintf("GET %s%s", c.prefix, "all"), c.service.GetUsers)
}

func (c UserController) getUserByConditions() {
	c.handleFunc(fmt.Sprintf("GET %s", c.prefix), c.service.GetUserByConditions)
}

func (c UserController) getUserById() {
	c.handleFunc(fmt.Sprintf("GET %s{id}", c.prefix), c.service.GetUserById)
}

func (c UserController) addUser() {
	c.handleFunc(fmt.Sprintf("POST %sadd", c.prefix), c.service.AddUser)
}

func (c UserController) updateUser() {
	c.handleFunc(fmt.Sprintf("PUT %supdate", c.prefix), c.service.UpdateUser)
}

func (c UserController) deleteUser() {
	c.handleFunc(fmt.Sprintf("DELETE %sdelete", c.prefix), c.service.DeleteUser)
}
