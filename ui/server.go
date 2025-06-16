package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", serveStaticFiles)
	http.HandleFunc("/ws", handleWebSocket)

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func serveStaticFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "ui/static/index.html")
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return
	}
	defer conn.Close()

	for {
		var msg string
		err := conn.ReadMessage(&msg)
		if err != nil {
			log.Println("Error while reading message:", err)
			break
		}

		// Here you would handle the message and interact with the chat agent
		// For now, we will just echo the message back
		err = conn.WriteMessage(websocket.TextMessage, []byte("Echo: "+msg))
		if err != nil {
			log.Println("Error while writing message:", err)
			break
		}
	}
}