package category

type Service interface {
	CreateCategory(category *Category) error
	GetAllCategory() ([]Category, error)
	GetCategoryByID(id uint) (*Category, error)
	DeleteCategory(id uint) error
}
