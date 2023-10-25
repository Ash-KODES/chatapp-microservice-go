// websocket_handler.go

package handler

import (
    "github.com/gorilla/websocket"
    "net/http"
)

// WebSocketHandler manages WebSocket connections.
type WebSocketHandler struct {
    // Store individual WebSocket connections
    conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// NewWebSocketHandler initializes a new WebSocketHandler.
func NewWebSocketHandler() *WebSocketHandler {
    return &WebSocketHandler{}
}

// HandleWebSocket upgrades a WebSocket connection and manages communication.
func (h *WebSocketHandler) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        // Handle upgrade error
        return
    }
    defer conn.Close()

    // Assign the connection to the handler
    h.conn = conn

    // Handle WebSocket messages
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            // Handle read error or client disconnect
            break
        }
        // Process the message, e.g., send it to the chat partner
        h.handleIncomingMessage(messageType, p)
    }
}

// Send a message to the connected chat partner.
func (h *WebSocketHandler) SendToChatPartner(messageType int, payload []byte) {
    if h.conn != nil {
        if err := h.conn.WriteMessage(messageType, payload); err != nil {
            // Handle write error
        }
    }
}

// Handle incoming messages, e.g., send them to the chat partner.
func (h *WebSocketHandler) handleIncomingMessage(messageType int, payload []byte) {
    // Handle incoming messages, e.g., send them to the chat partner
    // You need to implement the logic for sending messages to the chat partner.
}