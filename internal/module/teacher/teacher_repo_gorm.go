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

func (r *GormRepository) GetByUserID(userID uint) (*Teacher, error) {
	var teacher Teacher
	err := r.db.
		Preload("User").
		Preload("Category").
		Where("user_id = ?", userID).
		First(&teacher).Error

	if err != nil {
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

func (r *GormRepository) List() ([]Teacher, error) {
	var teachers []Teacher
	if err := r.db.Find(&teachers).Error; err != nil {
		return nil, err
	}
	return teachers, nil
}

func (r *GormRepository) Update(teacher *Teacher) error {
	return r.db.Save(teacher).Error
}

func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Delete(&Teacher{}, id).Error; err != nil {
		return err
	}
	return nil
}
