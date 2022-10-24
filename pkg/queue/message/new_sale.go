package message

type NewSaleMessage struct {
	MarketplacePlatformId uint64
	CollectionAddress     string
	NftId                 string
	TxHash                string
}
