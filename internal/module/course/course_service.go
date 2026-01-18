package course

type Service interface {
	AddTeacherToCourse(courseID uint, teacherID uint) error
	RemoveTeacherFromCourse(courseID uint, teacherID uint) error
}
