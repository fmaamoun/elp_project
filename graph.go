package main

import "fmt"

// Edge represents a connection from one stop/node to another.
// 'Type' can be "trip" or "transfer", and 'Time' is the cost/weight (in seconds).
type Edge struct {
	To   string
	Type string
	Time int
}

// Graph holds adjacency information in a map: each node ID maps to a list of edges.
type Graph struct {
	data map[string][]Edge
}

// NewGraph creates an empty Graph.
func NewGraph() *Graph {
	return &Graph{data: make(map[string][]Edge)}
}

// AddStop ensures a stop/node is in the graph's adjacency list (even if it has no edges yet).
func (g *Graph) AddStop(stopID string) {
	if _, exists := g.data[stopID]; !exists {
		g.data[stopID] = []Edge{}
	}
}

// AddEdge appends a single directed edge to the adjacency list.
// If you need an undirected edge, remember to call AddEdge in both directions.
func (g *Graph) AddEdge(fromStopID, toStopID, edgeType string, time int) {
	g.data[fromStopID] = append(g.data[fromStopID], Edge{To: toStopID, Type: edgeType, Time: time})
}

// PrintGraph prints all nodes and their outgoing edges.
func (g *Graph) PrintGraph() {
	fmt.Println("Graph representation:")
	for stop, edges := range g.data {
		fmt.Printf("%s -> ", stop)
		for _, edge := range edges {
			fmt.Printf("{To: %s, Type: %s, Time: %d} ", edge.To, edge.Type, edge.Time)
		}
		fmt.Println()
	}
}
