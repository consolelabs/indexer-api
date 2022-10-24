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

func (s *store) UpdateCollection(collection *model.NftCollection) error {
	return s.db.Table("nft_collection").Where("address = ?", collection.Address).Save(&collection).Error
}

func (s *store) SaveOwner(owner *model.NftOwner) error {
	return s.db.Create(owner).Error
}

func (s *store) GetNftOwners(q NftOwnerQuery) ([]model.NftOwner, int64, error) {
	db := s.db.Table("nft_owner")
	var owners []model.NftOwner
	var total int64
	if q.CollectionAddress != nil {
		db = db.Where("collection_address = ?", *q.CollectionAddress)
	}
	if q.TokenId != nil {
		db = db.Where("token_id = ?", *q.TokenId)
	}
	if q.ChainId != nil {
		db = db.Where("chain_id = ?", *q.ChainId)
	}

	err := db.Count(&total).Offset(q.Offset).Limit(q.Limit).Find(&owners).Error
	if err != nil {
		return nil, 0, err
	}
	return owners, total, nil
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

func (s *store) AttributeUpsertOne(tokenAttribute model.NftTokenAttribute) error {
	tx := s.db.Begin()
	if err := tx.Table("nft_token_attribute").Where("collection_address = ? AND token_id = ? AND trait_type = ?", tokenAttribute.CollectionAddress, tokenAttribute.TokenId, tokenAttribute.TraitType).Save(&tokenAttribute).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (s *store) UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId string, ownerAddress string) error {
	return s.db.Table("nft_owner").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenId).Updates(map[string]interface{}{"owner_address": ownerAddress, "last_updated_time": time.Now()}).Error
}

func (s *store) SaveTransfer(transfer *model.NftTransfer) error {
	return s.db.Create(transfer).Error
}

