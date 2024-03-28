package user

import (
	"testing"
)

func TestUserRepository(t *testing.T) {
	repo := NewUserRepository(map[string]User{
		"1": {
			Id:    "1",
			Name:  "John",
			Email: "j@j.com",
		},
		"2": {
			Id:    "2",
			Name:  "Jane",
			Email: "jane@j.com",
		},
		"3": {
			Id:    "3",
			Name:  "Joe",
			Email: "joe@j.com",
		},
	})
	t.Run("get users", func(t *testing.T) {
		_, err := repo.GetUsers()
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("get user by conditions", func(t *testing.T) {
		_, err := repo.GetUserByConditions("", "j@j.com")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("get User By Id", func(t *testing.T) {
		_, err := repo.GetUserById("1")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("add user", func(t *testing.T) {
		err := repo.AddUser(User{
			Id:    "4",
			Name:  "Joe",
			Email: "joe@j.com",
		})
		if err != nil {
			t.Fatal(err)
		}
	})
}
