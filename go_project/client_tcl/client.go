package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"go_project/graph"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

// TCP server address
const serverAddr = "localhost:8000"

// LoadStops loads stop data from stops.csv and adds the stops to the graph
func LoadStops(filePath string, graph *graph.Graph) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		return fmt.Errorf("failed to read header row: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}
		stopID := record[0] // stop_id
		graph.AddNode(stopID)
	}

	return nil
}

// LoadTrips loads trip data from trips.csv and adds the trips to the graph
func LoadTrips(filePath string, graph *graph.Graph) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		return fmt.Errorf("failed to read header row: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}

		fromStopID := record[1]                               // from_stop_id
		toStopID := record[2]                                 // to_stop_id
		timeSeconds, err := strconv.ParseFloat(record[3], 64) // time (in seconds)
		if err != nil {
			return fmt.Errorf("invalid time value: %w", err)
		}

		graph.AddEdge(fromStopID, toStopID, timeSeconds)
		graph.AddEdge(toStopID, fromStopID, timeSeconds)
	}

	return nil
}

// LoadTransfers loads transfer data from transfers.csv and adds the transfers to the graph
func LoadTransfers(filePath string, graph *graph.Graph) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filePath, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read() // Skip the header row
	if err != nil {
		return fmt.Errorf("failed to read header row: %w", err)
	}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}

		fromStopID := record[1]                               // from_stop_id
		toStopID := record[2]                                 // to_stop_id
		timeSeconds, err := strconv.ParseFloat(record[3], 64) // min_transfer_time (in seconds)
		if err != nil {
			return fmt.Errorf("invalid transfer time value: %w", err)
		}

		graph.AddEdge(fromStopID, toStopID, timeSeconds)
		graph.AddEdge(toStopID, fromStopID, timeSeconds)
	}

	return nil
}

// sendGraphAndReceiveResults sends the graph to the server and prints the received results
func sendGraphAndReceiveResults(g *graph.Graph) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to connect to server at %s: %v", serverAddr, err)
	}
	defer conn.Close()

	// Serialize the graph to JSON
	data, err := json.Marshal(g)
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}

	// Send the JSON-encoded graph (append newline to mark the end)
	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		log.Fatalf("Error sending graph to server: %v", err)
	}

	// Prepare a buffer to read the response
	var response bytes.Buffer
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n > 0 {
			response.Write(buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading server response: %v", err)
		}
	}

	// Deserialize the server response
	var results map[string]map[string]graph.PathInfo
	err = json.Unmarshal(response.Bytes(), &results)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	// Print the results grouped by source node
	fmt.Println("Received All-Pairs Shortest Paths:")
	for src, dstMap := range results {
		fmt.Printf("Source: %s\n", src)
		for dst, info := range dstMap {
			if info.Distance == -1 {
				fmt.Printf("  -> %s: Temps (en sec) : Unreachable, Path: N/A\n", dst)
			} else {
				fmt.Printf("  -> %s: Temps (en sec) : %.2f, Path: %v\n", dst, info.Distance, info.Path)
			}
		}
		fmt.Println() // Add a blank line for readability
	}
}

func main() {
	// Init graph
	g := graph.NewGraph()

	// Load graph data from CSV files
	if err := LoadStops("./data/stops.csv", g); err != nil {
		log.Fatalf("Error loading stops: %v", err)
	}
	if err := LoadTrips("./data/trips.csv", g); err != nil {
		log.Fatalf("Error loading trips: %v", err)
	}
	if err := LoadTransfers("./data/transfers.csv", g); err != nil {
		log.Fatalf("Error loading transfers: %v", err)
	}

	// Send the graph to the server and process the results
	sendGraphAndReceiveResults(g)
}
