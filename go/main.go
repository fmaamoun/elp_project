package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

func main() {
	// 1. Create an empty graph
	graph := NewGraph()

	// 2. Load stops, trips, and transfers from CSV files
	//    (Adjust file paths to match your actual directory structure!)
	err := loadStops("data/stops.csv", graph)
	if err != nil {
		log.Fatalf("Error loading stops: %v", err)
	}
	err = loadTrips("data/trips.csv", graph)
	if err != nil {
		log.Fatalf("Error loading trips: %v", err)
	}
	err = loadTransfers("data/transfers.csv", graph)
	if err != nil {
		log.Fatalf("Error loading transfers: %v", err)
	}

	// 3. (Optional) Print the full graph to verify data
	// fmt.Println("---- Loaded Graph ----")
	// graph.PrintGraph()

	/*
	   -----------------------------------------------------------
	   Measure time for each of the Dijkstra variations:

	   A) Single-Pair from "startNode" to "endNode"
	   B) All-Pairs Shortest Paths (Concurrent)
	   C) All-Pairs Shortest Paths (Sequential)
	   -----------------------------------------------------------
	*/

	startNode := "MA01"
	endNode := "T101"

	// A) Single-Pair
	fmt.Printf("\n--- Single Pair: Dijkstra from %s to %s ---\n", startNode, endNode)
	t0 := time.Now()
	dist, path := DijkstraSinglePair(graph, startNode, endNode)
	t1 := time.Since(t0)
	if dist == math.MaxInt {
		fmt.Printf("No path found from %s to %s\n", startNode, endNode)
	} else {
		fmt.Printf("Distance: %d\n", dist)
		fmt.Printf("Path: %v\n", path)
	}
	fmt.Printf("Single-Pair run time: %v\n", t1)

	// B) All-Pairs (Concurrent)
	fmt.Println("\n--- All-Pairs Shortest Paths (Concurrent) ---")
	startTime := time.Now()
	apspConcurrent := AllPairsShortestPathsConcurrent(graph)
	concurrentTime := time.Since(startTime)
	fmt.Printf("Concurrent APSP completed in %v\n", concurrentTime)

	// Demonstrate reading from the map (optional)
	if distances, ok := apspConcurrent[startNode]; ok {
		fmt.Printf("Distances from node %s (concurrent):\n", startNode)
		for node, d := range distances {
			fmt.Printf("  %s: %d\n", node, d)
		}
	}

	// C) All-Pairs (Sequential)
	fmt.Println("\n--- All-Pairs Shortest Paths (Sequential) ---")
	startTime = time.Now()
	apspSequential := AllPairsShortestPathsSequential(graph)
	sequentialTime := time.Since(startTime)
	fmt.Printf("Sequential APSP completed in %v\n", sequentialTime)

	if distances, ok := apspSequential[startNode]; ok {
		fmt.Printf("Distances from node %s (sequential):\n", startNode)
		for node, d := range distances {
			fmt.Printf("  %s: %d\n", node, d)
		}
	}
}
