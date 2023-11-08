// message_http_handler.go

package handler

import (
	models "chat-app-microservice/message-service/model"
	"chat-app-microservice/message-service/repository"
	"database/sql"
	"encoding/json"
	"net/http"
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
	h.sendToWebSocketChannel(message)

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

// sendToWebSocketChannel sends a message to the WebSocket channel
func (h *MessageHandler) sendToWebSocketChannel(message models.Message) {
	// Assuming you have a WebSocket connection associated with the channel
	// You may need to implement the logic for sending the message to the WebSocket connection
	// Here, we use the WebSocket handler to send the message to the WebSocket connection
	if h.ws != nil {
		messagePayload, err := json.Marshal(message)
		if err != nil {
			// Handle JSON encoding error
			return
		}
		h.ws.SendToChatPartner(websocket.TextMessage, messagePayload)
	}
}

// GetMessagesHandler handles retrieving messages from a channel
func (h *MessageHandler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the channel ID from the request (e.g., from URL parameters)
	vars := mux.Vars(r)
	channelID := vars["channelID"]

	// Validate user access to the channel and perform necessary checks
	if !h.isUserAuthorized(r, channelID) {
		http.Error(w, "Access denied", http.StatusForbidden)
		return
	}

	// Retrieve messages for the specified channel from the database
	messages, err := h.repo.GetMessagesForChannel(channelID)
	if err != nil {
		http.Error(w, "Error retrieving messages", http.StatusInternalServerError)
		return
	}

	// Return the messages as a JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(messages); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
	}
}

// isUserAuthorized checks if the user is authorized to access the channel
func (h *MessageHandler) isUserAuthorized(r *http.Request, channelID string) bool {
	// You should implement the logic to check if the user is authorized to access the channel.
	// This may involve checking if the user is a participant in the channel or has the required permissions.

	// For example, if your application uses authentication and sessions, you can check if the user is authenticated.
	if !userIsAuthenticated(r) {
		return false
	}

	// You may also need to check if the user is a participant in the channel.
	participants, err := h.repo.GetParticipantsForChannel(channelID)
	if err != nil {
		return false
	}

	user := getCurrentUserFromRequest(r)

	for _, participant := range participants {
		if participant == user {
			return true
		}
	}

	return false
}

// userIsAuthenticated checks if the user is authenticated
func userIsAuthenticated(r *http.Request) bool {
	// Implement your logic to check if the user is authenticated.
	// You can access user information from the request or session.
	// Replace this with your actual authentication logic.

	// Example: Check if a user is logged in by checking for a session or token.
	// user := getCurrentUserFromRequest(r)
	// return user != nil

	return true // Replace with your actual authentication logic
}

// getCurrentUserFromRequest retrieves user information from the request
func getCurrentUserFromRequest(r *http.Request) string {
	// Implement your logic to extract user information from the request.
	// This could involve checking session data, tokens, or other authentication mechanisms.
	// Replace this with your actual logic.

	// Example: Extract user information from a token or session.
	// userToken := extractUserTokenFromRequest(r)
	// user, err := validateUserToken(userToken)
	// if err != nil {
	//     return ""
	// }
	// return user

	return "sample_user" // Replace with your actual user retrieval logic
}
