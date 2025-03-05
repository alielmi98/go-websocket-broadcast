package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	var addr = "localhost:8080"
	fmt.Println("Connecting to server...")

	conn, _, err := websocket.DefaultDialer.Dial("ws://"+addr+"/ws", nil)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	go func() {
		for {
			var msg string
			err := conn.ReadJSON(&msg)
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("Received: %s\n", msg)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Println("write:", err)
			return
		}
	}
}
