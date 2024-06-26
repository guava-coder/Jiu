package example

import (
	"fmt"
	"log"
	"net/http"

	controller "guavacoder/jiu/example/controller"
	db "guavacoder/jiu/example/db"
	user "guavacoder/jiu/example/user"
)

func Init() {
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
