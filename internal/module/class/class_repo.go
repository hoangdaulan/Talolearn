package class

type Repository interface {
	Create(cls *Class) error
	Get(id uint) (*Class, error)
	Update(cls *Class) error
	Delete(id uint) error
	List() ([]Class, error)
}
