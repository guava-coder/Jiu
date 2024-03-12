package user

import (
	. "guavacoder/jiu/data"
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
