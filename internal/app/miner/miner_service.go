package miner

import (
	"GoMsgMiner/internal/app/port"
	"fmt"
)

type minerService struct {
	adapters map[string]port.LiveChatServicePort
}

func NewMinerService(adapters ...port.LiveChatServicePort) Miner {

	adapterMap := make(map[string]port.LiveChatServicePort)

	for _, adapter := range adapters {
		adapterMap[adapter.GetPlatformName()] = adapter
	}

	return &minerService{
		adapters: adapterMap,
	}
}

func (m minerService) StreamLiveMessages(platformName string, channelId string) {
	fmt.Printf("Started Streaming on %s for %s", platformName, channelId)

	messageChan, errorChan := m.adapters[platformName].StreamLiveMessages(channelId)

	go func() {
		for msg := range messageChan {
			fmt.Printf("%s | %s | %s: %s\n", msg.Platform, msg.Channel, msg.UserName, msg.Message)
		}
	}()

	go func() {
		for err := range errorChan {
			fmt.Printf("Error occurred: %s\n", err)
		}
	}()
}

func (m minerService) StopStreaming(platformName string, channelId string) {
	m.adapters[platformName].StopStreaming(channelId)
}
