// calling-service/main.go

package main

import (
	"chat-app-microservice/calling-service/db"
	"chat-app-microservice/calling-service/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize the database connection
	database, err := db.Initialize()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Conn.Close()

	// Create a new router
	r := mux.NewRouter()

	// Set up CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Replace with your allowed origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Use CORS middleware
	r.Use(corsHandler.Handler)

	// Set up video calling routes
	router.SetupRoutes(r, database.Conn)

	// Start the server
	Serverport := ":8082"
	log.Printf("Video Calling Service is running on port %s\n", Serverport)
	log.Fatal(http.ListenAndServe(Serverport, r))
}
