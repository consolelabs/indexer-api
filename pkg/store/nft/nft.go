package nft

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/consolelabs/indexer-api/pkg/model"
)

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB) INft {
	return &store{
		db: db,
	}
}

func (s *store) UpsertToken(token *model.NftToken, upsertAttributes bool) error {
	tx := s.db.Begin()
	if err := tx.Table("nft_token").Where("collection_address = ? AND token_id = ?", token.CollectionAddress, token.TokenId).Save(&token).Error; err != nil {
		tx.Rollback()
		return err
	}

	if !upsertAttributes {
		return tx.Commit().Error
	}

	for _, attr := range token.Attributes {
		attr.TokenId = token.TokenId
		attr.CollectionAddress = token.CollectionAddress
		err := tx.Table("nft_token_attribute").Where("collection_address = ? AND token_id = ? AND trait_type = ?", attr.CollectionAddress, attr.TokenId, attr.TraitType).Save(&attr).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (s *store) GetCollectionByAddress(address string) (*model.NftCollection, error) {
	collection := model.NftCollection{}
	row := s.db.Table("view_nft_collections").Select(
		"id",
		"address",
		"name",
		"symbol",
		"chain_id",
		"erc_format",
		"supply",
		"is_rarity_calculated",
		"image",
		"created_time",
		"last_updated_time",
		"description",
		"contract_scan",
		"discord",
		"twitter",
		"website",
		"owners",
		"chain",
		"marketplace_collections",
	).Where("address = ?", address).Row()
	var chainRaw, marketplaceRaw []byte
	err := row.Scan(
		&collection.ID,
		&collection.Address,
		&collection.Name,
		&collection.Symbol,
		&collection.ChainId,
		&collection.ERCFormat,
		&collection.Supply,
		&collection.IsRarityCalculated,
		&collection.Image,
		&collection.CreatedTime,
		&collection.LastUpdatedTime,
		&collection.Description,
		&collection.ContractScan,
		&collection.Discord,
		&collection.Twitter,
		&collection.Website,
		&collection.Owners,
		&chainRaw,
		&marketplaceRaw,
	)
	if err != nil {
		return nil, err
	}
	if chainRaw != nil {
		if err := json.Unmarshal(chainRaw, &collection.Chain); err != nil {
			return nil, err
		}
	}
	if marketplaceRaw != nil {
		if err := json.Unmarshal(marketplaceRaw, &collection.MarketplaceCollections); err != nil {
			return nil, err
		}
	}
	return &collection, nil
}

func (s *store) SaveOwner(owner *model.NftOwner) error {
	return s.db.Create(owner).Error
}

func (s *store) DeleteOwnerByCollectionAddressTokenId(collectionAddress, tokenId string) error {
	return s.db.Table("nft_owner").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenId).Delete(&model.NftOwner{}).Error
}

func (s *store) GetAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error) {
	var tokenAttributes []model.NftTokenAttribute
	err := s.db.Table("nft_token_attribute").Where("collection_address = ?", collectionAddress).Find(&tokenAttributes).Error
	if err != nil {
		return nil, err
	}
	return tokenAttributes, nil
}

func (s *store) UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId string, ownerAddress string) error {
	return s.db.Table("nft_owner").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenId).Updates(map[string]interface{}{"owner_address": ownerAddress, "last_updated_time": time.Now()}).Error
}

func (s *store) SaveTransfer(transfer *model.NftTransfer) error {
	return s.db.Create(transfer).Error
}

