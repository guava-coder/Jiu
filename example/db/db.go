package db

import (
	. "guavacoder/jiu/example/data"
)

func GetUserStorage() []User {
	users := make([]User, 0)
	users = append(users, User{
		Id:    1,
		Name:  "John",
		Email: "j@j.com",
	})
	users = append(users, User{
		Id:    2,
		Name:  "Jane",
		Email: "jane@j.com",
	})
	users = append(users, User{
		Id:    3,
		Name:  "Joe",
		Email: "joe@j.com",
	})
	return users
}
