# Golang All-Pair Shortest Paths Project

This project implements a **graph-based shortest path algorithm** using a server-client architecture in Golang. The server computes **all-pairs shortest paths (N x Dijkstra)** for graphs provided by clients. Clients can generate random graphs or load graphs from CSV files.


## How to Run

### 1. Start the Server
Run the server to listen for client connections:
```bash
cd server
go run server.go
```

### 2. Run a Client

#### Random Graph Client
Run a client that generates random graphs:
```bash
cd client
go run random_client.go
```

#### CSV-Based Client
Run a client that loads graph data from CSV files:
```bash
cd client
go run csv_client.go
```

## Code Overview

### Graph Library 
- `NewGraph()`: Creates a new graph.
- `AddNode(node string)`: Adds a node to the graph.
- `AddEdge(from, to string, weight float64)`: Adds a directed edge.
- `DijkstraWithPath(source string)`: Computes shortest paths from a source node.
- `AllPairsShortestPaths()`: Computes shortest paths between all pairs of nodes.

### Server (`server/server.go`)
- Listens on `localhost:8000` for incoming client connections.
- Uses Goroutines to handle multiple clients concurrently.
- Accepts JSON-encoded graphs, computes all-pairs shortest paths, and sends the results back.

### Clients
- **Random Client** (`client/random_client.go`):
  - Generates a random graph with:
    - `numNodes`: Number of nodes.
    - `maxEdges`: Maximum edges per node.
  - Sends the graph to the server and displays results.
- **CSV-Based Client** (`client/csv_client.go`):
  - Loads graph data from `stops.csv`, `trips.csv`, and `transfers.csv`.
  - Sends the graph to the server and displays results.
