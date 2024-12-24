package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
) // used for creating persistent connection - gorilla websockets

type Message struct {
	Counter      int    `json:"counter"`
	RandomString string `json:"randomString"`
}

var (
	counter  int
	clients  = make(map[*websocket.Conn]bool)
	upgrader = websocket.Upgrader{ // upgrader which allows websocket connection
		ReadBufferSize:  32,
		WriteBufferSize: 32,
	}
	lock sync.Mutex
)

func main() {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// WebSocket endpoint
	http.HandleFunc("/ws", handleConnections)

	// REST API endpoint
	http.HandleFunc("/data", handleData)

	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Frontend origin
		handlers.AllowedMethods([]string{"GET", "POST"}),
	)

	go handleMessages()

	fmt.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", corsOptions(http.DefaultServeMux)); err != nil {
		fmt.Println(err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error upgrading to WebSocket:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New client connected")
	clients[conn] = true

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			delete(clients, conn)
			break
		}
	}
}

func handleMessages() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop() //defer -> Used for closing the ticker when the function returns

	for range ticker.C {
		lock.Lock() // lock -> used for mutual exclusion (independence of processes through avoiding switching)
		counter++
		message := Message{
			Counter:      counter,
			RandomString: generateRandomString(10),
		}
		lock.Unlock()

		// Send message to all clients
		lock.Lock()
		for client := range clients {
			err := client.WriteJSON(message)
			if err != nil {
				fmt.Println("Error writing JSON:", err)
				client.Close()
				delete(clients, client)
			}
		}
		lock.Unlock()
	}
}

func generateRandomString(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length) // make strings
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func handleData(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	message := Message{
		Counter:      counter,
		RandomString: generateRandomString(10),
	}
	lock.Unlock()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
