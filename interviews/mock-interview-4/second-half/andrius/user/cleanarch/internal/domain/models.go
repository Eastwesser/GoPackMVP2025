package domain

type Emp struct {
	ID    int
	Name  string
	Year  int
	Phone string
}

type IEmpRepo interface {
	Add(employee *Emp) error
	GetAll() ([]Emp, error)
}
