package student

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(student *Student) (*Student, error) {
	if err := r.db.Create(student).Error; err != nil {
		return nil, err
	}
	return student, nil
}

func (r *GormRepository) GetById(studentId uint) (*Student, error) {
	var student Student
	if err := r.db.First(&student, studentId).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *GormRepository) GetByCode(code string) (*Student, error) {
	var student Student
	if err := r.db.First(&student, "code = ?", code).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *GormRepository) FindByName(name string) (*[]Student, error) {
	var students []Student
	if err := r.db.Where("name LIKE ?", "%"+name+"%").Find(&students).Error; err != nil {
		return nil, err
	}
	return &students, nil
}

func (r *GormRepository) GetAll() ([]Student, error) {
	var students []Student
	if err := r.db.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (r *GormRepository) UpdatePartial(id uint, updates map[string]interface{}) error {
	if err := r.db.Model(&Student{}).Where("user_id = ?", id).Updates(updates).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) Delete(code string) error {
	if err := r.db.Delete(&Student{}, "code = ?", code).Error; err != nil {
		return err
	}
	return nil
}
