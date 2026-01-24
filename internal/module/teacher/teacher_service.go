package teacher

type Service interface {
	CreateTeacher(userID uint, bio string, categoryID uint) (*Teacher, error)
	GetTeacherByID(id uint) (*Teacher, error)
	GetTeacherByUserID(userID uint) (*Teacher, error)
	FindTeacherByName(name string) (*[]Teacher, error)
	GetAllTeachers() ([]Teacher, error)
	UpdateTeacher(id uint, updates map[string]interface{}) error
	DeleteTeacher(id uint) error
}
