package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Weight, Value int
}

func main() {
	items := []Item{
		{10, 60},
		{20, 100},
		{30, 120},
	}
	capacity := 50 // Вместимость рюкзака

	maxValue := fractionalKnapsack(items, capacity)
	fmt.Println("Максимальная ценность:", maxValue)
}

// Жадный алгоритм для задачи о рюкзаке (дробный вариант)
func fractionalKnapsack(items []Item, capacity int) float64 {
	// Сортируем предметы по убыванию удельной ценности (Value / Weight)
	sort.Slice(items, func(i, j int) bool {
		return float64(items[i].Value)/float64(items[i].Weight) > float64(items[j].Value)/float64(items[j].Weight)
	})

	totalValue := 0.0
	for _, item := range items {
		if capacity >= item.Weight {
			capacity -= item.Weight
			totalValue += float64(item.Value)
		} else {
			totalValue += float64(item.Value) * (float64(capacity) / float64(item.Weight))
			break
		}
	}

	return totalValue
}
