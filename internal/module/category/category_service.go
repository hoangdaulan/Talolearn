package category

type Service interface {
	CreateCategory(name string) (*Category, error)
	GetCategoryByID(id uint) (*Category, error)
	DeleteCategory(id uint) error
}
