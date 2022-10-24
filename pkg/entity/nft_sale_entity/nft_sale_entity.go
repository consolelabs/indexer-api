package nftsaleentity

import (
	"github.com/consolelabs/indexer-api/pkg/model"
	"github.com/consolelabs/indexer-api/pkg/service"
	"github.com/consolelabs/indexer-api/pkg/store"
)

type nftSaleEntity struct {
	store   *store.Store
	service *service.Service
}

func New(store *store.Store, service *service.Service) INftSaleEntity {
	return &nftSaleEntity{
		store:   store,
		service: service,
	}
}

func (m nftSaleEntity) SalesWebhook(nftListing *model.NftListing) error {
	// collection, err := m.store.Nft.GetCollectionByAddress(nftListing.ContractAddress)
	// if err != nil {
	// 	return err
	// }

	// token, err := m.store.Nft.GetTokenByCollectionAddressTokenId(nftListing.ContractAddress, nftListing.TokenId)
	// if err != nil {
	// 	return err
	// }

	// req := model.NftSales{
	// 	NftSales: []model.NftSale{{
	// 		CollectionAddress: nftListing.ContractAddress,
	// 		CollectionName:    collection.Name,
	// 		CollectionImage:   collection.Image,
	// 		TokenName:         token.Name,
	// 		TokenImage:        token.Image,
	// 		Rarity:            token.RarityTier,
	// 		Rank:              int64(token.RarityRank),
	// 		Marketplace:       "a",
	// 		Transaction:       nftListing.TransactionHash,
	// 		From:              nftListing.FromAddress,
	// 		To:                nftListing.ToAddress,
	// 		Price:             nftListing.ListingPrice,
	// 		Bought:            "0.012",
	// 		Sold:              nftListing.SoldPrice,
	// 		Hodl:              "44 days",
	// 		Gain:              "0.012",
	// 		Pnl:               "2.19",
	// 		SubPnl:            "72.66",
	// 		PaymentToken:      nftListing.PaymentToken,
	// 	}},
	// }
	// err = m.service.Mochi.SendMochiSalesWebhook(&req)
	// if err != nil {
	// 	return err
	// }
	return nil
}