func (s *store) DeleteNftToken(q NftTokenQuery) error {
	db := s.db.Table("nft_token")
	if q.CollectionAddress != nil {
		db = db.Where("collection_address = ?", *q.CollectionAddress)
	}
	if q.TokenId != nil {
		db = db.Where("token_id = ?", *q.TokenId)
	}
	return db.Delete(&model.NftToken{}).Error
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

func (s *store) GetListingByTokenIdAndContractAddress(tokenId, contractAddress string) (*model.NftListing, error) {
	var listing model.NftListing
	return &listing, s.db.Table("nft_listing").Where("token_id = ? AND contract_address = ?", tokenId, contractAddress).First(&listing).Error
}

func (pg *store) UpdateListing(listing *model.NftListing) error {
	return pg.db.Table("nft_listing").Where("id = ?", listing.Id).Save(&listing).Error
}

func (pg *store) SaveListing(listing *model.NftListing) error {
	return pg.db.Create(listing).Error
}

func (pg *store) GetPlatformByName(name string) (*model.MarketplacePlatform, error) {
	var platform model.MarketplacePlatform
	return &platform, pg.db.Table("marketplace_platform").Where("lower(name) = lower(?)", name).First(&platform).Error
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

func (s *store) GetDistinctAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error) {
	var attributes []model.NftTokenAttribute
	return attributes, s.db.Table("view_nft_collections_attributes").
		Select("collection_address, trait_type, value, count, frequency, rarity").
		Where("collection_address = ?", collectionAddress).
		Order("count DESC").Find(&attributes).Error
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

func (s *store) DeleteNftTokenAttributesByCollectionAddress(collecionAddress string) error {
	return s.db.Table("nft_token_attribute").Where("collection_address = ?", collecionAddress).Delete(&model.NftTokenAttribute{}).Error
}

func (s *store) DeleteNftOwnerByCollectionAddress(collecionAddress string) error {
	return s.db.Table("nft_owner").Where("collection_address = ?", collecionAddress).Delete(&model.NftOwner{}).Error
}

func (s *store) DeleteNftTransferByContractAddress(collecionAddress string) error {
	return s.db.Table("nft_transfer").Where("contract_address = ?", collecionAddress).Delete(&model.NftTransfer{}).Error
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

func (s *store) UpsertNftMarketplaceCollection(model *model.NftMarketplaceCollection) error {
	return s.db.Table("nft_marketplace_collection").Where("collection_address = ? AND platform_id = ?", model.CollectionAddress, model.PlatformID).Save(model).Error
}

func (s *store) RefreshViewNFTCollections() error {
	return s.db.Exec("REFRESH MATERIALIZED VIEW CONCURRENTLY view_nft_collections").Error
}

func (s *store) GetCollectionTraitsCount(collectionAddress string) ([]model.NftTokenAttributeCount, error) {
	var attrCount []model.NftTokenAttributeCount
	subQ := s.db.Table("nft_token_attribute").
		Select("count(*) AS trait_count").
		Where("collection_address = ?", collectionAddress).
		Group("token_id")
	db := s.db.Table("(?) AS temp", subQ)
	return attrCount, db.
		Select(
			"trait_count AS id",
			"count(*) AS count",
			"(CASE WHEN MAX(nc.supply) = 0 THEN 1 ELSE (count(*)::numeric / MAX(nc.supply)) END) AS ratio",
		).
		Joins("JOIN nft_collection nc ON nc.address = ?", collectionAddress).
		Group("trait_count").
		Order("count DESC").
		Find(&attrCount).Error
}

func (pg *store) GetAttributeByCollectionAddressTokenID(collectionAddress, tokenID string) ([]model.NftTokenAttribute, error) {
	var tokenAttributes []model.NftTokenAttribute
	err := pg.db.Table("nft_token_attribute").Where("collection_address = ? AND token_id = ?", collectionAddress, tokenID).Find(&tokenAttributes).Error
	if err != nil {
		return nil, err
	}
	return tokenAttributes, nil
}

func (s *store) GetMarketplaceWithListingStats(collectionAddress string) ([]model.MarketplaceWithListingStats, error) {
	var metadatas []model.MarketplaceWithListingStats
	rows, err := s.db.Table("nft_listing AS nl").Raw(
		`WITH lgn AS (
			SELECT DISTINCT ON (token_id, platform_id) platform_id, listing_status
			FROM nft_listing
			WHERE contract_address = @collectionAddress
			ORDER BY token_id, platform_id, created_time DESC
		)
		SELECT
       mp.id,
       mp.name,
       mp.url,
       count(*)                                                                            AS listing_number,
       (CASE WHEN MAX(nc.supply) = 0 THEN 1 ELSE (count(*)::numeric / MAX(nc.supply)) END) AS lisiting_percentage
		FROM
			lgn
			JOIN nft_collection nc ON nc.address = @collectionAddress
			JOIN marketplace_platform mp ON mp.id = lgn.platform_id
		WHERE lgn.listing_status = 'listed'
		GROUP BY mp.id, mp.name, mp.url
		ORDER BY
			listing_number DESC;
	`, map[string]interface{}{
			"collectionAddress": collectionAddress,
		}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		metadata := model.MarketplaceWithListingStats{}
		err = rows.Scan(
			&metadata.ID,
			&metadata.Name,
			&metadata.URL,
			&metadata.ListingNumber,
			&metadata.ListingPercentage,
		)
		if err != nil {
			return nil, err
		}
		metadatas = append(metadatas, metadata)
	}

	return metadatas, nil
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
	defer rows.Close()
	for rows.Next() {
		listing := model.NftListing{Marketplace: &model.MarketplacePlatform{}}
		token := model.Token{}
		rows.Scan(
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
		listing.ListingPriceObj = &model.Price{Amount: listing.ListingPrice, Token: token}
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

func (s *store) GetNftTransfers(q NftTransferQuery) ([]model.NftTransfer, int64, error) {
	var transfers []model.NftTransfer
	var total int64
	db := s.db.Table("nft_transfer")
	if q.CollectionAddress != nil {
		db = db.Where("contract_address = ?", *q.CollectionAddress)
	}
	if q.TokenId != nil {
		db = db.Where("token_id = ?", *q.TokenId)
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
	return transfers, total, db.Find(&transfers).Error
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

func (s *store) GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingSoldPrices []*model.NftListingSoldPrice, err error) {
	rows, err := s.db.Raw(
		`SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(COALESCE(AVG(nl.sold_price::DOUBLE PRECISION / nl.quantity::DOUBLE PRECISION), 0), 'FM9999999999999999999999999999') AS avg_sold_price
		FROM
			generate_records_by_time_range (@timeDuration, @timeInterval, @timeRoundUp) tr
		LEFT JOIN token t ON t.symbol = @tokenSymbol
		LEFT JOIN nft_listing nl ON nl.created_time < tr.t
			AND nl.created_time > (tr.t - (@timeInterval)::INTERVAL)
			AND nl.contract_address = @contractAddress
			AND nl.token_id = @tokenId
			AND nl.listing_status = 'sold'
			AND nl.payment_token = t.id
		GROUP BY
			(tr.t,
				t.symbol,
				t.is_native,
				t.address,
				t.decimals,
				t.icon_url)
		ORDER BY
			tr.t ASC;`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
			"tokenId":         tokenId,
		}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		listingSoldPrice := &model.NftListingSoldPrice{}
		var tokenRaw []byte
		tokenObj := model.Token{}
		soldPriceObj := &model.Price{}
		err = rows.Scan(
			&listingSoldPrice.TimeRangeStart,
			&tokenRaw,
			&soldPriceObj.Amount,
		)
		if err != nil {
			return nil, err
		}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &tokenObj); err != nil {
				return nil, err
			}
		}

		soldPriceObj.Token = tokenObj

		listingSoldPrice.SoldPriceObj = soldPriceObj
		listingSoldPrice.TimeInterval = timeInterval

		listingSoldPrices = append(listingSoldPrices, listingSoldPrice)
	}
	return listingSoldPrices, nil
}

func (s *store) GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingAvgPricesAndFloors []*model.NftListingAvgPriceAndFloor, err error) {
	rows, err := s.db.Raw(
		`SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(COALESCE(AVG(nl.sold_price::DOUBLE PRECISION / nl.quantity::DOUBLE PRECISION), 0), 'FM9999999999999999999999999999') AS avg_sold_price,
			nmcs.floor_price
		FROM
			generate_records_by_time_range (@timeDuration,
				@timeInterval,
				@timeRoundUp) tr
			LEFT JOIN token t ON t.symbol = @tokenSymbol
			LEFT JOIN nft_listing nl ON nl.created_time < tr.t
				AND nl.created_time > (tr.t - (@timeInterval)::INTERVAL)
				AND contract_address = @contractAddress
				AND nl.listing_status = 'sold'
				AND nl.payment_token = t.id
				AND nl.token_id = @tokenId
			LEFT JOIN nft_token nt ON nt.collection_address = @contractAddress
			JOIN LATERAL (
				SELECT
					nmcs.floor_price
				FROM
					nft_marketplace_collection_snapshot nmcs
				WHERE
					nmcs.collection_address = @contractAddress
					AND nmcs.created_time < tr.t
					AND nmcs.token_id = t.id
				ORDER BY
					nmcs.created_time DESC
				LIMIT 1) nmcs ON TRUE
		GROUP BY
			(tr.t,
				nmcs.floor_price,
				t.symbol,
				t.is_native,
				t.address,
				t.decimals,
				t.icon_url)
		ORDER BY
			tr.t ASC;`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
			"tokenId":         tokenId,
		}).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		listingAvgPriceAndFloor := &model.NftListingAvgPriceAndFloor{}
		var tokenRaw []byte
		tokenObj := model.Token{}
		avgPriceObj := &model.Price{}
		floorPriceObj := &model.Price{}
		err = rows.Scan(
			&listingAvgPriceAndFloor.TimeRangeStart,
			&tokenRaw,
			&avgPriceObj.Amount,
			&floorPriceObj.Amount,
		)
		if err != nil {
			return nil, err
		}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &tokenObj); err != nil {
				return nil, err
			}
		}

		avgPriceObj.Token = tokenObj
		floorPriceObj.Token = tokenObj

		listingAvgPriceAndFloor.AvgPriceObj = avgPriceObj
		listingAvgPriceAndFloor.FloorPriceObj = floorPriceObj

		listingAvgPriceAndFloor.TimeInterval = timeInterval

		listingAvgPricesAndFloors = append(listingAvgPricesAndFloors, listingAvgPriceAndFloor)
	}
	return listingAvgPricesAndFloors, nil
}

