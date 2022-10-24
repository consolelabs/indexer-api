package nft

type NftListingQuery struct {
	ContractAddress string
	TokenId         string
	Marketplace     string
	ListingStatus   string
	Limit           int
	Offset          int
	Status          string
	Sort            string
	From            int64
	To              int64
}

type NftTokenQuery struct {
	Limit             int
	Offset            int
	Sort              string
	CollectionAddress *string
	ChainId           *int64
	IsSelfHostImage   *bool
	TokenId           *string
	Name              *string
	Currency          *string
	ListingType       *string
	ListingStatus     *string
	PriceMin          *float64
	PriceMax          *float64
	RarityRankMin     *uint64
	RarityRankMax     *uint64
	Marketplaces      []string
	Traits            []string
	TraitsCount       []uint64
}

type WalletTokenQuery struct {
	Limit               int
	Offset              int
	Sort                string
	WalletAddress       string
	ChainId             *int64
	CollectionAddresses []string
}

type NftCollectionQuery struct {
	Name        *string
	Address     *string
	Marketplace *string
	Synced      *bool
	Limit       int
	Offset      int
	Sort        string
}

type NftTickerQuery struct {
	CollectionAddress string
	From              int64
	To                int64
}

type NftAttributeIconQuery struct {
	TraitTypes []string
	Limit      int
	Offset     int
}

type NftOwnerQuery struct {
	CollectionAddress *string
	TokenId           *string
	ChainId           *int64
	Limit             int
	Offset            int
}

type NftTransferQuery struct {
	CollectionAddress *string
	TokenId           *string
	Sort              string
	Limit             int
	Offset            int
}

type WalletCollectionQuery struct {
	Limit         int
	Offset        int
	Sort          string
	WalletAddress string
}
