package graph

import "math"

// PathInfo represents the shortest path and its total distance.
type PathInfo struct {
	Path     []string `json:"path"`
	Distance float64  `json:"distance"`
}

// AllPairsShortestPaths computes the shortest paths and total distances for all pairs of nodes.
// It returns a map of (source -> destination -> PathInfo).
// Unreachable destinations have Distance set to -1 and an empty Path.
func (g *Graph) AllPairsShortestPaths() map[string]map[string]PathInfo {
	results := make(map[string]map[string]PathInfo)

	for src := range g.AdjacencyList {
		dist, prev := g.DijkstraWithPath(src)
		results[src] = make(map[string]PathInfo)
		for dst := range g.AdjacencyList {
			if src == dst {
				continue // Optionally, skip or handle self-pairs
			}
			if dist[dst] == math.Inf(1) {
				// Destination is unreachable from source
				results[src][dst] = PathInfo{
					Path:     []string{}, // Empty path
					Distance: -1,         // Sentinel value indicating no path
				}
			} else {
				path := ReconstructPath(prev, src, dst)
				results[src][dst] = PathInfo{
					Path:     path,
					Distance: dist[dst],
				}
			}
		}
	}

	return results
}
