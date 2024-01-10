package outbound

import (
	"GoMsgMiner/internal/domain"
	"time"
)

type MockChatAdapter struct{}

func (m *MockChatAdapter) GetPlatformName() string {
	return "Mock Chat"
}

func NewMockChatAdapter() *MockChatAdapter {
	return &MockChatAdapter{}
}

func (m *MockChatAdapter) StreamLiveMessages(channelID string) (<-chan domain.ChatMessage, <-chan error) {

	messageChan := make(chan domain.ChatMessage)

	go func() {
		for {
			messageChan <- domain.ChatMessage{
				Message: "Message",
				Channel: channelID,
			}
			time.Sleep(time.Second)
		}
	}()

	return messageChan, nil
}

func (m *MockChatAdapter) StopStreaming(channelID string) {

}
