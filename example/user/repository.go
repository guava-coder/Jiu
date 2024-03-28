package user

type User struct {
	Id    string
	Name  string
	Email string
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
		_, err = users, NewQueryError(User{}).UserNotFound
	} else {
		_, err = users, nil
	}
	return
}

func (r UserRepository) GetUserById(id string) (user User, err error) {
	if r.users[id].Id == "" {
		user, err = User{}, NewQueryError(User{Id: id}).UserNotExist
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
		_, err = users, NewQueryError(User{}).UserNotFound
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
		err = NewQueryError(u).UserExist
	}
	return
}

func (r UserRepository) UpdateUser(data []User) (err error) {
	for _, v := range data {
		if r.users[v.Id].Id == "" {
			err = NewQueryError(User{Id: v.Id}).UserNotExist
			return
		} else {
			r.users[v.Id] = v
		}
	}
	return
}

func (r UserRepository) DeleteUser(id string) (err error) {
	if r.users[id].Id == "" {
		err = NewQueryError(User{Id: id}).UserNotExist
	} else {
		delete(r.users, id)
		err = nil
	}
	return
}
