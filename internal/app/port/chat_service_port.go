package port

import "GoMsgMiner/internal/domain"

type LiveChatServicePort interface {
	GetPlatformName() string
	StreamLiveMessages(channelID string) (<-chan domain.ChatMessage, <-chan error)
	StopStreaming(channelID string)
}
