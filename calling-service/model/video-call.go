// model/video_call.go
// Define the video call data model

package model

type VideoCall struct {
    ID      int    `json:"id"`
    CallID  string `json:"call_id"`
    ChannelID int  `json:"channel_id"`
    // Add more fields as needed, e.g., timestamps, user IDs, etc.
}
