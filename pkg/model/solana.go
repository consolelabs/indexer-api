package model

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

type SolanaCollectionMetadata struct {
	Success bool        `json:"success"`
	Data    SolscanData `json:"data"`
}

type SolscanData struct {
	Success bool                 `json:"success"`
	Data    SolanaCollectionInfo `json:"data"`
}

type SolanaCollectionInfo struct {
	Avatar       string `json:"avatar"`
	Collection   string `json:"collection"`
	Symbol       string `json:"symbol"`
	CollectionId string `json:"collectionId"`
}
