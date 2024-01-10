package outbound

import (
	"GoMsgMiner/internal/domain"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
	"time"
)

type TwitchChatAdapter struct {
	twitchClient *twitch.Client
}

func (m *TwitchChatAdapter) GetPlatformName() string {
	return "Twitch"
}

func NewTwitchChatAdapter() *TwitchChatAdapter {

	client := twitch.NewAnonymousClient()

	client.OnConnect(func() {
		fmt.Println("Connected to Twitch")
	})

	go func() {
		err := client.Connect()
		if err != nil {
			panic("Error Connection to Twitch")
		}
	}()

	return &TwitchChatAdapter{
		twitchClient: client,
	}
}

func (m *TwitchChatAdapter) StreamLiveMessages(channelID string) (<-chan domain.ChatMessage, <-chan error) {
	// Use
	messageChan := make(chan domain.ChatMessage)
	errorChan := make(chan error)

	m.twitchClient.OnPrivateMessage(func(message twitch.PrivateMessage) {
		chatMessage := domain.ChatMessage{
			Channel: message.Channel,

			UserName: message.User.DisplayName,
			UserID:   message.User.ID,

			Message:  message.Message,
			Platform: m.GetPlatformName(),

			Timestamp: time.Now().Unix(),
		}

		// Send the ChatMessage struct through the message channel
		messageChan <- chatMessage
	})

	m.twitchClient.Join(channelID)

	return messageChan, errorChan
}

func (m *TwitchChatAdapter) StopStreaming(channelID string) {
	m.twitchClient.Depart(channelID)
}
