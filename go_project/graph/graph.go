package graph

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
	g.AdjacencyList[from][to] = weight
}

// AddNode adds a node to the graph. If the node already exists, it does nothing.
func (g *Graph) AddNode(node string) {
	if _, exists := g.AdjacencyList[node]; !exists {
		g.AdjacencyList[node] = make(map[string]float64)
	}
}