func (s *store) GetCollections(q NftCollectionQuery) ([]model.NftCollection, int64, error) {
	collections := make([]model.NftCollection, 0, q.Limit)
	var total int64
	db := s.db.Table("view_nft_collections AS vnc").Select(
		"vnc.id",
		"vnc.address",
		"vnc.name",
		"vnc.symbol",
		"vnc.chain_id",
		"vnc.erc_format",
		"vnc.supply",
		"vnc.is_rarity_calculated",
		"vnc.image",
		"vnc.created_time",
		"vnc.last_updated_time",
		"vnc.description",
		"vnc.contract_scan",
		"vnc.discord",
		"vnc.twitter",
		"vnc.website",
		"vnc.owners",
		"vnc.chain",
		"vnc.marketplace_collections",
		"stats.floor_price",
		"stats.one_day_volume",
		"stats.one_day_volume_change",
		"stats.seven_day_volume",
		"stats.seven_day_volume_change",
		"stats.thirty_day_volume",
		"stats.thirty_day_volume_change",
		"stats.all_time_volume",
		"stats.all_time_volume_change",
		"token.symbol",
		"token.address",
		"token.is_native",
		"token.decimals",
		"token.icon_url",
	).
		Joins("LEFT JOIN view_nft_collection_stats stats ON vnc.address = stats.address").
		Joins("LEFT JOIN token ON stats.payment_token = token.id").
		Joins("LEFT JOIN contract ON vnc.address = contract.address")

	if q.Synced != nil && *q.Synced {
		db = db.Where("contract.last_updated_block > 0")
	}
	if q.Name != nil {
		db = db.Where("to_tsvector(vnc.name) @@ to_tsquery(?)", *q.Name)
	}
	if q.Address != nil {
		db = db.Where("vnc.address = ?", *q.Address)
	}
	if q.Marketplace != nil {
		db = db.Where("to_tsvector(vnc.marketplace_names) @@ to_tsquery(?)", *q.Marketplace)
	}
	if len(q.Sort) > 0 {
		db = db.Order(q.Sort)
	}

	rows, err := db.Count(&total).Offset(q.Offset).Limit(q.Limit).Rows()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return nil, total, err
	}

	for rows.Next() {
		collection := model.NftCollection{Stats: &model.ViewNftCollectionStats{Token: &model.Token{}}}
		var chainRaw, marketplaceRaw, floorPriceRaw, oneDayVolRaw, sevenDayVolRaw, thirtyDayVolRaw, allTimeVolRaw []byte
		rows.Scan(
			&collection.ID,
			&collection.Address,
			&collection.Name,
			&collection.Symbol,
			&collection.ChainId,
			&collection.ERCFormat,
			&collection.Supply,
			&collection.IsRarityCalculated,
			&collection.Image,
			&collection.CreatedTime,
			&collection.LastUpdatedTime,
			&collection.Description,
			&collection.ContractScan,
			&collection.Discord,
			&collection.Twitter,
			&collection.Website,
			&collection.Owners,
			&chainRaw,
			&marketplaceRaw,
			&floorPriceRaw,
			&oneDayVolRaw,
			&collection.Stats.OneDayVolumeChange,
			&sevenDayVolRaw,
			&collection.Stats.SevenDayVolumeChange,
			&thirtyDayVolRaw,
			&collection.Stats.ThirtyDayVolumeChange,
			&allTimeVolRaw,
			&collection.Stats.AllTimeVolumeChange,
			&collection.Stats.Token.Symbol,
			&collection.Stats.Token.Address,
			&collection.Stats.Token.IsNative,
			&collection.Stats.Token.Decimals,
			&collection.Stats.Token.IconUrl,
		)
		if chainRaw != nil {
			if err := json.Unmarshal(chainRaw, &collection.Chain); err != nil {
				return nil, total, err
			}
		}
		if marketplaceRaw != nil {
			if err := json.Unmarshal(marketplaceRaw, &collection.MarketplaceCollections); err != nil {
				return nil, total, err
			}
		}
		if floorPriceRaw != nil {
			if err := json.Unmarshal(floorPriceRaw, &collection.Stats.FloorPrice); err != nil {
				return nil, total, err
			}
		}
		if oneDayVolRaw != nil {
			if err := json.Unmarshal(oneDayVolRaw, &collection.Stats.OneDayVolume); err != nil {
				return nil, total, err
			}
		}
		if sevenDayVolRaw != nil {
			if err := json.Unmarshal(sevenDayVolRaw, &collection.Stats.SevenDayVolume); err != nil {
				return nil, total, err
			}
		}
		if thirtyDayVolRaw != nil {
			if err := json.Unmarshal(thirtyDayVolRaw, &collection.Stats.ThirtyDayVolume); err != nil {
				return nil, total, err
			}
		}
		if allTimeVolRaw != nil {
			if err := json.Unmarshal(allTimeVolRaw, &collection.Stats.AllTimeVolume); err != nil {
				return nil, total, err
			}
		}
		collections = append(collections, collection)
	}

	return collections, total, nil
}

func (s *store) SaveNftCollection(nftCollection *model.NftCollection) error {
	return s.db.Create(nftCollection).Error
}

func (pg *store) UpdateListing(listing *model.NftListing) error {
	return pg.db.Table("nft_listing").Where("id = ?", listing.Id).Save(&listing).Error
}

func (pg *store) SaveListing(listing *model.NftListing) error {
	return pg.db.Create(listing).Error
}

