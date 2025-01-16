package main

import (
	"container/heap"
	"fmt"
	"math"
	"sync"
	"time"
)

type Edge struct {
	to       string
	edgeType string
	time     int
}

type Graph struct {
	nodes map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]Edge)}
}

func (g *Graph) AddEdge(from, to, edgeType string, time int) {
	g.nodes[from] = append(g.nodes[from], Edge{to, edgeType, time})
	g.nodes[to] = append(g.nodes[to], Edge{from, edgeType, time}) // Add reverse edge for undirected graph
}

type Item struct {
	node string
	dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq *PriorityQueue) Swap(i, j int)     { (*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(g *Graph, start string) map[string]int {
	dist := make(map[string]int)
	for node := range g.nodes {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{start, 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(Item)
		for _, edge := range g.nodes[current.node] {
			newDist := current.dist + edge.time
			if newDist < dist[edge.to] {
				dist[edge.to] = newDist
				heap.Push(pq, Item{edge.to, newDist})
			}
		}
	}
	return dist
}

func AllPairsShortestPathsConcurrent(g *Graph) map[string]map[string]int {
	result := make(map[string]map[string]int)
	var wg sync.WaitGroup

	for node := range g.nodes {
		wg.Add(1)
		go func(node string) {
			defer wg.Done()
			result[node] = Dijkstra(g, node)
		}(node)
	}
	wg.Wait()
	return result
}

func AllPairsShortestPathsSequential(g *Graph) map[string]map[string]int {
	result := make(map[string]map[string]int)
	for node := range g.nodes {
		result[node] = Dijkstra(g, node)
	}
	return result
}

func main() {
	graph := NewGraph()

	// Define the graph based on the provided structure
	graph.AddEdge("A", "B", "trip", 2)
	graph.AddEdge("A", "C", "trip", 4)
	graph.AddEdge("B", "A", "trip", 2)
	graph.AddEdge("B", "C", "transfer", 1)
	graph.AddEdge("B", "D", "trip", 7)
	graph.AddEdge("C", "A", "trip", 4)
	graph.AddEdge("C", "B", "transfer", 1)
	graph.AddEdge("C", "D", "trip", 3)
	graph.AddEdge("D", "B", "trip", 7)
	graph.AddEdge("D", "C", "trip", 3)

	fmt.Println("Calculating All-Pairs Shortest Paths...")

	start := time.Now()
	resultConcurrent := AllPairsShortestPathsConcurrent(graph)
	concurrentDuration := time.Since(start)

	start = time.Now()
	resultSequential := AllPairsShortestPathsSequential(graph)
	sequentialDuration := time.Since(start)

	fmt.Println("Concurrent Result:")
	for start, distances := range resultConcurrent {
		fmt.Printf("From node %s:\n", start)
		for dest, dist := range distances {
			fmt.Printf("  To node %s: %d\n", dest, dist)
		}
	}
	fmt.Printf("Concurrent execution time: %v\n", concurrentDuration)

	fmt.Println("Sequential Result:")
	for start, distances := range resultSequential {
		fmt.Printf("From node %s:\n", start)
		for dest, dist := range distances {
			fmt.Printf("  To node %s: %d\n", dest, dist)
		}
	}
	fmt.Printf("Sequential execution time: %v\n", sequentialDuration)
}
