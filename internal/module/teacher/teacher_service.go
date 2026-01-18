package teacher

type Service interface {
	CreateTeacher(name, email string) (*Teacher, error)
	GetTeacherByID(id uint) (*Teacher, error)
	FindTeacherByName(name string) (*Teacher, error)
	GetAllTeachers() ([]Teacher, error)
	UpdateTeacher(teacher *Teacher) error
	DeleteTeacher(id uint) error
}
