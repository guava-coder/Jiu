package user

import (
	db "guavacoder/jiu/example/db"
	"testing"
)

func TestGetUsers(t *testing.T) {
	repo := NewUserRepository(db.GetUserStorage())
	if len(repo.GetUsers()) > 0 {
		t.Log("GetUsers test passed")
	} else {
		t.Fatal()
	}
}

func TestGetUsersByConditions(t *testing.T) {
	repo := NewUserRepository(db.GetUserStorage())
	if v := repo.GetUserByConditions("", "j@j.com"); len(v) > 0 {
		t.Log(v)
		t.Log("GetUsersByConditions test passed")
	} else {
		t.Fatal()
	}
}
