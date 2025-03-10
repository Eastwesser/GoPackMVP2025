package main

import "fmt"

/*
5. BFS -Breadth First Search- (поиск в ширину)
Пример: Поиск кратчайшего пути до укрытия.

Объяснение: Игрок ищет кратчайший путь до ближайшего укрытия.
*/

func bfs(graph map[string][]string, start string) {
	queue := []string{start}
	visited := make(map[string]bool)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if !visited[node] {
			fmt.Println("Visited:", node)
			visited[node] = true
			for _, neighbor := range graph[node] {
				queue = append(queue, neighbor)
			}
		}
	}
}

func main() {
	mapGraph := map[string][]string{
		"House1": {"House2", "House3"},
		"House2": {"Bunker"},
		"House3": {"Bunker"},
	}
	bfs(mapGraph, "House1")
}
