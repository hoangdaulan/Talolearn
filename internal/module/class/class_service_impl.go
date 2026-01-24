package class

type ClassService struct {
	repo Repository
}

func NewClassService(repo Repository) *ClassService {
	return &ClassService{repo: repo}
}

func (s *ClassService) CreateClass(class Class) error {
	return s.repo.Create(&class)
}

func (s *ClassService) GetClass(id uint) (*Class, error) {
	return s.repo.Get(id)
}

func (s *ClassService) UpdateClass(class Class) error {
	return s.repo.Update(&class)
}

func (s *ClassService) DeleteClass(id uint) error {
	return s.repo.Delete(id)
}

func (s *ClassService) ListClasses() ([]Class, error) {
	return s.repo.List()
}