func (s *store) GetPlatformsByCollectionAddress(collectionAddress string) ([]model.MarketplacePlatform, error) {
	var platformIds []uint64
	err := s.db.Table("nft_marketplace_collection").
		Joins("LEFT JOIN nft_collection c ON nft_marketplace_collection.collection_address = c.address").
		Where("c.address = ?", collectionAddress).Distinct().
		Pluck("nft_marketplace_collection.platform_id", &platformIds).Error
	if err != nil {
		return nil, err
	}

	if len(platformIds) == 0 {
		return []model.MarketplacePlatform{}, nil
	}

	var marketplaces []model.MarketplacePlatform
	return marketplaces, s.db.Table("marketplace_platform").Where("id IN (?)", platformIds).Find(&marketplaces).Error
}

func (s *store) UpdateAttributeCount(count uint64, address, traitType, value string) error {
	return s.db.Table("nft_token_attribute").Where("collection_address = ? AND trait_type = ? AND value = ?", address, traitType, value).Update("count", count).Error
}

func (s *store) GetNftMarketplaceCollectionSnapshots(q NftTickerQuery) ([]model.NftMarketplaceCollectionSnapshot, error) {
	var snapshots []model.NftMarketplaceCollectionSnapshot
	rows, err := s.db.Table("nft_marketplace_collection_snapshot AS snapshot").
		Select("snapshot.created_time", "snapshot.floor_price", "t.symbol", "t.address", "t.decimals, mp.name").
		Joins("LEFT JOIN token t ON snapshot.token_id = t.id").
		Joins("LEFT JOIN marketplace_platform mp ON snapshot.platform_id = mp.id").
		Where("snapshot.collection_address = ?", q.CollectionAddress).
		Where("extract(epoch from snapshot.created_time) * 1000 >= ?", q.From).
		Where("extract(epoch from snapshot.created_time) * 1000 <= ?", q.To).
		Order("snapshot.created_time").Rows()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		snapshot := model.NftMarketplaceCollectionSnapshot{Marketplace: &model.MarketplacePlatform{}}
		token := model.Token{}
		err := rows.Scan(&snapshot.CreatedTime, &snapshot.FloorPrice, &token.Symbol, &token.Address, &token.Decimals, &snapshot.Marketplace.Name)
		if err != nil {
			return nil, err
		}
		snapshot.FloorPriceObj = &model.Price{Amount: snapshot.FloorPrice, Token: token}
		snapshots = append(snapshots, snapshot)
	}
	return snapshots, nil
}

func (s *store) GetNftMetadataAttributesIcon(q NftAttributeIconQuery) ([]model.NftMetadataAttributesIcon, int64, error) {
	var nftMetadataAttributesIcon []model.NftMetadataAttributesIcon
	var total int64
	query := s.db.Table("nft_token_attribute_icon")

	if len(q.TraitTypes) > 0 {
		query = query.Where("lower(trait_type) in (?)", q.TraitTypes)
	}

	if q.Limit > 0 {
		query = query.Limit(q.Limit).Offset(q.Offset)
	}

	return nftMetadataAttributesIcon, total, query.Count(&total).Find(&nftMetadataAttributesIcon).Error
}

func (s *store) RefreshViewNFTCollections() error {
	return s.db.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY view_nft_collections").Error
}

func (s *store) RefreshViewNFTCollectionStats() error {
	return s.db.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY view_nft_collection_stats").Error
}

