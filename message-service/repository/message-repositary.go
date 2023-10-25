package repository

import (
    "database/sql"
    
)

type MessageRepository struct {
    db *sql.DB
}

func NewMessageRepository(db *sql.DB) *MessageRepository {
    return &MessageRepository{db}
}

func (r *MessageRepository) CreateMessage(message models.Message) (int, error) {
    // Implement the logic to insert a message into the database
    // You need to adapt this to your database schema
    // For example, if you have a "messages" table with fields: message_id, channel_id, sender_id, content, created_at, etc.
    insertQuery := `
        INSERT INTO messages (channel_id, sender_id, content)
        VALUES ($1, $2, $3)
        RETURNING message_id
    `

    var messageID int
    err := r.db.QueryRow(insertQuery, message.ChannelID, message.SenderID, message.Content).Scan(&messageID)
    if err != nil {
        return 0, err
    }

    return messageID, nil
}

func (r *MessageRepository) RetrieveMessages(channelID int) ([]models.Message, error) {
    // Implement the logic to retrieve messages for a specific channel from the database
    // You may retrieve messages based on the channel_id and apply any necessary filters or limits
    // For example:
    query := "SELECT message_id, sender_id, content, created_at FROM messages WHERE channel_id = $1"
    rows, err := r.db.Query(query, channelID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var messages []models.Message
    for rows.Next() {
        var message models.Message
        err := rows.Scan(&message.MessageID, &message.SenderID, &message.Content, &message.CreatedAt)
        if err != nil {
            return nil, err
        }
        messages = append(messages, message)
    }

    return messages, nil
}

// Implement methods for message history, searching for messages, and other message-related actions
