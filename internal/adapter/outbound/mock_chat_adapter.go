package outbound

import (
	"GoMsgMiner/internal/domain"
	"fmt"
	"strconv"
	"time"
)

type MockChatAdapter struct {
}

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
				Message:   strconv.FormatInt(time.Now().Unix(), 10),
				Channel:   channelID,
				Platform:  m.GetPlatformName(),
				Timestamp: time.Now().Unix(),
				UserID:    fmt.Sprintf("UserID-%d", time.Now().Unix()),
				UserName:  fmt.Sprintf("UserName-%d", time.Now().Unix()),
			}
			time.Sleep(time.Second)
		}
	}()

	return messageChan, nil
}

func (m *MockChatAdapter) StopStreaming(channelID string) {

}
