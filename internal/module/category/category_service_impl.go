package category

type CategoryService struct {
	repo Repository
}

func NewCategoryService(repo Repository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *Category) error {
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetAllCategory() ([]Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetCategoryByID(id uint) (*Category, error) {
	return s.repo.GetByID(id)
}

func (s *CategoryService) DeleteCategory(id uint) error {
	return s.repo.Delete(id)
}

func (s *CategoryService) UpdateCategory(id uint, c *Category) interface{} {
	return s.repo.Update(id, c)
}
