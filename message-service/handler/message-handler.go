// message_http_handler.go

package handler

import (
    "encoding/json"
    "net/http"
    "database/sql"
    "mychatapp/models"
    "mychatapp/repository"
)

type MessageHandler struct {
    repo *repository.MessageRepository
}

func NewMessageHandler(db *sql.DB) *MessageHandler {
    return &MessageHandler{repository.NewMessageRepository(db)}
}

// message_http_handler.go

func (h *MessageHandler) CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
    var message models.Message
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&message); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // You should validate and authenticate the user, obtain the channel ID, and perform necessary checks.

    // Send the message to the WebSocket channel for real-time delivery
    // For example, if you have a WebSocket connection for the channel, you can send the message to that connection.

    // Insert the message into the database with the associated channel ID
    // You should extract the channel ID from the message or the user's session.

    messageID, err := h.repo.CreateMessage(message)
    if err != nil {
        http.Error(w, "Error creating message", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(struct {
        MessageID int `json:"message_id"`
    }{MessageID: messageID})
}

// message_http_handler.go

func (h *MessageHandler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the channel ID from the request, validate user access, and obtain messages for that channel from the database.

    // Retrieve messages for the specified channel
    // You should adjust the logic to match your data schema and security requirements.

    // Return the messages as a JSON response
}
