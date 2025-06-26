package models

type Customers *[]struct{}

type Customer struct {
	Name          string
	Age           int
	Creditability bool
	Phone         string
	Email         string
}
