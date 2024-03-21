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

	controller.NewUserController(
		mux,
		user.NewUserSerivice(user.NewUserRepository(users)),
	).Run()

	port := 8080
	fmt.Printf("Server started on localhost:%d\n", port)
	fmt.Println()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
