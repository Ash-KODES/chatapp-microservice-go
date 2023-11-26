// handler/handler.go

package handler

import (
	"chat-app-microservice/calling-service/model"
	"chat-app-microservice/calling-service/db"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// Create a call (video or voice)
func CreateCallHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var call model.Call
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&call); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Validate user information, call details, and call type
		// Ensure that user authentication and authorization checks are performed

		// Check if a 1-1 channel exists between the users
		// You may use the channel service API to verify this
		// For example, check if a channel with the specified users exists

		// Create a unique call ID for the call
		// You can use a UUID or any other method to ensure uniqueness
		call.CallID = generateUniqueCallID() // Implement this function to generate a unique call ID

		// Associate the call ID with the channel
		call.ChannelID = 1 // Replace with the actual channel ID

		// Notify the other user about the incoming call using WebRTC signaling
		// This can be done through WebSocket or another real-time communication method
		// Send a notification to the other user indicating an incoming call

		// Insert the call details into the database
		// Store the call details, including call ID, user IDs, timestamps, and other relevant information
		if err := db.CreateCall(call); err != nil {
			http.Error(w, "Error creating call", http.StatusInternalServerError)
			return
		}

		// Return a response, indicating success and providing the call ID
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct {
			CallID string `json:"call_id"`
		}{CallID: call.CallID})
	}
}

// Join a call (video or voice)
func JoinCallHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var joinRequest model.JoinRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&joinRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Parse and validate the request payload
		// Extract user information, call ID, and channel ID
		userID := joinRequest.UserID
		callID := joinRequest.CallID
		channelID := joinRequest.ChannelID

		// Identify the existing channel associated with the call
		// Retrieve the call details from the database
		call, err := db.GetCall(callID)
		if err != nil {
			http.Error(w, "Error retrieving call details", http.StatusInternalServerError)
			return
		}

		// Establish the appropriate connection based on the call type
		// If it's a video call, set up WebRTC connections
		// If it's a voice call, set up audio connections
		// Implement your WebRTC setup logic here

		// Update the call status, e.g., set it to "in-progress"
		call.Status = model.CallStatusInProgress
		if err := db.UpdateCall(call); err != nil {
			http.Error(w, "Error updating call status", http.StatusInternalServerError)
			return
		}

		// Return a response, indicating success or failure
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: "Successfully joined the call"})
	}
}

// Leave a call (video or voice)
func LeaveCallHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var leaveRequest model.LeaveRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&leaveRequest); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// Parse and validate the request payload
		// Extract user information, call ID, and channel ID
		userID := leaveRequest.UserID
		callID := leaveRequest.CallID
		channelID := leaveRequest.ChannelID

		// Close the appropriate connection based on the call type
		// If it's a video call, close WebRTC connections
		// If it's a voice call, close audio connections
		// Implement your WebRTC close logic here

		// Update the call status, e.g., set it to "ended"
		call, err := db.GetCall(callID)
		if err != nil {
			http.Error(w, "Error retrieving call details", http.StatusInternalServerError)
			return
		}

		call.Status = model.CallStatusEnded
		if err := db.UpdateCall(call); err != nil {
			http.Error(w, "Error updating call status", http.StatusInternalServerError)
			return
		}

		// If both users leave the call, close the call and update the call status
		if err := db.CloseCallIfEmpty(channelID); err != nil {
			http.Error(w, "Error closing call", http.StatusInternalServerError)
			return
		}

		// Return a response, indicating success or failure
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Message string `json:"message"`
		}{Message: "Successfully left the call"})
	}
}

func generateUniqueCallID() string {
	id := uuid.New()
	return id.String()
}
