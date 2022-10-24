// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nftkey

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

// INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty struct {
	Recipient   common.Address
	FeeFraction *big.Int
	SetBy       common.Address
}

// INFTKEYMarketplaceV2Bid is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceV2Bid struct {
	TokenId         *big.Int
	Value           *big.Int
	Bidder          common.Address
	ExpireTimestamp *big.Int
}

// INFTKEYMarketplaceV2Listing is an auto generated low-level Go binding around an user-defined struct.
type INFTKEYMarketplaceV2Listing struct {
	TokenId         *big.Int
	Value           *big.Int
	Seller          common.Address
	ExpireTimestamp *big.Int
}

// NFTkeyMetaData contains all meta data concerning the NFTkey contract.
var NFTkeyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_paymentTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"SetRoyalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyFee\",\"type\":\"uint256\"}],\"name\":\"TokenBidAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"TokenBidEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"bid\",\"type\":\"tuple\"}],\"name\":\"TokenBidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"serviceFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royaltyFee\",\"type\":\"uint256\"}],\"name\":\"TokenBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"TokenDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"listing\",\"type\":\"tuple\"}],\"name\":\"TokenListed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"acceptBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionTimeOutRangeMax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"actionTimeOutRangeMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"}],\"name\":\"changeMarketplaceStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeInSec\",\"type\":\"uint256\"}],\"name\":\"changeMaxActionTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timeInSec\",\"type\":\"uint256\"}],\"name\":\"changeMinActionTimeLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"serviceFeeFraction_\",\"type\":\"uint8\"}],\"name\":\"changeSeriveFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"defaultRoyaltyFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"delistToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"name\":\"enterBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getBidderBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"bidderBids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getBidderTokenBid\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"validBid\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"bids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenHighestBid\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid\",\"name\":\"highestBid\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getTokenHighestBids\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Bid[]\",\"name\":\"highestBids\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getTokenListing\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Listing\",\"name\":\"validListing\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"getTokenListings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"internalType\":\"structINFTKEYMarketplaceV2.Listing[]\",\"name\":\"listings\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTradingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expireTimestamp\",\"type\":\"uint256\"}],\"name\":\"listToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"numTokenListings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"numTokenWithBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"}],\"name\":\"royalty\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"internalType\":\"structINFTKEYMarketplaceRoyalty.ERC721CollectionRoyalty\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyUpperLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceFee\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeFraction\",\"type\":\"uint256\"}],\"name\":\"setRoyaltyForCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newUpperLimit\",\"type\":\"uint256\"}],\"name\":\"updateRoyaltyUpperLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"erc721Address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"withdrawBidForToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NFTkeyABI is the input ABI used to generate the binding from.
// Deprecated: Use NFTkeyMetaData.ABI instead.
var NFTkeyABI = NFTkeyMetaData.ABI

// NFTkey is an auto generated Go binding around an Ethereum contract.
type NFTkey struct {
	NFTkeyCaller     // Read-only binding to the contract
	NFTkeyTransactor // Write-only binding to the contract
	NFTkeyFilterer   // Log filterer for contract events
}

