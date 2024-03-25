package user

type User struct {
	Id    int
	Name  string
	Email string
}

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
