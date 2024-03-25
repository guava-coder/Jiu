package user

import (
	"testing"
)

func TestUserRepository(t *testing.T) {
	repo := NewUserRepository([]User{
		{
			Id:    1,
			Name:  "John",
			Email: "j@j.com",
		},
		{
			Id:    2,
			Name:  "Jane",
			Email: "jane@j.com",
		},
		{
			Id:    3,
			Name:  "Joe",
			Email: "joe@j.com",
		},
	})
	t.Run("GetUsers", func(t *testing.T) {
		if len(repo.GetUsers()) > 0 {
			t.Log("GetUsers test passed")
		} else {
			t.Fatal()
		}
	})
	t.Run("GetUserByConditions", func(t *testing.T) {
		if v := repo.GetUserByConditions("", "j@j.com"); len(v) > 0 {
			t.Log(v)
			t.Log("GetUsersByConditions test passed")
		} else {
			t.Fatal()
		}
	})
}