// NFTkeyCaller is an auto generated read-only Go binding around an Ethereum contract.
type NFTkeyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTkeyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NFTkeyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTkeyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NFTkeyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NFTkeySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NFTkeySession struct {
	Contract     *NFTkey           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTkeyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NFTkeyCallerSession struct {
	Contract *NFTkeyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NFTkeyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NFTkeyTransactorSession struct {
	Contract     *NFTkeyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NFTkeyRaw is an auto generated low-level Go binding around an Ethereum contract.
type NFTkeyRaw struct {
	Contract *NFTkey // Generic contract binding to access the raw methods on
}

// NFTkeyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NFTkeyCallerRaw struct {
	Contract *NFTkeyCaller // Generic read-only contract binding to access the raw methods on
}

// NFTkeyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NFTkeyTransactorRaw struct {
	Contract *NFTkeyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNFTkey creates a new instance of NFTkey, bound to a specific deployed contract.
func NewNFTkey(address common.Address, backend bind.ContractBackend) (*NFTkey, error) {
	contract, err := bindNFTkey(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NFTkey{NFTkeyCaller: NFTkeyCaller{contract: contract}, NFTkeyTransactor: NFTkeyTransactor{contract: contract}, NFTkeyFilterer: NFTkeyFilterer{contract: contract}}, nil
}

// NewNFTkeyCaller creates a new read-only instance of NFTkey, bound to a specific deployed contract.
func NewNFTkeyCaller(address common.Address, caller bind.ContractCaller) (*NFTkeyCaller, error) {
	contract, err := bindNFTkey(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NFTkeyCaller{contract: contract}, nil
}

// NewNFTkeyTransactor creates a new write-only instance of NFTkey, bound to a specific deployed contract.
func NewNFTkeyTransactor(address common.Address, transactor bind.ContractTransactor) (*NFTkeyTransactor, error) {
	contract, err := bindNFTkey(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTransactor{contract: contract}, nil
}

// NewNFTkeyFilterer creates a new log filterer instance of NFTkey, bound to a specific deployed contract.
func NewNFTkeyFilterer(address common.Address, filterer bind.ContractFilterer) (*NFTkeyFilterer, error) {
	contract, err := bindNFTkey(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NFTkeyFilterer{contract: contract}, nil
}

// bindNFTkey binds a generic wrapper to an already deployed contract.
func bindNFTkey(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NFTkeyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTkey *NFTkeyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTkey.Contract.NFTkeyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTkey *NFTkeyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTkey.Contract.NFTkeyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTkey *NFTkeyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTkey.Contract.NFTkeyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NFTkey *NFTkeyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NFTkey.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NFTkey *NFTkeyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTkey.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NFTkey *NFTkeyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NFTkey.Contract.contract.Transact(opts, method, params...)
}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_NFTkey *NFTkeyCaller) ActionTimeOutRangeMax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "actionTimeOutRangeMax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_NFTkey *NFTkeySession) ActionTimeOutRangeMax() (*big.Int, error) {
	return _NFTkey.Contract.ActionTimeOutRangeMax(&_NFTkey.CallOpts)
}

// ActionTimeOutRangeMax is a free data retrieval call binding the contract method 0x453dfc50.
//
// Solidity: function actionTimeOutRangeMax() view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) ActionTimeOutRangeMax() (*big.Int, error) {
	return _NFTkey.Contract.ActionTimeOutRangeMax(&_NFTkey.CallOpts)
}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_NFTkey *NFTkeyCaller) ActionTimeOutRangeMin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "actionTimeOutRangeMin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_NFTkey *NFTkeySession) ActionTimeOutRangeMin() (*big.Int, error) {
	return _NFTkey.Contract.ActionTimeOutRangeMin(&_NFTkey.CallOpts)
}

// ActionTimeOutRangeMin is a free data retrieval call binding the contract method 0x33549d3d.
//
// Solidity: function actionTimeOutRangeMin() view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) ActionTimeOutRangeMin() (*big.Int, error) {
	return _NFTkey.Contract.ActionTimeOutRangeMin(&_NFTkey.CallOpts)
}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_NFTkey *NFTkeyCaller) DefaultRoyaltyFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "defaultRoyaltyFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_NFTkey *NFTkeySession) DefaultRoyaltyFraction() (*big.Int, error) {
	return _NFTkey.Contract.DefaultRoyaltyFraction(&_NFTkey.CallOpts)
}

// DefaultRoyaltyFraction is a free data retrieval call binding the contract method 0x6d0042b8.
//
// Solidity: function defaultRoyaltyFraction() view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) DefaultRoyaltyFraction() (*big.Int, error) {
	return _NFTkey.Contract.DefaultRoyaltyFraction(&_NFTkey.CallOpts)
}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_NFTkey *NFTkeyCaller) GetBidderBids(opts *bind.CallOpts, erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getBidderBids", erc721Address, bidder, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_NFTkey *NFTkeySession) GetBidderBids(erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetBidderBids(&_NFTkey.CallOpts, erc721Address, bidder, from, size)
}

// GetBidderBids is a free data retrieval call binding the contract method 0xcf6ac953.
//
// Solidity: function getBidderBids(address erc721Address, address bidder, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] bidderBids)
func (_NFTkey *NFTkeyCallerSession) GetBidderBids(erc721Address common.Address, bidder common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetBidderBids(&_NFTkey.CallOpts, erc721Address, bidder, from, size)
}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_NFTkey *NFTkeyCaller) GetBidderTokenBid(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getBidderTokenBid", erc721Address, tokenId, bidder)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Bid)).(*INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_NFTkey *NFTkeySession) GetBidderTokenBid(erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetBidderTokenBid(&_NFTkey.CallOpts, erc721Address, tokenId, bidder)
}

