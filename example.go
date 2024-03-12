package main

import (
	"fmt"
	"log"
	"net/http"

	controller "guavacoder/jiu/controller"
	db "guavacoder/jiu/db"
	user "guavacoder/jiu/user"
)

type Example struct{}

func (e Example) init() {
	mux := http.NewServeMux()

	users := db.GetUserStorage()

	userController := controller.NewUserController(
		mux,
		user.NewUserSerivice(user.NewUserRepository(users)),
	)
	userController.Run()

	port := 8080
	fmt.Printf("Server started on localhost:%d\n", port)
	fmt.Println()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
