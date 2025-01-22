# **Graph and Dijkstra Algorithm Project**

## **Overview**

This project implements a graph structure and provides utilities to:

1. Load data from CSV files to build a transportation graph.
2. Compute shortest paths using Dijkstra's algorithm:
   - Single-pair shortest path.
   - All-pairs shortest paths (concurrent).
   - All-pairs shortest paths (sequential).

The program demonstrates the efficiency and use of these algorithms by timing each computation.

## **Project Structure**

The project is divided into four key components:

### 1. `graph.go`
- Defines the `Graph` and `Edge` structs.
- Provides methods to:
  - Add nodes (`AddStop`).
  - Add edges (`AddEdge`).
  - Print the graph (`PrintGraph`).

### 2. `build_graph.go`
- Loads graph data from CSV files:
  - `stops.csv`: List of stops/nodes.
  - `trips.csv`: Direct edges between stops with trip times.
  - `transfers.csv`: Transfers between stops with transfer times.

### 3. `dijkstra.go`
- Implements Dijkstra's algorithm:
  - **`DijkstraSinglePair`**: Calculates the shortest path between two specific nodes.
  - **`AllPairsShortestPathsConcurrent`**: Calculates shortest paths between all nodes using concurrency.
  - **`AllPairsShortestPathsSequential`**: Calculates shortest paths between all nodes sequentially.

### 4. `main.go`
- Entry point of the program:
  - Builds the graph by loading CSV data.
  - Demonstrates:
    1. Single-pair shortest path computation.
    2. Concurrent all-pairs computation.
    3. Sequential all-pairs computation.
  - Measures and compares execution times.


## **How to Run**

Run the program with:

```bash
go run main.go graph.go dijkstra.go build_graph.go
```