package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"sync"
)

var graph = map[string][]string{
	"A": {"B", "C"},
	"B": {"A", "D", "E"},
	"C": {"A", "F"},
	"D": {"B"},
	"E": {"B", "F"},
	"F": {"C", "E"},
}

// Fonction pour calculer l'itinéraire entre deux stations
func findRoute(start, end string) string {
	// Utilisation de la recherche en largeur (BFS) pour trouver le chemin le plus court
	visited := make(map[string]bool) // Stations visitées
	queue := []string{start}         // File d'attente pour BFS
	prev := make(map[string]string)  // Pour reconstruire le chemin

	// Parcours BFS
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Si on atteint la station d'arrivée
		if current == end {
			path := []string{}
			for at := end; at != ""; at = prev[at] {
				path = append([]string{at}, path...)
			}
			return strings.Join(path, " -> ") // Retourne le chemin sous forme de chaîne
		}

		// Parcourir les voisins
		for _, neighbor := range graph[current] {
			if !visited[neighbor] {
				visited[neighbor] = true
				prev[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	// Si aucun chemin trouvé
	return "Itinéraire introuvable"
}

// Structure pour le message
type Message struct {
	To   string `json:"to"`
	From string `json:"from"`
}

// Fonction pour gérer un client
func handleClient(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()    // Indique au WaitGroup que cette goroutine est terminée
	defer conn.Close() // Ferme la connexion une fois le traitement terminé

	// Crée un lecteur pour lire les données envoyées par le client
	reader := bufio.NewReader(conn)

	// Lire une ligne envoyée par le client (jusqu'à '\n')
	data, err := reader.ReadString('\n')
	if err != nil { // Vérifie s'il y a une erreur lors de la lecture
		fmt.Println("Erreur de lecture :", err)
		return
	}

	// Supprime les espaces inutiles autour de la chaîne reçue
	data = strings.TrimSpace(data)

	// Déclare une variable pour stocker le message décodé
	var msg Message

	// Décoder les données JSON reçues dans la structure Message
	err = json.Unmarshal([]byte(data), &msg)
	if err != nil { // Vérifie si le format JSON est invalide
		conn.Write([]byte("Erreur : format JSON invalide\n")) // Envoie un message d'erreur au client
		fmt.Println("Erreur de décodage JSON :", err)
		return
	}

	// Affiche dans la console le contenu du message reçu
	fmt.Printf("Message reçu : De %s à %s\n", msg.From, msg.To)

	// Utiliser la fonction findRoute pour calculer l'itinéraire
	route := findRoute(msg.From, msg.To)

	// Envoyer le résultat de l'itinéraire au client
	conn.Write([]byte(route + "\n"))
}

// Fonction principale du programme (point d'entrée)
func main() {
	port := ":8000"                          // Définit le port sur lequel le serveur écoute les connexions
	listener, err := net.Listen("tcp", port) // Démarre un serveur TCP sur le port spécifié
	if err != nil {                          // Vérifie s'il y a une erreur lors de l'initialisation
		fmt.Println("Erreur lors de l'écoute :", err)
		return
	}
	defer listener.Close() // Ferme le serveur lorsque le programme se termine
	fmt.Println("Serveur en écoute sur le port", port)

	var wg sync.WaitGroup // Crée un WaitGroup pour suivre les goroutines actives

	for {
		// Accepte une connexion entrante d'un client
		conn, err := listener.Accept()
		if err != nil { // Vérifie s'il y a une erreur lors de l'acceptation
			fmt.Println("Erreur lors de l'acceptation :", err)
			continue
		}

		// Ajoute une goroutine au WaitGroup
		wg.Add(1)

		// Lance une goroutine pour gérer la connexion client
		go handleClient(conn, &wg)
	}
}
