package main

import (
	"flag"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan string)
	mu        sync.Mutex
)

func main() {
	port := flag.String("port", "8080", "Port to run the server on")
	flag.Parse()

	http.HandleFunc("/ws", handleConnections)
	go handleBroadcast()

	log.Printf("Server started on :%s\n", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading connection: ", err)
		return
	}
	defer conn.Close()

	mu.Lock()
	clients[conn] = true
	mu.Unlock()

	for {
		var msg string
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading message: ", err)
			break
		}
		broadcast <- msg
	}

	mu.Lock()
	delete(clients, conn)
	mu.Unlock()
}

func handleBroadcast() {
	for {
		msg := <-broadcast
		mu.Lock()
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("Error broadcasting message: ", err)
				client.Close()
				delete(clients, client)
			}
		}
		mu.Unlock()
	}
}
