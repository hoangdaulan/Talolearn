package teacher

import "Talolearn-Project/internal/module/user"

type TeacherService struct {
	repo     Repository
	userRepo user.Repository
}

func NewTeacherService(repo Repository, userRepo user.Repository) *TeacherService {
	return &TeacherService{repo: repo, userRepo: userRepo}
}

func (s *TeacherService) CreateTeacher(userID uint, bio string, categoryID uint) (*Teacher, error) {
	_, err := s.repo.GetByUserID(userID)
	if err == nil {
		return nil, err
	}
	t := &Teacher{
		UserID:     userID,
		Bio:        bio,
		CategoryID: categoryID,
	}
	if err := s.repo.Create(t); err != nil {
		return nil, err
	}
	if err := s.userRepo.UpdateRole(userID, map[string]interface{}{"role_id": 2}); err != nil {
		return nil, err
	}
	return t, nil
}

func (s *TeacherService) GetTeacherByID(id uint) (*Teacher, error) {
	return s.repo.GetByID(id)
}

func (s *TeacherService) GetTeacherByUserID(userID uint) (*Teacher, error) {
	return s.repo.GetByUserID(userID)
}

func (s *TeacherService) FindTeacherByName(name string) (*[]Teacher, error) {
	return s.repo.FindByName(name)
}

func (s *TeacherService) GetAllTeachers() ([]Teacher, error) {
	return s.repo.List()
}

func (s *TeacherService) UpdateTeacher(id uint, updates map[string]interface{}) error {
	teacher, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	for key, value := range updates {
		switch key {
		case "Bio":
			if v, ok := value.(string); ok {
				teacher.Bio = v
			}
		case "CategoryID":
			if v, ok := value.(uint); ok {
				teacher.CategoryID = v
			}
		}
	}
	err = s.repo.Update(teacher)
	if err != nil {
		return err
	}
	return nil
}

func (s *TeacherService) DeleteTeacher(id uint) error {
	return s.repo.Delete(id)
}
