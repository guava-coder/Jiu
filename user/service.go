package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserService struct {
	repo *UserRepository
}

func NewUserSerivice(r *UserRepository) UserService {
	return UserService{repo: r}
}

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) {
	if res, err := json.Marshal(serv.repo.GetUsers()); err == nil {
		w.Write(res)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{Response: %s}", err)
	}

}
