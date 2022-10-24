// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/store/nft/interface.go

// Package mock_nft is a generated GoMock package.
package mock_nft

import (
	reflect "reflect"

	model "github.com/consolelabs/indexer-api/pkg/model"
	nft "github.com/consolelabs/indexer-api/pkg/store/nft"
	gomock "github.com/golang/mock/gomock"
)

// MockINft is a mock of INft interface.
type MockINft struct {
	ctrl     *gomock.Controller
	recorder *MockINftMockRecorder
}

// MockINftMockRecorder is the mock recorder for MockINft.
type MockINftMockRecorder struct {
	mock *MockINft
}

// NewMockINft creates a new mock instance.
func NewMockINft(ctrl *gomock.Controller) *MockINft {
	mock := &MockINft{ctrl: ctrl}
	mock.recorder = &MockINftMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINft) EXPECT() *MockINftMockRecorder {
	return m.recorder
}

// AttributeUpsertOne mocks base method.
func (m *MockINft) AttributeUpsertOne(tokenAttribute model.NftTokenAttribute) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AttributeUpsertOne", tokenAttribute)
	ret0, _ := ret[0].(error)
	return ret0
}

// AttributeUpsertOne indicates an expected call of AttributeUpsertOne.
func (mr *MockINftMockRecorder) AttributeUpsertOne(tokenAttribute interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AttributeUpsertOne", reflect.TypeOf((*MockINft)(nil).AttributeUpsertOne), tokenAttribute)
}

