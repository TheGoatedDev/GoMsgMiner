package domain

type ChatMessage struct {
	UserID    string
	UserName  string
	Platform  string
	Channel   string
	Message   string
	Timestamp int64 // Unix timestamp
}
