package contracts

type Pleasure struct {
	Price  int
	Salary int
	Tips   int
}

type BaristaInterface interface {
	MakeCoffee(b Barista1) (Pleasure, error)
	MakeCream(b Barista2) (Pleasure, error)
	ServeCoffee(b Barista1) (Pleasure, error)
	ServeCream(b Barista2) (Pleasure, error)
	AddPosition()
	DeleteOrder()
}
