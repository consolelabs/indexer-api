// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package paintswap

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// PaintSwapMarketplaceV2nftDetails is an auto generated low-level Go binding around an user-defined struct.
type PaintSwapMarketplaceV2nftDetails struct {
	Nfts               []common.Address
	TokenIds           []*big.Int
	AmountBatches      []*big.Int
	Seller             common.Address
	Price              *big.Int
	StartTime          *big.Int
	EndTime            *big.Int
	MaxBidOrOffer      *big.Int
	MaxBidderOrOfferer common.Address
	IsAuction          bool
	Amount             *big.Int
	AmountRemaining    *big.Int
	PaymentToken       common.Address
	Complete           bool
	DevFeePercentage   *big.Int
}

// PaintswapMetaData contains all meta data concerning the Paintswap contract.
var PaintswapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_officialNFTDiscount\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_marketplaceWhitelist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_collectionRoyalties\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokens\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_paintRouter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_wftm\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"}],\"name\":\"CancelledSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"bid\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextMinimum\",\"type\":\"uint256\"}],\"name\":\"NewBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"offer\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nextMinimum\",\"type\":\"uint256\"}],\"name\":\"NewOffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isAuction\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isNSFW\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"marketplaceURI\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"searchKeywords\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"NewSale\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"failedSellAll\",\"type\":\"bool\"}],\"name\":\"SaleFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sold\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"UpdateEndTime\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"UpdatePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marketplaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"UpdateStartTime\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"acceptBestOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOfferAfterDeadlineEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amountBatches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"},{\"internalType\":\"uint256[]\",\"name\":\"_times\",\"type\":\"uint256[]\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"bool[]\",\"name\":\"_flags\",\"type\":\"bool[]\"},{\"internalType\":\"string\",\"name\":\"_marketplaceURI\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_searchKeywords\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"addSale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"addWhitelistedSellerBuyerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionEndTimeIncreaseCutOff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bundlesEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_amount\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cancelCutoff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"cancelSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"completeMarketplaceEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentMarketplaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devFeePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"devSalesFeeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"}],\"name\":\"editPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTimeCutOff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"excessNFTAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"excessTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"getSaleDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"nfts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountBatches\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"maxBidOrOffer\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"maxBidderOrOfferer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isAuction\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountRemaining\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"devFeePercentage\",\"type\":\"uint256\"}],\"internalType\":\"structPaintSwapMarketplaceV2.nftDetails\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSellingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_bid\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"makeBid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"_offer\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"makeOffer\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceWhitelist\",\"outputs\":[{\"internalType\":\"contractMarketplaceWhitelist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxBundleNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStartTimeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minIncreasedBidOfferPercent\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modifyInitialStartTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_marketplaceId\",\"type\":\"uint256\"}],\"name\":\"nextMinimumBidOrOffer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"officialNFTDiscount\",\"outputs\":[{\"internalType\":\"contractOfficialNFTDiscount\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onlySellerAndHighestOfferCanEndSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"removeWhitelistedSellerBuyerContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRoyalty\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collectionRoyalties\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleGracePeriodForCancelling\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setAcceptOfferAfterDeadlineEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_auctionEndTimeIncreaseCutOff\",\"type\":\"uint256\"}],\"name\":\"setAuctionEndTimeIncreaseCutOff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setBundlesSupported\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_callGasLimit\",\"type\":\"uint256\"}],\"name\":\"setCallGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_cancelCutoff\",\"type\":\"uint256\"}],\"name\":\"setCancelCutoff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dev\",\"type\":\"address\"}],\"name\":\"setDevAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_feePercentage\",\"type\":\"uint16\"}],\"name\":\"setDevFeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dev\",\"type\":\"address\"}],\"name\":\"setDevSalesAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setEnableSelling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_endTimeCutoff\",\"type\":\"uint256\"}],\"name\":\"setEndTimeCutoff\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_maxBundleNumber\",\"type\":\"uint128\"}],\"name\":\"setMaxBundleNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxDuration\",\"type\":\"uint256\"}],\"name\":\"setMaxDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxRoyalty\",\"type\":\"uint256\"}],\"name\":\"setMaxRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxStartTimeIncrement\",\"type\":\"uint256\"}],\"name\":\"setMaxStartTimeIncrement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_minIncreasedBidOfferPercent\",\"type\":\"uint128\"}],\"name\":\"setMinIncreasedBidOfferPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDuration\",\"type\":\"uint256\"}],\"name\":\"setMinimumDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_modifyInitialStartTime\",\"type\":\"bool\"}],\"name\":\"setModifyInitialStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_officialNFTDiscount\",\"type\":\"address\"}],\"name\":\"setOfficialNFTDiscount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setOnlySellerAndHighestOfferCanEndSale\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paintRouter\",\"type\":\"address\"}],\"name\":\"setRouter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_saleGracePeriodForCancellingInSeconds\",\"type\":\"uint256\"}],\"name\":\"setSaleGracePeriodForCancelling\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_enable\",\"type\":\"bool\"}],\"name\":\"setSlidingEndTimeForBidsEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slidingEndTimeForBidsEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"contractTokensV2\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelistedSellerBuyerContract\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawExcessNFTs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"withdrawExcessTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// PaintswapABI is the input ABI used to generate the binding from.
// Deprecated: Use PaintswapMetaData.ABI instead.
var PaintswapABI = PaintswapMetaData.ABI

// Paintswap is an auto generated Go binding around an Ethereum contract.
type Paintswap struct {
	PaintswapCaller     // Read-only binding to the contract
	PaintswapTransactor // Write-only binding to the contract
	PaintswapFilterer   // Log filterer for contract events
}

// PaintswapCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaintswapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaintswapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaintswapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaintswapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaintswapSession struct {
	Contract     *Paintswap        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaintswapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaintswapCallerSession struct {
	Contract *PaintswapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PaintswapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaintswapTransactorSession struct {
	Contract     *PaintswapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PaintswapRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaintswapRaw struct {
	Contract *Paintswap // Generic contract binding to access the raw methods on
}

// PaintswapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaintswapCallerRaw struct {
	Contract *PaintswapCaller // Generic read-only contract binding to access the raw methods on
}

// PaintswapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaintswapTransactorRaw struct {
	Contract *PaintswapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaintswap creates a new instance of Paintswap, bound to a specific deployed contract.
func NewPaintswap(address common.Address, backend bind.ContractBackend) (*Paintswap, error) {
	contract, err := bindPaintswap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Paintswap{PaintswapCaller: PaintswapCaller{contract: contract}, PaintswapTransactor: PaintswapTransactor{contract: contract}, PaintswapFilterer: PaintswapFilterer{contract: contract}}, nil
}

// NewPaintswapCaller creates a new read-only instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapCaller(address common.Address, caller bind.ContractCaller) (*PaintswapCaller, error) {
	contract, err := bindPaintswap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaintswapCaller{contract: contract}, nil
}

// NewPaintswapTransactor creates a new write-only instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapTransactor(address common.Address, transactor bind.ContractTransactor) (*PaintswapTransactor, error) {
	contract, err := bindPaintswap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaintswapTransactor{contract: contract}, nil
}

// NewPaintswapFilterer creates a new log filterer instance of Paintswap, bound to a specific deployed contract.
func NewPaintswapFilterer(address common.Address, filterer bind.ContractFilterer) (*PaintswapFilterer, error) {
	contract, err := bindPaintswap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaintswapFilterer{contract: contract}, nil
}

// bindPaintswap binds a generic wrapper to an already deployed contract.
func bindPaintswap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaintswapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paintswap *PaintswapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paintswap.Contract.PaintswapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paintswap *PaintswapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.Contract.PaintswapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paintswap *PaintswapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paintswap.Contract.PaintswapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Paintswap *PaintswapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Paintswap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Paintswap *PaintswapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Paintswap *PaintswapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Paintswap.Contract.contract.Transact(opts, method, params...)
}

// AcceptOfferAfterDeadlineEnabled is a free data retrieval call binding the contract method 0x19050d20.
//
// Solidity: function acceptOfferAfterDeadlineEnabled() view returns(bool)
func (_Paintswap *PaintswapCaller) AcceptOfferAfterDeadlineEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "acceptOfferAfterDeadlineEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AcceptOfferAfterDeadlineEnabled is a free data retrieval call binding the contract method 0x19050d20.
//
// Solidity: function acceptOfferAfterDeadlineEnabled() view returns(bool)
func (_Paintswap *PaintswapSession) AcceptOfferAfterDeadlineEnabled() (bool, error) {
	return _Paintswap.Contract.AcceptOfferAfterDeadlineEnabled(&_Paintswap.CallOpts)
}

// AcceptOfferAfterDeadlineEnabled is a free data retrieval call binding the contract method 0x19050d20.
//
// Solidity: function acceptOfferAfterDeadlineEnabled() view returns(bool)
func (_Paintswap *PaintswapCallerSession) AcceptOfferAfterDeadlineEnabled() (bool, error) {
	return _Paintswap.Contract.AcceptOfferAfterDeadlineEnabled(&_Paintswap.CallOpts)
}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint256)
func (_Paintswap *PaintswapCaller) AuctionEndTimeIncreaseCutOff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "auctionEndTimeIncreaseCutOff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint256)
func (_Paintswap *PaintswapSession) AuctionEndTimeIncreaseCutOff() (*big.Int, error) {
	return _Paintswap.Contract.AuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// AuctionEndTimeIncreaseCutOff is a free data retrieval call binding the contract method 0x7cdc8e2c.
//
// Solidity: function auctionEndTimeIncreaseCutOff() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) AuctionEndTimeIncreaseCutOff() (*big.Int, error) {
	return _Paintswap.Contract.AuctionEndTimeIncreaseCutOff(&_Paintswap.CallOpts)
}

// BundlesEnabled is a free data retrieval call binding the contract method 0x34098c74.
//
// Solidity: function bundlesEnabled() view returns(bool)
func (_Paintswap *PaintswapCaller) BundlesEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "bundlesEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BundlesEnabled is a free data retrieval call binding the contract method 0x34098c74.
//
// Solidity: function bundlesEnabled() view returns(bool)
func (_Paintswap *PaintswapSession) BundlesEnabled() (bool, error) {
	return _Paintswap.Contract.BundlesEnabled(&_Paintswap.CallOpts)
}

// BundlesEnabled is a free data retrieval call binding the contract method 0x34098c74.
//
// Solidity: function bundlesEnabled() view returns(bool)
func (_Paintswap *PaintswapCallerSession) BundlesEnabled() (bool, error) {
	return _Paintswap.Contract.BundlesEnabled(&_Paintswap.CallOpts)
}

