// message.go

package models

type Message struct {
    ID        int    `json:"id"`
    Content   string `json:"content"`
    UserID    int    `json:"user_id"`
    Timestamp string `json:"timestamp"`

    // Fields for other message types
    ImageURL  string `json:"image_url"`
    AudioURL  string `json:"audio_url"`
    VideoURL  string `json:"video_url"`
    FileURL   string `json:"file_url"`
    // Add more fields as needed
}