// GetBidderTokenBid is a free data retrieval call binding the contract method 0x90bc4e37.
//
// Solidity: function getBidderTokenBid(address erc721Address, uint256 tokenId, address bidder) view returns((uint256,uint256,address,uint256) validBid)
func (_NFTkey *NFTkeyCallerSession) GetBidderTokenBid(erc721Address common.Address, tokenId *big.Int, bidder common.Address) (INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetBidderTokenBid(&_NFTkey.CallOpts, erc721Address, tokenId, bidder)
}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_NFTkey *NFTkeyCaller) GetTokenBids(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getTokenBids", erc721Address, tokenId)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_NFTkey *NFTkeySession) GetTokenBids(erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenBids(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenBids is a free data retrieval call binding the contract method 0x66c1e8bf.
//
// Solidity: function getTokenBids(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256)[] bids)
func (_NFTkey *NFTkeyCallerSession) GetTokenBids(erc721Address common.Address, tokenId *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenBids(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_NFTkey *NFTkeyCaller) GetTokenHighestBid(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getTokenHighestBid", erc721Address, tokenId)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Bid)).(*INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_NFTkey *NFTkeySession) GetTokenHighestBid(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenHighestBid(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBid is a free data retrieval call binding the contract method 0x1f77a820.
//
// Solidity: function getTokenHighestBid(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) highestBid)
func (_NFTkey *NFTkeyCallerSession) GetTokenHighestBid(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenHighestBid(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_NFTkey *NFTkeyCaller) GetTokenHighestBids(opts *bind.CallOpts, erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getTokenHighestBids", erc721Address, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Bid), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Bid)).(*[]INFTKEYMarketplaceV2Bid)

	return out0, err

}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_NFTkey *NFTkeySession) GetTokenHighestBids(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenHighestBids(&_NFTkey.CallOpts, erc721Address, from, size)
}

// GetTokenHighestBids is a free data retrieval call binding the contract method 0x2c20d1b4.
//
// Solidity: function getTokenHighestBids(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] highestBids)
func (_NFTkey *NFTkeyCallerSession) GetTokenHighestBids(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Bid, error) {
	return _NFTkey.Contract.GetTokenHighestBids(&_NFTkey.CallOpts, erc721Address, from, size)
}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_NFTkey *NFTkeyCaller) GetTokenListing(opts *bind.CallOpts, erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getTokenListing", erc721Address, tokenId)

	if err != nil {
		return *new(INFTKEYMarketplaceV2Listing), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceV2Listing)).(*INFTKEYMarketplaceV2Listing)

	return out0, err

}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_NFTkey *NFTkeySession) GetTokenListing(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	return _NFTkey.Contract.GetTokenListing(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenListing is a free data retrieval call binding the contract method 0x96f97164.
//
// Solidity: function getTokenListing(address erc721Address, uint256 tokenId) view returns((uint256,uint256,address,uint256) validListing)
func (_NFTkey *NFTkeyCallerSession) GetTokenListing(erc721Address common.Address, tokenId *big.Int) (INFTKEYMarketplaceV2Listing, error) {
	return _NFTkey.Contract.GetTokenListing(&_NFTkey.CallOpts, erc721Address, tokenId)
}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_NFTkey *NFTkeyCaller) GetTokenListings(opts *bind.CallOpts, erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "getTokenListings", erc721Address, from, size)

	if err != nil {
		return *new([]INFTKEYMarketplaceV2Listing), err
	}

	out0 := *abi.ConvertType(out[0], new([]INFTKEYMarketplaceV2Listing)).(*[]INFTKEYMarketplaceV2Listing)

	return out0, err

}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_NFTkey *NFTkeySession) GetTokenListings(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	return _NFTkey.Contract.GetTokenListings(&_NFTkey.CallOpts, erc721Address, from, size)
}

// GetTokenListings is a free data retrieval call binding the contract method 0x45105cb1.
//
// Solidity: function getTokenListings(address erc721Address, uint256 from, uint256 size) view returns((uint256,uint256,address,uint256)[] listings)
func (_NFTkey *NFTkeyCallerSession) GetTokenListings(erc721Address common.Address, from *big.Int, size *big.Int) ([]INFTKEYMarketplaceV2Listing, error) {
	return _NFTkey.Contract.GetTokenListings(&_NFTkey.CallOpts, erc721Address, from, size)
}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_NFTkey *NFTkeyCaller) IsTradingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "isTradingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_NFTkey *NFTkeySession) IsTradingEnabled() (bool, error) {
	return _NFTkey.Contract.IsTradingEnabled(&_NFTkey.CallOpts)
}

// IsTradingEnabled is a free data retrieval call binding the contract method 0x064a59d0.
//
// Solidity: function isTradingEnabled() view returns(bool)
func (_NFTkey *NFTkeyCallerSession) IsTradingEnabled() (bool, error) {
	return _NFTkey.Contract.IsTradingEnabled(&_NFTkey.CallOpts)
}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeyCaller) NumTokenListings(opts *bind.CallOpts, erc721Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "numTokenListings", erc721Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeySession) NumTokenListings(erc721Address common.Address) (*big.Int, error) {
	return _NFTkey.Contract.NumTokenListings(&_NFTkey.CallOpts, erc721Address)
}

// NumTokenListings is a free data retrieval call binding the contract method 0xf8e7b936.
//
// Solidity: function numTokenListings(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) NumTokenListings(erc721Address common.Address) (*big.Int, error) {
	return _NFTkey.Contract.NumTokenListings(&_NFTkey.CallOpts, erc721Address)
}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeyCaller) NumTokenWithBids(opts *bind.CallOpts, erc721Address common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "numTokenWithBids", erc721Address)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeySession) NumTokenWithBids(erc721Address common.Address) (*big.Int, error) {
	return _NFTkey.Contract.NumTokenWithBids(&_NFTkey.CallOpts, erc721Address)
}

