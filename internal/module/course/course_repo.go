package course

type Repository interface {
	CreateCourse(course *Course) error
	GetCourseByID(id uint) (*Course, error)
	UpdateCourse(id uint, updates map[string]interface{}) error
	DeleteCourse(id uint) error
	ListCourses() ([]Course, error)
	IsTeacherInCourse(id uint) bool
	AddTeacherToCourse(courseID uint, teacherID uint) error
	RemoveTeacherFromCourse(courseID uint, teacherID uint) error
}
