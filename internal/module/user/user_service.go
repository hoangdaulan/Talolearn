package user

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	GetUserByName(email string) (*User, error)
	ListUsers() ([]User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}
