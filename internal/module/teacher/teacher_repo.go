package teacher

type Repository interface {
	Create(teacher *Teacher) (*Teacher, error)
	GetByID(id uint) (*Teacher, error)
	FindByName(name string) (*[]Teacher, error)
	GetAll() ([]Teacher, error)
	UpdatePartial(id uint, updates map[string]interface{}) error
	Delete(id uint) error
}