func (s *store) GetNftTokens(q NftTokenQuery) ([]model.NftToken, int64, error) {
	db := s.db.Table("view_nft_tokens").Select(
		"token_id",
		"collection_address",
		"name",
		"description",
		"amount",
		"image",
		"image_cdn",
		"thumbnail_cdn",
		"image_content_type",
		"rarity_rank",
		"rarity_score",
		"rarity_tier",
		"listing",
		"attributes",
		"owner",
		"COALESCE((listing::json #>> '{listing_price_obj,amount}')::numeric, 0) as price", // price sorting
	)
	res := make([]model.NftToken, 0, q.Limit)
	var total int64

	if q.CollectionAddress != nil {
		db = db.Where("collection_address = ?", *q.CollectionAddress)
	}
	if q.IsSelfHostImage != nil {
		db = db.Where("is_self_hosted = ?", *q.IsSelfHostImage)
	}
	if q.TokenId != nil {
		db = db.Where("token_id = ?", *q.TokenId)
	}
	if q.Name != nil {
		db = db.Where("to_tsvector(name) @@ to_tsquery(?)", *q.Name)
	}
	if q.Currency != nil {
		db = db.Where("lower(listing::json #>> '{listing_price_obj,token,symbol}') = lower(?)", *q.Currency)
	}
	if q.ListingType != nil {
		db = db.Where("lower(listing::json ->> 'listing_type') = lower(?)", *q.ListingType)
	}
	if q.ListingStatus != nil {
		db = db.Where("lower(listing::json ->> 'listing_status') = lower(?)", *q.ListingStatus)
	}
	if q.PriceMin != nil {
		db = db.Where("(listing::json #>> '{listing_price_obj,amount}')::numeric >= ?", *q.PriceMin)
	}
	if q.PriceMax != nil {
		db = db.Where("(listing::json #>> '{listing_price_obj,amount}')::numeric <= ?", *q.PriceMax)
	}
	if q.RarityRankMin != nil {
		db = db.Where("rarity_rank >= ?", q.RarityRankMin)
	}
	if q.RarityRankMax != nil {
		db = db.Where("rarity_rank <= ?", q.RarityRankMax)
	}
	if q.Marketplaces != nil && len(q.Marketplaces) > 0 {
		db = db.Where("lower(listing::json #>> '{marketplace,name}') IN ?", q.Marketplaces)
	}
	if q.TraitsCount != nil && len(q.TraitsCount) > 0 {
		db = db.Where("trait_count IN ?", q.TraitsCount)
	}
	if q.Traits != nil {
		for _, trait := range q.Traits {
			kv := strings.Split(trait, "=")
			vals := strings.Split(kv[1], ",")
			if len(vals) == 0 {
				continue
			}
			traitType := kv[0]
			db = db.Where("lower(attributes -> ? ->> 'value') IN ?", traitType, vals)
		}
	}
	// TODO: add this back when we add chainid into this
	// if q.ChainId != 0 {
	// 	db.Where("chain_id = ?", q.ChainId)
	// }

	if len(q.Sort) > 0 {
		db = db.Order(q.Sort)
	}
	rows, err := db.Count(&total).Offset(q.Offset).Limit(q.Limit).Rows()
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		token := model.NftToken{}
		var listingRaw, attributesRaw, ownerRaw []byte
		var attributeInfo map[string]model.NftTokenAttribute
		var price *float64
		err := rows.Scan(
			&token.TokenId,
			&token.CollectionAddress,
			&token.Name,
			&token.Description,
			&token.Amount,
			&token.Image,
			&token.ImageCDN,
			&token.ThumbnailCDN,
			&token.ImageContentType,
			&token.RarityRank,
			&token.RarityScore,
			&token.RarityTier,
			&listingRaw,
			&attributesRaw,
			&ownerRaw,
			&price,
		)
		if err != nil {
			return nil, 0, err
		}
		if listingRaw != nil {
			if err := json.Unmarshal(listingRaw, &token.Listing); err != nil {
				return nil, 0, err
			}
		}
		if attributesRaw != nil {
			if err := json.Unmarshal(attributesRaw, &attributeInfo); err != nil {
				return nil, 0, err
			}
			for k, v := range attributeInfo {
				token.Attributes = append(token.Attributes, model.NftTokenAttribute{
					TraitType: k,
					Value:     v.Value,
					Count:     v.Count,
					Rarity:    v.Rarity,
					Frequency: v.Frequency,
					ChainId:   v.ChainId,
				})
			}
		}
		if ownerRaw != nil {
			if err := json.Unmarshal(ownerRaw, &token.Owner); err != nil {
				return nil, 0, err
			}
		}
		res = append(res, token)
	}
	return res, total, nil
}

func (s *store) RefreshViewNFTTokens() error {
	return s.db.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY view_nft_tokens").Error
}

