// model/video_call.go
// Define the video call data model

package model

type CallType string

const (
    VideoCall CallType = "video"
    VoiceCall CallType = "voice"
)

type Call struct {
    ID      int    `json:"id"`
    CallID  string `json:"call_id"`
    ChannelID int  `json:"channel_id"`
    Type    CallType `json:"call_type"`
    // Add more fields as needed, e.g., timestamps, user IDs, etc.
}