func (s *store) GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration string) (saleVolumeAndFloorPrices []*model.NftCollectionSaleVolumeAndFloorPrice, err error) {
	rows, err := s.db.Raw(
		`SELECT
			nmcs.collection_address,
			nmcs.created_time,
			json_build_object('symbol', t.symbol, 'address', t.address, 'is_native', t.is_native, 'decimals', t.decimals) as token,
			nmcs.floor_price,
			to_char(SUM((
					CASE WHEN nl.created_time < (nmcs.created_time + INTERVAL '1 days')
						AND nl.created_time >= nmcs.created_time THEN
						COALESCE(nl.sold_price, '0')
					ELSE
						'0'
					END)::double precision), 'FM9999999999999999999999999999') AS sold_volume
		FROM nft_marketplace_collection_snapshot nmcs
		LEFT JOIN nft_listing nl ON nl.contract_address = nmcs.collection_address
		LEFT JOIN token t ON nmcs.token_id = t.id
		WHERE
			nmcs.created_time < CURRENT_DATE AND
			nmcs.created_time > CASE
                WHEN @timeDuration = '7 Days' THEN CURRENT_DATE - interval '7 days'
                WHEN @timeDuration = '30 Days' THEN CURRENT_DATE - interval '30 days'
                WHEN @timeDuration = '3 Months' THEN CURRENT_DATE - interval '3 Months'
                WHEN @timeDuration = '12 Months' THEN CURRENT_DATE - interval '12 Months'
                ELSE '2021-01-01'
            END AND
			nmcs.collection_address = @collectionAddress AND
			nl.listing_status = 'sold' GROUP BY(nmcs.id, t.id) ORDER BY nmcs.created_time ASC
	`, map[string]interface{}{
			"timeDuration":      timeDuration,
			"collectionAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		saleVolumeAndFloorPrice := &model.NftCollectionSaleVolumeAndFloorPrice{}
		var tokenRaw []byte

		tokenObj := model.Token{}
		floorPriceObj := &model.Price{}
		soldVolumeObj := &model.Price{}
		err = rows.Scan(
			&saleVolumeAndFloorPrice.CollectionAddress,
			&saleVolumeAndFloorPrice.CreatedTime,
			&tokenRaw,
			&floorPriceObj.Amount,
			&soldVolumeObj.Amount,
		)

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &tokenObj); err != nil {
				return nil, err
			}
		}
		floorPriceObj.Token = tokenObj
		soldVolumeObj.Token = tokenObj

		saleVolumeAndFloorPrice.FloorPriceObj = floorPriceObj
		saleVolumeAndFloorPrice.SaleVolume = soldVolumeObj
		if err != nil {
			return nil, err
		}
		saleVolumeAndFloorPrices = append(saleVolumeAndFloorPrices, saleVolumeAndFloorPrice)
	}
	return saleVolumeAndFloorPrices, nil
}

