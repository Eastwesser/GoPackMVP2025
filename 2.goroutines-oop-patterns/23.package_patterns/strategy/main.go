package main

import "fmt"

// main strategy iface
type RouteStrategy interface {
	BuildRoute(start, end string)
}

// by car strat
type CarRouteStrategy struct{}

func (c CarRouteStrategy) BuildRoute(start, end string) {
	fmt.Printf("Car route strategy is built from %s to %s\n", start, end)
}

// by bike strat
type BikeRouteStrategy struct{}

func (c BikeRouteStrategy) BuildRoute(start, end string) {
	fmt.Printf("Bike route strategy is built from %s to %s\n", start, end)
}

// on foot strat
type WalkRouteStrategy struct{}

func (c WalkRouteStrategy) BuildRoute(start, end string) {
	fmt.Printf("Walk route strategy is built from %s to %s\n", start, end)
}

// context navi with original strategy (it can't be strat itself, but it must hold the strat)
type Navigator struct {
	strategy RouteStrategy
}

// set strat dynamically
func (n Navigator) SetStrategy(strategy RouteStrategy) {
	n.strategy = strategy
}

// call BuildRoute method
func (n Navigator) BuildRoute(start, end string) {
	fmt.Printf("Маршрут навигатора построен от %s до %s\n", start, end)
}

// optional: show all possible strats
func (n Navigator) GetRouteStrategies() []RouteStrategy {
	return []RouteStrategy{
		&CarRouteStrategy{},
		&BikeRouteStrategy{},
		&WalkRouteStrategy{},
	}
}

func main() {
	navi := &Navigator{}

	// by car strat
	navi.SetStrategy(&CarRouteStrategy{})
	navi.BuildRoute("Дом", "Ашан")

	// by bike strat
	navi.SetStrategy(&BikeRouteStrategy{})
	navi.BuildRoute("Парк", "Дом")

	// on foot strat
	navi.SetStrategy(&WalkRouteStrategy{})
	navi.BuildRoute("Мак", "KFC")

	// opt: get all strats
	strats := navi.GetRouteStrategies()
	for _, strat := range strats {
		strat.BuildRoute("Start", "Finish")
	}

}
