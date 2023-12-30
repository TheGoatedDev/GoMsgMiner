package miner

type Miner interface {
	FetchAndStoreMessages(channelId string) error
	StreamLiveMessages(channelId string)
	StopStreaming(channelId string)
}
