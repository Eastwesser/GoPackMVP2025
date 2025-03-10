package main

import "fmt"

/*
6. DFS -Depth First Search- (поиск в глубину)
Пример: Поиск всех возможных путей до укрытия.

Объяснение: Игрок исследует все возможные пути.
*/

func dfs(graph map[string][]string, node string, visited map[string]bool) {
	if !visited[node] {
		fmt.Println("Visited:", node)
		visited[node] = true
		for _, neighbor := range graph[node] {
			dfs(graph, neighbor, visited)
		}
	}
}

func main() {
	mapGraph := map[string][]string{
		"House1": {"House2", "House3"},
		"House2": {"Bunker"},
		"House3": {"Bunker"},
	}
	visited := make(map[string]bool)
	dfs(mapGraph, "House1", visited)
}
