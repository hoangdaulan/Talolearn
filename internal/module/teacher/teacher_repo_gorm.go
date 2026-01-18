package teacher

import (
	"gorm.io/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(teacher *Teacher) error {
	if err := r.db.Create(teacher).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) GetByID(id uint) (*Teacher, error) {
	var teacher Teacher
	if err := r.db.First(&teacher, id).Error; err != nil {
		return nil, err
	}
	return &teacher, nil
}

func (r *GormRepository) FindByName(name string) (*[]Teacher, error) {
	var teacher []Teacher
	if err := r.db.Where("name = ?", name).Find(&teacher).Error; err != nil {
		return nil, err
	}
	if len(teacher) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return &teacher, nil
}

func (r *GormRepository) GetAll() ([]Teacher, error) {
	var teachers []Teacher
	if err := r.db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *GormRepository) UpdatePartial(id uint, updates map[string]interface{}) error {
	return r.db.
		Model(&Teacher{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Delete(&Teacher{}, id).Error; err != nil {
		return err
	}
	return nil
}
