// main.go
// Entry point for the video calling service

package main

import (
    "database/sql"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "chat-app-microservice/calling-service/router"
)

func main() {
    // Replace with your database connection string
    db, err := sql.Open("postgres", "user=username dbname=mychatapp sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create a new router
    router := mux.NewRouter()

    // Set up CORS middleware
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"*"}, // Replace with your allowed origins
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    })

    // Use CORS middleware
    router.Use(corsHandler.Handler)

    // Set up video calling routes
    router.SetupRoutes(router, db)

    // Start the server
    port := ":8082" // You can use a different port if needed
    log.Printf("Video Calling Service is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(port, routerSetup))
}
