package user

type Repository interface {
	Create(user *User) error
	GetByID(id uint) (*User, error)
	GetByName(email string) (*User, error)
	List() ([]User, error)
	Update(id uint, updates map[string]interface{}) error
	UpdateRole(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}