func (s *store) GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange string) (listingPriceRanges []*model.NftListingPriceRange, err error) {
	rows, err := s.db.Raw(
		`SELECT
		t.symbol,
		t.address,
		t.is_native,
		t.decimals,
		nl.price_range_start,
		json_agg(nl.total_volume)
	FROM (
		SELECT
			*
		FROM
			get_nft_listing_price_range (?,
				?,
				?) src (payment_token,
				price_range_start,
				total_volume)) nl
		LEFT JOIN token t ON nl.payment_token = t.id
	GROUP BY
		(t.symbol,
			t.address,
			t.is_native,
			t.decimals,
			nl.price_range_start)
	ORDER BY
		price_range_start::double precision ASC;`, collectionAddress, paymentToken, priceRange).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		listingPriceRange := &model.NftListingPriceRange{}
		listingPriceRangeByMarketplacePlatforms := []model.NftListingPriceRangeByMarketplacePlatform{}
		var listingPriceRangeByMarketplacePlatformsRaw []byte
		priceRangeStartObj := &model.Price{}
		err = rows.Scan(
			&priceRangeStartObj.Token.Symbol,
			&priceRangeStartObj.Token.Address,
			&priceRangeStartObj.Token.IsNative,
			&priceRangeStartObj.Token.Decimals,
			&priceRangeStartObj.Amount,
			&listingPriceRangeByMarketplacePlatformsRaw,
		)
		listingPriceRange.CollectionAddress = collectionAddress
		listingPriceRange.PriceRangeStart = priceRangeStartObj
		listingPriceRange.PriceRange = priceRange

		if listingPriceRangeByMarketplacePlatformsRaw != nil {
			if err := json.Unmarshal(listingPriceRangeByMarketplacePlatformsRaw, &listingPriceRangeByMarketplacePlatforms); err != nil {
				return nil, err
			}
		}

		listingPriceRange.MarketplacePlatformListings = listingPriceRangeByMarketplacePlatforms
		if err != nil {
			return nil, err
		}
		listingPriceRanges = append(listingPriceRanges, listingPriceRange)
	}
	return listingPriceRanges, nil
}