// CancelCutoff is a free data retrieval call binding the contract method 0x8a79feeb.
//
// Solidity: function cancelCutoff() view returns(uint256)
func (_Paintswap *PaintswapCaller) CancelCutoff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "cancelCutoff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelCutoff is a free data retrieval call binding the contract method 0x8a79feeb.
//
// Solidity: function cancelCutoff() view returns(uint256)
func (_Paintswap *PaintswapSession) CancelCutoff() (*big.Int, error) {
	return _Paintswap.Contract.CancelCutoff(&_Paintswap.CallOpts)
}

// CancelCutoff is a free data retrieval call binding the contract method 0x8a79feeb.
//
// Solidity: function cancelCutoff() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) CancelCutoff() (*big.Int, error) {
	return _Paintswap.Contract.CancelCutoff(&_Paintswap.CallOpts)
}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapCaller) CurrentMarketplaceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "currentMarketplaceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapSession) CurrentMarketplaceId() (*big.Int, error) {
	return _Paintswap.Contract.CurrentMarketplaceId(&_Paintswap.CallOpts)
}

// CurrentMarketplaceId is a free data retrieval call binding the contract method 0x9e9109fe.
//
// Solidity: function currentMarketplaceId() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) CurrentMarketplaceId() (*big.Int, error) {
	return _Paintswap.Contract.CurrentMarketplaceId(&_Paintswap.CallOpts)
}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapCaller) DevFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "devFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapSession) DevFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevFeeAddress(&_Paintswap.CallOpts)
}

// DevFeeAddress is a free data retrieval call binding the contract method 0xc0907099.
//
// Solidity: function devFeeAddress() view returns(address)
func (_Paintswap *PaintswapCallerSession) DevFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevFeeAddress(&_Paintswap.CallOpts)
}

// DevFeePercentage is a free data retrieval call binding the contract method 0xe884269a.
//
// Solidity: function devFeePercentage() view returns(uint16)
func (_Paintswap *PaintswapCaller) DevFeePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "devFeePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DevFeePercentage is a free data retrieval call binding the contract method 0xe884269a.
//
// Solidity: function devFeePercentage() view returns(uint16)
func (_Paintswap *PaintswapSession) DevFeePercentage() (uint16, error) {
	return _Paintswap.Contract.DevFeePercentage(&_Paintswap.CallOpts)
}

// DevFeePercentage is a free data retrieval call binding the contract method 0xe884269a.
//
// Solidity: function devFeePercentage() view returns(uint16)
func (_Paintswap *PaintswapCallerSession) DevFeePercentage() (uint16, error) {
	return _Paintswap.Contract.DevFeePercentage(&_Paintswap.CallOpts)
}

// DevSalesFeeAddress is a free data retrieval call binding the contract method 0xbe165089.
//
// Solidity: function devSalesFeeAddress() view returns(address)
func (_Paintswap *PaintswapCaller) DevSalesFeeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "devSalesFeeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DevSalesFeeAddress is a free data retrieval call binding the contract method 0xbe165089.
//
// Solidity: function devSalesFeeAddress() view returns(address)
func (_Paintswap *PaintswapSession) DevSalesFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevSalesFeeAddress(&_Paintswap.CallOpts)
}

// DevSalesFeeAddress is a free data retrieval call binding the contract method 0xbe165089.
//
// Solidity: function devSalesFeeAddress() view returns(address)
func (_Paintswap *PaintswapCallerSession) DevSalesFeeAddress() (common.Address, error) {
	return _Paintswap.Contract.DevSalesFeeAddress(&_Paintswap.CallOpts)
}

// EndTimeCutOff is a free data retrieval call binding the contract method 0xfbd6ef36.
//
// Solidity: function endTimeCutOff() view returns(uint256)
func (_Paintswap *PaintswapCaller) EndTimeCutOff(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "endTimeCutOff")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTimeCutOff is a free data retrieval call binding the contract method 0xfbd6ef36.
//
// Solidity: function endTimeCutOff() view returns(uint256)
func (_Paintswap *PaintswapSession) EndTimeCutOff() (*big.Int, error) {
	return _Paintswap.Contract.EndTimeCutOff(&_Paintswap.CallOpts)
}

// EndTimeCutOff is a free data retrieval call binding the contract method 0xfbd6ef36.
//
// Solidity: function endTimeCutOff() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) EndTimeCutOff() (*big.Int, error) {
	return _Paintswap.Contract.EndTimeCutOff(&_Paintswap.CallOpts)
}

// ExcessNFTAmount is a free data retrieval call binding the contract method 0x9e4af935.
//
// Solidity: function excessNFTAmount(address _nft, uint256 _tokenId) view returns(uint256)
func (_Paintswap *PaintswapCaller) ExcessNFTAmount(opts *bind.CallOpts, _nft common.Address, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "excessNFTAmount", _nft, _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExcessNFTAmount is a free data retrieval call binding the contract method 0x9e4af935.
//
// Solidity: function excessNFTAmount(address _nft, uint256 _tokenId) view returns(uint256)
func (_Paintswap *PaintswapSession) ExcessNFTAmount(_nft common.Address, _tokenId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.ExcessNFTAmount(&_Paintswap.CallOpts, _nft, _tokenId)
}

// ExcessNFTAmount is a free data retrieval call binding the contract method 0x9e4af935.
//
// Solidity: function excessNFTAmount(address _nft, uint256 _tokenId) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) ExcessNFTAmount(_nft common.Address, _tokenId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.ExcessNFTAmount(&_Paintswap.CallOpts, _nft, _tokenId)
}

// ExcessTokenAmount is a free data retrieval call binding the contract method 0x12f83520.
//
// Solidity: function excessTokenAmount(address _token) view returns(uint256)
func (_Paintswap *PaintswapCaller) ExcessTokenAmount(opts *bind.CallOpts, _token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "excessTokenAmount", _token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExcessTokenAmount is a free data retrieval call binding the contract method 0x12f83520.
//
// Solidity: function excessTokenAmount(address _token) view returns(uint256)
func (_Paintswap *PaintswapSession) ExcessTokenAmount(_token common.Address) (*big.Int, error) {
	return _Paintswap.Contract.ExcessTokenAmount(&_Paintswap.CallOpts, _token)
}

// ExcessTokenAmount is a free data retrieval call binding the contract method 0x12f83520.
//
// Solidity: function excessTokenAmount(address _token) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) ExcessTokenAmount(_token common.Address) (*big.Int, error) {
	return _Paintswap.Contract.ExcessTokenAmount(&_Paintswap.CallOpts, _token)
}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],address,uint128,uint256,uint256,uint128,address,bool,uint128,uint128,address,bool,uint256))
func (_Paintswap *PaintswapCaller) GetSaleDetails(opts *bind.CallOpts, _marketplaceId *big.Int) (PaintSwapMarketplaceV2nftDetails, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "getSaleDetails", _marketplaceId)

	if err != nil {
		return *new(PaintSwapMarketplaceV2nftDetails), err
	}

	out0 := *abi.ConvertType(out[0], new(PaintSwapMarketplaceV2nftDetails)).(*PaintSwapMarketplaceV2nftDetails)

	return out0, err

}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],address,uint128,uint256,uint256,uint128,address,bool,uint128,uint128,address,bool,uint256))
func (_Paintswap *PaintswapSession) GetSaleDetails(_marketplaceId *big.Int) (PaintSwapMarketplaceV2nftDetails, error) {
	return _Paintswap.Contract.GetSaleDetails(&_Paintswap.CallOpts, _marketplaceId)
}

// GetSaleDetails is a free data retrieval call binding the contract method 0x8c746d8b.
//
// Solidity: function getSaleDetails(uint256 _marketplaceId) view returns((address[],uint256[],uint256[],address,uint128,uint256,uint256,uint128,address,bool,uint128,uint128,address,bool,uint256))
func (_Paintswap *PaintswapCallerSession) GetSaleDetails(_marketplaceId *big.Int) (PaintSwapMarketplaceV2nftDetails, error) {
	return _Paintswap.Contract.GetSaleDetails(&_Paintswap.CallOpts, _marketplaceId)
}

// IsSellingEnabled is a free data retrieval call binding the contract method 0x0433b1a9.
//
// Solidity: function isSellingEnabled() view returns(bool)
func (_Paintswap *PaintswapCaller) IsSellingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "isSellingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSellingEnabled is a free data retrieval call binding the contract method 0x0433b1a9.
//
// Solidity: function isSellingEnabled() view returns(bool)
func (_Paintswap *PaintswapSession) IsSellingEnabled() (bool, error) {
	return _Paintswap.Contract.IsSellingEnabled(&_Paintswap.CallOpts)
}

// IsSellingEnabled is a free data retrieval call binding the contract method 0x0433b1a9.
//
// Solidity: function isSellingEnabled() view returns(bool)
func (_Paintswap *PaintswapCallerSession) IsSellingEnabled() (bool, error) {
	return _Paintswap.Contract.IsSellingEnabled(&_Paintswap.CallOpts)
}

// MarketplaceWhitelist is a free data retrieval call binding the contract method 0x3ba845ef.
//
// Solidity: function marketplaceWhitelist() view returns(address)
func (_Paintswap *PaintswapCaller) MarketplaceWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "marketplaceWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceWhitelist is a free data retrieval call binding the contract method 0x3ba845ef.
//
// Solidity: function marketplaceWhitelist() view returns(address)
func (_Paintswap *PaintswapSession) MarketplaceWhitelist() (common.Address, error) {
	return _Paintswap.Contract.MarketplaceWhitelist(&_Paintswap.CallOpts)
}

// MarketplaceWhitelist is a free data retrieval call binding the contract method 0x3ba845ef.
//
// Solidity: function marketplaceWhitelist() view returns(address)
func (_Paintswap *PaintswapCallerSession) MarketplaceWhitelist() (common.Address, error) {
	return _Paintswap.Contract.MarketplaceWhitelist(&_Paintswap.CallOpts)
}

// MaxBundleNumber is a free data retrieval call binding the contract method 0xce55073b.
//
// Solidity: function maxBundleNumber() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxBundleNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxBundleNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxBundleNumber is a free data retrieval call binding the contract method 0xce55073b.
//
// Solidity: function maxBundleNumber() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxBundleNumber() (*big.Int, error) {
	return _Paintswap.Contract.MaxBundleNumber(&_Paintswap.CallOpts)
}

// MaxBundleNumber is a free data retrieval call binding the contract method 0xce55073b.
//
// Solidity: function maxBundleNumber() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxBundleNumber() (*big.Int, error) {
	return _Paintswap.Contract.MaxBundleNumber(&_Paintswap.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxDuration(&_Paintswap.CallOpts)
}

