package user

import (
	. "guavacoder/jiu/example/data"
)

type UserRepository struct {
	users []User
}

func NewUserRepository(us []User) *UserRepository {
	return &UserRepository{
		users: us,
	}
}

func (r *UserRepository) GetUsers() []User {
	return r.users
}

func (r *UserRepository) GetUserByConditions(name string, email string) []User {
	temp := make([]User, 0)
	for _, u := range r.users {
		if u.Name == name || u.Email == email {
			return append(temp, u)
		}
	}
	return temp
}