// NumTokenWithBids is a free data retrieval call binding the contract method 0xeb635ab8.
//
// Solidity: function numTokenWithBids(address erc721Address) view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) NumTokenWithBids(erc721Address common.Address) (*big.Int, error) {
	return _NFTkey.Contract.NumTokenWithBids(&_NFTkey.CallOpts, erc721Address)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTkey *NFTkeyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTkey *NFTkeySession) Owner() (common.Address, error) {
	return _NFTkey.Contract.Owner(&_NFTkey.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NFTkey *NFTkeyCallerSession) Owner() (common.Address, error) {
	return _NFTkey.Contract.Owner(&_NFTkey.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_NFTkey *NFTkeyCaller) PaymentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "paymentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_NFTkey *NFTkeySession) PaymentToken() (common.Address, error) {
	return _NFTkey.Contract.PaymentToken(&_NFTkey.CallOpts)
}

// PaymentToken is a free data retrieval call binding the contract method 0x3013ce29.
//
// Solidity: function paymentToken() view returns(address)
func (_NFTkey *NFTkeyCallerSession) PaymentToken() (common.Address, error) {
	return _NFTkey.Contract.PaymentToken(&_NFTkey.CallOpts)
}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_NFTkey *NFTkeyCaller) Royalty(opts *bind.CallOpts, erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "royalty", erc721Address)

	if err != nil {
		return *new(INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty), err
	}

	out0 := *abi.ConvertType(out[0], new(INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty)).(*INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty)

	return out0, err

}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_NFTkey *NFTkeySession) Royalty(erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	return _NFTkey.Contract.Royalty(&_NFTkey.CallOpts, erc721Address)
}

// Royalty is a free data retrieval call binding the contract method 0x861b69d6.
//
// Solidity: function royalty(address erc721Address) view returns((address,uint256,address))
func (_NFTkey *NFTkeyCallerSession) Royalty(erc721Address common.Address) (INFTKEYMarketplaceRoyaltyERC721CollectionRoyalty, error) {
	return _NFTkey.Contract.Royalty(&_NFTkey.CallOpts, erc721Address)
}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_NFTkey *NFTkeyCaller) RoyaltyUpperLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "royaltyUpperLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_NFTkey *NFTkeySession) RoyaltyUpperLimit() (*big.Int, error) {
	return _NFTkey.Contract.RoyaltyUpperLimit(&_NFTkey.CallOpts)
}

// RoyaltyUpperLimit is a free data retrieval call binding the contract method 0xee8ef34d.
//
// Solidity: function royaltyUpperLimit() view returns(uint256)
func (_NFTkey *NFTkeyCallerSession) RoyaltyUpperLimit() (*big.Int, error) {
	return _NFTkey.Contract.RoyaltyUpperLimit(&_NFTkey.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_NFTkey *NFTkeyCaller) ServiceFee(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NFTkey.contract.Call(opts, &out, "serviceFee")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_NFTkey *NFTkeySession) ServiceFee() (uint8, error) {
	return _NFTkey.Contract.ServiceFee(&_NFTkey.CallOpts)
}

// ServiceFee is a free data retrieval call binding the contract method 0x8abdf5aa.
//
// Solidity: function serviceFee() view returns(uint8)
func (_NFTkey *NFTkeyCallerSession) ServiceFee() (uint8, error) {
	return _NFTkey.Contract.ServiceFee(&_NFTkey.CallOpts)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_NFTkey *NFTkeyTransactor) AcceptBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "acceptBidForToken", erc721Address, tokenId, bidder, value)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_NFTkey *NFTkeySession) AcceptBidForToken(erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.AcceptBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId, bidder, value)
}

// AcceptBidForToken is a paid mutator transaction binding the contract method 0xeb3e87b9.
//
// Solidity: function acceptBidForToken(address erc721Address, uint256 tokenId, address bidder, uint256 value) returns()
func (_NFTkey *NFTkeyTransactorSession) AcceptBidForToken(erc721Address common.Address, tokenId *big.Int, bidder common.Address, value *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.AcceptBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId, bidder, value)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_NFTkey *NFTkeyTransactor) BuyToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "buyToken", erc721Address, tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_NFTkey *NFTkeySession) BuyToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.BuyToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// BuyToken is a paid mutator transaction binding the contract method 0x68f8fc10.
//
// Solidity: function buyToken(address erc721Address, uint256 tokenId) payable returns()
func (_NFTkey *NFTkeyTransactorSession) BuyToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.BuyToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_NFTkey *NFTkeyTransactor) ChangeMarketplaceStatus(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "changeMarketplaceStatus", enabled)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_NFTkey *NFTkeySession) ChangeMarketplaceStatus(enabled bool) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMarketplaceStatus(&_NFTkey.TransactOpts, enabled)
}

// ChangeMarketplaceStatus is a paid mutator transaction binding the contract method 0xb6be53ba.
//
// Solidity: function changeMarketplaceStatus(bool enabled) returns()
func (_NFTkey *NFTkeyTransactorSession) ChangeMarketplaceStatus(enabled bool) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMarketplaceStatus(&_NFTkey.TransactOpts, enabled)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeyTransactor) ChangeMaxActionTimeLimit(opts *bind.TransactOpts, timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "changeMaxActionTimeLimit", timeInSec)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeySession) ChangeMaxActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMaxActionTimeLimit(&_NFTkey.TransactOpts, timeInSec)
}

