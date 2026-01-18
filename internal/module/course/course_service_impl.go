package course

import (
	"Talolearn-Project/internal/module/teacher"
	"errors"
)

type CourseService struct {
	courseRepo  Repository
	teacherRepo teacher.Repository
}

func NewService(courseRepo Repository, teacherRepo teacher.Repository) *CourseService {
	return &CourseService{
		courseRepo:  courseRepo,
		teacherRepo: teacherRepo,
	}
}

func (s *CourseService) AddTeacherToCourse(courseID uint, teacherID uint) error {
	// Check if the teacher exists
	_, err := s.teacherRepo.GetByID(teacherID)
	if err != nil {
		return err
	}

	// Check if the course exists
	_, err = s.courseRepo.GetCourseByID(courseID)
	if err != nil {
		return err
	}

	// Check if the teacher is already assigned to the course
	if s.courseRepo.IsTeacherInCourse(teacherID) {
		return errors.New("Teacher has already in this course")
	}

	// Add the teacher to the course
	return s.courseRepo.AddTeacherToCourse(courseID, teacherID)
}

func (s *CourseService) RemoveTeacherFromCourse(courseID uint, teacherID uint) error {
	// Check if the teacher exists
	_, err := s.teacherRepo.GetByID(teacherID)
	if err != nil {
		return err
	}

	// Check if the course exists
	_, err = s.courseRepo.GetCourseByID(courseID)
	if err != nil {
		return err
	}

	// Remove the teacher from the course
	return s.courseRepo.RemoveTeacherFromCourse(courseID, teacherID)
}