func (s *store) GetNftListing(q NftListingQuery) (listings []*model.NftListing, total int64, err error) {
	db := s.db.Table("nft_listing AS nl").Select(
		"nl.platform_id",
		"nl.token_id",
		"nl.contract_address",
		"nl.chain_id",
		"nl.quantity",
		"nl.payment_token",
		"nl.transaction_hash",
		"nl.from_address",
		"nl.to_address",
		"nl.listing_type",
		"nl.listing_status",
		"nl.created_time",
		"nl.last_updated_time",
		"nl.listing_price",
		"nl.sold_price",
		"t.symbol",
		"t.address",
		"t.is_native",
		"t.decimals",
		"t.icon_url AS token_icon",
		"mp.name",
		"mp.url",
		"mp.icon_url AS marketplace_icon",
	).
		Joins("LEFT JOIN token t ON nl.payment_token = t.id").
		Joins("LEFT JOIN marketplace_platform mp ON nl.platform_id = mp.id")

	if len(q.ContractAddress) > 0 {
		db = db.Where("nl.contract_address = ?", q.ContractAddress)
	}
	if len(q.TokenId) > 0 {
		db = db.Where("nl.token_id = ?", q.TokenId)
	}
	if len(q.Status) > 0 {
		db = db.Where("nl.listing_status = ?", q.Status)
	}
	if len(q.Marketplace) > 0 {
		db = db.Where("mp.name = ?", q.Marketplace)
	}
	if len(q.ListingStatus) > 0 {
		db = db.Where("nl.listing_status = ?", q.ListingStatus)
	}
	if len(q.Sort) > 0 {
		db = db.Order(q.Sort)
	}
	db = db.Count(&total)
	if q.Offset != 0 {
		db = db.Offset(q.Offset)
	}
	if q.Limit != 0 {
		db = db.Limit(q.Limit)
	}
	if q.From != 0 {
		db = db.Where("extract(epoch from nl.created_time) * 1000 >= ?", q.From)
	}
	if q.To != 0 {
		db = db.Where("extract(epoch from nl.created_time) * 1000 >= ?", q.From)
	}
	rows, err := db.Order("created_time DESC").Rows()
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	if err != nil {
		return nil, 0, err
	}
	for rows.Next() {
		listing := model.NftListing{Marketplace: &model.MarketplacePlatform{}}
		token := model.Token{}
		err := rows.Scan(
			&listing.PlatformId,
			&listing.TokenId,
			&listing.ContractAddress,
			&listing.ChainId,
			&listing.Quantity,
			&listing.PaymentToken,
			&listing.TransactionHash,
			&listing.FromAddress,
			&listing.ToAddress,
			&listing.ListingType,
			&listing.ListingStatus,
			&listing.CreatedTime,
			&listing.LastUpdatedTime,
			&listing.ListingPrice,
			&listing.SoldPrice,
			&token.Symbol,
			&token.Address,
			&token.IsNative,
			&token.Decimals,
			&token.IconUrl,
			&listing.Marketplace.Name,
			&listing.Marketplace.URL,
			&listing.Marketplace.IconUrl,
		)
		if err != nil {
			return nil, 0, err
		}

		listingPrice := ""
		if listing.ListingPrice != nil {
			listingPrice = *listing.ListingPrice
		}

		listing.ListingPriceObj = &model.Price{Amount: listingPrice, Token: token}
		listing.SoldPriceObj = &model.Price{Amount: listing.SoldPrice, Token: token}
		listings = append(listings, &listing)
	}

	return listings, total, err
}

func (s *store) GetNftListingByTokenID(collectionAddress, tokenID string) ([]model.NftListingMarketplace, error) {
	rows, err := s.db.Raw(
		`
		SELECT
			list.contract_address,
			list.token_id,
			list.platform_id,
			list.listing_status,
			list.listing_price,
			list.created_time,
			plat.name,
			token.symbol,
			token.decimals
		FROM
			nft_listing AS LIST
			LEFT JOIN marketplace_platform AS plat ON list.platform_id = plat.id
			LEFT JOIN token ON list.payment_token = token.id
		WHERE
			list.contract_address = $1
			AND list.token_id = $2
			AND list.listing_status = 'listed'
		ORDER BY
			created_time DESC
		LIMIT 1;
		`, collectionAddress, tokenID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var listingMarketplace []model.NftListingMarketplace
	for rows.Next() {
		h := model.NftListingMarketplace{}
		rows.Scan(
			&h.ContractAddress,
			&h.TokenId,
			&h.PlatformId,
			&h.ListingStatus,
			&h.ListingPrice,
			&h.CreatedTime,
			&h.PlatformName,
			&h.PaymentToken,
			&h.PaymentTokenDecimals,
		)
		listingMarketplace = append(listingMarketplace, h)
	}
	return listingMarketplace, nil
}

func (s *store) GetNftTokenTxHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error) {
	rows, err := s.db.Raw(
		`
		SELECT
			transf."from",
			transf."to",
			transf.token_id,
			transf.transaction_hash,
			transf.contract_address,
			transf.created_time,
			listing.listing_status
		FROM
			nft_transfer AS transf
			LEFT JOIN nft_listing as listing ON transf.transaction_hash = listing.transaction_hash
		WHERE
			transf.contract_address = $1 and transf.token_id = $2
		LIMIT 5;
		`, collectionAddress, tokenId).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var history []model.NftTxHistory
	for rows.Next() {
		h := model.NftTxHistory{}
		rows.Scan(
			&h.From,
			&h.To,
			&h.TokenId,
			&h.TransactionHash,
			&h.ContractAddress,
			&h.CreatedTime,
			&h.ListingStatus,
		)
		history = append(history, h)
	}
	return history, nil
}

