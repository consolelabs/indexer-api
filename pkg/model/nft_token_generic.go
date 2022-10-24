package model

import "strconv"

// NftTokenGeneric is an generic type to make sure we can parse all collection
// goal to convert to NftToken
// please don't use this type directly in API/parser
type NftTokenGeneric struct {
	TokenId           uint64 `json:"token_id"`
	CollectionAddress string `json:"collection_address,omitempty"`
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Amount            uint64 `json:"amount,omitempty"`
	Image             string `json:"image,omitempty"`
	ImageCDN          string `json:"image_cdn,omitempty"`
	ThumbnailCDN      string `json:"thumbnail_cdn,omitempty"`
	ImageContentType  string `json:"image_content_type,omitempty"`
	RarityRank        uint64 `json:"rarity_rank,omitempty"`
	RarityScore       string `json:"rarity_score,omitempty"`
	RarityTier        string `json:"rarity_tier"`

	Attributes []NftTokenAttributeGeneric `json:"attributes" gorm:"-"`
	// Rarity     *NftTokenRarityGeneric     `json:"rarity" gorm:"-"`
}

// NftTokenAttributeGeneric is an generic type to make sure we can parse all collection
// goal to convert to NftTokenAttribute
// please don't use this type directly in API/parser
type NftTokenAttributeGeneric struct {
	CollectionAddress string      `json:"collection_address"`
	TokenId           uint64      `json:"token_id"`
	TraitType         string      `json:"trait_type"`
	Value             interface{} `json:"value"`
	Count             interface{} `json:"count"`
	Rarity            string      `json:"rarity"`
	Frequency         interface{} `json:"frequency"`
	ChainId           int64       `json:"chain_id"`
}

func (n NftTokenAttributeGeneric) Convert() NftTokenAttribute {
	nft := NftTokenAttribute{
		CollectionAddress: n.CollectionAddress,
		TokenId:           strconv.FormatUint(n.TokenId, 10),
		TraitType:         n.TraitType,
		Rarity:            n.Rarity,
		ChainId:           n.ChainId,
	}
	s, ok := n.Value.(string)
	if ok {
		nft.Value = s
	}
	i, ok := n.Value.(int)
	if ok {
		nft.Value = strconv.Itoa(i)
	}
	f, ok := n.Value.(float64)
	if ok {
		nft.Value = strconv.FormatFloat(f, 'f', -1, 64)
	}

	// frequency
	s, ok = n.Frequency.(string)
	if ok {
		nft.Frequency = s
	}
	i, ok = n.Frequency.(int)
	if ok {
		nft.Frequency = strconv.Itoa(i)
	}
	f, ok = n.Frequency.(float64)
	if ok {
		nft.Frequency = strconv.FormatFloat(f, 'f', -1, 64)
	}

	// count
	s, ok = n.Count.(string)
	if ok {
		i, _ := strconv.Atoi(s)
		nft.Count = uint64(i)
	}
	i, ok = n.Count.(int)
	if ok {
		nft.Count = uint64(i)
	}
	return nft
}

// NftTokenRarityGeneric is an generic type to make sure we can parse all collection
// goal to convert to NftTokenRarity
// please don't use this type directly in API/parser
type NftTokenRarityGeneric struct {
	Rank   uint64 `json:"rank"`
	Score  string `json:"score"`
	Total  uint64 `json:"total"`
	Rarity string `json:"rarity,omitempty"`
}

func (n NftTokenRarityGeneric) Convert() *NftTokenRarity {
	return &NftTokenRarity{
		Rank:   n.Rank,
		Score:  n.Score,
		Total:  n.Total,
		Rarity: n.Rarity,
	}
}

func (n *NftTokenGeneric) Convert() *NftToken {
	nft := &NftToken{
		TokenId:           strconv.FormatUint(n.TokenId, 10),
		CollectionAddress: n.CollectionAddress,
		Name:              n.Name,
		Description:       n.Description,
		Amount:            strconv.FormatUint(n.Amount, 10),
		Image:             n.Image,
		ImageCDN:          n.ImageCDN,
		ThumbnailCDN:      n.ThumbnailCDN,
		ImageContentType:  n.ImageContentType,
		RarityRank:        n.RarityRank,
		RarityScore:       n.RarityScore,
		RarityTier:        n.RarityTier,
	}

	for i := range n.Attributes {
		nft.Attributes = append(nft.Attributes, n.Attributes[i].Convert())
	}
	return nft
}