func (s *store) GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange string) (listingRarityRankRanges []*model.NftListingRarityRankRange, err error) {
	rows, err := s.db.Raw(
		`SELECT
		rarity_rank_range_start,
		json_agg(marketplace_platform_average_price) AS marketplace_platform_average_prices
	FROM (
		SELECT
			rarity_rank_range_start,
			json_build_object('marketplace_platform', json_build_object('name', mp.name, 'url', mp.url, 'icon_url', mp.icon_url), 'average_price', json_build_object('amount', nl.average_price, 'token', json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url))) AS marketplace_platform_average_price
		FROM (
			SELECT
				*
			FROM
				get_nft_listing_rarity_rank_range (?, ?, ?) src (payment_token,
					platform_id,
					average_price,
					rarity_rank_range_start)) AS nl
		LEFT JOIN token t ON t.id = nl.payment_token
		LEFT JOIN marketplace_platform mp ON mp.id = nl.platform_id) AS nl
	GROUP BY
		(nl.rarity_rank_range_start)
	ORDER BY
		nl.rarity_rank_range_start::int4 ASC;`, collectionAddress, paymentToken, rarityRankRange).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		listingRarityRankRange := &model.NftListingRarityRankRange{}
		listingRarityRankRangeByMarketplacePlatforms := []model.NftListingRarityRankRangeByMarketplacePlatform{}
		var listingPriceRangeByMarketplacePlatformsRaw []byte
		err = rows.Scan(
			&listingRarityRankRange.RarityRankRangeStart,
			&listingPriceRangeByMarketplacePlatformsRaw,
		)
		listingRarityRankRange.CollectionAddress = collectionAddress
		listingRarityRankRange.RarityRankRange = rarityRankRange

		if listingPriceRangeByMarketplacePlatformsRaw != nil {
			if err := json.Unmarshal(listingPriceRangeByMarketplacePlatformsRaw, &listingRarityRankRangeByMarketplacePlatforms); err != nil {
				return nil, err
			}
		}

		listingRarityRankRange.MarketplacePlatformListings = listingRarityRankRangeByMarketplacePlatforms
		if err != nil {
			return nil, err
		}
		listingRarityRankRanges = append(listingRarityRankRanges, listingRarityRankRange)
	}
	return listingRarityRankRanges, nil
}

func (s *store) GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingMarketCapAndVolumeTimeRanges []*model.NftListingMarketCapAndVolumeTimeRange, err error) {
	rows, err := s.db.Raw(
		`SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(nmcs.total_volume::DOUBLE PRECISION * COUNT(nt), 'FM9999999999999999999999999999') AS market_cap,
			to_char(COALESCE(SUM(nl.sold_price::DOUBLE PRECISION), 0), 'FM9999999999999999999999999999') AS total_volume
		FROM
			generate_records_by_time_range (@timeDuration,
				@timeInterval,
				@timeRoundUp) tr
			LEFT JOIN token t ON t.symbol = @tokenSymbol
			LEFT JOIN nft_listing nl ON nl.created_time < tr.t
				AND nl.created_time > (tr.t - (@timeInterval)::INTERVAL)
				AND contract_address = @contractAddress
				AND nl.listing_status = 'sold'
				AND nl.payment_token = t.id
			LEFT JOIN nft_token nt ON nt.collection_address = @contractAddress
			JOIN LATERAL (
				SELECT
					nmcs.total_volume
				FROM
					nft_marketplace_collection_snapshot nmcs
				WHERE
					nmcs.collection_address = @contractAddress
					AND nmcs.created_time < tr.t
					AND nmcs.token_id = t.id
				ORDER BY
					nmcs.created_time DESC
				LIMIT 1) nmcs ON TRUE
		GROUP BY
			(tr.t,
				nmcs.total_volume,
				t.symbol,
				t.is_native,
				t.address,
				t.decimals,
				t.icon_url)
		ORDER BY
			tr.t ASC;`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		listingMarketCapAndVolumeTimeRange := &model.NftListingMarketCapAndVolumeTimeRange{}
		tokenRaw := []byte{}
		marketCapPrice := &model.Price{}
		volumePrice := &model.Price{}
		err = rows.Scan(
			&listingMarketCapAndVolumeTimeRange.TimeRangeStart,
			&tokenRaw,
			&marketCapPrice.Amount,
			&volumePrice.Amount,
		)
		if err != nil {
			return nil, err
		}

		listingMarketCapAndVolumeTimeRange.TimeInterval = timeInterval

		token := &model.Token{}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &token); err != nil {
				return nil, err
			}
		}

		marketCapPrice.Token = *token
		volumePrice.Token = *token

		listingMarketCapAndVolumeTimeRange.MarketCap = marketCapPrice
		listingMarketCapAndVolumeTimeRange.Volume = volumePrice

		listingMarketCapAndVolumeTimeRanges = append(listingMarketCapAndVolumeTimeRanges, listingMarketCapAndVolumeTimeRange)
	}
	return listingMarketCapAndVolumeTimeRanges, nil
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

