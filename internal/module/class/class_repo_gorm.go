package class

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(cls *Class) error {
	return r.db.Create(cls).Error
}

func (r *GormRepository) Get(id uint) (*Class, error) {
	var cls Class
	if err := r.db.First(&cls, id).Error; err != nil {
		return nil, err
	}
	return &cls, nil
}

func (r *GormRepository) Update(cls *Class) error {
	return r.db.Save(cls).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&Class{}, id).Error
}

func (r *GormRepository) List() ([]Class, error) {
	var classes []Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}
