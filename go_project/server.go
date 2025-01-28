package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"sync"
	"time"

	"example/go_project/structure" // Updated import path
)

// handleConnection processes individual client connections.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Decode JSON graph from client
	var graph structure.Graph
	err := json.NewDecoder(conn).Decode(&graph)
	if err != nil {
		fmt.Fprintf(conn, "Error decoding JSON: %s\n", err)
		return
	}

	// Compute all-pairs shortest paths
	start := time.Now()
	results := allPairsShortestPaths(&graph)
	elapsed := time.Since(start)

	// Encode and send results back to client
	err = json.NewEncoder(conn).Encode(results)
	if err != nil {
		fmt.Fprintf(conn, "Error encoding JSON: %s\n", err)
		return
	}

	fmt.Printf("Processed request in %v\n", elapsed)
}

// allPairsShortestPaths calculates all pairs shortest paths using Dijkstra's algorithm.
func allPairsShortestPaths(g *structure.Graph) map[string]map[string]int {
	const MaxWorkers = 10 // limit the number of workers
	var wg sync.WaitGroup
	results := make(map[string]map[string]int)
	mu := sync.Mutex{}

	jobs := make(chan string, len(g.Nodes)) // Job queue

	// Worker function
	worker := func() {
		for node := range jobs { // Pick tasks from the job queue
			dist := dijkstra(g, node)
			mu.Lock()
			results[node] = dist
			mu.Unlock()
			wg.Done()
		}
	}

	// Start worker pool
	for i := 0; i < MaxWorkers; i++ {
		go worker()
	}

	// Add jobs (nodes) to the queue
	for node := range g.Nodes {
		wg.Add(1)
		jobs <- node
	}

	close(jobs) // Close channel after sending all jobs
	wg.Wait()   // Wait for all workers to complete

	return results
}

// dijkstra implements Dijkstra's algorithm for a single source.
func dijkstra(g *structure.Graph, src string) map[string]int {
	dist := make(map[string]int)
	for node := range g.Nodes {
		dist[node] = math.MaxInt
	}
	dist[src] = 0

	pq := &structure.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, structure.Item{Node: src, Dist: 0})

	for pq.Len() > 0 {
		current := heap.Pop(pq).(structure.Item)
		currentNode := current.Node
		currentDist := current.Dist

		for _, edge := range g.Nodes[currentNode] {
			newDist := currentDist + edge.Time
			if newDist < dist[edge.To] {
				dist[edge.To] = newDist
				heap.Push(pq, structure.Item{Node: edge.To, Dist: newDist})
			}
		}
	}

	return dist
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is running on port 8000...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