func (s *store) RefreshViewNFTCollectionAttributes() error {
	return s.db.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY view_nft_collections_attributes").Error
}

func (s *store) GetCollectionsByWalletAddress(q WalletCollectionQuery) ([]model.NftCollection, int64, error) {
	collections := make([]model.NftCollection, 0, q.Limit)
	var total int64
	db := s.db.Table("view_nft_collections AS vnc").
		Joins("JOIN nft_owner no ON vnc.address = no.collection_address").
		Where("no.owner_address = ?", q.WalletAddress)

	if len(q.Sort) > 0 {
		db = db.Order(q.Sort)
	}
	// count query
	err := s.db.Table("nft_owner").Where("owner_address = ?", q.WalletAddress).Distinct("collection_address").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	// data query
	err = db.Select("DISTINCT ON (vnc.address) vnc.*").Offset(q.Offset).Limit(q.Limit).Find(&collections).Error
	if err != nil {
		return nil, 0, err
	}
	return collections, total, nil
}

func (s *store) GetNftMarketplaceCollection(collectionAddress string) ([]model.NftListingMarketplace, error) {
	rows, err := s.db.Raw(
		`
		SELECT DISTINCT
			nft_listing.contract_address,
			marketplace_platform."name"
		FROM
			nft_listing
			LEFT JOIN marketplace_platform ON nft_listing.platform_id = marketplace_platform.id
		WHERE
			nft_listing.contract_address = $1;
		`, collectionAddress).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var marketplaceCollection []model.NftListingMarketplace
	for rows.Next() {
		h := model.NftListingMarketplace{}
		rows.Scan(
			&h.ContractAddress,
			&h.PlatformName,
		)
		marketplaceCollection = append(marketplaceCollection, h)
	}
	return marketplaceCollection, nil
}

