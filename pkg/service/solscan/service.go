package solscan

import "github.com/consolelabs/indexer-api/pkg/model"

type IService interface {
	GetSolanaCollection(address string) (*model.SolanaCollectionMetadata, error)
}
