package graph

import (
	"container/heap"
	"math"
)

// nodeItem is used in the priority queue.
type nodeItem struct {
	Node     string
	Priority float64
	Index    int
}

// priorityQueue implements the heap.Interface for nodeItems.
type priorityQueue []*nodeItem

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*nodeItem)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	(*pq)[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// DijkstraWithPath returns two maps:
// 1. dist[node] = shortest distance from source to node
// 2. prev[node] = predecessor of node on the path from source
func (g *Graph) DijkstraWithPath(source string) (map[string]float64, map[string]string) {
	dist := make(map[string]float64)
	prev := make(map[string]string)

	// Initialize distances
	for node := range g.AdjacencyList {
		dist[node] = math.Inf(1)
		prev[node] = ""
	}
	dist[source] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &nodeItem{Node: source, Priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*nodeItem)
		currentNode := current.Node

		for neighbor, weight := range g.AdjacencyList[currentNode] {
			alt := dist[currentNode] + weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				prev[neighbor] = currentNode
				heap.Push(pq, &nodeItem{Node: neighbor, Priority: alt})
			}
		}
	}

	return dist, prev
}

// ReconstructPath builds the path from source to target using the predecessor map.
func ReconstructPath(prev map[string]string, source, target string) []string {
	var path []string
	for at := target; at != ""; at = prev[at] {
		path = append([]string{at}, path...)
		if at == source {
			break
		}
	}
	return path
}
