package main

import (
	"chat-app-microservice/calling-service/router"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
    const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "12345"
		dbname   = "mychatapp"
	)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
    defer db.Close()

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
    router.SetupRoutes(r, db)

    // Start the server
    Serverport := ":8082" 
    log.Printf("Video Calling Service is running on port %s\n", port)
    log.Fatal(http.ListenAndServe(Serverport, r))
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}