// MaxDuration is a free data retrieval call binding the contract method 0x6db5c8fd.
//
// Solidity: function maxDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxDuration() (*big.Int, error) {
	return _Paintswap.Contract.MaxDuration(&_Paintswap.CallOpts)
}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint256)
func (_Paintswap *PaintswapCaller) MaxStartTimeIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "maxStartTimeIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint256)
func (_Paintswap *PaintswapSession) MaxStartTimeIncrement() (*big.Int, error) {
	return _Paintswap.Contract.MaxStartTimeIncrement(&_Paintswap.CallOpts)
}

// MaxStartTimeIncrement is a free data retrieval call binding the contract method 0x01a81265.
//
// Solidity: function maxStartTimeIncrement() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MaxStartTimeIncrement() (*big.Int, error) {
	return _Paintswap.Contract.MaxStartTimeIncrement(&_Paintswap.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapCaller) MinDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapSession) MinDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinDuration(&_Paintswap.CallOpts)
}

// MinDuration is a free data retrieval call binding the contract method 0x56715761.
//
// Solidity: function minDuration() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) MinDuration() (*big.Int, error) {
	return _Paintswap.Contract.MinDuration(&_Paintswap.CallOpts)
}

// MinIncreasedBidOfferPercent is a free data retrieval call binding the contract method 0x332686e1.
//
// Solidity: function minIncreasedBidOfferPercent() view returns(uint128)
func (_Paintswap *PaintswapCaller) MinIncreasedBidOfferPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "minIncreasedBidOfferPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinIncreasedBidOfferPercent is a free data retrieval call binding the contract method 0x332686e1.
//
// Solidity: function minIncreasedBidOfferPercent() view returns(uint128)
func (_Paintswap *PaintswapSession) MinIncreasedBidOfferPercent() (*big.Int, error) {
	return _Paintswap.Contract.MinIncreasedBidOfferPercent(&_Paintswap.CallOpts)
}

// MinIncreasedBidOfferPercent is a free data retrieval call binding the contract method 0x332686e1.
//
// Solidity: function minIncreasedBidOfferPercent() view returns(uint128)
func (_Paintswap *PaintswapCallerSession) MinIncreasedBidOfferPercent() (*big.Int, error) {
	return _Paintswap.Contract.MinIncreasedBidOfferPercent(&_Paintswap.CallOpts)
}

// ModifyInitialStartTime is a free data retrieval call binding the contract method 0xa579b9c3.
//
// Solidity: function modifyInitialStartTime() view returns(bool)
func (_Paintswap *PaintswapCaller) ModifyInitialStartTime(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "modifyInitialStartTime")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ModifyInitialStartTime is a free data retrieval call binding the contract method 0xa579b9c3.
//
// Solidity: function modifyInitialStartTime() view returns(bool)
func (_Paintswap *PaintswapSession) ModifyInitialStartTime() (bool, error) {
	return _Paintswap.Contract.ModifyInitialStartTime(&_Paintswap.CallOpts)
}

// ModifyInitialStartTime is a free data retrieval call binding the contract method 0xa579b9c3.
//
// Solidity: function modifyInitialStartTime() view returns(bool)
func (_Paintswap *PaintswapCallerSession) ModifyInitialStartTime() (bool, error) {
	return _Paintswap.Contract.ModifyInitialStartTime(&_Paintswap.CallOpts)
}

// NextMinimumBidOrOffer is a free data retrieval call binding the contract method 0xa519824d.
//
// Solidity: function nextMinimumBidOrOffer(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapCaller) NextMinimumBidOrOffer(opts *bind.CallOpts, _marketplaceId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "nextMinimumBidOrOffer", _marketplaceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextMinimumBidOrOffer is a free data retrieval call binding the contract method 0xa519824d.
//
// Solidity: function nextMinimumBidOrOffer(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapSession) NextMinimumBidOrOffer(_marketplaceId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.NextMinimumBidOrOffer(&_Paintswap.CallOpts, _marketplaceId)
}

// NextMinimumBidOrOffer is a free data retrieval call binding the contract method 0xa519824d.
//
// Solidity: function nextMinimumBidOrOffer(uint256 _marketplaceId) view returns(uint256)
func (_Paintswap *PaintswapCallerSession) NextMinimumBidOrOffer(_marketplaceId *big.Int) (*big.Int, error) {
	return _Paintswap.Contract.NextMinimumBidOrOffer(&_Paintswap.CallOpts, _marketplaceId)
}

// OfficialNFTDiscount is a free data retrieval call binding the contract method 0x0a785cd6.
//
// Solidity: function officialNFTDiscount() view returns(address)
func (_Paintswap *PaintswapCaller) OfficialNFTDiscount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "officialNFTDiscount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OfficialNFTDiscount is a free data retrieval call binding the contract method 0x0a785cd6.
//
// Solidity: function officialNFTDiscount() view returns(address)
func (_Paintswap *PaintswapSession) OfficialNFTDiscount() (common.Address, error) {
	return _Paintswap.Contract.OfficialNFTDiscount(&_Paintswap.CallOpts)
}

// OfficialNFTDiscount is a free data retrieval call binding the contract method 0x0a785cd6.
//
// Solidity: function officialNFTDiscount() view returns(address)
func (_Paintswap *PaintswapCallerSession) OfficialNFTDiscount() (common.Address, error) {
	return _Paintswap.Contract.OfficialNFTDiscount(&_Paintswap.CallOpts)
}

// OnlySellerAndHighestOfferCanEndSale is a free data retrieval call binding the contract method 0x47ac36be.
//
// Solidity: function onlySellerAndHighestOfferCanEndSale() view returns(bool)
func (_Paintswap *PaintswapCaller) OnlySellerAndHighestOfferCanEndSale(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "onlySellerAndHighestOfferCanEndSale")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OnlySellerAndHighestOfferCanEndSale is a free data retrieval call binding the contract method 0x47ac36be.
//
// Solidity: function onlySellerAndHighestOfferCanEndSale() view returns(bool)
func (_Paintswap *PaintswapSession) OnlySellerAndHighestOfferCanEndSale() (bool, error) {
	return _Paintswap.Contract.OnlySellerAndHighestOfferCanEndSale(&_Paintswap.CallOpts)
}

