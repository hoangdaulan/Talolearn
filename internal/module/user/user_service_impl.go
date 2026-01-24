package user

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.Create(user)
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetByID(id)
}

func (s *Service) GetUserByName(email string) (*User, error) {
	return s.repo.GetByName(email)
}

func (s *Service) ListUsers() ([]User, error) {
	return s.repo.List()
}

func (s *Service) UpdateUser(user *User) error {
	updates := map[string]interface{}{
		"name":          user.Name,
		"password_hash": user.PasswordHash,
		"email":         user.Email,
		"avatar":        user.Avatar,
		"phone":         user.Phone,
		"role_id":       user.RoleID,
	}
	return s.repo.Update(user.ID, updates)
}

func (s *Service) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
