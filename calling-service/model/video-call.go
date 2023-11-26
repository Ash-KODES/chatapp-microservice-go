// model/video_call.go
// Define the video call data model

package model

type CallType string

const (
	VideoCall CallType = "video"
	VoiceCall CallType = "voice"
)

type CallStatus string

const (
	CallStatusPending   CallStatus = "pending"
	CallStatusInProgress CallStatus = "in-progress"
	CallStatusEnded     CallStatus = "ended"
)

type Call struct {
	ID        int       `json:"id"`
	CallID    string    `json:"call_id"`
	ChannelID int       `json:"channel_id"`
	Type      CallType  `json:"call_type"`
	Status    CallStatus `json:"status"`
	// Add more fields as needed, e.g., timestamps, user IDs, etc.
}

type JoinRequest struct {
	UserID    int    `json:"user_id"`
	CallID    string `json:"call_id"`
	ChannelID int    `json:"channel_id"`
}

type LeaveRequest struct {
	UserID    int    `json:"user_id"`
	CallID    string `json:"call_id"`
	ChannelID int    `json:"channel_id"`
}