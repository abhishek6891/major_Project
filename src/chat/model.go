package chat

import "time"

type ChatMessage struct {
	SenderID    string    `json:"sender_id"`
	SenderName  string    `json:"sender_name"`
	SenderType  string    `json:"sender_type"`
	ProjectName string    `json:"project_name"`
	Message     string    `json:"message"`
	Time        time.Time `json:"time"`
}

type ChatSession struct {
	ClientID    string        `json:"client_id"`
	DeveloperID string        `json:"developer_id"`
	ProjectName string        `json:"project_name"`
	StartTime   time.Time     `json:"start_time"`
	EndTime     time.Time     `json:"end_time"`
	Duration    string        `json:"duration"`
	Messages    []ChatMessage `json:"messages"`
}