// DeleteNftOwnerByCollectionAddress mocks base method.
func (m *MockINft) DeleteNftOwnerByCollectionAddress(collecionAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNftOwnerByCollectionAddress", collecionAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNftOwnerByCollectionAddress indicates an expected call of DeleteNftOwnerByCollectionAddress.
func (mr *MockINftMockRecorder) DeleteNftOwnerByCollectionAddress(collecionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNftOwnerByCollectionAddress", reflect.TypeOf((*MockINft)(nil).DeleteNftOwnerByCollectionAddress), collecionAddress)
}

// DeleteNftToken mocks base method.
func (m *MockINft) DeleteNftToken(query nft.NftTokenQuery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNftToken", query)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNftToken indicates an expected call of DeleteNftToken.
func (mr *MockINftMockRecorder) DeleteNftToken(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNftToken", reflect.TypeOf((*MockINft)(nil).DeleteNftToken), query)
}

// DeleteNftTokenAttributesByCollectionAddress mocks base method.
func (m *MockINft) DeleteNftTokenAttributesByCollectionAddress(collecionAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNftTokenAttributesByCollectionAddress", collecionAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNftTokenAttributesByCollectionAddress indicates an expected call of DeleteNftTokenAttributesByCollectionAddress.
func (mr *MockINftMockRecorder) DeleteNftTokenAttributesByCollectionAddress(collecionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNftTokenAttributesByCollectionAddress", reflect.TypeOf((*MockINft)(nil).DeleteNftTokenAttributesByCollectionAddress), collecionAddress)
}

// DeleteNftTransferByContractAddress mocks base method.
func (m *MockINft) DeleteNftTransferByContractAddress(collecionAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNftTransferByContractAddress", collecionAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNftTransferByContractAddress indicates an expected call of DeleteNftTransferByContractAddress.
func (mr *MockINftMockRecorder) DeleteNftTransferByContractAddress(collecionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNftTransferByContractAddress", reflect.TypeOf((*MockINft)(nil).DeleteNftTransferByContractAddress), collecionAddress)
}

// DeleteOwnerByCollectionAddressTokenId mocks base method.
func (m *MockINft) DeleteOwnerByCollectionAddressTokenId(collectionAddress, tokenId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOwnerByCollectionAddressTokenId", collectionAddress, tokenId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOwnerByCollectionAddressTokenId indicates an expected call of DeleteOwnerByCollectionAddressTokenId.
func (mr *MockINftMockRecorder) DeleteOwnerByCollectionAddressTokenId(collectionAddress, tokenId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOwnerByCollectionAddressTokenId", reflect.TypeOf((*MockINft)(nil).DeleteOwnerByCollectionAddressTokenId), collectionAddress, tokenId)
}

// GetAttributeByCollectionAddressTokenID mocks base method.
func (m *MockINft) GetAttributeByCollectionAddressTokenID(collectionAddress, tokenID string) ([]model.NftTokenAttribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttributeByCollectionAddressTokenID", collectionAddress, tokenID)
	ret0, _ := ret[0].([]model.NftTokenAttribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttributeByCollectionAddressTokenID indicates an expected call of GetAttributeByCollectionAddressTokenID.
func (mr *MockINftMockRecorder) GetAttributeByCollectionAddressTokenID(collectionAddress, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributeByCollectionAddressTokenID", reflect.TypeOf((*MockINft)(nil).GetAttributeByCollectionAddressTokenID), collectionAddress, tokenID)
}

// GetAttributesByCollectionAddress mocks base method.
func (m *MockINft) GetAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAttributesByCollectionAddress", collectionAddress)
	ret0, _ := ret[0].([]model.NftTokenAttribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAttributesByCollectionAddress indicates an expected call of GetAttributesByCollectionAddress.
func (mr *MockINftMockRecorder) GetAttributesByCollectionAddress(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAttributesByCollectionAddress", reflect.TypeOf((*MockINft)(nil).GetAttributesByCollectionAddress), collectionAddress)
}

// GetCollectionByAddress mocks base method.
func (m *MockINft) GetCollectionByAddress(address string) (*model.NftCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionByAddress", address)
	ret0, _ := ret[0].(*model.NftCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionByAddress indicates an expected call of GetCollectionByAddress.
func (mr *MockINftMockRecorder) GetCollectionByAddress(address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionByAddress", reflect.TypeOf((*MockINft)(nil).GetCollectionByAddress), address)
}

// GetCollectionSaleVolumeAndFloorPrice mocks base method.
func (m *MockINft) GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration string) ([]*model.NftCollectionSaleVolumeAndFloorPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionSaleVolumeAndFloorPrice", collectionAddress, timeDuration)
	ret0, _ := ret[0].([]*model.NftCollectionSaleVolumeAndFloorPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionSaleVolumeAndFloorPrice indicates an expected call of GetCollectionSaleVolumeAndFloorPrice.
func (mr *MockINftMockRecorder) GetCollectionSaleVolumeAndFloorPrice(collectionAddress, timeDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionSaleVolumeAndFloorPrice", reflect.TypeOf((*MockINft)(nil).GetCollectionSaleVolumeAndFloorPrice), collectionAddress, timeDuration)
}

// GetCollectionTraitsCount mocks base method.
func (m *MockINft) GetCollectionTraitsCount(collectionAddress string) ([]model.NftTokenAttributeCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionTraitsCount", collectionAddress)
	ret0, _ := ret[0].([]model.NftTokenAttributeCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCollectionTraitsCount indicates an expected call of GetCollectionTraitsCount.
func (mr *MockINftMockRecorder) GetCollectionTraitsCount(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionTraitsCount", reflect.TypeOf((*MockINft)(nil).GetCollectionTraitsCount), collectionAddress)
}

// GetCollections mocks base method.
func (m *MockINft) GetCollections(query nft.NftCollectionQuery) ([]model.NftCollection, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollections", query)
	ret0, _ := ret[0].([]model.NftCollection)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCollections indicates an expected call of GetCollections.
func (mr *MockINftMockRecorder) GetCollections(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollections", reflect.TypeOf((*MockINft)(nil).GetCollections), query)
}

// GetCollectionsByWalletAddress mocks base method.
func (m *MockINft) GetCollectionsByWalletAddress(query nft.WalletCollectionQuery) ([]model.NftCollection, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCollectionsByWalletAddress", query)
	ret0, _ := ret[0].([]model.NftCollection)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCollectionsByWalletAddress indicates an expected call of GetCollectionsByWalletAddress.
func (mr *MockINftMockRecorder) GetCollectionsByWalletAddress(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCollectionsByWalletAddress", reflect.TypeOf((*MockINft)(nil).GetCollectionsByWalletAddress), query)
}

// GetDistinctAttributesByCollectionAddress mocks base method.
func (m *MockINft) GetDistinctAttributesByCollectionAddress(collectionAddress string) ([]model.NftTokenAttribute, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDistinctAttributesByCollectionAddress", collectionAddress)
	ret0, _ := ret[0].([]model.NftTokenAttribute)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDistinctAttributesByCollectionAddress indicates an expected call of GetDistinctAttributesByCollectionAddress.
func (mr *MockINftMockRecorder) GetDistinctAttributesByCollectionAddress(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDistinctAttributesByCollectionAddress", reflect.TypeOf((*MockINft)(nil).GetDistinctAttributesByCollectionAddress), collectionAddress)
}

// GetListingAcrossPriceRanges mocks base method.
func (m *MockINft) GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange string) ([]*model.NftListingPriceRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListingAcrossPriceRanges", collectionAddress, paymentToken, priceRange)
	ret0, _ := ret[0].([]*model.NftListingPriceRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListingAcrossPriceRanges indicates an expected call of GetListingAcrossPriceRanges.
func (mr *MockINftMockRecorder) GetListingAcrossPriceRanges(collectionAddress, paymentToken, priceRange interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListingAcrossPriceRanges", reflect.TypeOf((*MockINft)(nil).GetListingAcrossPriceRanges), collectionAddress, paymentToken, priceRange)
}

// GetListingAcrossRarityRankRanges mocks base method.
func (m *MockINft) GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange string) ([]*model.NftListingRarityRankRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListingAcrossRarityRankRanges", collectionAddress, paymentToken, rarityRankRange)
	ret0, _ := ret[0].([]*model.NftListingRarityRankRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListingAcrossRarityRankRanges indicates an expected call of GetListingAcrossRarityRankRanges.
func (mr *MockINftMockRecorder) GetListingAcrossRarityRankRanges(collectionAddress, paymentToken, rarityRankRange interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListingAcrossRarityRankRanges", reflect.TypeOf((*MockINft)(nil).GetListingAcrossRarityRankRanges), collectionAddress, paymentToken, rarityRankRange)
}

// GetListingMarketCapAndVolumeByTimeRanges mocks base method.
func (m *MockINft) GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingMarketCapAndVolumeTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListingMarketCapAndVolumeByTimeRanges", collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftListingMarketCapAndVolumeTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListingMarketCapAndVolumeByTimeRanges indicates an expected call of GetListingMarketCapAndVolumeByTimeRanges.
func (mr *MockINftMockRecorder) GetListingMarketCapAndVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListingMarketCapAndVolumeByTimeRanges", reflect.TypeOf((*MockINft)(nil).GetListingMarketCapAndVolumeByTimeRanges), collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetListingTotalVolumeByTimeRanges mocks base method.
func (m *MockINft) GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingTotalVolumeTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListingTotalVolumeByTimeRanges", collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftListingTotalVolumeTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListingTotalVolumeByTimeRanges indicates an expected call of GetListingTotalVolumeByTimeRanges.
func (mr *MockINftMockRecorder) GetListingTotalVolumeByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListingTotalVolumeByTimeRanges", reflect.TypeOf((*MockINft)(nil).GetListingTotalVolumeByTimeRanges), collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetListingUniqueOwnersByTimeRange mocks base method.
func (m *MockINft) GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingUniqueOwnersTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListingUniqueOwnersByTimeRange", collectionAddress, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftListingUniqueOwnersTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListingUniqueOwnersByTimeRange indicates an expected call of GetListingUniqueOwnersByTimeRange.
func (mr *MockINftMockRecorder) GetListingUniqueOwnersByTimeRange(collectionAddress, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListingUniqueOwnersByTimeRange", reflect.TypeOf((*MockINft)(nil).GetListingUniqueOwnersByTimeRange), collectionAddress, timeDuration, timeInterval, timeRoundUp)
}

// GetMarketSnapshotFloorPriceByTimeRanges mocks base method.
func (m *MockINft) GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketSnapshotFloorPriceByTimeRanges", collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftMarketplaceCollectionSnapshotFloorPriceTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketSnapshotFloorPriceByTimeRanges indicates an expected call of GetMarketSnapshotFloorPriceByTimeRanges.
func (mr *MockINftMockRecorder) GetMarketSnapshotFloorPriceByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketSnapshotFloorPriceByTimeRanges", reflect.TypeOf((*MockINft)(nil).GetMarketSnapshotFloorPriceByTimeRanges), collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetMarketSnapshotMarketCapByTimeRanges mocks base method.
func (m *MockINft) GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketSnapshotMarketCapByTimeRanges", collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftMarketplaceCollectionSnapshotMarketCapTimeRange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketSnapshotMarketCapByTimeRanges indicates an expected call of GetMarketSnapshotMarketCapByTimeRanges.
func (mr *MockINftMockRecorder) GetMarketSnapshotMarketCapByTimeRanges(collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketSnapshotMarketCapByTimeRanges", reflect.TypeOf((*MockINft)(nil).GetMarketSnapshotMarketCapByTimeRanges), collectionAddress, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetMarketplaceWithListingStats mocks base method.
func (m *MockINft) GetMarketplaceWithListingStats(collectionAddress string) ([]model.MarketplaceWithListingStats, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarketplaceWithListingStats", collectionAddress)
	ret0, _ := ret[0].([]model.MarketplaceWithListingStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarketplaceWithListingStats indicates an expected call of GetMarketplaceWithListingStats.
func (mr *MockINftMockRecorder) GetMarketplaceWithListingStats(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarketplaceWithListingStats", reflect.TypeOf((*MockINft)(nil).GetMarketplaceWithListingStats), collectionAddress)
}

// GetNftListing mocks base method.
func (m *MockINft) GetNftListing(query nft.NftListingQuery) ([]*model.NftListing, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftListing", query)
	ret0, _ := ret[0].([]*model.NftListing)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNftListing indicates an expected call of GetNftListing.
func (mr *MockINftMockRecorder) GetNftListing(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftListing", reflect.TypeOf((*MockINft)(nil).GetNftListing), query)
}

// GetNftListingAvgPricesAndFloor mocks base method.
func (m *MockINft) GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingAvgPriceAndFloor, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftListingAvgPricesAndFloor", collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftListingAvgPriceAndFloor)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftListingAvgPricesAndFloor indicates an expected call of GetNftListingAvgPricesAndFloor.
func (mr *MockINftMockRecorder) GetNftListingAvgPricesAndFloor(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftListingAvgPricesAndFloor", reflect.TypeOf((*MockINft)(nil).GetNftListingAvgPricesAndFloor), collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetNftListingByTokenID mocks base method.
func (m *MockINft) GetNftListingByTokenID(collectionAddress, tokenID string) ([]model.NftListingMarketplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftListingByTokenID", collectionAddress, tokenID)
	ret0, _ := ret[0].([]model.NftListingMarketplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftListingByTokenID indicates an expected call of GetNftListingByTokenID.
func (mr *MockINftMockRecorder) GetNftListingByTokenID(collectionAddress, tokenID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftListingByTokenID", reflect.TypeOf((*MockINft)(nil).GetNftListingByTokenID), collectionAddress, tokenID)
}

// GetNftListingSoldPrices mocks base method.
func (m *MockINft) GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp string) ([]*model.NftListingSoldPrice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftListingSoldPrices", collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
	ret0, _ := ret[0].([]*model.NftListingSoldPrice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftListingSoldPrices indicates an expected call of GetNftListingSoldPrices.
func (mr *MockINftMockRecorder) GetNftListingSoldPrices(collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftListingSoldPrices", reflect.TypeOf((*MockINft)(nil).GetNftListingSoldPrices), collectionAddress, tokenId, paymentToken, timeDuration, timeInterval, timeRoundUp)
}

// GetNftMarketplaceCollection mocks base method.
func (m *MockINft) GetNftMarketplaceCollection(collectionAddress string) ([]model.NftListingMarketplace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftMarketplaceCollection", collectionAddress)
	ret0, _ := ret[0].([]model.NftListingMarketplace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftMarketplaceCollection indicates an expected call of GetNftMarketplaceCollection.
func (mr *MockINftMockRecorder) GetNftMarketplaceCollection(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftMarketplaceCollection", reflect.TypeOf((*MockINft)(nil).GetNftMarketplaceCollection), collectionAddress)
}

// GetNftMarketplaceCollectionSnapshots mocks base method.
func (m *MockINft) GetNftMarketplaceCollectionSnapshots(query nft.NftTickerQuery) ([]model.NftMarketplaceCollectionSnapshot, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftMarketplaceCollectionSnapshots", query)
	ret0, _ := ret[0].([]model.NftMarketplaceCollectionSnapshot)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftMarketplaceCollectionSnapshots indicates an expected call of GetNftMarketplaceCollectionSnapshots.
func (mr *MockINftMockRecorder) GetNftMarketplaceCollectionSnapshots(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftMarketplaceCollectionSnapshots", reflect.TypeOf((*MockINft)(nil).GetNftMarketplaceCollectionSnapshots), query)
}

// GetNftMetadataAttributesIcon mocks base method.
func (m *MockINft) GetNftMetadataAttributesIcon(query nft.NftAttributeIconQuery) ([]model.NftMetadataAttributesIcon, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftMetadataAttributesIcon", query)
	ret0, _ := ret[0].([]model.NftMetadataAttributesIcon)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNftMetadataAttributesIcon indicates an expected call of GetNftMetadataAttributesIcon.
func (mr *MockINftMockRecorder) GetNftMetadataAttributesIcon(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftMetadataAttributesIcon", reflect.TypeOf((*MockINft)(nil).GetNftMetadataAttributesIcon), query)
}

// GetNftOwners mocks base method.
func (m *MockINft) GetNftOwners(query nft.NftOwnerQuery) ([]model.NftOwner, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftOwners", query)
	ret0, _ := ret[0].([]model.NftOwner)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNftOwners indicates an expected call of GetNftOwners.
func (mr *MockINftMockRecorder) GetNftOwners(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftOwners", reflect.TypeOf((*MockINft)(nil).GetNftOwners), query)
}

// GetNftTokenTxHistory mocks base method.
func (m *MockINft) GetNftTokenTxHistory(collectionAddress, tokenId string) ([]model.NftTxHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftTokenTxHistory", collectionAddress, tokenId)
	ret0, _ := ret[0].([]model.NftTxHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNftTokenTxHistory indicates an expected call of GetNftTokenTxHistory.
func (mr *MockINftMockRecorder) GetNftTokenTxHistory(collectionAddress, tokenId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftTokenTxHistory", reflect.TypeOf((*MockINft)(nil).GetNftTokenTxHistory), collectionAddress, tokenId)
}

// GetNftTokens mocks base method.
func (m *MockINft) GetNftTokens(query nft.NftTokenQuery) ([]model.NftToken, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftTokens", query)
	ret0, _ := ret[0].([]model.NftToken)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNftTokens indicates an expected call of GetNftTokens.
func (mr *MockINftMockRecorder) GetNftTokens(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftTokens", reflect.TypeOf((*MockINft)(nil).GetNftTokens), query)
}

// GetNftTransfers mocks base method.
func (m *MockINft) GetNftTransfers(query nft.NftTransferQuery) ([]model.NftTransfer, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNftTransfers", query)
	ret0, _ := ret[0].([]model.NftTransfer)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNftTransfers indicates an expected call of GetNftTransfers.
func (mr *MockINftMockRecorder) GetNftTransfers(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNftTransfers", reflect.TypeOf((*MockINft)(nil).GetNftTransfers), query)
}

// GetPlatformByName mocks base method.
func (m *MockINft) GetPlatformByName(name string) (*model.MarketplacePlatform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatformByName", name)
	ret0, _ := ret[0].(*model.MarketplacePlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatformByName indicates an expected call of GetPlatformByName.
func (mr *MockINftMockRecorder) GetPlatformByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatformByName", reflect.TypeOf((*MockINft)(nil).GetPlatformByName), name)
}

// GetPlatformsByCollectionAddress mocks base method.
func (m *MockINft) GetPlatformsByCollectionAddress(collectionAddress string) ([]model.MarketplacePlatform, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlatformsByCollectionAddress", collectionAddress)
	ret0, _ := ret[0].([]model.MarketplacePlatform)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlatformsByCollectionAddress indicates an expected call of GetPlatformsByCollectionAddress.
func (mr *MockINftMockRecorder) GetPlatformsByCollectionAddress(collectionAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlatformsByCollectionAddress", reflect.TypeOf((*MockINft)(nil).GetPlatformsByCollectionAddress), collectionAddress)
}

// GetTokensByWalletAddress mocks base method.
func (m *MockINft) GetTokensByWalletAddress(query nft.WalletTokenQuery) ([]model.NftToken, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTokensByWalletAddress", query)
	ret0, _ := ret[0].([]model.NftToken)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetTokensByWalletAddress indicates an expected call of GetTokensByWalletAddress.
func (mr *MockINftMockRecorder) GetTokensByWalletAddress(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTokensByWalletAddress", reflect.TypeOf((*MockINft)(nil).GetTokensByWalletAddress), query)
}

// RefreshViewNFTCollectionAttributes mocks base method.
func (m *MockINft) RefreshViewNFTCollectionAttributes() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshViewNFTCollectionAttributes")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshViewNFTCollectionAttributes indicates an expected call of RefreshViewNFTCollectionAttributes.
func (mr *MockINftMockRecorder) RefreshViewNFTCollectionAttributes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshViewNFTCollectionAttributes", reflect.TypeOf((*MockINft)(nil).RefreshViewNFTCollectionAttributes))
}

// RefreshViewNFTCollectionStats mocks base method.
func (m *MockINft) RefreshViewNFTCollectionStats() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshViewNFTCollectionStats")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshViewNFTCollectionStats indicates an expected call of RefreshViewNFTCollectionStats.
func (mr *MockINftMockRecorder) RefreshViewNFTCollectionStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshViewNFTCollectionStats", reflect.TypeOf((*MockINft)(nil).RefreshViewNFTCollectionStats))
}

// RefreshViewNFTCollections mocks base method.
func (m *MockINft) RefreshViewNFTCollections() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshViewNFTCollections")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshViewNFTCollections indicates an expected call of RefreshViewNFTCollections.
func (mr *MockINftMockRecorder) RefreshViewNFTCollections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshViewNFTCollections", reflect.TypeOf((*MockINft)(nil).RefreshViewNFTCollections))
}

// RefreshViewNFTTokens mocks base method.
func (m *MockINft) RefreshViewNFTTokens() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshViewNFTTokens")
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshViewNFTTokens indicates an expected call of RefreshViewNFTTokens.
func (mr *MockINftMockRecorder) RefreshViewNFTTokens() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshViewNFTTokens", reflect.TypeOf((*MockINft)(nil).RefreshViewNFTTokens))
}

// SaveListing mocks base method.
func (m *MockINft) SaveListing(listing *model.NftListing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveListing", listing)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveListing indicates an expected call of SaveListing.
func (mr *MockINftMockRecorder) SaveListing(listing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveListing", reflect.TypeOf((*MockINft)(nil).SaveListing), listing)
}

// SaveNftCollection mocks base method.
func (m *MockINft) SaveNftCollection(nftCollection *model.NftCollection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNftCollection", nftCollection)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNftCollection indicates an expected call of SaveNftCollection.
func (mr *MockINftMockRecorder) SaveNftCollection(nftCollection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNftCollection", reflect.TypeOf((*MockINft)(nil).SaveNftCollection), nftCollection)
}

// SaveOwner mocks base method.
func (m *MockINft) SaveOwner(owner *model.NftOwner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveOwner", owner)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveOwner indicates an expected call of SaveOwner.
func (mr *MockINftMockRecorder) SaveOwner(owner interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveOwner", reflect.TypeOf((*MockINft)(nil).SaveOwner), owner)
}

// SaveTransfer mocks base method.
func (m *MockINft) SaveTransfer(transfer *model.NftTransfer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTransfer", transfer)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTransfer indicates an expected call of SaveTransfer.
func (mr *MockINftMockRecorder) SaveTransfer(transfer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTransfer", reflect.TypeOf((*MockINft)(nil).SaveTransfer), transfer)
}

// UpdateAttributeCount mocks base method.
func (m *MockINft) UpdateAttributeCount(count uint64, address, traitType, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAttributeCount", count, address, traitType, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAttributeCount indicates an expected call of UpdateAttributeCount.
func (mr *MockINftMockRecorder) UpdateAttributeCount(count, address, traitType, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAttributeCount", reflect.TypeOf((*MockINft)(nil).UpdateAttributeCount), count, address, traitType, value)
}

// UpdateCollection mocks base method.
func (m *MockINft) UpdateCollection(collection *model.NftCollection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCollection", collection)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCollection indicates an expected call of UpdateCollection.
func (mr *MockINftMockRecorder) UpdateCollection(collection interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCollection", reflect.TypeOf((*MockINft)(nil).UpdateCollection), collection)
}

// UpdateListing mocks base method.
func (m *MockINft) UpdateListing(listing *model.NftListing) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateListing", listing)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateListing indicates an expected call of UpdateListing.
func (mr *MockINftMockRecorder) UpdateListing(listing interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateListing", reflect.TypeOf((*MockINft)(nil).UpdateListing), listing)
}

// UpdateOwnerByCollectionAddressTokenId mocks base method.
func (m *MockINft) UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId, ownerAddress string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOwnerByCollectionAddressTokenId", collectionAddress, tokenId, ownerAddress)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOwnerByCollectionAddressTokenId indicates an expected call of UpdateOwnerByCollectionAddressTokenId.
func (mr *MockINftMockRecorder) UpdateOwnerByCollectionAddressTokenId(collectionAddress, tokenId, ownerAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOwnerByCollectionAddressTokenId", reflect.TypeOf((*MockINft)(nil).UpdateOwnerByCollectionAddressTokenId), collectionAddress, tokenId, ownerAddress)
}

// UpsertNftMarketplaceCollection mocks base method.
func (m *MockINft) UpsertNftMarketplaceCollection(model *model.NftMarketplaceCollection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNftMarketplaceCollection", model)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNftMarketplaceCollection indicates an expected call of UpsertNftMarketplaceCollection.
func (mr *MockINftMockRecorder) UpsertNftMarketplaceCollection(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNftMarketplaceCollection", reflect.TypeOf((*MockINft)(nil).UpsertNftMarketplaceCollection), model)
}

// UpsertToken mocks base method.
func (m *MockINft) UpsertToken(token *model.NftToken, upsertAttributes bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertToken", token, upsertAttributes)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertToken indicates an expected call of UpsertToken.
func (mr *MockINftMockRecorder) UpsertToken(token, upsertAttributes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertToken", reflect.TypeOf((*MockINft)(nil).UpsertToken), token, upsertAttributes)
}
