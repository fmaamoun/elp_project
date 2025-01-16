package main

import "fmt"

type Edge struct {
	To   string
	Type string
	Time int
}

type Graph struct {
	data map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{data: make(map[string][]Edge)}
}

func (g *Graph) AddStop(stopID string) {
	if _, exists := g.data[stopID]; !exists {
		g.data[stopID] = []Edge{}
	}
}

func (g *Graph) AddEdge(fromStopID, toStopID, edgeType string, time int) {
	g.data[fromStopID] = append(g.data[fromStopID], Edge{To: toStopID, Type: edgeType, Time: time})
}

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
