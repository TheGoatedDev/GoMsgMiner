package miner

type Miner interface {
	StreamLiveMessages(platformName string, channelId string)
	StopStreaming(platformName string, channelId string)
}
