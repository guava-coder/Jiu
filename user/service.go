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

func (serv UserService) GetUsers(w http.ResponseWriter, r *http.Request) (statusCode int) {
	if res, err := json.Marshal(serv.repo.GetUsers()); err == nil {
		statusCode = http.StatusOK
		w.Write(res)
	} else {
		statusCode = http.StatusInternalServerError
		w.WriteHeader(statusCode)
		w.Write([]byte(fmt.Sprintf("{Response: %s}", err)))
	}
	return
}
