package category

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) CreateCategory(category *Category) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *GormRepository) GetByID(id uint) (*Category, error) {
	var category Category
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *GormRepository) GetAll() ([]Category, error) {
	var categories []Category
	if err := r.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *GormRepository) Update(id uint, category *Category) error {
	return r.db.
		Model(&Category{}).
		Where("id = ?", id).
		Updates(category).Error
}

func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Delete(&Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
