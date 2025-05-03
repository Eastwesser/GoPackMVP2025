package main

import (
	"container/heap"
	"fmt"
)

/*
6. Кучи

Пример:
Приоритетная очередь для выбора целей (например, ближайший противник).

Объяснение:
Куча позволяет быстро находить ближайшего противника.
*/

type Enemy struct {
	distance int
	name     string
}

type PriorityQueue []Enemy

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Enemy)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func main() {
	pq := &PriorityQueue{}
	heap.Push(pq, Enemy{100, "Enemy1"})
	heap.Push(pq, Enemy{50, "Enemy2"})
	heap.Push(pq, Enemy{75, "Enemy3"})
	fmt.Println("Closest Enemy:", (*pq)[0].name)
}
