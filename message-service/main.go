// main.go

package main

import (
	"chat-app-microservice/message-service/handler"
	"chat-app-microservice/message-service/router"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	router.Use(corsHandler.Handler)

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Handle WebSocket connections here
		// Initialize your WebSocketHandler and manage WebSocket connections
		webSocketHandler := handler.NewWebSocketHandler()
		webSocketHandler.HandleWebSocket(w, r)
	})

	// Set up routes using your custom router package
	router.SetupRoutes(r, db)

	// Start the server
	port := ":8081" // You can use a different port if needed
	log.Printf("Server is running on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
