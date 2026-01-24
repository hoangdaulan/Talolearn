package class

type Service interface {
	CreateClass(class Class) error
	GetClass(id uint) (*Class, error)
	UpdateClass(class Class) error
	DeleteClass(id uint) error
	ListClasses() ([]Class, error)
}
