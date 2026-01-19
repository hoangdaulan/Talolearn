package teacher

type TeacherService struct {
	repo Repository
}

func NewTeacherService(repo Repository) *TeacherService {
	return &TeacherService{repo: repo}
}

func (s *TeacherService) CreateTeacher(teacher *Teacher) (*Teacher, error) {
	return s.repo.Create(teacher)
}

func (s *TeacherService) GetTeacherByID(id uint) (*Teacher, error) {
	return s.repo.GetByID(id)
}

func (s *TeacherService) FindTeacherByName(name string) (*[]Teacher, error) {
	return s.repo.FindByName(name)
}

func (s *TeacherService) GetAllTeachers() ([]Teacher, error) {
	return s.repo.GetAll()
}

func (s *TeacherService) UpdateTeacher(id uint, updates map[string]interface{}) error {
	return s.repo.UpdatePartial(id, updates)
}

func (s *TeacherService) DeleteTeacher(id uint) error {
	return s.repo.Delete(id)
}