// OnlySellerAndHighestOfferCanEndSale is a free data retrieval call binding the contract method 0x47ac36be.
//
// Solidity: function onlySellerAndHighestOfferCanEndSale() view returns(bool)
func (_Paintswap *PaintswapCallerSession) OnlySellerAndHighestOfferCanEndSale() (bool, error) {
	return _Paintswap.Contract.OnlySellerAndHighestOfferCanEndSale(&_Paintswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapSession) Owner() (common.Address, error) {
	return _Paintswap.Contract.Owner(&_Paintswap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Paintswap *PaintswapCallerSession) Owner() (common.Address, error) {
	return _Paintswap.Contract.Owner(&_Paintswap.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x46e80720.
//
// Solidity: function royaltyInfo() view returns(uint256 maxRoyalty, address collectionRoyalties)
func (_Paintswap *PaintswapCaller) RoyaltyInfo(opts *bind.CallOpts) (struct {
	MaxRoyalty          *big.Int
	CollectionRoyalties common.Address
}, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "royaltyInfo")

	outstruct := new(struct {
		MaxRoyalty          *big.Int
		CollectionRoyalties common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxRoyalty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollectionRoyalties = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x46e80720.
//
// Solidity: function royaltyInfo() view returns(uint256 maxRoyalty, address collectionRoyalties)
func (_Paintswap *PaintswapSession) RoyaltyInfo() (struct {
	MaxRoyalty          *big.Int
	CollectionRoyalties common.Address
}, error) {
	return _Paintswap.Contract.RoyaltyInfo(&_Paintswap.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x46e80720.
//
// Solidity: function royaltyInfo() view returns(uint256 maxRoyalty, address collectionRoyalties)
func (_Paintswap *PaintswapCallerSession) RoyaltyInfo() (struct {
	MaxRoyalty          *big.Int
	CollectionRoyalties common.Address
}, error) {
	return _Paintswap.Contract.RoyaltyInfo(&_Paintswap.CallOpts)
}

// SaleGracePeriodForCancelling is a free data retrieval call binding the contract method 0xb2612614.
//
// Solidity: function saleGracePeriodForCancelling() view returns(uint256)
func (_Paintswap *PaintswapCaller) SaleGracePeriodForCancelling(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "saleGracePeriodForCancelling")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SaleGracePeriodForCancelling is a free data retrieval call binding the contract method 0xb2612614.
//
// Solidity: function saleGracePeriodForCancelling() view returns(uint256)
func (_Paintswap *PaintswapSession) SaleGracePeriodForCancelling() (*big.Int, error) {
	return _Paintswap.Contract.SaleGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// SaleGracePeriodForCancelling is a free data retrieval call binding the contract method 0xb2612614.
//
// Solidity: function saleGracePeriodForCancelling() view returns(uint256)
func (_Paintswap *PaintswapCallerSession) SaleGracePeriodForCancelling() (*big.Int, error) {
	return _Paintswap.Contract.SaleGracePeriodForCancelling(&_Paintswap.CallOpts)
}

// SlidingEndTimeForBidsEnabled is a free data retrieval call binding the contract method 0xb55e0e65.
//
// Solidity: function slidingEndTimeForBidsEnabled() view returns(bool)
func (_Paintswap *PaintswapCaller) SlidingEndTimeForBidsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "slidingEndTimeForBidsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlidingEndTimeForBidsEnabled is a free data retrieval call binding the contract method 0xb55e0e65.
//
// Solidity: function slidingEndTimeForBidsEnabled() view returns(bool)
func (_Paintswap *PaintswapSession) SlidingEndTimeForBidsEnabled() (bool, error) {
	return _Paintswap.Contract.SlidingEndTimeForBidsEnabled(&_Paintswap.CallOpts)
}

// SlidingEndTimeForBidsEnabled is a free data retrieval call binding the contract method 0xb55e0e65.
//
// Solidity: function slidingEndTimeForBidsEnabled() view returns(bool)
func (_Paintswap *PaintswapCallerSession) SlidingEndTimeForBidsEnabled() (bool, error) {
	return _Paintswap.Contract.SlidingEndTimeForBidsEnabled(&_Paintswap.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Paintswap *PaintswapCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Paintswap *PaintswapSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Paintswap.Contract.SupportsInterface(&_Paintswap.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Paintswap *PaintswapCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Paintswap.Contract.SupportsInterface(&_Paintswap.CallOpts, interfaceId)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address)
func (_Paintswap *PaintswapCaller) Tokens(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "tokens")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address)
func (_Paintswap *PaintswapSession) Tokens() (common.Address, error) {
	return _Paintswap.Contract.Tokens(&_Paintswap.CallOpts)
}

// Tokens is a free data retrieval call binding the contract method 0x9d63848a.
//
// Solidity: function tokens() view returns(address)
func (_Paintswap *PaintswapCallerSession) Tokens() (common.Address, error) {
	return _Paintswap.Contract.Tokens(&_Paintswap.CallOpts)
}

// WhitelistedSellerBuyerContract is a free data retrieval call binding the contract method 0xf83c7264.
//
// Solidity: function whitelistedSellerBuyerContract(address ) view returns(bool)
func (_Paintswap *PaintswapCaller) WhitelistedSellerBuyerContract(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Paintswap.contract.Call(opts, &out, "whitelistedSellerBuyerContract", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WhitelistedSellerBuyerContract is a free data retrieval call binding the contract method 0xf83c7264.
//
// Solidity: function whitelistedSellerBuyerContract(address ) view returns(bool)
func (_Paintswap *PaintswapSession) WhitelistedSellerBuyerContract(arg0 common.Address) (bool, error) {
	return _Paintswap.Contract.WhitelistedSellerBuyerContract(&_Paintswap.CallOpts, arg0)
}

// WhitelistedSellerBuyerContract is a free data retrieval call binding the contract method 0xf83c7264.
//
// Solidity: function whitelistedSellerBuyerContract(address ) view returns(bool)
func (_Paintswap *PaintswapCallerSession) WhitelistedSellerBuyerContract(arg0 common.Address) (bool, error) {
	return _Paintswap.Contract.WhitelistedSellerBuyerContract(&_Paintswap.CallOpts, arg0)
}

// AcceptBestOffer is a paid mutator transaction binding the contract method 0x13e0156c.
//
// Solidity: function acceptBestOffer(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) AcceptBestOffer(opts *bind.TransactOpts, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "acceptBestOffer", _marketplaceId)
}

// AcceptBestOffer is a paid mutator transaction binding the contract method 0x13e0156c.
//
// Solidity: function acceptBestOffer(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) AcceptBestOffer(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptBestOffer(&_Paintswap.TransactOpts, _marketplaceId)
}

// AcceptBestOffer is a paid mutator transaction binding the contract method 0x13e0156c.
//
// Solidity: function acceptBestOffer(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) AcceptBestOffer(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.AcceptBestOffer(&_Paintswap.TransactOpts, _marketplaceId)
}

// AddSale is a paid mutator transaction binding the contract method 0x9856bfbe.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint256[] _times, uint128 _amount, bool[] _flags, string _marketplaceURI, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactor) AddSale(opts *bind.TransactOpts, _nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []*big.Int, _amount *big.Int, _flags []bool, _marketplaceURI string, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "addSale", _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _marketplaceURI, _searchKeywords, _addresses)
}

// AddSale is a paid mutator transaction binding the contract method 0x9856bfbe.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint256[] _times, uint128 _amount, bool[] _flags, string _marketplaceURI, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapSession) AddSale(_nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []*big.Int, _amount *big.Int, _flags []bool, _marketplaceURI string, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSale(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _marketplaceURI, _searchKeywords, _addresses)
}

// AddSale is a paid mutator transaction binding the contract method 0x9856bfbe.
//
// Solidity: function addSale(address[] _nfts, uint256[] _tokenIds, uint256[] _amountBatches, uint128 _price, uint256[] _times, uint128 _amount, bool[] _flags, string _marketplaceURI, string _searchKeywords, address[] _addresses) payable returns(uint256)
func (_Paintswap *PaintswapTransactorSession) AddSale(_nfts []common.Address, _tokenIds []*big.Int, _amountBatches []*big.Int, _price *big.Int, _times []*big.Int, _amount *big.Int, _flags []bool, _marketplaceURI string, _searchKeywords string, _addresses []common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddSale(&_Paintswap.TransactOpts, _nfts, _tokenIds, _amountBatches, _price, _times, _amount, _flags, _marketplaceURI, _searchKeywords, _addresses)
}

// AddWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0x04fd41df.
//
// Solidity: function addWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapTransactor) AddWhitelistedSellerBuyerContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "addWhitelistedSellerBuyerContract", _address)
}

// AddWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0x04fd41df.
//
// Solidity: function addWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapSession) AddWhitelistedSellerBuyerContract(_address common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddWhitelistedSellerBuyerContract(&_Paintswap.TransactOpts, _address)
}

// AddWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0x04fd41df.
//
// Solidity: function addWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapTransactorSession) AddWhitelistedSellerBuyerContract(_address common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.AddWhitelistedSellerBuyerContract(&_Paintswap.TransactOpts, _address)
}

// Buy is a paid mutator transaction binding the contract method 0x37bd97b9.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactor) Buy(opts *bind.TransactOpts, _marketplaceId *big.Int, _amount *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "buy", _marketplaceId, _amount, _buyer)
}

// Buy is a paid mutator transaction binding the contract method 0x37bd97b9.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount, address _buyer) payable returns()
func (_Paintswap *PaintswapSession) Buy(_marketplaceId *big.Int, _amount *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.Buy(&_Paintswap.TransactOpts, _marketplaceId, _amount, _buyer)
}

// Buy is a paid mutator transaction binding the contract method 0x37bd97b9.
//
// Solidity: function buy(uint256 _marketplaceId, uint128 _amount, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactorSession) Buy(_marketplaceId *big.Int, _amount *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.Buy(&_Paintswap.TransactOpts, _marketplaceId, _amount, _buyer)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) CancelSale(opts *bind.TransactOpts, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "cancelSale", _marketplaceId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) CancelSale(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSale(&_Paintswap.TransactOpts, _marketplaceId)
}

// CancelSale is a paid mutator transaction binding the contract method 0xbd94b005.
//
// Solidity: function cancelSale(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) CancelSale(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CancelSale(&_Paintswap.TransactOpts, _marketplaceId)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactor) CompleteMarketplaceEntry(opts *bind.TransactOpts, _marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "completeMarketplaceEntry", _marketplaceId)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapSession) CompleteMarketplaceEntry(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CompleteMarketplaceEntry(&_Paintswap.TransactOpts, _marketplaceId)
}

// CompleteMarketplaceEntry is a paid mutator transaction binding the contract method 0xfc60698b.
//
// Solidity: function completeMarketplaceEntry(uint256 _marketplaceId) returns()
func (_Paintswap *PaintswapTransactorSession) CompleteMarketplaceEntry(_marketplaceId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.CompleteMarketplaceEntry(&_Paintswap.TransactOpts, _marketplaceId)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapTransactor) EditPrice(opts *bind.TransactOpts, _marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "editPrice", _marketplaceId, _price)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapSession) EditPrice(_marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPrice(&_Paintswap.TransactOpts, _marketplaceId, _price)
}

// EditPrice is a paid mutator transaction binding the contract method 0xa9a0c655.
//
// Solidity: function editPrice(uint256 _marketplaceId, uint128 _price) returns()
func (_Paintswap *PaintswapTransactorSession) EditPrice(_marketplaceId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.EditPrice(&_Paintswap.TransactOpts, _marketplaceId, _price)
}

// MakeBid is a paid mutator transaction binding the contract method 0xf2a960ba.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactor) MakeBid(opts *bind.TransactOpts, _marketplaceId *big.Int, _bid *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeBid", _marketplaceId, _bid, _buyer)
}

// MakeBid is a paid mutator transaction binding the contract method 0xf2a960ba.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid, address _buyer) payable returns()
func (_Paintswap *PaintswapSession) MakeBid(_marketplaceId *big.Int, _bid *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBid(&_Paintswap.TransactOpts, _marketplaceId, _bid, _buyer)
}

// MakeBid is a paid mutator transaction binding the contract method 0xf2a960ba.
//
// Solidity: function makeBid(uint256 _marketplaceId, uint128 _bid, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeBid(_marketplaceId *big.Int, _bid *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeBid(&_Paintswap.TransactOpts, _marketplaceId, _bid, _buyer)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xf1bbe386.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _offer, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactor) MakeOffer(opts *bind.TransactOpts, _marketplaceId *big.Int, _offer *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "makeOffer", _marketplaceId, _offer, _buyer)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xf1bbe386.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _offer, address _buyer) payable returns()
func (_Paintswap *PaintswapSession) MakeOffer(_marketplaceId *big.Int, _offer *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOffer(&_Paintswap.TransactOpts, _marketplaceId, _offer, _buyer)
}