// ChangeMaxActionTimeLimit is a paid mutator transaction binding the contract method 0xa3c0b5f0.
//
// Solidity: function changeMaxActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeyTransactorSession) ChangeMaxActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMaxActionTimeLimit(&_NFTkey.TransactOpts, timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeyTransactor) ChangeMinActionTimeLimit(opts *bind.TransactOpts, timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "changeMinActionTimeLimit", timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeySession) ChangeMinActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMinActionTimeLimit(&_NFTkey.TransactOpts, timeInSec)
}

// ChangeMinActionTimeLimit is a paid mutator transaction binding the contract method 0x2426fc24.
//
// Solidity: function changeMinActionTimeLimit(uint256 timeInSec) returns()
func (_NFTkey *NFTkeyTransactorSession) ChangeMinActionTimeLimit(timeInSec *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeMinActionTimeLimit(&_NFTkey.TransactOpts, timeInSec)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_NFTkey *NFTkeyTransactor) ChangeSeriveFee(opts *bind.TransactOpts, serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "changeSeriveFee", serviceFeeFraction_)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_NFTkey *NFTkeySession) ChangeSeriveFee(serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeSeriveFee(&_NFTkey.TransactOpts, serviceFeeFraction_)
}

// ChangeSeriveFee is a paid mutator transaction binding the contract method 0xf8ad6f62.
//
// Solidity: function changeSeriveFee(uint8 serviceFeeFraction_) returns()
func (_NFTkey *NFTkeyTransactorSession) ChangeSeriveFee(serviceFeeFraction_ uint8) (*types.Transaction, error) {
	return _NFTkey.Contract.ChangeSeriveFee(&_NFTkey.TransactOpts, serviceFeeFraction_)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeyTransactor) DelistToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "delistToken", erc721Address, tokenId)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeySession) DelistToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.DelistToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// DelistToken is a paid mutator transaction binding the contract method 0xfeb88406.
