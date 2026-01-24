package category

type Service interface {
	CreateCategory(category *Category) error
	GetAllCategory() ([]Category, error)
	GetCategoryByID(id uint) (*Category, error)
	UpdateCategory(id uint, category *Category) error
	DeleteCategory(id uint) error
}
