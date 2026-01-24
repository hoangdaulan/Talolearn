package category

type Repository interface {
	CreateCategory(category *Category) error
	GetByID(id uint) (*Category, error)
	GetAll() ([]Category, error)
	Update(id uint, category *Category) error
	Delete(id uint) error
}
