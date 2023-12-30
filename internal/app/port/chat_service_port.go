package port

type ChatMessage struct {
	UserID    string
	UserName  string
	Message   string
	Timestamp int64 // Unix timestamp
}

type LiveChatServicePort interface {
	GetPlatformName() string
	FetchHistoricalMessages(channelID string) ([]ChatMessage, error)
	StreamLiveMessages(channelID string) (<-chan ChatMessage, <-chan error)
	StopStreaming(channelID string)
}