//
// Solidity: function delistToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeyTransactorSession) DelistToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.DelistToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeyTransactor) EnterBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "enterBidForToken", erc721Address, tokenId, value, expireTimestamp)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeySession) EnterBidForToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.EnterBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// EnterBidForToken is a paid mutator transaction binding the contract method 0x4313e9bd.
//
// Solidity: function enterBidForToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeyTransactorSession) EnterBidForToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.EnterBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeyTransactor) ListToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "listToken", erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeySession) ListToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ListToken(&_NFTkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// ListToken is a paid mutator transaction binding the contract method 0xb43d901d.
//
// Solidity: function listToken(address erc721Address, uint256 tokenId, uint256 value, uint256 expireTimestamp) returns()
func (_NFTkey *NFTkeyTransactorSession) ListToken(erc721Address common.Address, tokenId *big.Int, value *big.Int, expireTimestamp *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.ListToken(&_NFTkey.TransactOpts, erc721Address, tokenId, value, expireTimestamp)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTkey *NFTkeyTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTkey *NFTkeySession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTkey.Contract.RenounceOwnership(&_NFTkey.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NFTkey *NFTkeyTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NFTkey.Contract.RenounceOwnership(&_NFTkey.TransactOpts)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeyTransactor) SetRoyalty(opts *bind.TransactOpts, erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "setRoyalty", erc721Address, newRecipient, feeFraction)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeySession) SetRoyalty(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.SetRoyalty(&_NFTkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeyTransactorSession) SetRoyalty(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.SetRoyalty(&_NFTkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeyTransactor) SetRoyaltyForCollection(opts *bind.TransactOpts, erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "setRoyaltyForCollection", erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeySession) SetRoyaltyForCollection(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.SetRoyaltyForCollection(&_NFTkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// SetRoyaltyForCollection is a paid mutator transaction binding the contract method 0x19d334cb.
//
// Solidity: function setRoyaltyForCollection(address erc721Address, address newRecipient, uint256 feeFraction) returns()
func (_NFTkey *NFTkeyTransactorSession) SetRoyaltyForCollection(erc721Address common.Address, newRecipient common.Address, feeFraction *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.SetRoyaltyForCollection(&_NFTkey.TransactOpts, erc721Address, newRecipient, feeFraction)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTkey *NFTkeyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTkey *NFTkeySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTkey.Contract.TransferOwnership(&_NFTkey.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NFTkey *NFTkeyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NFTkey.Contract.TransferOwnership(&_NFTkey.TransactOpts, newOwner)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_NFTkey *NFTkeyTransactor) UpdateRoyaltyUpperLimit(opts *bind.TransactOpts, _newUpperLimit *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "updateRoyaltyUpperLimit", _newUpperLimit)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_NFTkey *NFTkeySession) UpdateRoyaltyUpperLimit(_newUpperLimit *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.UpdateRoyaltyUpperLimit(&_NFTkey.TransactOpts, _newUpperLimit)
}

// UpdateRoyaltyUpperLimit is a paid mutator transaction binding the contract method 0xcdea490d.
//
// Solidity: function updateRoyaltyUpperLimit(uint256 _newUpperLimit) returns()
func (_NFTkey *NFTkeyTransactorSession) UpdateRoyaltyUpperLimit(_newUpperLimit *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.UpdateRoyaltyUpperLimit(&_NFTkey.TransactOpts, _newUpperLimit)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeyTransactor) WithdrawBidForToken(opts *bind.TransactOpts, erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.contract.Transact(opts, "withdrawBidForToken", erc721Address, tokenId)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeySession) WithdrawBidForToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.WithdrawBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// WithdrawBidForToken is a paid mutator transaction binding the contract method 0x0609d095.
//
// Solidity: function withdrawBidForToken(address erc721Address, uint256 tokenId) returns()
func (_NFTkey *NFTkeyTransactorSession) WithdrawBidForToken(erc721Address common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NFTkey.Contract.WithdrawBidForToken(&_NFTkey.TransactOpts, erc721Address, tokenId)
}

// NFTkeyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NFTkey contract.
type NFTkeyOwnershipTransferredIterator struct {
	Event *NFTkeyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NFTkeyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyOwnershipTransferred)
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
		it.Event = new(NFTkeyOwnershipTransferred)
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
func (it *NFTkeyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyOwnershipTransferred represents a OwnershipTransferred event raised by the NFTkey contract.
type NFTkeyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTkey *NFTkeyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NFTkeyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyOwnershipTransferredIterator{contract: _NFTkey.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NFTkey *NFTkeyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NFTkeyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyOwnershipTransferred)
				if err := _NFTkey.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NFTkey *NFTkeyFilterer) ParseOwnershipTransferred(log types.Log) (*NFTkeyOwnershipTransferred, error) {
	event := new(NFTkeyOwnershipTransferred)
	if err := _NFTkey.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeySetRoyaltyIterator is returned from FilterSetRoyalty and is used to iterate over the raw logs and unpacked data for SetRoyalty events raised by the NFTkey contract.
type NFTkeySetRoyaltyIterator struct {
	Event *NFTkeySetRoyalty // Event containing the contract specifics and raw log

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
func (it *NFTkeySetRoyaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeySetRoyalty)
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
		it.Event = new(NFTkeySetRoyalty)
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
func (it *NFTkeySetRoyaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeySetRoyaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeySetRoyalty represents a SetRoyalty event raised by the NFTkey contract.
type NFTkeySetRoyalty struct {
	Erc721Address common.Address
	Recipient     common.Address
	FeeFraction   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSetRoyalty is a free log retrieval operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_NFTkey *NFTkeyFilterer) FilterSetRoyalty(opts *bind.FilterOpts, erc721Address []common.Address, recipient []common.Address) (*NFTkeySetRoyaltyIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "SetRoyalty", erc721AddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeySetRoyaltyIterator{contract: _NFTkey.contract, event: "SetRoyalty", logs: logs, sub: sub}, nil
}

// WatchSetRoyalty is a free log subscription operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_NFTkey *NFTkeyFilterer) WatchSetRoyalty(opts *bind.WatchOpts, sink chan<- *NFTkeySetRoyalty, erc721Address []common.Address, recipient []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "SetRoyalty", erc721AddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeySetRoyalty)
				if err := _NFTkey.contract.UnpackLog(event, "SetRoyalty", log); err != nil {
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

// ParseSetRoyalty is a log parse operation binding the contract event 0xb6b61d78ac5372b51940cf3b322ea21839c456cade69acdf1b9fb9a6f79d6ff7.
//
// Solidity: event SetRoyalty(address indexed erc721Address, address indexed recipient, uint256 feeFraction)
func (_NFTkey *NFTkeyFilterer) ParseSetRoyalty(log types.Log) (*NFTkeySetRoyalty, error) {
	event := new(NFTkeySetRoyalty)
	if err := _NFTkey.contract.UnpackLog(event, "SetRoyalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenBidAcceptedIterator is returned from FilterTokenBidAccepted and is used to iterate over the raw logs and unpacked data for TokenBidAccepted events raised by the NFTkey contract.
type NFTkeyTokenBidAcceptedIterator struct {
	Event *NFTkeyTokenBidAccepted // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenBidAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenBidAccepted)
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
		it.Event = new(NFTkeyTokenBidAccepted)
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
func (it *NFTkeyTokenBidAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenBidAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenBidAccepted represents a TokenBidAccepted event raised by the NFTkey contract.
type NFTkeyTokenBidAccepted struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Seller        common.Address
	Bid           INFTKEYMarketplaceV2Bid
	ServiceFee    *big.Int
	RoyaltyFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidAccepted is a free log retrieval operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) FilterTokenBidAccepted(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int, seller []common.Address) (*NFTkeyTokenBidAcceptedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenBidAccepted", erc721AddressRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenBidAcceptedIterator{contract: _NFTkey.contract, event: "TokenBidAccepted", logs: logs, sub: sub}, nil
}

// WatchTokenBidAccepted is a free log subscription operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) WatchTokenBidAccepted(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenBidAccepted, erc721Address []common.Address, tokenId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenBidAccepted", erc721AddressRule, tokenIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenBidAccepted)
				if err := _NFTkey.contract.UnpackLog(event, "TokenBidAccepted", log); err != nil {
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

// ParseTokenBidAccepted is a log parse operation binding the contract event 0xcc8a6de82515ca1ae870ff05651038b858e8bcd1b5143c342987b6dc3c343453.
//
// Solidity: event TokenBidAccepted(address indexed erc721Address, uint256 indexed tokenId, address indexed seller, (uint256,uint256,address,uint256) bid, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) ParseTokenBidAccepted(log types.Log) (*NFTkeyTokenBidAccepted, error) {
	event := new(NFTkeyTokenBidAccepted)
	if err := _NFTkey.contract.UnpackLog(event, "TokenBidAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenBidEnteredIterator is returned from FilterTokenBidEntered and is used to iterate over the raw logs and unpacked data for TokenBidEntered events raised by the NFTkey contract.
type NFTkeyTokenBidEnteredIterator struct {
	Event *NFTkeyTokenBidEntered // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenBidEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenBidEntered)
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
		it.Event = new(NFTkeyTokenBidEntered)
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
func (it *NFTkeyTokenBidEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenBidEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenBidEntered represents a TokenBidEntered event raised by the NFTkey contract.
type NFTkeyTokenBidEntered struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Bid           INFTKEYMarketplaceV2Bid
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidEntered is a free log retrieval operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) FilterTokenBidEntered(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NFTkeyTokenBidEnteredIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenBidEntered", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenBidEnteredIterator{contract: _NFTkey.contract, event: "TokenBidEntered", logs: logs, sub: sub}, nil
}

// WatchTokenBidEntered is a free log subscription operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) WatchTokenBidEntered(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenBidEntered, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenBidEntered", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenBidEntered)
				if err := _NFTkey.contract.UnpackLog(event, "TokenBidEntered", log); err != nil {
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

// ParseTokenBidEntered is a log parse operation binding the contract event 0xc547e24584f4dd2da70009d5141bf4344b59a0ce26cd8f7c5d5a28151a11f219.
//
// Solidity: event TokenBidEntered(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) ParseTokenBidEntered(log types.Log) (*NFTkeyTokenBidEntered, error) {
	event := new(NFTkeyTokenBidEntered)
	if err := _NFTkey.contract.UnpackLog(event, "TokenBidEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenBidWithdrawnIterator is returned from FilterTokenBidWithdrawn and is used to iterate over the raw logs and unpacked data for TokenBidWithdrawn events raised by the NFTkey contract.
type NFTkeyTokenBidWithdrawnIterator struct {
	Event *NFTkeyTokenBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenBidWithdrawn)
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
		it.Event = new(NFTkeyTokenBidWithdrawn)
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
func (it *NFTkeyTokenBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenBidWithdrawn represents a TokenBidWithdrawn event raised by the NFTkey contract.
type NFTkeyTokenBidWithdrawn struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Bid           INFTKEYMarketplaceV2Bid
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBidWithdrawn is a free log retrieval operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) FilterTokenBidWithdrawn(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NFTkeyTokenBidWithdrawnIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenBidWithdrawn", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenBidWithdrawnIterator{contract: _NFTkey.contract, event: "TokenBidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenBidWithdrawn is a free log subscription operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) WatchTokenBidWithdrawn(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenBidWithdrawn, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenBidWithdrawn", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenBidWithdrawn)
				if err := _NFTkey.contract.UnpackLog(event, "TokenBidWithdrawn", log); err != nil {
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

// ParseTokenBidWithdrawn is a log parse operation binding the contract event 0xef9d84bc41d1a54361c4bf46f5e11b7c90a3fcb4e604b1b24e0e35d0fa371604.
//
// Solidity: event TokenBidWithdrawn(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) bid)
func (_NFTkey *NFTkeyFilterer) ParseTokenBidWithdrawn(log types.Log) (*NFTkeyTokenBidWithdrawn, error) {
	event := new(NFTkeyTokenBidWithdrawn)
	if err := _NFTkey.contract.UnpackLog(event, "TokenBidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenBoughtIterator is returned from FilterTokenBought and is used to iterate over the raw logs and unpacked data for TokenBought events raised by the NFTkey contract.
type NFTkeyTokenBoughtIterator struct {
	Event *NFTkeyTokenBought // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenBought)
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
		it.Event = new(NFTkeyTokenBought)
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
func (it *NFTkeyTokenBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenBought represents a TokenBought event raised by the NFTkey contract.
type NFTkeyTokenBought struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Buyer         common.Address
	Listing       INFTKEYMarketplaceV2Listing
	ServiceFee    *big.Int
	RoyaltyFee    *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenBought is a free log retrieval operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) FilterTokenBought(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int, buyer []common.Address) (*NFTkeyTokenBoughtIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenBought", erc721AddressRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenBoughtIterator{contract: _NFTkey.contract, event: "TokenBought", logs: logs, sub: sub}, nil
}

// WatchTokenBought is a free log subscription operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) WatchTokenBought(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenBought, erc721Address []common.Address, tokenId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenBought", erc721AddressRule, tokenIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenBought)
				if err := _NFTkey.contract.UnpackLog(event, "TokenBought", log); err != nil {
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

// ParseTokenBought is a log parse operation binding the contract event 0x50a3cf2b1df7bd2994e752563ce6f85769fb50da66bbb9a9912d2d6066a6b4da.
//
// Solidity: event TokenBought(address indexed erc721Address, uint256 indexed tokenId, address indexed buyer, (uint256,uint256,address,uint256) listing, uint256 serviceFee, uint256 royaltyFee)
func (_NFTkey *NFTkeyFilterer) ParseTokenBought(log types.Log) (*NFTkeyTokenBought, error) {
	event := new(NFTkeyTokenBought)
	if err := _NFTkey.contract.UnpackLog(event, "TokenBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenDelistedIterator is returned from FilterTokenDelisted and is used to iterate over the raw logs and unpacked data for TokenDelisted events raised by the NFTkey contract.
type NFTkeyTokenDelistedIterator struct {
	Event *NFTkeyTokenDelisted // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenDelisted)
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
		it.Event = new(NFTkeyTokenDelisted)
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
func (it *NFTkeyTokenDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenDelisted represents a TokenDelisted event raised by the NFTkey contract.
type NFTkeyTokenDelisted struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Listing       INFTKEYMarketplaceV2Listing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenDelisted is a free log retrieval operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) FilterTokenDelisted(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NFTkeyTokenDelistedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenDelisted", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenDelistedIterator{contract: _NFTkey.contract, event: "TokenDelisted", logs: logs, sub: sub}, nil
}

// WatchTokenDelisted is a free log subscription operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) WatchTokenDelisted(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenDelisted, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenDelisted", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenDelisted)
				if err := _NFTkey.contract.UnpackLog(event, "TokenDelisted", log); err != nil {
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

// ParseTokenDelisted is a log parse operation binding the contract event 0x835a0a426c94e53ab00dd6cf617ba68d1fa6c9162ef7a036b80be930c9a32c53.
//
// Solidity: event TokenDelisted(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) ParseTokenDelisted(log types.Log) (*NFTkeyTokenDelisted, error) {
	event := new(NFTkeyTokenDelisted)
	if err := _NFTkey.contract.UnpackLog(event, "TokenDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NFTkeyTokenListedIterator is returned from FilterTokenListed and is used to iterate over the raw logs and unpacked data for TokenListed events raised by the NFTkey contract.
type NFTkeyTokenListedIterator struct {
	Event *NFTkeyTokenListed // Event containing the contract specifics and raw log

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
func (it *NFTkeyTokenListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NFTkeyTokenListed)
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
		it.Event = new(NFTkeyTokenListed)
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
func (it *NFTkeyTokenListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NFTkeyTokenListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NFTkeyTokenListed represents a TokenListed event raised by the NFTkey contract.
type NFTkeyTokenListed struct {
	Erc721Address common.Address
	TokenId       *big.Int
	Listing       INFTKEYMarketplaceV2Listing
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTokenListed is a free log retrieval operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) FilterTokenListed(opts *bind.FilterOpts, erc721Address []common.Address, tokenId []*big.Int) (*NFTkeyTokenListedIterator, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.FilterLogs(opts, "TokenListed", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NFTkeyTokenListedIterator{contract: _NFTkey.contract, event: "TokenListed", logs: logs, sub: sub}, nil
}

// WatchTokenListed is a free log subscription operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) WatchTokenListed(opts *bind.WatchOpts, sink chan<- *NFTkeyTokenListed, erc721Address []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var erc721AddressRule []interface{}
	for _, erc721AddressItem := range erc721Address {
		erc721AddressRule = append(erc721AddressRule, erc721AddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NFTkey.contract.WatchLogs(opts, "TokenListed", erc721AddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NFTkeyTokenListed)
				if err := _NFTkey.contract.UnpackLog(event, "TokenListed", log); err != nil {
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

// ParseTokenListed is a log parse operation binding the contract event 0xfc8bd63d1c4480fdf7501f95b3dd0ba53542a02abbab92c0f579468341741abd.
//
// Solidity: event TokenListed(address indexed erc721Address, uint256 indexed tokenId, (uint256,uint256,address,uint256) listing)
func (_NFTkey *NFTkeyFilterer) ParseTokenListed(log types.Log) (*NFTkeyTokenListed, error) {
	event := new(NFTkeyTokenListed)
	if err := _NFTkey.contract.UnpackLog(event, "TokenListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
