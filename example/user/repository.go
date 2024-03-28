package user

import "errors"

type User struct {
	Id    string
	Name  string
	Email string
}

type UserQueryError struct {
	UserNotFound error
	UserExist    error
}

func NewUserQueryError() UserQueryError {
	return UserQueryError{
		UserNotFound: errors.New("no user founded"),
		UserExist:    errors.New("user already exist"),
	}
}

type UserRepository struct {
	users map[string]User
}

func NewUserRepository(us map[string]User) *UserRepository {
	return &UserRepository{
		users: us,
	}
}

func (r UserRepository) GetUsers() (users []User, err error) {
	for _, v := range r.users {
		users = append(users, v)
	}
	if len(users) == 0 {
		_, err = users, NewUserQueryError().UserNotFound
	} else {
		_, err = users, nil
	}
	return
}

func (r UserRepository) GetUserById(id string) (user User, err error) {
	if r.users[id].Id == "" {
		user, err = User{}, NewUserQueryError().UserNotFound
	} else {
		user, err = r.users[id], nil
	}
	return
}

func (r UserRepository) GetUserByConditions(name string, email string) (users []User, err error) {
	for _, u := range r.users {
		if u.Name == name || u.Email == email {
			users = append(users, u)
		}
	}
	if len(users) == 0 {
		_, err = users, NewUserQueryError().UserNotFound
	} else {
		_, err = users, nil
	}
	return
}

func (r UserRepository) AddUser(u User) (err error) {
	if r.users[u.Id].Id == "" {
		r.users[u.Id] = u
		err = nil
	} else {
		err = NewUserQueryError().UserExist
	}
	return
}
