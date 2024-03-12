package user

import (
	"guavacoder/jiu/db"
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
