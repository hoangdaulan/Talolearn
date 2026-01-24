package teacher

type Repository interface {
	Create(teacher *Teacher) error
	GetByID(id uint) (*Teacher, error)
	GetByUserID(userID uint) (*Teacher, error)
	FindByName(name string) (*[]Teacher, error)
	List() ([]Teacher, error)
	Update(teacher *Teacher) error
	Delete(id uint) error
}