// TODO(trkhoi): because view_nft_collection_stats not working, we need to use this function to get collection stats
// remove asap
func (s *store) GetNftCollectionStats(collectionAddress string) ([]model.ViewNftCollectionStats, error) {
	rows, err := s.db.Raw(
		`
		WITH v1 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS one_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now())
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '1 day'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		v7 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS seven_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now())
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '7 days'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		v30 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS thirty_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now())
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '30 days'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		v AS (
			SELECT
				nl.contract_address,
				sum(COALESCE(NULLIF(nl.sold_price,
							''::text),
						'0'::text)::numeric) AS all_time_volume,
				max(nl.payment_token) AS payment_token
			FROM
				nft_listing nl
			WHERE
				nl.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nl.contract_address
		),
		prev_v1 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS prev_one_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now() - '1 day'::interval)
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '2 days'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		prev_v7 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS prev_seven_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now() - '7 days'::interval)
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '14 days'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		prev_v30 AS (
			SELECT
				nft_listing.contract_address,
				sum(COALESCE(NULLIF(nft_listing.sold_price,
							''::text),
						'0'::text)::numeric) AS prev_thirty_day_volume
			FROM
				nft_listing
			WHERE
				date_part('epoch'::text,
					nft_listing.created_time) <= date_part('epoch'::text,
					now() - '30 days'::interval)
				AND date_part('epoch'::text,
					nft_listing.created_time) >= date_part('epoch'::text,
					now() - '60 days'::interval)
				AND nft_listing.listing_status = 'sold'::nft_listing_status
			GROUP BY
				nft_listing.contract_address
		),
		mcollection AS (
			SELECT
				max(nft_marketplace_collection_snapshot.created_time) AS created_time,
				nft_marketplace_collection_snapshot.collection_address,
				nft_marketplace_collection_snapshot.platform_id
			FROM
				nft_marketplace_collection_snapshot
			GROUP BY
				nft_marketplace_collection_snapshot.collection_address,
				nft_marketplace_collection_snapshot.platform_id
		)
		SELECT
			t.address,
			t.payment_token,
			t.floor_price,
			t.one_day_volume,
			COALESCE(NULLIF(t.one_day_volume_change, 0::numeric)::double precision, t.one_day_volume::double precision / pow(10::double precision, t.decimals::double precision)) AS one_day_volume_change,
			t.seven_day_volume,
			COALESCE(NULLIF(t.seven_day_volume_change, 0::numeric)::double precision, t.seven_day_volume::double precision / pow(10::double precision, t.decimals::double precision)) AS seven_day_volume_change,
			t.thirty_day_volume,
			COALESCE(NULLIF(t.thirty_day_volume_change, 0::numeric)::double precision, t.thirty_day_volume::double precision / pow(10::double precision, t.decimals::double precision)) AS thirty_day_volume_change,
			t.all_time_volume,
			t.all_time_volume::double precision / pow(10::double precision, t.decimals::double precision) AS all_time_volume_change
		FROM (
			SELECT
				nc.address,
				(
					SELECT
						min(nft_marketplace_collection_snapshot.floor_price::numeric) AS min
					FROM
						nft_marketplace_collection_snapshot
					WHERE
						nft_marketplace_collection_snapshot.created_time = mcollection.created_time
						AND nft_marketplace_collection_snapshot.collection_address = mcollection.collection_address
						AND nft_marketplace_collection_snapshot.platform_id = mcollection.platform_id) AS floor_price, COALESCE(v1.one_day_volume, 0::numeric) AS one_day_volume, COALESCE((v1.one_day_volume - prev_v1.prev_one_day_volume) / NULLIF(prev_v1.prev_one_day_volume, 0::numeric), 0::numeric) AS one_day_volume_change, COALESCE(v7.seven_day_volume, 0::numeric) AS seven_day_volume, COALESCE((v7.seven_day_volume - prev_v7.prev_seven_day_volume) / NULLIF(prev_v7.prev_seven_day_volume, 0::numeric), 0::numeric) AS seven_day_volume_change, COALESCE(v30.thirty_day_volume, 0::numeric) AS thirty_day_volume, COALESCE((v30.thirty_day_volume - prev_v30.prev_thirty_day_volume) / NULLIF(prev_v30.prev_thirty_day_volume, 0::numeric), 0::numeric) AS thirty_day_volume_change, COALESCE(v.all_time_volume, 0::numeric) AS all_time_volume, token.id AS payment_token, token.decimals
				FROM
					nft_collection nc
				LEFT JOIN mcollection ON nc.address = mcollection.collection_address
				LEFT JOIN v1 ON nc.address = v1.contract_address
				LEFT JOIN prev_v1 ON nc.address = prev_v1.contract_address
				LEFT JOIN v7 ON nc.address = v7.contract_address
				LEFT JOIN prev_v7 ON nc.address = prev_v7.contract_address
				LEFT JOIN v30 ON nc.address = v30.contract_address
				LEFT JOIN prev_v30 ON nc.address = prev_v30.contract_address
				LEFT JOIN v ON nc.address = v.contract_address
				LEFT JOIN token ON v.payment_token = token.id) t
		WHERE
			t.address = $1;
		`, collectionAddress).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var stats []model.ViewNftCollectionStats
	for rows.Next() {
		h := model.ViewNftCollectionStats{}
		rows.Scan(
			&h.Address,
			&h.PaymentToken,
			&h.FloorPrice,
			&h.OneDayVolume,
			&h.OneDayVolumeChange,
			&h.SevenDayVolume,
			&h.SevenDayVolumeChange,
			&h.ThirtyDayVolume,
			&h.ThirtyDayVolumeChange,
			&h.AllTimeVolume,
			&h.AllTimeVolumeChange,
		)
		stats = append(stats, h)
	}
	return stats, nil
}

