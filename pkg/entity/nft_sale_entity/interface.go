package nftsaleentity

import "github.com/consolelabs/indexer-api/pkg/model"

type INftSaleEntity interface {
	SalesWebhook(nftListing *model.NftListing) error
}
