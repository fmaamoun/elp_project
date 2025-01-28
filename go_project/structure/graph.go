package structure

// Graph represents a graph structure with an adjacency list.
type Graph struct {
	Nodes map[string][]Edge
}

// Edge represents a single edge in the graph.
type Edge struct {
	To       string `json:"to"`
	EdgeType string `json:"edgeType"`
	Time     int    `json:"time"`
}

// NewGraph initializes a new empty Graph.
func NewGraph() *Graph {
	return &Graph{Nodes: make(map[string][]Edge)}
}

// AddEdge adds a bidirectional edge to the graph.
func (g *Graph) AddEdge(from, to, edgeType string, time int) {
	g.Nodes[from] = append(g.Nodes[from], Edge{To: to, EdgeType: edgeType, Time: time})
	g.Nodes[to] = append(g.Nodes[to], Edge{To: from, EdgeType: edgeType, Time: time}) // Ensuring bidirectional
}

// PriorityQueue and Item definitions for Dijkstra's algorithm.
type Item struct {
	Node string
	Dist int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}
func (pq PriorityQueue) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }

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
