package graph

import "fmt"

// Graph represents a directed weighted graph.
// AdjacencyList[node] = map[targetNode]weight
type Graph struct {
	AdjacencyList map[string]map[string]float64 `json:"adjacency_list"`
}

// NewGraph creates and returns a new empty Graph.
func NewGraph() *Graph {
	return &Graph{
		AdjacencyList: make(map[string]map[string]float64),
	}
}

// AddEdge adds a directed edge from `from` to `to` with the given `weight`.
func (g *Graph) AddEdge(from, to string, weight float64) {
	if _, exists := g.AdjacencyList[from]; !exists {
		g.AdjacencyList[from] = make(map[string]float64)
	}
	g.AdjacencyList[from][to] = weight
}

// AddNode adds a node to the graph. If the node already exists, it does nothing.
func (g *Graph) AddNode(node string) {
	if _, exists := g.AdjacencyList[node]; !exists {
		g.AdjacencyList[node] = make(map[string]float64)
	}
}

// PrintGraph prints the graph's adjacency list in a readable format.
func (g *Graph) PrintGraph() {
	fmt.Println("Graph Structure:")
	for from, edges := range g.AdjacencyList {
		for to, weight := range edges {
			fmt.Printf("  %s -> %s [Weight: %.2f]\n", from, to, weight)
		}
	}
}
