package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

// Graph and Edge definitions
type Graph struct {
	Nodes map[string][]Edge `json:"nodes"`
}

type Edge struct {
	To       string `json:"to"`
	EdgeType string `json:"edgeType"`
	Time     int    `json:"time"`
}

// Generates a unique graph for each client
func generateGraph(clientID, numNodes, maxEdgesPerNode int) *Graph {
	graph := &Graph{Nodes: make(map[string][]Edge)}
	rand.Seed(time.Now().UnixNano() + int64(clientID)) // Ensure different random values per client

	for i := 0; i < numNodes; i++ {
		node := fmt.Sprintf("%d", i)
		numEdges := rand.Intn(maxEdgesPerNode) + 1 // Random number of edges per node

		for j := 1; j <= numEdges && i+j < numNodes; j++ {
			target := fmt.Sprintf("%d", i+j)
			weight := rand.Intn(10) + 1 // Random weight between 1 and 10
			graph.Nodes[node] = append(graph.Nodes[node], Edge{To: target, EdgeType: "direct", Time: weight})
			graph.Nodes[target] = append(graph.Nodes[target], Edge{To: node, EdgeType: "direct", Time: weight}) // Bidirectional
		}
	}

	return graph
}

// Runs a client that connects to the server, sends a graph, and receives results
func runClient(clientID, numNodes, maxEdgesPerNode int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Connect to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Printf("[Client %d] Error connecting to server: %v\n", clientID, err)
		return
	}
	defer conn.Close()

	// Generate a unique graph
	graph := generateGraph(clientID, numNodes, maxEdgesPerNode)
	graphJSON, err := json.Marshal(graph)
	if err != nil {
		fmt.Printf("[Client %d] Error encoding graph: %v\n", clientID, err)
		return
	}

	// Send the graph to the server
	fmt.Fprintf(conn, "%s\n", graphJSON)

	// Read the response
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("[Client %d] Error reading response: %v\n", clientID, err)
		return
	}

	// Check for server errors
	if strings.HasPrefix(response, "Error") {
		fmt.Printf("[Client %d] Server error: %s\n", clientID, response)
		return
	}

	// Parse the response
	var results map[string]map[string]int
	if err := json.Unmarshal([]byte(response), &results); err != nil {
		fmt.Printf("[Client %d] Error decoding response: %v\n", clientID, err)
		fmt.Printf("[Client %d] Raw response: %s\n", clientID, response)
		return
	}

	// Print the results
	fmt.Printf("\n[Client %d] Received results:\n", clientID)
	for src, dists := range results {
		fmt.Printf("  From Node %s:\n", src)
		for dest, dist := range dists {
			if dist == int(^uint(0)>>1) { // math.MaxInt equivalent
				fmt.Printf("    To Node %s: Unreachable\n", dest)
			} else {
				fmt.Printf("    To Node %s: %d\n", dest, dist)
			}
		}
	}
}

func main() {
	const numClients = 5 // Number of clients to run
	const maxNodes = 30  // Maximum number of nodes in graphs
	const maxEdges = 5   // Maximum number of edges per node

	var wg sync.WaitGroup

	// Start multiple clients concurrently
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go runClient(i, rand.Intn(maxNodes-5)+5, maxEdges, &wg) // Each client gets a different graph size
	}

	wg.Wait() // Wait for all clients to finish
	fmt.Println("\nAll clients have finished processing.")
}
