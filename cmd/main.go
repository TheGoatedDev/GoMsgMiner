package main

import (
	"GoMsgMiner/internal/adapter/outbound"
	"GoMsgMiner/internal/app/miner"
	"log"
)

func main() {
	// Initialize the MockChatAdapter
	mockAdapter := outbound.NewMockChatAdapter()
	mock2Adapter := outbound.NewMockChat2Adapter()

	// Initialize the miner service with the MockChatAdapter
	minerService := miner.NewMinerService(mockAdapter, mock2Adapter)

	// Fetch and print historical messages
	err := minerService.FetchAndStoreMessages("test_channel")
	if err != nil {
		log.Fatalf("could not fetch and store messages: %v", err)
	}

}
