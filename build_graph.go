package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// loadStops loads stop data from stops.csv and adds the stops to the graph
func LoadStops(filePath string, graph *Graph) error {
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
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}
		stopID := record[0] // stop_id
		graph.AddStop(stopID)
	}

	return nil
}

// loadTrips loads trip data from trips.csv and adds the trips to the graph
func LoadTrips(filePath string, graph *Graph) error {
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
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}

		fromStopID := record[1]                     // from_stop_id
		toStopID := record[2]                       // to_stop_id
		timeSeconds, err := strconv.Atoi(record[3]) // time (in seconds)
		if err != nil {
			return fmt.Errorf("invalid time value: %w", err)
		}

		// Add edge in both directions
		graph.AddEdge(fromStopID, toStopID, "trip", timeSeconds)
		graph.AddEdge(toStopID, fromStopID, "trip", timeSeconds)
	}

	return nil
}

// loadTransfers loads transfer data from transfers.csv and adds the transfers to the graph
func LoadTransfers(filePath string, graph *Graph) error {
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
			if err.Error() == "EOF" {
				break
			}
			return fmt.Errorf("failed to read record: %w", err)
		}

		fromStopID := record[1]                     // from_stop_id
		toStopID := record[2]                       // to_stop_id
		timeSeconds, err := strconv.Atoi(record[3]) // min_transfer_time (in seconds)
		if err != nil {
			return fmt.Errorf("invalid transfer time value: %w", err)
		}

		// Add edge in both directions
		graph.AddEdge(fromStopID, toStopID, "transfer", timeSeconds)
		graph.AddEdge(toStopID, fromStopID, "transfer", timeSeconds)
	}

	return nil
}
