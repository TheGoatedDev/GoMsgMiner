package miner

import (
	"GoMsgMiner/internal/app/port"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockChatServicePort struct {
	mock.Mock
}

func (m *MockChatServicePort) GetPlatformName() string {
	return "Mock"
}

func (m *MockChatServicePort) FetchHistoricalMessages(channelID string) ([]port.ChatMessage, error) {
	args := m.Called(channelID)
	return args.Get(0).([]port.ChatMessage), args.Error(1)
}

func (m *MockChatServicePort) StreamLiveMessages(channelID string) (<-chan port.ChatMessage, <-chan error) {
	args := m.Called(channelID)
	return args.Get(0).(chan port.ChatMessage), args.Get(1).(chan error)
}

func (m *MockChatServicePort) StopStreaming(channelID string) {
	m.Called(channelID)
}

func TestNewMinerService(t *testing.T) {
	mockChatServicePort := new(MockChatServicePort)
	newMinerService := NewMinerService(mockChatServicePort)
	assert.NotNil(t, newMinerService)
}

func TestMinerService_FetchAndStoreMessages(t *testing.T) {
	mockChatServicePort := new(MockChatServicePort)
	mockChatServicePort.On("FetchHistoricalMessages", "testChannel").Return([]port.ChatMessage{}, nil)
	service := NewMinerService(mockChatServicePort)

	err := service.FetchAndStoreMessages("Mock", "testChannel") // Correct Call
	assert.Nil(t, err)
}

func TestMinerService_StreamLiveMessages(t *testing.T) {
	mockChatServicePort := new(MockChatServicePort)
	mockChatServicePort.On("StreamLiveMessages", "testChannel").Return(make(chan port.ChatMessage), make(chan error))
	service := NewMinerService(mockChatServicePort)

	service.StreamLiveMessages("Mock", "testChannel")
}

func TestMinerService_StopStreaming(t *testing.T) {
	mockChatServicePort := new(MockChatServicePort)
	mockChatServicePort.On("StopStreaming", "testChannel").Return()
	service := NewMinerService(mockChatServicePort)

	// calling StopStreaming
	service.StopStreaming("Mock", "testChannel")
}
