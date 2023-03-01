package model

import "time"

type Element struct {
	Data                Data            `json:"data"`
	IsMutable           bool            `json:"is_mutable"`
	Mint                string          `json:"mint"`
	PrimarySaleHappened bool            `json:"primary_sale_happened"`
	UpdateAuthority     string          `json:"update_authority"`
	MetadataAccount     string          `json:"metadata_account"`
	EditionNounce       int64           `json:"edition_nounce"`
	TokenStandard       int64           `json:"token_standard"`
	Collection          Collection1     `json:"collection"`
	Uses                int64           `json:"uses"`
	MetadataFromUri     MetadataFromUri `json:"MetadataFromUri"`
}

// Data
type Data struct {
	Creators             []string `json:"creators"`
	Name                 string   `json:"name"`
	SellerFeeBasisPoints int64    `json:"seller_fee_basis_points"`
	Share                []int64  `json:"share"`
	Symbol               string   `json:"symbol"`
	Uri                  string   `json:"uri"`
	Verified             []int64  `json:"verified"`
}

// Collection
type Collection1 struct {
	Verified int64  `json:"verified"`
	Key      string `json:"key"`
}

// MetadataFromUri
type MetadataFromUri struct {
	Name                 string      `json:"name"`
	Symbol               string      `json:"symbol"`
	Image                string      `json:"image"`
	Description          string      `json:"description"`
	ExternalUrl          string      `json:"external_url"`
	SellerFeeBasisPoints int64       `json:"seller_fee_basis_points"`
	Collection           Collection  `json:"collection"`
	Properties           Properties  `json:"properties"`
	Attributes           []Attribute `json:"attributes"`
}
type Collection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}
type Properties struct {
	Files    []File    `json:"files"`
	Creators []Creator `json:"creators"`
	Category string    `json:"category"`
}
type File struct {
	Uri  string `json:"uri"`
	Type string `json:"type"`
}

type Creator struct {
	Address string `json:"address"`
	Share   int64  `json:"share"`
}

type Attribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

type NftCollectionResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Success bool                 `json:"success"`
		Data    NftCollectionSolScan `json:"data"`
	} `json:"data"`
}

type NftCollectionSolScan struct {
	ID                  string  `json:"_id"`
	CollectionID        string  `json:"collectionId"`
	Collection          string  `json:"collection"`
	CollectionOnchainID *string `json:"collectionOnchainId"`
}

type SolanaMappingAddress struct {
	ID             string    `json:"id"`
	OnchainAddress string    `json:"onchain_address"`
	SolscanId      string    `json:"solscan_id"`
	Slug           string    `json:"slug"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type SolscanCollection struct {
	Data []struct {
		CollectionID   string `json:"collection_id"`
		CollectionName string `json:"collection_name"`
	} `json:"data"`
}

type NftTokenSolScanResponse struct {
	Data struct {
		ListNfts []struct {
			NftAddress string `json:"nft_address"`
		} `json:"list_nfts"`
	} `json:"data"`
}

type NftTokenMagicedenReponse struct {
	MintAddress string `json:"mintAddress"`
	Collection  string `json:"collection"`
	Name        string `json:"name"`
}
