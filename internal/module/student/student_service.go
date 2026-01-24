package student

type Service interface {
	CreateStudent(student *Student) (*Student, error)
	GetStudentByID(id uint) (*Student, error)
	GetStudentByCode(code string) (*Student, error)
	FindStudentsByName(name string) (*[]Student, error)
	GetAllStudents() ([]Student, error)
	UpdateStudent(code string, updates map[string]interface{}) error
	DeleteStudent(code string) error
}
