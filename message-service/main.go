// main.go

package main

import (
	"chat-app-microservice/message-service/handler"
	"chat-app-microservice/message-service/router"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

func main() {
	// Replace with your database connection string
	db, err := sql.Open("postgres", "user=username dbname=mychatapp sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Set up message routes
	r := mux.NewRouter()

	// Setting up CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Replace with your allowed origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Use CORS middleware
	r.Use(corsHandler.Handler)
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading connection to WebSocket: %v", err)
			return
		}
		defer conn.Close()
		webSocketHandler := handler.NewWebSocketHandler()
		webSocketHandler.HandleWebSocket(conn)
	})

	// Set up routes using your custom router package
	router.SetupRoutes(r, db)

	// Start the server
	port := ":8081" // You can use a different port if needed
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
