package main

import (
	"GoMsgMiner/internal/adapter/outbound"
	"GoMsgMiner/internal/app/miner"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Initialize the TwitchChatAdapter
	twitchAdapter := outbound.NewTwitchChatAdapter()

	// Initialize the MockChatAdapter
	// mockAdapter := outbound.NewMockChatAdapter()

	// Initialize the miner service with the Adapters
	minerService := miner.NewMinerService(twitchAdapter)

	minerService.StreamLiveMessages("Twitch", "vedal987")
	minerService.StreamLiveMessages("Twitch", "PirateSoftware")
	minerService.StreamLiveMessages("Twitch", "ThePrimeagen")

	//minerService.StreamLiveMessages("Mock Chat", "Mock Channel 1")

	// Wait for a SIGINT or SIGTERM signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
}
