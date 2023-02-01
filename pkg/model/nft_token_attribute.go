package model

type NftTokenAttribute struct {
	CollectionAddress string `json:"collection_address"`
	TokenId           string `json:"token_id"`
	TraitType         string `json:"trait_type"`
	Value             string `json:"value"`
	Count             uint64 `json:"count"`
	Rarity            string `json:"rarity"`
	Frequency         string `json:"frequency"`
	ChainId           int64  `json:"chain_id"`
}

type NftTokenRarity struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

type NftTokenAttributeCount struct {
	Id    int64   `json:"id"`
	Count int64   `json:"count"`
	Ratio float64 `json:"ratio"`
}

type NftTokenAttrSoulBound struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	Count     uint64 `json:"count"`
}