// MakeOffer is a paid mutator transaction binding the contract method 0xf1bbe386.
//
// Solidity: function makeOffer(uint256 _marketplaceId, uint128 _offer, address _buyer) payable returns()
func (_Paintswap *PaintswapTransactorSession) MakeOffer(_marketplaceId *big.Int, _offer *big.Int, _buyer common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.MakeOffer(&_Paintswap.TransactOpts, _marketplaceId, _offer, _buyer)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Paintswap *PaintswapSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC1155BatchReceived(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC1155BatchReceived(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC1155Received(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC1155Received(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC721Received(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_Paintswap *PaintswapTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _Paintswap.Contract.OnERC721Received(&_Paintswap.TransactOpts, arg0, arg1, arg2, arg3)
}

// RemoveWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0xf8f08fab.
//
// Solidity: function removeWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapTransactor) RemoveWhitelistedSellerBuyerContract(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "removeWhitelistedSellerBuyerContract", _address)
}

// RemoveWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0xf8f08fab.
//
// Solidity: function removeWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapSession) RemoveWhitelistedSellerBuyerContract(_address common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.RemoveWhitelistedSellerBuyerContract(&_Paintswap.TransactOpts, _address)
}

// RemoveWhitelistedSellerBuyerContract is a paid mutator transaction binding the contract method 0xf8f08fab.
//
// Solidity: function removeWhitelistedSellerBuyerContract(address _address) returns()
func (_Paintswap *PaintswapTransactorSession) RemoveWhitelistedSellerBuyerContract(_address common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.RemoveWhitelistedSellerBuyerContract(&_Paintswap.TransactOpts, _address)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paintswap.Contract.RenounceOwnership(&_Paintswap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Paintswap *PaintswapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Paintswap.Contract.RenounceOwnership(&_Paintswap.TransactOpts)
}

// SetAcceptOfferAfterDeadlineEnabled is a paid mutator transaction binding the contract method 0xaffb7265.
//
// Solidity: function setAcceptOfferAfterDeadlineEnabled(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetAcceptOfferAfterDeadlineEnabled(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setAcceptOfferAfterDeadlineEnabled", _enable)
}

// SetAcceptOfferAfterDeadlineEnabled is a paid mutator transaction binding the contract method 0xaffb7265.
//
// Solidity: function setAcceptOfferAfterDeadlineEnabled(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetAcceptOfferAfterDeadlineEnabled(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetAcceptOfferAfterDeadlineEnabled(&_Paintswap.TransactOpts, _enable)
}

// SetAcceptOfferAfterDeadlineEnabled is a paid mutator transaction binding the contract method 0xaffb7265.
//
// Solidity: function setAcceptOfferAfterDeadlineEnabled(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetAcceptOfferAfterDeadlineEnabled(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetAcceptOfferAfterDeadlineEnabled(&_Paintswap.TransactOpts, _enable)
}

// SetAuctionEndTimeIncreaseCutOff is a paid mutator transaction binding the contract method 0xd3e03e1a.
//
// Solidity: function setAuctionEndTimeIncreaseCutOff(uint256 _auctionEndTimeIncreaseCutOff) returns()
func (_Paintswap *PaintswapTransactor) SetAuctionEndTimeIncreaseCutOff(opts *bind.TransactOpts, _auctionEndTimeIncreaseCutOff *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setAuctionEndTimeIncreaseCutOff", _auctionEndTimeIncreaseCutOff)
}

// SetAuctionEndTimeIncreaseCutOff is a paid mutator transaction binding the contract method 0xd3e03e1a.
//
// Solidity: function setAuctionEndTimeIncreaseCutOff(uint256 _auctionEndTimeIncreaseCutOff) returns()
func (_Paintswap *PaintswapSession) SetAuctionEndTimeIncreaseCutOff(_auctionEndTimeIncreaseCutOff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetAuctionEndTimeIncreaseCutOff(&_Paintswap.TransactOpts, _auctionEndTimeIncreaseCutOff)
}

// SetAuctionEndTimeIncreaseCutOff is a paid mutator transaction binding the contract method 0xd3e03e1a.
//
// Solidity: function setAuctionEndTimeIncreaseCutOff(uint256 _auctionEndTimeIncreaseCutOff) returns()
func (_Paintswap *PaintswapTransactorSession) SetAuctionEndTimeIncreaseCutOff(_auctionEndTimeIncreaseCutOff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetAuctionEndTimeIncreaseCutOff(&_Paintswap.TransactOpts, _auctionEndTimeIncreaseCutOff)
}

// SetBundlesSupported is a paid mutator transaction binding the contract method 0x650946c3.
//
// Solidity: function setBundlesSupported(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetBundlesSupported(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setBundlesSupported", _enable)
}

// SetBundlesSupported is a paid mutator transaction binding the contract method 0x650946c3.
//
// Solidity: function setBundlesSupported(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetBundlesSupported(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetBundlesSupported(&_Paintswap.TransactOpts, _enable)
}

// SetBundlesSupported is a paid mutator transaction binding the contract method 0x650946c3.
//
// Solidity: function setBundlesSupported(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetBundlesSupported(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetBundlesSupported(&_Paintswap.TransactOpts, _enable)
}

// SetCallGasLimit is a paid mutator transaction binding the contract method 0x2acda292.
//
// Solidity: function setCallGasLimit(uint256 _callGasLimit) returns()
func (_Paintswap *PaintswapTransactor) SetCallGasLimit(opts *bind.TransactOpts, _callGasLimit *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setCallGasLimit", _callGasLimit)
}

// SetCallGasLimit is a paid mutator transaction binding the contract method 0x2acda292.
//
// Solidity: function setCallGasLimit(uint256 _callGasLimit) returns()
func (_Paintswap *PaintswapSession) SetCallGasLimit(_callGasLimit *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetCallGasLimit(&_Paintswap.TransactOpts, _callGasLimit)
}

// SetCallGasLimit is a paid mutator transaction binding the contract method 0x2acda292.
//
// Solidity: function setCallGasLimit(uint256 _callGasLimit) returns()
func (_Paintswap *PaintswapTransactorSession) SetCallGasLimit(_callGasLimit *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetCallGasLimit(&_Paintswap.TransactOpts, _callGasLimit)
}

// SetCancelCutoff is a paid mutator transaction binding the contract method 0x2d07ff63.
//
// Solidity: function setCancelCutoff(uint256 _cancelCutoff) returns()
func (_Paintswap *PaintswapTransactor) SetCancelCutoff(opts *bind.TransactOpts, _cancelCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setCancelCutoff", _cancelCutoff)
}

// SetCancelCutoff is a paid mutator transaction binding the contract method 0x2d07ff63.
//
// Solidity: function setCancelCutoff(uint256 _cancelCutoff) returns()
func (_Paintswap *PaintswapSession) SetCancelCutoff(_cancelCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetCancelCutoff(&_Paintswap.TransactOpts, _cancelCutoff)
}

// SetCancelCutoff is a paid mutator transaction binding the contract method 0x2d07ff63.
//
// Solidity: function setCancelCutoff(uint256 _cancelCutoff) returns()
func (_Paintswap *PaintswapTransactorSession) SetCancelCutoff(_cancelCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetCancelCutoff(&_Paintswap.TransactOpts, _cancelCutoff)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_Paintswap *PaintswapTransactor) SetDevAddress(opts *bind.TransactOpts, _dev common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setDevAddress", _dev)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_Paintswap *PaintswapSession) SetDevAddress(_dev common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevAddress(&_Paintswap.TransactOpts, _dev)
}

// SetDevAddress is a paid mutator transaction binding the contract method 0xd0d41fe1.
//
// Solidity: function setDevAddress(address _dev) returns()
func (_Paintswap *PaintswapTransactorSession) SetDevAddress(_dev common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevAddress(&_Paintswap.TransactOpts, _dev)
}

// SetDevFeePercentage is a paid mutator transaction binding the contract method 0x673a0d71.
//
// Solidity: function setDevFeePercentage(uint16 _feePercentage) returns()
func (_Paintswap *PaintswapTransactor) SetDevFeePercentage(opts *bind.TransactOpts, _feePercentage uint16) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setDevFeePercentage", _feePercentage)
}

// SetDevFeePercentage is a paid mutator transaction binding the contract method 0x673a0d71.
//
// Solidity: function setDevFeePercentage(uint16 _feePercentage) returns()
func (_Paintswap *PaintswapSession) SetDevFeePercentage(_feePercentage uint16) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevFeePercentage(&_Paintswap.TransactOpts, _feePercentage)
}

// SetDevFeePercentage is a paid mutator transaction binding the contract method 0x673a0d71.
//
// Solidity: function setDevFeePercentage(uint16 _feePercentage) returns()
func (_Paintswap *PaintswapTransactorSession) SetDevFeePercentage(_feePercentage uint16) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevFeePercentage(&_Paintswap.TransactOpts, _feePercentage)
}

// SetDevSalesAddress is a paid mutator transaction binding the contract method 0x57cd987e.
//
// Solidity: function setDevSalesAddress(address _dev) returns()
func (_Paintswap *PaintswapTransactor) SetDevSalesAddress(opts *bind.TransactOpts, _dev common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setDevSalesAddress", _dev)
}

// SetDevSalesAddress is a paid mutator transaction binding the contract method 0x57cd987e.
//
// Solidity: function setDevSalesAddress(address _dev) returns()
func (_Paintswap *PaintswapSession) SetDevSalesAddress(_dev common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevSalesAddress(&_Paintswap.TransactOpts, _dev)
}

// SetDevSalesAddress is a paid mutator transaction binding the contract method 0x57cd987e.
//
// Solidity: function setDevSalesAddress(address _dev) returns()
func (_Paintswap *PaintswapTransactorSession) SetDevSalesAddress(_dev common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetDevSalesAddress(&_Paintswap.TransactOpts, _dev)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetEnableSelling(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setEnableSelling", _enable)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetEnableSelling(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEnableSelling(&_Paintswap.TransactOpts, _enable)
}

// SetEnableSelling is a paid mutator transaction binding the contract method 0x8642683f.
//
// Solidity: function setEnableSelling(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetEnableSelling(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEnableSelling(&_Paintswap.TransactOpts, _enable)
}

// SetEndTimeCutoff is a paid mutator transaction binding the contract method 0xb1b3dbbe.
//
// Solidity: function setEndTimeCutoff(uint256 _endTimeCutoff) returns()
func (_Paintswap *PaintswapTransactor) SetEndTimeCutoff(opts *bind.TransactOpts, _endTimeCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setEndTimeCutoff", _endTimeCutoff)
}

// SetEndTimeCutoff is a paid mutator transaction binding the contract method 0xb1b3dbbe.
//
// Solidity: function setEndTimeCutoff(uint256 _endTimeCutoff) returns()
func (_Paintswap *PaintswapSession) SetEndTimeCutoff(_endTimeCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEndTimeCutoff(&_Paintswap.TransactOpts, _endTimeCutoff)
}

// SetEndTimeCutoff is a paid mutator transaction binding the contract method 0xb1b3dbbe.
//
// Solidity: function setEndTimeCutoff(uint256 _endTimeCutoff) returns()
func (_Paintswap *PaintswapTransactorSession) SetEndTimeCutoff(_endTimeCutoff *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetEndTimeCutoff(&_Paintswap.TransactOpts, _endTimeCutoff)
}

// SetMaxBundleNumber is a paid mutator transaction binding the contract method 0xa54fd9fc.
//
// Solidity: function setMaxBundleNumber(uint128 _maxBundleNumber) returns()
func (_Paintswap *PaintswapTransactor) SetMaxBundleNumber(opts *bind.TransactOpts, _maxBundleNumber *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMaxBundleNumber", _maxBundleNumber)
}

// SetMaxBundleNumber is a paid mutator transaction binding the contract method 0xa54fd9fc.
//
// Solidity: function setMaxBundleNumber(uint128 _maxBundleNumber) returns()
func (_Paintswap *PaintswapSession) SetMaxBundleNumber(_maxBundleNumber *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxBundleNumber(&_Paintswap.TransactOpts, _maxBundleNumber)
}

