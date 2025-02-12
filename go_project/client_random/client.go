package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"go_project/graph"
	"math/rand"
	"net"
	"sync"
	"time"
)

// Constants for random graph generation.
const (
	numClients = 5 // Number of parallel client connections
	numNodes   = 5 // Exact number of nodes in each graph
	maxEdges   = 5 // Maximum edges per node
	serverAddr = "localhost:8000"
)

// generateRandomGraph creates a random graph with an exact number of nodes (numNodes)
// and up to maxEdges edges per node, with random weights.
func generateRandomGraph() *graph.Graph {
	g := graph.NewGraph()

	// Generate node IDs and add them to the graph.
	nodeIDs := make([]string, 0, numNodes)
	for i := 0; i < numNodes; i++ {
		nodeIDs = append(nodeIDs, fmt.Sprintf("Node%d", i+1))
		g.AddNode(fmt.Sprintf("Node%d", i+1))
	}

	// For each node, add up to maxEdges random edges to other nodes.
	for _, from := range nodeIDs {
		numEdgesForNode := rand.Intn(maxEdges + 1)
		for j := 0; j < numEdgesForNode; j++ {
			to := nodeIDs[rand.Intn(numNodes)]
			if to == from {
				continue // Skip self-loop for simplicity
			}
			weight := float64(rand.Intn(10) + 1) // Random weight 1..10
			g.AddEdge(from, to, weight)
		}
	}
	return g
}

// sendGraphAndReceiveResults connects to the server, sends the graph,
// and appends the graph and results to the provided buffer.
func sendGraphAndReceiveResults(g *graph.Graph, clientID int, buffer *bytes.Buffer) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Fprintf(buffer, "[Client %d] Failed to connect to server at %s: %v\n", clientID, serverAddr, err)
		return
	}
	defer conn.Close()

	// Convert graph to JSON.
	data, err := json.Marshal(g)
	if err != nil {
		fmt.Fprintf(buffer, "[Client %d] JSON marshal error: %v\n", clientID, err)
		return
	}

	// Send the JSON-encoded graph (append newline to mark the end).
	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		fmt.Fprintf(buffer, "[Client %d] Error sending graph to server: %v\n", clientID, err)
		return
	}

	// Read the response (All-Pairs Shortest Paths).
	reader := bufio.NewReader(conn)
	response, err := reader.ReadBytes('\n') // Read until newline
	if err != nil {
		fmt.Fprintf(buffer, "[Client %d] Error reading response: %v\n", clientID, err)
	}

	// Unmarshal the results.
	var results map[string]map[string]graph.PathInfo
	err = json.Unmarshal(response, &results)
	if err != nil {
		fmt.Fprintf(buffer, "[Client %d] Error unmarshaling response: %v\n", clientID, err)
		return
	}

	// Print the results grouped by source node.
	fmt.Fprintf(buffer, "[Client %d] Received All-Pairs Shortest Paths:\n", clientID)
	for src, dstMap := range results {
		fmt.Fprintf(buffer, "Source: %s\n", src)
		for dst, info := range dstMap {
			if info.Distance == -1 {
				fmt.Fprintf(buffer, "  -> %s: Distance: Unreachable, Path: N/A\n", dst)
			} else {
				fmt.Fprintf(buffer, "  -> %s: Distance: %.2f, Path: %v\n", dst, info.Distance, info.Path)
			}
		}
		fmt.Fprintln(buffer) // Add a blank line for readability
	}
}

// main starts multiple clients to send graphs to the server.
func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the global random number generator

	var wg sync.WaitGroup
	for i := 1; i <= numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			// Use a buffer to collect output for this client
			var buffer bytes.Buffer

			// Generate the random graph
			g := generateRandomGraph()

			// Print the generated graph
			fmt.Fprintf(&buffer, "[Client %d] Printing the random graph...\n", clientID)
			for from, edges := range g.AdjacencyList {
				for to, weight := range edges {
					fmt.Fprintf(&buffer, "  %s -> %s [Weight: %.2f]\n", from, to, weight)
				}
			}

			// Send the graph and process the results
			sendGraphAndReceiveResults(g, clientID, &buffer)

			// Print all buffered output for this client
			fmt.Println(buffer.String())
		}(i)
	}

	wg.Wait() // Wait for all clients to finish
	fmt.Println("All clients finished.")
}
