package message

type KafkaMessage struct {
	// EnrichNftToken           *EnrichNftTokenMessage
	// NewSale                  *NewSaleMessage
	// EnrichNftCollection      *EnrichNftCollectionMessage
	// EnrichNftRarityCalculate *EnrichNftRarityCalculateMessage
	// UploadingImage           *ImageUploadingMessage
	Topic   string `json:"topic"`
	Address string `json:"address"`
	ChainId int64  `json:"chain_id"`
}
