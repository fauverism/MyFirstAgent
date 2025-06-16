package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/yourusername/code-editing-agent/internal" // Adjust the import path as necessary
)

var upgrader = websocket.Upgrader{}

func main() {
	agent := internal.NewAgent() // Assuming NewAgent initializes the agent
	http.HandleFunc("/ws", handleWebSocket(agent))
	http.Handle("/", http.FileServer(http.Dir("./ui/static")))

	fmt.Println("Server started at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handleWebSocket(agent *internal.Agent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				break
			}

			response, err := agent.Run(context.Background(), string(msg))
			if err != nil {
				fmt.Println("Error running agent:", err)
				break
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(response)); err != nil {
				fmt.Println("Error writing message:", err)
				break
			}
		}
	}
}