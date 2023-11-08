package router

import (
	"chat-app-microservice/message-service/handler"
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router, db *sql.DB) *mux.Router {
    messageHandler := handler.NewMessageHandler(db)
    webSocketHandler := handler.NewWebSocketHandler()

    r.HandleFunc("/messages", messageHandler.CreateMessageHandler).Methods("POST")
    r.HandleFunc("/ws", webSocketHandler.HandleWebSocket)

    // Define routes for retrieving messages, message history, and searching for messages

    return r
}
