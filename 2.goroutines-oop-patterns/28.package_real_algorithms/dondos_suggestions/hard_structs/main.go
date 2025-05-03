package main

import (
	"container/heap"
	"container/list"
	"fmt"
)

// 1. Динамический массив (слайс)
func dynamicArrayExample() {
	cache := make([]string, 0, 100) // Кэш последних 100 запросов
	cache = append(cache, "request1")
	fmt.Println("Dynamic Array:", cache)
}

// 2. Связанный список
func linkedListExample() {
	history := list.New() // История действий (undo/redo)
	history.PushBack("action1")
	history.PushBack("action2")
	fmt.Println("Linked List:")
	for e := history.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// 3. Стек
func stackExample() {
	stack := []string{} // Проверка скобок
	brackets := map[rune]rune{')': '(', '}': '{', ']': '['}
	expression := "({[]})"
	for _, char := range expression {
		if char == '(' || char == '{' || char == '[' {
			stack = append(stack, string(char))
		} else if char == ')' || char == '}' || char == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != string(brackets[char]) {
				fmt.Println("Invalid expression")
				return
			}
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println("Stack: Expression is valid")
}

// 4. Очередь/Дек
func queueExample() {
	queue := list.New() // Очередь задач
	queue.PushBack("task1")
	queue.PushBack("task2")
	fmt.Println("Queue:")
	for e := queue.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// 5. Куча (приоритетная очередь)
type Task struct {
	priority int
	name     string
}

type PriorityQueue []Task

func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Task)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func heapExample() {
	pq := &PriorityQueue{}
	heap.Push(pq, Task{2, "Low priority"})
	heap.Push(pq, Task{1, "High priority"})
	fmt.Println("Heap:")
	for pq.Len() > 0 {
		task := heap.Pop(pq).(Task)
		fmt.Println(task.name)
	}
}

// 6. Дерево
type TreeNode struct {
	name     string
	children []*TreeNode
}

func (n *TreeNode) addChild(child *TreeNode) {
	n.children = append(n.children, child)
}

func treeExample() {
	root := &TreeNode{name: "root"}
	child1 := &TreeNode{name: "child1"}
	child2 := &TreeNode{name: "child2"}
	root.addChild(child1)
	root.addChild(child2)
	fmt.Println("Tree:")
	fmt.Println(root.name, "->", root.children[0].name, ",", root.children[1].name)
}

// 7. Граф
type Graph struct {
	nodes map[string][]string
}

func (g *Graph) addEdge(u, v string) {
	g.nodes[u] = append(g.nodes[u], v)
}

func graphExample() {
	g := &Graph{nodes: make(map[string][]string)}
	g.addEdge("A", "B")
	g.addEdge("A", "C")
	fmt.Println("Graph:")
	fmt.Println("A ->", g.nodes["A"])
}

func main() {
	dynamicArrayExample()
	linkedListExample()
	stackExample()
	queueExample()
	heapExample()
	treeExample()
	graphExample()
}
