package user

import (
	db "guavacoder/jiu/example/db"
	"testing"
)

func TestUserRepository(t *testing.T) {
	t.Run("NewUserRepository", func(t *testing.T) {
		repo := NewUserRepository(db.GetUserStorage())
		if len(repo.users) > 0 {
			t.Log("NewUserRepository test passed")
		} else {
			t.Fatal()
		}
	})
	t.Run("GetUsers", func(t *testing.T) {
		repo := NewUserRepository(db.GetUserStorage())
		if len(repo.GetUsers()) > 0 {
			t.Log("GetUsers test passed")
		} else {
			t.Fatal()
		}
	})
	t.Run("GetUserByConditions", func(t *testing.T) {
		repo := NewUserRepository(db.GetUserStorage())
		if v := repo.GetUserByConditions("", "j@j.com"); len(v) > 0 {
			t.Log(v)
			t.Log("GetUsersByConditions test passed")
		} else {
			t.Fatal()
		}
	})
}