// SetMaxBundleNumber is a paid mutator transaction binding the contract method 0xa54fd9fc.
//
// Solidity: function setMaxBundleNumber(uint128 _maxBundleNumber) returns()
func (_Paintswap *PaintswapTransactorSession) SetMaxBundleNumber(_maxBundleNumber *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxBundleNumber(&_Paintswap.TransactOpts, _maxBundleNumber)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 _maxDuration) returns()
func (_Paintswap *PaintswapTransactor) SetMaxDuration(opts *bind.TransactOpts, _maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMaxDuration", _maxDuration)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 _maxDuration) returns()
func (_Paintswap *PaintswapSession) SetMaxDuration(_maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxDuration(&_Paintswap.TransactOpts, _maxDuration)
}

// SetMaxDuration is a paid mutator transaction binding the contract method 0xcf0f34c4.
//
// Solidity: function setMaxDuration(uint256 _maxDuration) returns()
func (_Paintswap *PaintswapTransactorSession) SetMaxDuration(_maxDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxDuration(&_Paintswap.TransactOpts, _maxDuration)
}

// SetMaxRoyalty is a paid mutator transaction binding the contract method 0xc96e9f4b.
//
// Solidity: function setMaxRoyalty(uint256 _maxRoyalty) returns()
func (_Paintswap *PaintswapTransactor) SetMaxRoyalty(opts *bind.TransactOpts, _maxRoyalty *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMaxRoyalty", _maxRoyalty)
}

// SetMaxRoyalty is a paid mutator transaction binding the contract method 0xc96e9f4b.
//
// Solidity: function setMaxRoyalty(uint256 _maxRoyalty) returns()
func (_Paintswap *PaintswapSession) SetMaxRoyalty(_maxRoyalty *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxRoyalty(&_Paintswap.TransactOpts, _maxRoyalty)
}

// SetMaxRoyalty is a paid mutator transaction binding the contract method 0xc96e9f4b.
//
// Solidity: function setMaxRoyalty(uint256 _maxRoyalty) returns()
func (_Paintswap *PaintswapTransactorSession) SetMaxRoyalty(_maxRoyalty *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxRoyalty(&_Paintswap.TransactOpts, _maxRoyalty)
}

// SetMaxStartTimeIncrement is a paid mutator transaction binding the contract method 0xc509d231.
//
// Solidity: function setMaxStartTimeIncrement(uint256 _maxStartTimeIncrement) returns()
func (_Paintswap *PaintswapTransactor) SetMaxStartTimeIncrement(opts *bind.TransactOpts, _maxStartTimeIncrement *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMaxStartTimeIncrement", _maxStartTimeIncrement)
}

// SetMaxStartTimeIncrement is a paid mutator transaction binding the contract method 0xc509d231.
//
// Solidity: function setMaxStartTimeIncrement(uint256 _maxStartTimeIncrement) returns()
func (_Paintswap *PaintswapSession) SetMaxStartTimeIncrement(_maxStartTimeIncrement *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxStartTimeIncrement(&_Paintswap.TransactOpts, _maxStartTimeIncrement)
}

// SetMaxStartTimeIncrement is a paid mutator transaction binding the contract method 0xc509d231.
//
// Solidity: function setMaxStartTimeIncrement(uint256 _maxStartTimeIncrement) returns()
func (_Paintswap *PaintswapTransactorSession) SetMaxStartTimeIncrement(_maxStartTimeIncrement *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMaxStartTimeIncrement(&_Paintswap.TransactOpts, _maxStartTimeIncrement)
}

// SetMinIncreasedBidOfferPercent is a paid mutator transaction binding the contract method 0xbce846c9.
//
// Solidity: function setMinIncreasedBidOfferPercent(uint128 _minIncreasedBidOfferPercent) returns()
func (_Paintswap *PaintswapTransactor) SetMinIncreasedBidOfferPercent(opts *bind.TransactOpts, _minIncreasedBidOfferPercent *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMinIncreasedBidOfferPercent", _minIncreasedBidOfferPercent)
}

// SetMinIncreasedBidOfferPercent is a paid mutator transaction binding the contract method 0xbce846c9.
//
// Solidity: function setMinIncreasedBidOfferPercent(uint128 _minIncreasedBidOfferPercent) returns()
func (_Paintswap *PaintswapSession) SetMinIncreasedBidOfferPercent(_minIncreasedBidOfferPercent *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMinIncreasedBidOfferPercent(&_Paintswap.TransactOpts, _minIncreasedBidOfferPercent)
}

// SetMinIncreasedBidOfferPercent is a paid mutator transaction binding the contract method 0xbce846c9.
//
// Solidity: function setMinIncreasedBidOfferPercent(uint128 _minIncreasedBidOfferPercent) returns()
func (_Paintswap *PaintswapTransactorSession) SetMinIncreasedBidOfferPercent(_minIncreasedBidOfferPercent *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMinIncreasedBidOfferPercent(&_Paintswap.TransactOpts, _minIncreasedBidOfferPercent)
}

// SetMinimumDuration is a paid mutator transaction binding the contract method 0xed5ec648.
//
// Solidity: function setMinimumDuration(uint256 _minDuration) returns()
func (_Paintswap *PaintswapTransactor) SetMinimumDuration(opts *bind.TransactOpts, _minDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setMinimumDuration", _minDuration)
}

// SetMinimumDuration is a paid mutator transaction binding the contract method 0xed5ec648.
//
// Solidity: function setMinimumDuration(uint256 _minDuration) returns()
func (_Paintswap *PaintswapSession) SetMinimumDuration(_minDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMinimumDuration(&_Paintswap.TransactOpts, _minDuration)
}

// SetMinimumDuration is a paid mutator transaction binding the contract method 0xed5ec648.
//
// Solidity: function setMinimumDuration(uint256 _minDuration) returns()
func (_Paintswap *PaintswapTransactorSession) SetMinimumDuration(_minDuration *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetMinimumDuration(&_Paintswap.TransactOpts, _minDuration)
}

// SetModifyInitialStartTime is a paid mutator transaction binding the contract method 0x7b35e115.
//
// Solidity: function setModifyInitialStartTime(bool _modifyInitialStartTime) returns()
func (_Paintswap *PaintswapTransactor) SetModifyInitialStartTime(opts *bind.TransactOpts, _modifyInitialStartTime bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setModifyInitialStartTime", _modifyInitialStartTime)
}

// SetModifyInitialStartTime is a paid mutator transaction binding the contract method 0x7b35e115.
//
// Solidity: function setModifyInitialStartTime(bool _modifyInitialStartTime) returns()
func (_Paintswap *PaintswapSession) SetModifyInitialStartTime(_modifyInitialStartTime bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetModifyInitialStartTime(&_Paintswap.TransactOpts, _modifyInitialStartTime)
}

// SetModifyInitialStartTime is a paid mutator transaction binding the contract method 0x7b35e115.
//
// Solidity: function setModifyInitialStartTime(bool _modifyInitialStartTime) returns()
func (_Paintswap *PaintswapTransactorSession) SetModifyInitialStartTime(_modifyInitialStartTime bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetModifyInitialStartTime(&_Paintswap.TransactOpts, _modifyInitialStartTime)
}

// SetOfficialNFTDiscount is a paid mutator transaction binding the contract method 0x78963772.
//
// Solidity: function setOfficialNFTDiscount(address _officialNFTDiscount) returns()
func (_Paintswap *PaintswapTransactor) SetOfficialNFTDiscount(opts *bind.TransactOpts, _officialNFTDiscount common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setOfficialNFTDiscount", _officialNFTDiscount)
}

// SetOfficialNFTDiscount is a paid mutator transaction binding the contract method 0x78963772.
//
// Solidity: function setOfficialNFTDiscount(address _officialNFTDiscount) returns()
func (_Paintswap *PaintswapSession) SetOfficialNFTDiscount(_officialNFTDiscount common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOfficialNFTDiscount(&_Paintswap.TransactOpts, _officialNFTDiscount)
}

// SetOfficialNFTDiscount is a paid mutator transaction binding the contract method 0x78963772.
//
// Solidity: function setOfficialNFTDiscount(address _officialNFTDiscount) returns()
func (_Paintswap *PaintswapTransactorSession) SetOfficialNFTDiscount(_officialNFTDiscount common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOfficialNFTDiscount(&_Paintswap.TransactOpts, _officialNFTDiscount)
}

// SetOnlySellerAndHighestOfferCanEndSale is a paid mutator transaction binding the contract method 0xd888780c.
//
// Solidity: function setOnlySellerAndHighestOfferCanEndSale(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetOnlySellerAndHighestOfferCanEndSale(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setOnlySellerAndHighestOfferCanEndSale", _enable)
}

// SetOnlySellerAndHighestOfferCanEndSale is a paid mutator transaction binding the contract method 0xd888780c.
//
// Solidity: function setOnlySellerAndHighestOfferCanEndSale(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetOnlySellerAndHighestOfferCanEndSale(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOnlySellerAndHighestOfferCanEndSale(&_Paintswap.TransactOpts, _enable)
}

// SetOnlySellerAndHighestOfferCanEndSale is a paid mutator transaction binding the contract method 0xd888780c.
//
// Solidity: function setOnlySellerAndHighestOfferCanEndSale(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetOnlySellerAndHighestOfferCanEndSale(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetOnlySellerAndHighestOfferCanEndSale(&_Paintswap.TransactOpts, _enable)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _paintRouter) returns()
func (_Paintswap *PaintswapTransactor) SetRouter(opts *bind.TransactOpts, _paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setRouter", _paintRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _paintRouter) returns()
func (_Paintswap *PaintswapSession) SetRouter(_paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetRouter(&_Paintswap.TransactOpts, _paintRouter)
}

// SetRouter is a paid mutator transaction binding the contract method 0xc0d78655.
//
// Solidity: function setRouter(address _paintRouter) returns()
func (_Paintswap *PaintswapTransactorSession) SetRouter(_paintRouter common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.SetRouter(&_Paintswap.TransactOpts, _paintRouter)
}

// SetSaleGracePeriodForCancelling is a paid mutator transaction binding the contract method 0x4163e599.
//
// Solidity: function setSaleGracePeriodForCancelling(uint256 _saleGracePeriodForCancellingInSeconds) returns()
func (_Paintswap *PaintswapTransactor) SetSaleGracePeriodForCancelling(opts *bind.TransactOpts, _saleGracePeriodForCancellingInSeconds *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setSaleGracePeriodForCancelling", _saleGracePeriodForCancellingInSeconds)
}

// SetSaleGracePeriodForCancelling is a paid mutator transaction binding the contract method 0x4163e599.
//
// Solidity: function setSaleGracePeriodForCancelling(uint256 _saleGracePeriodForCancellingInSeconds) returns()
func (_Paintswap *PaintswapSession) SetSaleGracePeriodForCancelling(_saleGracePeriodForCancellingInSeconds *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetSaleGracePeriodForCancelling(&_Paintswap.TransactOpts, _saleGracePeriodForCancellingInSeconds)
}

// SetSaleGracePeriodForCancelling is a paid mutator transaction binding the contract method 0x4163e599.
//
// Solidity: function setSaleGracePeriodForCancelling(uint256 _saleGracePeriodForCancellingInSeconds) returns()
func (_Paintswap *PaintswapTransactorSession) SetSaleGracePeriodForCancelling(_saleGracePeriodForCancellingInSeconds *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.SetSaleGracePeriodForCancelling(&_Paintswap.TransactOpts, _saleGracePeriodForCancellingInSeconds)
}

// SetSlidingEndTimeForBidsEnabled is a paid mutator transaction binding the contract method 0xb3f2f251.
//
// Solidity: function setSlidingEndTimeForBidsEnabled(bool _enable) returns()
func (_Paintswap *PaintswapTransactor) SetSlidingEndTimeForBidsEnabled(opts *bind.TransactOpts, _enable bool) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "setSlidingEndTimeForBidsEnabled", _enable)
}

