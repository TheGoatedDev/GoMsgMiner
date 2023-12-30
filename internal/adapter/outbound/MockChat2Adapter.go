package outbound

import (
	"GoMsgMiner/internal/app/port"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type MockChat2Adapter struct{}

func (m *MockChat2Adapter) GetPlatformName() string {
	return "Mock Chat Two"
}

func NewMockChat2Adapter() *MockChat2Adapter {
	return &MockChat2Adapter{}
}

func (m *MockChat2Adapter) FetchHistoricalMessages(channelID string) ([]port.ChatMessage, error) {

	msg := port.ChatMessage{
		UserID:    strconv.Itoa(rand.Intn(100)),
		UserName:  "mock_user_2",
		Message:   fmt.Sprintf("This is a mock message from channel: %s", channelID),
		Timestamp: time.Now().Unix(),
	}

	// Return a slice containing a single mock message
	return []port.ChatMessage{msg}, nil
}

func (m *MockChat2Adapter) StreamLiveMessages(channelID string) (<-chan port.ChatMessage, <-chan error) {
	// For simplicity, we are not streaming live messages in the mock adapter
	return nil, nil
}

func (m *MockChat2Adapter) StopStreaming(channelID string) {

}
