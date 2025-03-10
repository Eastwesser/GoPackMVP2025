package main

import "fmt"

/*
8. Графы

Пример:
Карта местности с точками интереса (например, здания, укрытия).

Объяснение:
Граф позволяет моделировать карту и находить кратчайшие пути.
*/

type Graph struct {
	nodes map[string][]string
}

func (g *Graph) addEdge(u, v string) {
	g.nodes[u] = append(g.nodes[u], v)
}

func main() {
	g := &Graph{nodes: make(map[string][]string)}
	g.addEdge("House1", "House2")
	g.addEdge("House2", "Bunker")
	fmt.Println("Map Connections:")
	fmt.Println("House1 ->", g.nodes["House1"])
}