// SetSlidingEndTimeForBidsEnabled is a paid mutator transaction binding the contract method 0xb3f2f251.
//
// Solidity: function setSlidingEndTimeForBidsEnabled(bool _enable) returns()
func (_Paintswap *PaintswapSession) SetSlidingEndTimeForBidsEnabled(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetSlidingEndTimeForBidsEnabled(&_Paintswap.TransactOpts, _enable)
}

// SetSlidingEndTimeForBidsEnabled is a paid mutator transaction binding the contract method 0xb3f2f251.
//
// Solidity: function setSlidingEndTimeForBidsEnabled(bool _enable) returns()
func (_Paintswap *PaintswapTransactorSession) SetSlidingEndTimeForBidsEnabled(_enable bool) (*types.Transaction, error) {
	return _Paintswap.Contract.SetSlidingEndTimeForBidsEnabled(&_Paintswap.TransactOpts, _enable)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.TransferOwnership(&_Paintswap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Paintswap *PaintswapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.TransferOwnership(&_Paintswap.TransactOpts, newOwner)
}

// WithdrawExcessNFTs is a paid mutator transaction binding the contract method 0x531372ce.
//
// Solidity: function withdrawExcessNFTs(address _nft, uint256 _tokenId) returns()
func (_Paintswap *PaintswapTransactor) WithdrawExcessNFTs(opts *bind.TransactOpts, _nft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "withdrawExcessNFTs", _nft, _tokenId)
}

// WithdrawExcessNFTs is a paid mutator transaction binding the contract method 0x531372ce.
//
// Solidity: function withdrawExcessNFTs(address _nft, uint256 _tokenId) returns()
func (_Paintswap *PaintswapSession) WithdrawExcessNFTs(_nft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.WithdrawExcessNFTs(&_Paintswap.TransactOpts, _nft, _tokenId)
}

// WithdrawExcessNFTs is a paid mutator transaction binding the contract method 0x531372ce.
//
// Solidity: function withdrawExcessNFTs(address _nft, uint256 _tokenId) returns()
func (_Paintswap *PaintswapTransactorSession) WithdrawExcessNFTs(_nft common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Paintswap.Contract.WithdrawExcessNFTs(&_Paintswap.TransactOpts, _nft, _tokenId)
}

// WithdrawExcessTokens is a paid mutator transaction binding the contract method 0x0d842561.
//
// Solidity: function withdrawExcessTokens(address _token) returns()
func (_Paintswap *PaintswapTransactor) WithdrawExcessTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Paintswap.contract.Transact(opts, "withdrawExcessTokens", _token)
}

// WithdrawExcessTokens is a paid mutator transaction binding the contract method 0x0d842561.
//
// Solidity: function withdrawExcessTokens(address _token) returns()
func (_Paintswap *PaintswapSession) WithdrawExcessTokens(_token common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.WithdrawExcessTokens(&_Paintswap.TransactOpts, _token)
}

