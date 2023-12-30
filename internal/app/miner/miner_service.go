package miner

import (
	"GoMsgMiner/internal/app/port"
	"log"
)

type minerService struct {
	adapters []port.LiveChatServicePort
}

func NewMinerService(adapters ...port.LiveChatServicePort) Miner {
	return &minerService{
		adapters: adapters,
	}
}

func (m minerService) FetchAndStoreMessages(channelId string) error {
	for _, adapter := range m.adapters {
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

func (m minerService) StreamLiveMessages(channelId string) {
	//TODO implement me
	panic("implement me")
}

func (m minerService) StopStreaming(channelId string) {
	//TODO implement me
	panic("implement me")
}
