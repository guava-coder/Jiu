package db

import (
	user "guavacoder/jiu/example/user"
)

type User = user.User

func GetUserStorage() map[string]User {
	users := []User{
		{
			Id:    "1",
			Name:  "John",
			Email: "j@j.com",
		},
		{
			Id:    "2",
			Name:  "Jane",
			Email: "jane@j.com",
		},
		{
			Id:    "3",
			Name:  "Joe",
			Email: "joe@j.com",
		},
	}
	userMap := make(map[string]User, 0)
	for _, v := range users {
		userMap[v.Id] = v
	}
	return userMap
}
