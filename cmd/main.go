package main

import (
	"GoMsgMiner/internal/adapter/outbound"
	"GoMsgMiner/internal/app/miner"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Initialize the TwitchChatAdapter
	twitchAdapter := outbound.NewTwitchChatAdapter()

	// Initialize the miner service with the TwitchChatAdapter
	minerService := miner.NewMinerService(twitchAdapter)

	minerService.StreamLiveMessages("Twitch", "Sweet_Anita")

	minerService.StreamLiveMessages("Twitch", "vedal987")

	minerService.StreamLiveMessages("Twitch", "Emiru")

	minerService.StreamLiveMessages("Twitch", "PirateSoftware")

	<-stop
}