// WithdrawExcessTokens is a paid mutator transaction binding the contract method 0x0d842561.
//
// Solidity: function withdrawExcessTokens(address _token) returns()
func (_Paintswap *PaintswapTransactorSession) WithdrawExcessTokens(_token common.Address) (*types.Transaction, error) {
	return _Paintswap.Contract.WithdrawExcessTokens(&_Paintswap.TransactOpts, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Paintswap.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapSession) Receive() (*types.Transaction, error) {
	return _Paintswap.Contract.Receive(&_Paintswap.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Paintswap *PaintswapTransactorSession) Receive() (*types.Transaction, error) {
	return _Paintswap.Contract.Receive(&_Paintswap.TransactOpts)
}

// PaintswapCancelledSaleIterator is returned from FilterCancelledSale and is used to iterate over the raw logs and unpacked data for CancelledSale events raised by the Paintswap contract.
type PaintswapCancelledSaleIterator struct {
	Event *PaintswapCancelledSale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapCancelledSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapCancelledSale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapCancelledSale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapCancelledSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapCancelledSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapCancelledSale represents a CancelledSale event raised by the Paintswap contract.
type PaintswapCancelledSale struct {
	MarketplaceId *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	AmountBatches []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCancelledSale is a free log retrieval operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) FilterCancelledSale(opts *bind.FilterOpts) (*PaintswapCancelledSaleIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "CancelledSale")
	if err != nil {
		return nil, err
	}
	return &PaintswapCancelledSaleIterator{contract: _Paintswap.contract, event: "CancelledSale", logs: logs, sub: sub}, nil
}

// WatchCancelledSale is a free log subscription operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) WatchCancelledSale(opts *bind.WatchOpts, sink chan<- *PaintswapCancelledSale) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "CancelledSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapCancelledSale)
				if err := _Paintswap.contract.UnpackLog(event, "CancelledSale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCancelledSale is a log parse operation binding the contract event 0x91ee9eb513284d27f660ec3907ad4fd27415782b1faa9b5fa46898dcb2f6983a.
//
// Solidity: event CancelledSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches)
func (_Paintswap *PaintswapFilterer) ParseCancelledSale(log types.Log) (*PaintswapCancelledSale, error) {
	event := new(PaintswapCancelledSale)
	if err := _Paintswap.contract.UnpackLog(event, "CancelledSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewBidIterator is returned from FilterNewBid and is used to iterate over the raw logs and unpacked data for NewBid events raised by the Paintswap contract.
type PaintswapNewBidIterator struct {
	Event *PaintswapNewBid // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapNewBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewBid)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapNewBid)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapNewBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewBid represents a NewBid event raised by the Paintswap contract.
type PaintswapNewBid struct {
	MarketplaceId *big.Int
	Bidder        common.Address
	Bid           *big.Int
	NextMinimum   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewBid is a free log retrieval operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) FilterNewBid(opts *bind.FilterOpts, bidder []common.Address) (*PaintswapNewBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewBid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewBidIterator{contract: _Paintswap.contract, event: "NewBid", logs: logs, sub: sub}, nil
}

// WatchNewBid is a free log subscription operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) WatchNewBid(opts *bind.WatchOpts, sink chan<- *PaintswapNewBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewBid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewBid)
				if err := _Paintswap.contract.UnpackLog(event, "NewBid", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBid is a log parse operation binding the contract event 0xd911ff5c673055b244483530dcc26f6fb4089b5883f1aa26c97b3dcf954b29f0.
//
// Solidity: event NewBid(uint256 marketplaceId, address indexed bidder, uint128 bid, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) ParseNewBid(log types.Log) (*PaintswapNewBid, error) {
	event := new(PaintswapNewBid)
	if err := _Paintswap.contract.UnpackLog(event, "NewBid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewOfferIterator is returned from FilterNewOffer and is used to iterate over the raw logs and unpacked data for NewOffer events raised by the Paintswap contract.
type PaintswapNewOfferIterator struct {
	Event *PaintswapNewOffer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapNewOfferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewOffer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapNewOffer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapNewOfferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewOfferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewOffer represents a NewOffer event raised by the Paintswap contract.
type PaintswapNewOffer struct {
	MarketplaceId *big.Int
	Offerer       common.Address
	Offer         *big.Int
	NextMinimum   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewOffer is a free log retrieval operation binding the contract event 0x7b7527a3beed8a65165a094535948d1c0b3a29739992c0fcba2528012f6a26ec.
//
// Solidity: event NewOffer(uint256 marketplaceId, address indexed offerer, uint128 offer, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) FilterNewOffer(opts *bind.FilterOpts, offerer []common.Address) (*PaintswapNewOfferIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewOffer", offererRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapNewOfferIterator{contract: _Paintswap.contract, event: "NewOffer", logs: logs, sub: sub}, nil
}

// WatchNewOffer is a free log subscription operation binding the contract event 0x7b7527a3beed8a65165a094535948d1c0b3a29739992c0fcba2528012f6a26ec.
//
// Solidity: event NewOffer(uint256 marketplaceId, address indexed offerer, uint128 offer, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) WatchNewOffer(opts *bind.WatchOpts, sink chan<- *PaintswapNewOffer, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewOffer", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewOffer)
				if err := _Paintswap.contract.UnpackLog(event, "NewOffer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewOffer is a log parse operation binding the contract event 0x7b7527a3beed8a65165a094535948d1c0b3a29739992c0fcba2528012f6a26ec.
//
// Solidity: event NewOffer(uint256 marketplaceId, address indexed offerer, uint128 offer, uint256 nextMinimum)
func (_Paintswap *PaintswapFilterer) ParseNewOffer(log types.Log) (*PaintswapNewOffer, error) {
	event := new(PaintswapNewOffer)
	if err := _Paintswap.contract.UnpackLog(event, "NewOffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapNewSaleIterator is returned from FilterNewSale and is used to iterate over the raw logs and unpacked data for NewSale events raised by the Paintswap contract.
type PaintswapNewSaleIterator struct {
	Event *PaintswapNewSale // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapNewSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapNewSale)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapNewSale)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapNewSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapNewSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapNewSale represents a NewSale event raised by the Paintswap contract.
type PaintswapNewSale struct {
	MarketplaceId  *big.Int
	Nfts           []common.Address
	TokenIds       []*big.Int
	AmountBatches  []*big.Int
	Price          *big.Int
	Duration       *big.Int
	IsAuction      bool
	Amount         *big.Int
	IsNSFW         bool
	MarketplaceURI string
	SearchKeywords string
	Addresses      []common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewSale is a free log retrieval operation binding the contract event 0xb3cd69c2fae51b3e21f25978f3604c96acf27abbf4fcb05685b102b4a567cb1d.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, uint256 amount, bool isNSFW, string marketplaceURI, string searchKeywords, address[] addresses)
func (_Paintswap *PaintswapFilterer) FilterNewSale(opts *bind.FilterOpts) (*PaintswapNewSaleIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "NewSale")
	if err != nil {
		return nil, err
	}
	return &PaintswapNewSaleIterator{contract: _Paintswap.contract, event: "NewSale", logs: logs, sub: sub}, nil
}

// WatchNewSale is a free log subscription operation binding the contract event 0xb3cd69c2fae51b3e21f25978f3604c96acf27abbf4fcb05685b102b4a567cb1d.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, uint256 amount, bool isNSFW, string marketplaceURI, string searchKeywords, address[] addresses)
func (_Paintswap *PaintswapFilterer) WatchNewSale(opts *bind.WatchOpts, sink chan<- *PaintswapNewSale) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "NewSale")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapNewSale)
				if err := _Paintswap.contract.UnpackLog(event, "NewSale", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSale is a log parse operation binding the contract event 0xb3cd69c2fae51b3e21f25978f3604c96acf27abbf4fcb05685b102b4a567cb1d.
//
// Solidity: event NewSale(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, uint256 duration, bool isAuction, uint256 amount, bool isNSFW, string marketplaceURI, string searchKeywords, address[] addresses)
func (_Paintswap *PaintswapFilterer) ParseNewSale(log types.Log) (*PaintswapNewSale, error) {
	event := new(PaintswapNewSale)
	if err := _Paintswap.contract.UnpackLog(event, "NewSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Paintswap contract.
type PaintswapOwnershipTransferredIterator struct {
	Event *PaintswapOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapOwnershipTransferred represents a OwnershipTransferred event raised by the Paintswap contract.
type PaintswapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paintswap *PaintswapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaintswapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaintswapOwnershipTransferredIterator{contract: _Paintswap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paintswap *PaintswapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaintswapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapOwnershipTransferred)
				if err := _Paintswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Paintswap *PaintswapFilterer) ParseOwnershipTransferred(log types.Log) (*PaintswapOwnershipTransferred, error) {
	event := new(PaintswapOwnershipTransferred)
	if err := _Paintswap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapSaleFinishedIterator is returned from FilterSaleFinished and is used to iterate over the raw logs and unpacked data for SaleFinished events raised by the Paintswap contract.
type PaintswapSaleFinishedIterator struct {
	Event *PaintswapSaleFinished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapSaleFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapSaleFinished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapSaleFinished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapSaleFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapSaleFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapSaleFinished represents a SaleFinished event raised by the Paintswap contract.
type PaintswapSaleFinished struct {
	MarketplaceId *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	AmountBatches []*big.Int
	FailedSellAll bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSaleFinished is a free log retrieval operation binding the contract event 0xd1f4aab9b0d4b94e26beff68299512c1e11b250c148d243d4a9dd10bed2dc514.
//
// Solidity: event SaleFinished(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, bool failedSellAll)
func (_Paintswap *PaintswapFilterer) FilterSaleFinished(opts *bind.FilterOpts) (*PaintswapSaleFinishedIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "SaleFinished")
	if err != nil {
		return nil, err
	}
	return &PaintswapSaleFinishedIterator{contract: _Paintswap.contract, event: "SaleFinished", logs: logs, sub: sub}, nil
}

// WatchSaleFinished is a free log subscription operation binding the contract event 0xd1f4aab9b0d4b94e26beff68299512c1e11b250c148d243d4a9dd10bed2dc514.
//
// Solidity: event SaleFinished(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, bool failedSellAll)
func (_Paintswap *PaintswapFilterer) WatchSaleFinished(opts *bind.WatchOpts, sink chan<- *PaintswapSaleFinished) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "SaleFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapSaleFinished)
				if err := _Paintswap.contract.UnpackLog(event, "SaleFinished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSaleFinished is a log parse operation binding the contract event 0xd1f4aab9b0d4b94e26beff68299512c1e11b250c148d243d4a9dd10bed2dc514.
//
// Solidity: event SaleFinished(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, bool failedSellAll)
func (_Paintswap *PaintswapFilterer) ParseSaleFinished(log types.Log) (*PaintswapSaleFinished, error) {
	event := new(PaintswapSaleFinished)
	if err := _Paintswap.contract.UnpackLog(event, "SaleFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapSoldIterator is returned from FilterSold and is used to iterate over the raw logs and unpacked data for Sold events raised by the Paintswap contract.
type PaintswapSoldIterator struct {
	Event *PaintswapSold // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapSold)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapSold)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapSold represents a Sold event raised by the Paintswap contract.
type PaintswapSold struct {
	MarketplaceId *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	AmountBatches []*big.Int
	Price         *big.Int
	Buyer         common.Address
	Seller        common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSold is a free log retrieval operation binding the contract event 0x0cda439d506dbc3b73fe10f062cf285c4e75fe85d310decf4b8239841879ed61.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount)
func (_Paintswap *PaintswapFilterer) FilterSold(opts *bind.FilterOpts) (*PaintswapSoldIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return &PaintswapSoldIterator{contract: _Paintswap.contract, event: "Sold", logs: logs, sub: sub}, nil
}

// WatchSold is a free log subscription operation binding the contract event 0x0cda439d506dbc3b73fe10f062cf285c4e75fe85d310decf4b8239841879ed61.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount)
func (_Paintswap *PaintswapFilterer) WatchSold(opts *bind.WatchOpts, sink chan<- *PaintswapSold) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "Sold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapSold)
				if err := _Paintswap.contract.UnpackLog(event, "Sold", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSold is a log parse operation binding the contract event 0x0cda439d506dbc3b73fe10f062cf285c4e75fe85d310decf4b8239841879ed61.
//
// Solidity: event Sold(uint256 marketplaceId, address[] nfts, uint256[] tokenIds, uint256[] amountBatches, uint128 price, address buyer, address seller, uint256 amount)
func (_Paintswap *PaintswapFilterer) ParseSold(log types.Log) (*PaintswapSold, error) {
	event := new(PaintswapSold)
	if err := _Paintswap.contract.UnpackLog(event, "Sold", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateEndTimeIterator is returned from FilterUpdateEndTime and is used to iterate over the raw logs and unpacked data for UpdateEndTime events raised by the Paintswap contract.
type PaintswapUpdateEndTimeIterator struct {
	Event *PaintswapUpdateEndTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapUpdateEndTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateEndTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapUpdateEndTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapUpdateEndTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateEndTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateEndTime represents a UpdateEndTime event raised by the Paintswap contract.
type PaintswapUpdateEndTime struct {
	MarketplaceId *big.Int
	EndTime       *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateEndTime is a free log retrieval operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) FilterUpdateEndTime(opts *bind.FilterOpts) (*PaintswapUpdateEndTimeIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateEndTime")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateEndTimeIterator{contract: _Paintswap.contract, event: "UpdateEndTime", logs: logs, sub: sub}, nil
}

// WatchUpdateEndTime is a free log subscription operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) WatchUpdateEndTime(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateEndTime) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateEndTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateEndTime)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateEndTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateEndTime is a log parse operation binding the contract event 0x17d3c533548aa281cd4b573e7659f1b3269b772da8e4fc6425a0cf9e662c43e9.
//
// Solidity: event UpdateEndTime(uint256 marketplaceId, uint256 endTime)
func (_Paintswap *PaintswapFilterer) ParseUpdateEndTime(log types.Log) (*PaintswapUpdateEndTime, error) {
	event := new(PaintswapUpdateEndTime)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateEndTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdatePriceIterator is returned from FilterUpdatePrice and is used to iterate over the raw logs and unpacked data for UpdatePrice events raised by the Paintswap contract.
type PaintswapUpdatePriceIterator struct {
	Event *PaintswapUpdatePrice // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapUpdatePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdatePrice)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapUpdatePrice)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapUpdatePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdatePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdatePrice represents a UpdatePrice event raised by the Paintswap contract.
type PaintswapUpdatePrice struct {
	MarketplaceId *big.Int
	Price         *big.Int
	Nfts          []common.Address
	TokenIds      []*big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdatePrice is a free log retrieval operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) FilterUpdatePrice(opts *bind.FilterOpts) (*PaintswapUpdatePriceIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdatePriceIterator{contract: _Paintswap.contract, event: "UpdatePrice", logs: logs, sub: sub}, nil
}

// WatchUpdatePrice is a free log subscription operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) WatchUpdatePrice(opts *bind.WatchOpts, sink chan<- *PaintswapUpdatePrice) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdatePrice")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdatePrice)
				if err := _Paintswap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdatePrice is a log parse operation binding the contract event 0xcc2e03b72a9d3fb70264a784dade2978f8659ed6c85646480ae2ba15e8c5545b.
//
// Solidity: event UpdatePrice(uint256 marketplaceId, uint128 price, address[] nfts, uint256[] tokenIds)
func (_Paintswap *PaintswapFilterer) ParseUpdatePrice(log types.Log) (*PaintswapUpdatePrice, error) {
	event := new(PaintswapUpdatePrice)
	if err := _Paintswap.contract.UnpackLog(event, "UpdatePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaintswapUpdateStartTimeIterator is returned from FilterUpdateStartTime and is used to iterate over the raw logs and unpacked data for UpdateStartTime events raised by the Paintswap contract.
type PaintswapUpdateStartTimeIterator struct {
	Event *PaintswapUpdateStartTime // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PaintswapUpdateStartTimeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaintswapUpdateStartTime)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PaintswapUpdateStartTime)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PaintswapUpdateStartTimeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaintswapUpdateStartTimeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaintswapUpdateStartTime represents a UpdateStartTime event raised by the Paintswap contract.
type PaintswapUpdateStartTime struct {
	MarketplaceId *big.Int
	StartTime     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUpdateStartTime is a free log retrieval operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) FilterUpdateStartTime(opts *bind.FilterOpts) (*PaintswapUpdateStartTimeIterator, error) {

	logs, sub, err := _Paintswap.contract.FilterLogs(opts, "UpdateStartTime")
	if err != nil {
		return nil, err
	}
	return &PaintswapUpdateStartTimeIterator{contract: _Paintswap.contract, event: "UpdateStartTime", logs: logs, sub: sub}, nil
}

// WatchUpdateStartTime is a free log subscription operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) WatchUpdateStartTime(opts *bind.WatchOpts, sink chan<- *PaintswapUpdateStartTime) (event.Subscription, error) {

	logs, sub, err := _Paintswap.contract.WatchLogs(opts, "UpdateStartTime")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaintswapUpdateStartTime)
				if err := _Paintswap.contract.UnpackLog(event, "UpdateStartTime", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateStartTime is a log parse operation binding the contract event 0x1367fc9729096aae2ba55403f0591c3fb92012e1e32beaf69a680024f1016c5a.
//
// Solidity: event UpdateStartTime(uint256 marketplaceId, uint256 startTime)
func (_Paintswap *PaintswapFilterer) ParseUpdateStartTime(log types.Log) (*PaintswapUpdateStartTime, error) {
	event := new(PaintswapUpdateStartTime)
	if err := _Paintswap.contract.UnpackLog(event, "UpdateStartTime", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
