package outbound

import (
	"GoMsgMiner/internal/app/port"
	"fmt"
	"time"
)

type MockChatAdapter struct{}

func (m *MockChatAdapter) GetPlatformName() string {
	return "Mock Chat"
}

func NewMockChatAdapter() *MockChatAdapter {
	return &MockChatAdapter{}
}

func (m *MockChatAdapter) FetchHistoricalMessages(channelID string) ([]port.ChatMessage, error) {
	msg := port.ChatMessage{
		UserID:    "123",
		UserName:  "mock_user",
		Message:   fmt.Sprintf("This is a mock message from channel: %s", channelID),
		Timestamp: time.Now().Unix(),
	}

	// Return a slice containing a single mock message
	return []port.ChatMessage{msg}, nil
}

func (m *MockChatAdapter) StreamLiveMessages(channelID string) (<-chan port.ChatMessage, <-chan error) {

	messageChan := make(chan port.ChatMessage)

	go func() {
		for {
			messageChan <- port.ChatMessage{Message: "Message"}
			time.Sleep(time.Second)
		}
	}()

	return messageChan, nil
}

func (m *MockChatAdapter) StopStreaming(channelID string) {

}