func (s *store) GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp string) (listingUniqueOwnerTimeRanges []*model.NftListingUniqueOwnersTimeRange, err error) {
	rows, err := s.db.Raw(
		`
		SELECT
			t AS time_range,
			COALESCE(COUNT(DISTINCT(nl.to_address)), 0) AS unique_owners
		FROM
			generate_records_by_time_range (@timeDuration, @timeInterval, @timeRoundUp) tr
			LEFT JOIN nft_listing nl ON nl.created_time < tr.t
				AND nl.created_time > (tr.t - (@timeInterval)::INTERVAL)
				AND contract_address = @contractAddress
				AND nl.listing_status = 'sold'
		GROUP BY
			(tr.t)
		ORDER BY
			tr.t ASC;
			`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"contractAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		listingUniqueOwnerTimeRange := &model.NftListingUniqueOwnersTimeRange{}
		err = rows.Scan(
			&listingUniqueOwnerTimeRange.TimeRangeStart,
			&listingUniqueOwnerTimeRange.UniqueOwners,
		)
		listingUniqueOwnerTimeRange.TimeInterval = timeInterval
		listingUniqueOwnerTimeRanges = append(listingUniqueOwnerTimeRanges, listingUniqueOwnerTimeRange)
	}
	return listingUniqueOwnerTimeRanges, nil
}

func (s *store) GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (marketplaceCollectionSnapshotFloorPriceTimeRanges []*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange, err error) {
	rows, err := s.db.Raw(
		`
		SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(COALESCE(AVG(nmcs.floor_price::DOUBLE PRECISION), 0), 'FM9999999999999999999999999999') AS floor_price
		FROM
			generate_records_by_time_range (@timeDuration,
				@timeInterval,
				@timeRoundUp) tr
		LEFT JOIN token t ON t.symbol = @tokenSymbol
		LEFT JOIN nft_marketplace_collection_snapshot nmcs ON nmcs.created_time < tr.t
			AND nmcs.created_time > (tr.t - (@timeInterval)::INTERVAL)
			AND nmcs.collection_address = @contractAddress
			AND nmcs.token_id = t.id
		GROUP BY
			(tr.t,
				t.symbol,
				t.is_native,
				t.address,
				t.decimals,
				t.icon_url)
		ORDER BY
			tr.t ASC;
	`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		marketplaceCollectionSnapshotFloorPriceTimeRange := &model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange{}
		tokenRaw := []byte{}
		floorPrice := &model.Price{}
		err = rows.Scan(
			&marketplaceCollectionSnapshotFloorPriceTimeRange.TimeRangeStart,
			&tokenRaw,
			&floorPrice.Amount,
		)
		marketplaceCollectionSnapshotFloorPriceTimeRange.TimeInterval = timeInterval

		token := &model.Token{}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &token); err != nil {
				return nil, err
			}
		}
		floorPrice.Token = *token
		marketplaceCollectionSnapshotFloorPriceTimeRange.FloorPriceObj = floorPrice

		marketplaceCollectionSnapshotFloorPriceTimeRanges = append(marketplaceCollectionSnapshotFloorPriceTimeRanges, marketplaceCollectionSnapshotFloorPriceTimeRange)
	}
	return marketplaceCollectionSnapshotFloorPriceTimeRanges, nil
}