func (s *store) GetTokensByWalletAddress(q WalletTokenQuery) ([]model.NftToken, int64, error) {
	db := s.db.Table("nft_token")
	tokens := make([]model.NftToken, 0, q.Limit)
	var total int64
	var ownerToken []struct {
		CollectionAddress string
		TokenId           string
	}

	ownerQ := s.db.Table("nft_owner").Select("collection_address,token_id").Where("owner_address = ?", q.WalletAddress)
	if q.ChainId != nil {
		ownerQ = ownerQ.Where("chain_id = ?", q.ChainId)
	}
	if q.CollectionAddresses != nil {
		subq := ""
		for _, v := range q.CollectionAddresses {
			subq = subq + fmt.Sprintf("('%v'),", v)
		}
		subq = fmt.Sprintf("collection_address = ANY(VALUES %s)", subq[:len(subq)-1])
		ownerQ = ownerQ.Where(subq)
	}
	// check if owner have no tokens
	var tokensCount int64
	if err := ownerQ.Count(&tokensCount).Find(&ownerToken).Error; err != nil {
		return nil, 0, err
	}
	if tokensCount == 0 {
		return tokens, 0, nil
	}

	subq := ""
	for _, v := range ownerToken {
		subq = subq + fmt.Sprintf("('%v','%v'),", v.CollectionAddress, v.TokenId)
	}
	subq = fmt.Sprintf("(collection_address,token_id) = ANY(VALUES %s)", subq[:len(subq)-1])

	db = db.Where(subq)
	if len(q.Sort) > 0 {
		db = db.Order(q.Sort)
	}
	if err := db.Count(&total).Offset(q.Offset).Limit(q.Limit).Find(&tokens).Error; err != nil {
		return nil, 0, err
	}
	return tokens, total, nil
}

func (pg *store) GetAttributeByCollectionAddressTokenID(collectionAddress, tokenID string) ([]model.NftTokenAttribute, error) {
	var tokenAttributes []model.NftTokenAttribute
	err := pg.db.Table("nft_token_attribute").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenID).Find(&tokenAttributes).Error
	if err != nil {
		return nil, err
	}
	return tokenAttributes, nil
}

func (pg *store) GetAllMarketplacePlatform() (platforms []model.MarketplacePlatform, err error) {
	return platforms, pg.db.Table("marketplace_platform").Find(&platforms).Error
}

func (pg *store) SummarizeSnapshotCollection(platformId int64) error {
	query := `
		INSERT INTO nft_marketplace_collection_snapshot (collection_address, token_id, platform_id, total_volume, created_time, floor_price)
		SELECT
			contract_address,
			payment_token,
			platform_id,
			sum(sold_price::NUMERIC) AS total_volume,
			date(created_time) AS created_time,
			min(sold_price::NUMERIC) AS floor_price
		FROM
			nft_listing
		WHERE
			platform_id = %v AND listing_status = 'sold'
		GROUP BY
			contract_address,
			payment_token,
			date(created_time),
			platform_id
		ORDER BY
			date(created_time) ON CONFLICT (platform_id,
				collection_address,
				token_id,
				created_time) DO UPDATE SET total_volume = nft_marketplace_collection_snapshot.total_volume, floor_price = nft_marketplace_collection_snapshot.floor_price;
	`
	pg.db.Exec(fmt.Sprintf(query, platformId))

	return nil
}

func (pg *store) SummarizeSnapshotHolder() error {
	query := `
		INSERT INTO nft_holder_snapshot (collection_address, total_holder, created_time)
		SELECT
			collection_address,
			count(collection_address) AS total_holder,
			date(created_time) AS created_time
		FROM
			nft_owner
		GROUP BY
			collection_address,
			date(created_time)
		ORDER BY
			date(created_time)
	`
	pg.db.Exec(fmt.Sprintf(query))

	return nil
}

func (pg *store) GetNftTokenAttrWithSoulBound(collectionAddress string) ([]model.NftTokenAttrSoulBound, error) {
	rows, err := pg.db.Raw(
		`
		SELECT
			trait_type,
			value,
			count(value)
		FROM
			nft_token_attribute
		WHERE
			collection_address = $1
			AND trait_type = 'Character'
			AND token_id in( SELECT DISTINCT
					(token_id)
					FROM nft_token_attribute
				WHERE
					collection_address = $2
					AND trait_type = 'Soulbound'
					AND value = 'True')
		GROUP BY
			trait_type, value;
		`, collectionAddress, collectionAddress).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var soulBound []model.NftTokenAttrSoulBound
	for rows.Next() {
		h := model.NftTokenAttrSoulBound{}
		rows.Scan(
			&h.TraitType,
			&h.Value,
			&h.Count,
		)
		soulBound = append(soulBound, h)
	}
	return soulBound, nil
}

func (pg *store) GetSolanaMappingAddress(solscanId string) (model *model.SolanaMappingAddress, err error) {
	return model, pg.db.Table("solana_mapping_address").Where("solscan_id = ?", solscanId).First(&model).Error
}
