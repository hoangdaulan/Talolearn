package student

import (
	"errors"
	"strconv"
)

type StudentService struct {
	repo Repository
}

func NewStudentService(repo Repository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) CreateStudent(student *Student) (*Student, error) {
	return s.repo.Create(student)
}

func (s *StudentService) GetStudentByID(id uint) (*Student, error) {
	return s.repo.GetById(id)
}

func (s *StudentService) GetStudentByCode(code string) (*Student, error) {
	return s.repo.GetByCode(code)
}

func (s *StudentService) FindStudentsByName(name string) (*[]Student, error) {
	return s.repo.FindByName(name)
}

func (s *StudentService) GetAllStudents() ([]Student, error) {
	return s.repo.GetAll()
}

func (s *StudentService) UpdateStudent(idStr string, updates map[string]interface{}) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}
	_, err = s.repo.GetById(uint(id))
	if err != nil {
		return errors.New("student not found")
	}
	return s.repo.UpdatePartial(uint(id), updates)
}

func (s *StudentService) DeleteStudent(code string) error {
	return s.repo.Delete(code)
}