func (s *store) GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (marketplaceCollectionSnapshotMarketCapTimeRanges []*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange, err error) {
	rows, err := s.db.Raw(
		`SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(nmcs.total_volume::DOUBLE PRECISION * COUNT(nt), 'FM9999999999999999999999999999') AS market_cap
		FROM
			generate_records_by_time_range (@timeDuration,
					@timeInterval,
					@timeRoundUp) tr
			LEFT JOIN token t ON t.symbol = @tokenSymbol
			LEFT JOIN nft_token nt ON nt.collection_address = @contractAddress
			JOIN LATERAL (
					SELECT
							nmcs.total_volume
					FROM
							nft_marketplace_collection_snapshot nmcs
					WHERE
							nmcs.collection_address = @contractAddress
							AND nmcs.created_time < tr.t
							AND nmcs.token_id = t.id
					ORDER BY
							nmcs.created_time DESC
					LIMIT 1) nmcs ON TRUE
		GROUP BY
			(tr.t,
					nmcs.total_volume,
					t.symbol,
					t.is_native,
					t.address,
					t.decimals,
					t.icon_url)
		ORDER BY
			tr.t ASC;`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		marketplaceCollectionSnapshotMarketCapTimeRange := &model.NftMarketplaceCollectionSnapshotMarketCapTimeRange{}
		tokenRaw := []byte{}
		marketCap := &model.Price{}
		err = rows.Scan(
			&marketplaceCollectionSnapshotMarketCapTimeRange.TimeRangeStart,
			&tokenRaw,
			&marketCap.Amount,
		)
		marketplaceCollectionSnapshotMarketCapTimeRange.TimeInterval = timeInterval

		token := &model.Token{}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &token); err != nil {
				return nil, err
			}
		}
		marketCap.Token = *token
		marketplaceCollectionSnapshotMarketCapTimeRange.MarketCapObj = marketCap

		marketplaceCollectionSnapshotMarketCapTimeRanges = append(marketplaceCollectionSnapshotMarketCapTimeRanges, marketplaceCollectionSnapshotMarketCapTimeRange)
	}
	return marketplaceCollectionSnapshotMarketCapTimeRanges, nil
}

func (s *store) GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) (listingTotalVolumeTimeRanges []*model.NftListingTotalVolumeTimeRange, err error) {
	rows, err := s.db.Raw(
		`SELECT
			t AS time_range,
			json_build_object('symbol', t.symbol, 'is_native', t.is_native, 'address', t.address, 'decimals', t.decimals, 'icon_url', t.icon_url) AS token,
			to_char(COALESCE(SUM(nl.sold_price::DOUBLE PRECISION), 0), 'FM9999999999999999999999999999') AS total_volume
		FROM
			generate_records_by_time_range (@timeDuration,
				@timeInterval,
				@timeRoundUp) tr
			LEFT JOIN token t ON t.symbol = @tokenSymbol
			LEFT JOIN nft_listing nl ON nl.created_time < tr.t
				AND nl.created_time > (tr.t - (@timeInterval)::INTERVAL)
				AND contract_address = @contractAddress
				AND nl.listing_status = 'sold'
				AND nl.payment_token = t.id
			LEFT JOIN nft_token nt ON nt.collection_address = @contractAddress
		GROUP BY
			(tr.t,
				t.symbol,
				t.is_native,
				t.address,
				t.decimals,
				t.icon_url)
		ORDER BY
			tr.t ASC;`, map[string]interface{}{
			"timeDuration":    timeDuration,
			"timeInterval":    timeInterval,
			"timeRoundUp":     timeRoundUp,
			"tokenSymbol":     paymentToken,
			"contractAddress": collectionAddress,
		}).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		listingTotalVolumeTimeRange := &model.NftListingTotalVolumeTimeRange{}
		tokenRaw := []byte{}
		totalVolume := &model.Price{}
		err = rows.Scan(
			&listingTotalVolumeTimeRange.TimeRangeStart,
			&tokenRaw,
			&totalVolume.Amount,
		)
		listingTotalVolumeTimeRange.TimeInterval = timeInterval

		token := &model.Token{}

		if tokenRaw != nil {
			if err := json.Unmarshal(tokenRaw, &token); err != nil {
				return nil, err
			}
		}
		totalVolume.Token = *token
		listingTotalVolumeTimeRange.TotalVolumeObj = totalVolume

		listingTotalVolumeTimeRanges = append(listingTotalVolumeTimeRanges, listingTotalVolumeTimeRange)
	}
	return listingTotalVolumeTimeRanges, nil
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
			nft_listing.contract_address = $1
		LIMIT 1;
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
