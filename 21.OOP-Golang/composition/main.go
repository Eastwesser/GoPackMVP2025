package main

import "fmt"

type Engine struct {
	power int
}

type Wheel struct {
	size int
}

type Car struct {
	engine Engine
	wheel  Wheel
}

func main() {
	car := Car{
		engine: Engine{power: 200},
		wheel:  Wheel{size: 18},
	}

	fmt.Println("Мощность двигателя:", car.engine.power)
	fmt.Println("Размер колес:", car.wheel.size)
}
