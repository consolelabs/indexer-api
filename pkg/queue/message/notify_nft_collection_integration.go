package message

type NotifyNftCollectionIntegrationMessage struct {
	GuildId   string `json:"guild_id"`
	ChannelId string `json:"channel_id"`
	MessageId string `json:"message_id"`
}
