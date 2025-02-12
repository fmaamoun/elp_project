package graph

import (
	"math"
	"sync"
)

// PathInfo represents the shortest path and its total distance.
type PathInfo struct {
	Path     []string `json:"path"`
	Distance float64  `json:"distance"`
}

// AllPairsShortestPaths computes the shortest paths and total distances for all pairs of nodes
// using a worker pool with a fixed maximum number of workers. Unreachable destinations have
// Distance set to -1 and an empty Path.
func (g *Graph) AllPairsShortestPaths() map[string]map[string]PathInfo {
	const MaxWorkers = 5 // Limit the number of worker goroutines

	var wg sync.WaitGroup
	results := make(map[string]map[string]PathInfo)
	mu := sync.Mutex{}

	// jobs channel carries the source nodes (as strings) to process.
	jobs := make(chan string, len(g.AdjacencyList))

	// Worker function that processes nodes from the jobs channel.
	worker := func() {
		for src := range jobs {
			// Compute shortest paths from src using Dijkstra's algorithm.
			dist, prev := g.DijkstraWithPath(src)
			paths := make(map[string]PathInfo)

			for dst := range g.AdjacencyList {
				// Skip the source itself.
				if src == dst {
					continue
				}
				if dist[dst] == math.Inf(1) {
					// No path found: record unreachable destination.
					paths[dst] = PathInfo{
						Path:     []string{},
						Distance: -1,
					}
				} else {
					// Reconstruct the shortest path from src to dst.
					paths[dst] = PathInfo{
						Path:     ReconstructPath(prev, src, dst),
						Distance: dist[dst],
					}
				}
			}

			// Safely update the shared results map.
			mu.Lock()
			results[src] = paths
			mu.Unlock()

			// Signal completion for this job.
			wg.Done()
		}
	}

	// Start the worker pool.
	for i := 0; i < MaxWorkers; i++ {
		go worker()
	}

	// Enqueue all source nodes into the jobs channel.
	for src := range g.AdjacencyList {
		wg.Add(1)
		jobs <- src
	}
	close(jobs) // No more jobs will be sent.

	// Wait until all jobs have been processed.
	wg.Wait()

	return results
}
