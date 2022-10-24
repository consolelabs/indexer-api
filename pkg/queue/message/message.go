package message

type KafkaMessage struct {
	EnrichNftToken           *EnrichNftTokenMessage
	NewSale                  *NewSaleMessage
	EnrichNftCollection      *EnrichNftCollectionMessage
	EnrichNftRarityCalculate *EnrichNftRarityCalculateMessage
	UploadingImage           *ImageUploadingMessage
}
