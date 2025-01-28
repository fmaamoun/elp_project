package main

import (
	"container/heap"
	"math"
	"sync"
)

/*
------------------------------------
1) Single-Pair Dijkstra (start->end)
------------------------------------
*/

// pqItem is a node+distance pair for the priority queue.
type pqItem struct {
	node string
	dist int
}

// priorityQueue implements heap.Interface for pqItem.
type priorityQueue []pqItem

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(pqItem)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// DijkstraSinglePair calculates the shortest path between two specific nodes
// (start -> end). Returns (distance, path). If unreachable, distance == math.MaxInt, path = [].
func DijkstraSinglePair(g *Graph, start, end string) (int, []string) {
	// Distances and parents for path reconstruction
	dist := make(map[string]int)
	parent := make(map[string]string)

	for node := range g.data {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqItem{node: start, dist: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(pqItem)
		currNode := current.node
		currDist := current.dist

		// Optimization: if we've reached 'end', stop
		if currNode == end {
			break
		}

		// If we have already a better distance, skip
		if currDist > dist[currNode] {
			continue
		}

		// Explore neighbors
		for _, edge := range g.data[currNode] {
			newDist := currDist + edge.Time
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				parent[edge.To] = currNode
				heap.Push(pq, pqItem{node: edge.To, dist: newDist})
			}
		}
	}

	finalDist := dist[end]
	if finalDist == math.MaxInt {
		return finalDist, []string{} // unreachable
	}

	// Reconstruct path
	path := []string{}
	cur := end
	for {
		path = append([]string{cur}, path...)
		if cur == start {
			break
		}
		cur = parent[cur]
	}

	return finalDist, path
}

/*
------------------------------------------------
2) Helper: Single-Source Dijkstra for All-Pairs
------------------------------------------------
We need a single-source Dijkstra that returns the distance to ALL nodes.
This is used by the all-pairs functions below.
*/
func dijkstraSingleSource(g *Graph, start string) map[string]int {
	dist := make(map[string]int)
	for node := range g.data {
		dist[node] = math.MaxInt
	}
	dist[start] = 0

	pq := &priorityQueue{}
	heap.Init(pq)
	heap.Push(pq, pqItem{node: start, dist: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(pqItem)
		currNode := current.node
		currDist := current.dist

		if currDist > dist[currNode] {
			continue
		}

		for _, edge := range g.data[currNode] {
			newDist := currDist + edge.Time
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				heap.Push(pq, pqItem{node: edge.To, dist: newDist})
			}
		}
	}
	return dist
}

/*
--------------------------------
3) All-Pairs Dijkstra: Concurrent
--------------------------------
For each node, run dijkstraSingleSource in a goroutine.
*/
func AllPairsShortestPathsConcurrent(g *Graph) map[string]map[string]int {
	result := make(map[string]map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for node := range g.data {
		wg.Add(1)
		go func(n string) {
			defer wg.Done()
			distances := dijkstraSingleSource(g, n)

			// Safely store the result
			mu.Lock()
			result[n] = distances
			mu.Unlock()
		}(node)
	}
	wg.Wait()
	return result
}

/*
---------------------------------
4) All-Pairs Dijkstra: Sequential
---------------------------------
For each node, run dijkstraSingleSource in a simple for-loop (no concurrency).
*/
func AllPairsShortestPathsSequential(g *Graph) map[string]map[string]int {
	result := make(map[string]map[string]int)
	for node := range g.data {
		result[node] = dijkstraSingleSource(g, node)
	}
	return result
}
