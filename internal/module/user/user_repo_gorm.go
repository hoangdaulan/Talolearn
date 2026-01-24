package user

import "gorm.io/gorm"

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db: db}
}

func (r *GormRepository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *GormRepository) GetByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormRepository) GetByName(name string) (*User, error) {
	var user User
	if err := r.db.Where("name = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *GormRepository) List() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormRepository) Update(id uint, updates map[string]interface{}) error {
	return r.db.
		Model(&User{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormRepository) UpdateRole(id uint, updates map[string]interface{}) error {
	return r.db.
		Model(&User{}).
		Where("id = ?", id).
		Updates(updates).Error
}

func (r *GormRepository) Delete(id uint) error {
	return r.db.Delete(&User{}, id).Error
}
