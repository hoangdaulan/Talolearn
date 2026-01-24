package student

type Repository interface {
	Create(student *Student) (*Student, error)
	GetById(studentId uint) (*Student, error)
	GetByCode(code string) (*Student, error)
	FindByName(name string) (*[]Student, error)
	GetAll() ([]Student, error)
	UpdatePartial(id uint, updates map[string]interface{}) error
	Delete(code string) error
}
