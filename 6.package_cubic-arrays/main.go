package main

import (
	"fmt"
)

type Cube interface {
	ExpandVertical()
	FoldToCube()
	GetState() [][]int
}

type cube struct {
	state [][]int
}

func newAntiStressCube() Cube {
	return &cube{
		state: [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
		},
	}
}

func (c *cube) ExpandVertical() {
	c.state = [][]int{
		{1, 5},
		{2, 6},
		{3, 7},
		{4, 8},
	}
}

func (c *cube) FoldToCube() {
	c.state = [][]int{
		{1, 2, 5, 6},
		{3, 4, 7, 8},
	}
}

func (c *cube) GetState() [][]int {
	return c.state
}

func printState(state [][]int) {
	for _, row := range state {
		fmt.Println(row)
	}
	fmt.Println()
}

func compareStates(state1, state2 [][]int) bool {
	if len(state1) != len(state2) {
		return false
	}
	for i := range state1 {
		if len(state1[i]) != len(state2[i]) {
			return false
		}
		for j := range state1[i] {
			if state1[i][j] != state2[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	c := newAntiStressCube()

	fmt.Println("Изначальная конфигурация:")
	printState(c.GetState())

	c.ExpandVertical()
	printState(c.GetState())

	fmt.Println("Складывание в куб:")
	c.FoldToCube()
	printState(c.GetState())

	fmt.Println("Проверка стабильности после 10 циклов:")
	initialState := c.GetState()
	stable := true
	for i := 0; i < 10; i++ {
		c.ExpandVertical()
		c.FoldToCube()
		if !compareStates(initialState, c.GetState()) {
			stable = false
			break
		}
	}
	if stable {
		fmt.Println("Кубик возвращается в исходное состояние.")
	} else {
		fmt.Println("Кубик не возвращается в исходное состояние.")
	}
}
