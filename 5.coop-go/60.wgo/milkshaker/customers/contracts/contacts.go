package contracts

type CustomersInterface interface {
	MakeOrder()
	TakeOrder()
	RefuseToTakeOrder()
	PayForOrder()
	GiveTips()
}
