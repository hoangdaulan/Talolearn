package course

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) CreateCourse(course *Course) error {
	if err := r.db.Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) GetCourseByID(id uint) (*Course, error) {
	var course Course
	if err := r.db.First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *GormRepository) UpdateCourse(id uint, updates map[string]interface{}) error {
	return r.db.
		Model(&Course{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormRepository) DeleteCourse(id uint) error {
	if err := r.db.Delete(&Course{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) ListCourses() ([]Course, error) {
	var courses []Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *GormRepository) IsTeacherInCourse(id uint) bool {
	var count int64
	r.db.Model(&CourseTeacher{}).Where("teacher_id = ?", id).Count(&count)
	return count > 0
}

func (r *GormRepository) AddTeacherToCourse(courseID uint, teacherID uint) error {
	courseTeacher := CourseTeacher{
		CourseID:  courseID,
		TeacherID: teacherID,
	}
	if err := r.db.Create(&courseTeacher).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) RemoveTeacherFromCourse(courseID uint, teacherID uint) error {
	if err := r.db.
		Where("course_id = ? AND teacher_id = ?", courseID, teacherID).
		Delete(&CourseTeacher{}).Error; err != nil {
		return err
	}
	return nil
}
