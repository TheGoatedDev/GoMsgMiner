package miner

type Miner interface {
	FetchAndStoreMessages(platformName string, channelId string) error
	StreamLiveMessages(platformName string, channelId string)
	StopStreaming(platformName string, channelId string)
}
