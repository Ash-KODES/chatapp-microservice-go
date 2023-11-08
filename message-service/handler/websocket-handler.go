// websocket_handler.go

package handler

import (
	"github.com/gorilla/websocket"
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

// NewWebSocketHandler initializes a new WebSocketHandler instance.
func NewWebSocketHandler() *WebSocketHandler {
	return &WebSocketHandler{}
}

// HandleWebSocket upgrades a WebSocket connection and manages communication.
func (h *WebSocketHandler) HandleWebSocket(conn *websocket.Conn) {
	h.conn = conn
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
	if h.conn != nil {
		// Send the incoming message to the chat partner
		h.SendToChatPartner(messageType, payload)
	}
}
