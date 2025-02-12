package main

import (
	"bufio"
	"encoding/json"
	"go_project/graph"
	"log"
	"net"
	"sync"
)

// ServerPort is the TCP port the server will listen on.
const ServerPort = ":8000"

// handleClient handles communication with a single client.
func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	defer conn.Close()

	reader := bufio.NewReader(conn)

	// Read the raw data (JSON) that the client sends.
	rawData, err := reader.ReadBytes('\n')
	if err != nil {
		log.Printf("Error reading client data: %v\n", err)
		return
	}

	// Parse the received graph.
	var g graph.Graph
	err = json.Unmarshal(rawData, &g)
	if err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		return
	}

	// Compute All-Pairs Shortest Paths (N x Dijkstra).
	results := g.AllPairsShortestPaths()

	// Encode the results as JSON and send back to the client.
	response, err := json.Marshal(results)
	if err != nil {
		log.Printf("Error encoding JSON: %v\n", err)
		return
	}

	// Send the response followed by a newline character to indicate the end.
	_, err = conn.Write(append(response, '\n'))
	if err != nil {
		log.Printf("Error sending response: %v\n", err)
		return
	}
}

// main starts a TCP server that can handle multiple clients simultaneously.
func main() {
	listener, err := net.Listen("tcp", ServerPort)
	if err != nil {
		log.Fatalf("Failed to start server on %s: %v", ServerPort, err)
	}
	defer listener.Close()

	log.Printf("Server listening on %s...\n", ServerPort)

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
			continue
		}

		wg.Add(1)
		go handleClient(conn, &wg)
	}

	// Note: wg.Wait() is not reachable here.
}
