package miner

import (
	"GoMsgMiner/internal/app/port"
	"fmt"
	"log"
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

func (m minerService) FetchAndStoreMessages(platformName string, channelId string) error {
	for _, adapter := range m.adapters {
		if adapter.GetPlatformName() == platformName {

		}
		messages, err := adapter.FetchHistoricalMessages(channelId)
		if err != nil {
			return err
		}

		platformName := adapter.GetPlatformName()

		// Print messages for illustration.
		for _, msg := range messages {
			log.Printf("%v | %v | %v", platformName, channelId, msg)
		}

		// TODO: Add your logic of storing messages.
	}
	return nil
}

func (m minerService) StreamLiveMessages(platformName string, channelId string) {
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
	//TODO implement me
	panic("implement me")
}
