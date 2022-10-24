// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quixotic

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

// QuixoticMetaData contains all meta data concerning the Quixotic contract.
var QuixoticMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"DutchAuctionFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SellOrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPrice\",\"type\":\"uint256\"}],\"name\":\"calculateCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelBuyOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelPreviousSellOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"fillBuyOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"fillDutchAuctionOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAtBlockNumber\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"fillSellOrder\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getRoyaltyPayoutAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"getRoyaltyPayoutRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isOrderCancelled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newMakerWallet\",\"type\":\"address\"}],\"name\":\"setMakerWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"royaltyRegistryAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cancellationRegistryAddr\",\"type\":\"address\"}],\"name\":\"setRegistryContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_payoutAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_payoutPerMille\",\"type\":\"uint256\"}],\"name\":\"setRoyalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wethAddress\",\"type\":\"address\"}],\"name\":\"setWethContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// QuixoticABI is the input ABI used to generate the binding from.
// Deprecated: Use QuixoticMetaData.ABI instead.
var QuixoticABI = QuixoticMetaData.ABI

// Quixotic is an auto generated Go binding around an Ethereum contract.
type Quixotic struct {
	QuixoticCaller     // Read-only binding to the contract
	QuixoticTransactor // Write-only binding to the contract
	QuixoticFilterer   // Log filterer for contract events
}

// QuixoticCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuixoticCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuixoticTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuixoticTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuixoticFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuixoticFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuixoticSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuixoticSession struct {
	Contract     *Quixotic         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuixoticCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuixoticCallerSession struct {
	Contract *QuixoticCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QuixoticTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuixoticTransactorSession struct {
	Contract     *QuixoticTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QuixoticRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuixoticRaw struct {
	Contract *Quixotic // Generic contract binding to access the raw methods on
}

// QuixoticCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuixoticCallerRaw struct {
	Contract *QuixoticCaller // Generic read-only contract binding to access the raw methods on
}

// QuixoticTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuixoticTransactorRaw struct {
	Contract *QuixoticTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuixotic creates a new instance of Quixotic, bound to a specific deployed contract.
func NewQuixotic(address common.Address, backend bind.ContractBackend) (*Quixotic, error) {
	contract, err := bindQuixotic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quixotic{QuixoticCaller: QuixoticCaller{contract: contract}, QuixoticTransactor: QuixoticTransactor{contract: contract}, QuixoticFilterer: QuixoticFilterer{contract: contract}}, nil
}

// NewQuixoticCaller creates a new read-only instance of Quixotic, bound to a specific deployed contract.
func NewQuixoticCaller(address common.Address, caller bind.ContractCaller) (*QuixoticCaller, error) {
	contract, err := bindQuixotic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuixoticCaller{contract: contract}, nil
}

// NewQuixoticTransactor creates a new write-only instance of Quixotic, bound to a specific deployed contract.
func NewQuixoticTransactor(address common.Address, transactor bind.ContractTransactor) (*QuixoticTransactor, error) {
	contract, err := bindQuixotic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuixoticTransactor{contract: contract}, nil
}

// NewQuixoticFilterer creates a new log filterer instance of Quixotic, bound to a specific deployed contract.
func NewQuixoticFilterer(address common.Address, filterer bind.ContractFilterer) (*QuixoticFilterer, error) {
	contract, err := bindQuixotic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuixoticFilterer{contract: contract}, nil
}

// bindQuixotic binds a generic wrapper to an already deployed contract.
func bindQuixotic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuixoticABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quixotic *QuixoticRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quixotic.Contract.QuixoticCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quixotic *QuixoticRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.Contract.QuixoticTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quixotic *QuixoticRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quixotic.Contract.QuixoticTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quixotic *QuixoticCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Quixotic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quixotic *QuixoticTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quixotic *QuixoticTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quixotic.Contract.contract.Transact(opts, method, params...)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_Quixotic *QuixoticCaller) CalculateCurrentPrice(opts *bind.CallOpts, startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "calculateCurrentPrice", startTime, endTime, startPrice, endPrice)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_Quixotic *QuixoticSession) CalculateCurrentPrice(startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	return _Quixotic.Contract.CalculateCurrentPrice(&_Quixotic.CallOpts, startTime, endTime, startPrice, endPrice)
}

// CalculateCurrentPrice is a free data retrieval call binding the contract method 0xc3fbdb7c.
//
// Solidity: function calculateCurrentPrice(uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice) view returns(uint256)
func (_Quixotic *QuixoticCallerSession) CalculateCurrentPrice(startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int) (*big.Int, error) {
	return _Quixotic.Contract.CalculateCurrentPrice(&_Quixotic.CallOpts, startTime, endTime, startPrice, endPrice)
}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_Quixotic *QuixoticCaller) GetRoyaltyPayoutAddress(opts *bind.CallOpts, contractAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "getRoyaltyPayoutAddress", contractAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_Quixotic *QuixoticSession) GetRoyaltyPayoutAddress(contractAddress common.Address) (common.Address, error) {
	return _Quixotic.Contract.GetRoyaltyPayoutAddress(&_Quixotic.CallOpts, contractAddress)
}

// GetRoyaltyPayoutAddress is a free data retrieval call binding the contract method 0x2df5cb23.
//
// Solidity: function getRoyaltyPayoutAddress(address contractAddress) view returns(address)
func (_Quixotic *QuixoticCallerSession) GetRoyaltyPayoutAddress(contractAddress common.Address) (common.Address, error) {
	return _Quixotic.Contract.GetRoyaltyPayoutAddress(&_Quixotic.CallOpts, contractAddress)
}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_Quixotic *QuixoticCaller) GetRoyaltyPayoutRate(opts *bind.CallOpts, contractAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "getRoyaltyPayoutRate", contractAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_Quixotic *QuixoticSession) GetRoyaltyPayoutRate(contractAddress common.Address) (*big.Int, error) {
	return _Quixotic.Contract.GetRoyaltyPayoutRate(&_Quixotic.CallOpts, contractAddress)
}

// GetRoyaltyPayoutRate is a free data retrieval call binding the contract method 0xec4210ff.
//
// Solidity: function getRoyaltyPayoutRate(address contractAddress) view returns(uint256)
func (_Quixotic *QuixoticCallerSession) GetRoyaltyPayoutRate(contractAddress common.Address) (*big.Int, error) {
	return _Quixotic.Contract.GetRoyaltyPayoutRate(&_Quixotic.CallOpts, contractAddress)
}

// IsOrderCancelled is a free data retrieval call binding the contract method 0x563166a9.
//
// Solidity: function isOrderCancelled(bytes signature) view returns(bool)
func (_Quixotic *QuixoticCaller) IsOrderCancelled(opts *bind.CallOpts, signature []byte) (bool, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "isOrderCancelled", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderCancelled is a free data retrieval call binding the contract method 0x563166a9.
//
// Solidity: function isOrderCancelled(bytes signature) view returns(bool)
func (_Quixotic *QuixoticSession) IsOrderCancelled(signature []byte) (bool, error) {
	return _Quixotic.Contract.IsOrderCancelled(&_Quixotic.CallOpts, signature)
}

// IsOrderCancelled is a free data retrieval call binding the contract method 0x563166a9.
//
// Solidity: function isOrderCancelled(bytes signature) view returns(bool)
func (_Quixotic *QuixoticCallerSession) IsOrderCancelled(signature []byte) (bool, error) {
	return _Quixotic.Contract.IsOrderCancelled(&_Quixotic.CallOpts, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quixotic *QuixoticCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quixotic *QuixoticSession) Owner() (common.Address, error) {
	return _Quixotic.Contract.Owner(&_Quixotic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Quixotic *QuixoticCallerSession) Owner() (common.Address, error) {
	return _Quixotic.Contract.Owner(&_Quixotic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quixotic *QuixoticCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Quixotic.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quixotic *QuixoticSession) Paused() (bool, error) {
	return _Quixotic.Contract.Paused(&_Quixotic.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Quixotic *QuixoticCallerSession) Paused() (bool, error) {
	return _Quixotic.Contract.Paused(&_Quixotic.CallOpts)
}

// CancelBuyOrder is a paid mutator transaction binding the contract method 0xdc7e053f.
//
// Solidity: function cancelBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature) returns()
func (_Quixotic *QuixoticTransactor) CancelBuyOrder(opts *bind.TransactOpts, buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "cancelBuyOrder", buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature)
}

// CancelBuyOrder is a paid mutator transaction binding the contract method 0xdc7e053f.
//
// Solidity: function cancelBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature) returns()
func (_Quixotic *QuixoticSession) CancelBuyOrder(buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Quixotic.Contract.CancelBuyOrder(&_Quixotic.TransactOpts, buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature)
}

// CancelBuyOrder is a paid mutator transaction binding the contract method 0xdc7e053f.
//
// Solidity: function cancelBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature) returns()
func (_Quixotic *QuixoticTransactorSession) CancelBuyOrder(buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte) (*types.Transaction, error) {
	return _Quixotic.Contract.CancelBuyOrder(&_Quixotic.TransactOpts, buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_Quixotic *QuixoticTransactor) CancelPreviousSellOrders(opts *bind.TransactOpts, addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "cancelPreviousSellOrders", addr, tokenAddr, tokenId)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_Quixotic *QuixoticSession) CancelPreviousSellOrders(addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Quixotic.Contract.CancelPreviousSellOrders(&_Quixotic.TransactOpts, addr, tokenAddr, tokenId)
}

// CancelPreviousSellOrders is a paid mutator transaction binding the contract method 0x083d089d.
//
// Solidity: function cancelPreviousSellOrders(address addr, address tokenAddr, uint256 tokenId) returns()
func (_Quixotic *QuixoticTransactorSession) CancelPreviousSellOrders(addr common.Address, tokenAddr common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Quixotic.Contract.CancelPreviousSellOrders(&_Quixotic.TransactOpts, addr, tokenAddr, tokenId)
}

// FillBuyOrder is a paid mutator transaction binding the contract method 0x55e4a3fa.
//
// Solidity: function fillBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature, address seller) payable returns()
func (_Quixotic *QuixoticTransactor) FillBuyOrder(opts *bind.TransactOpts, buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte, seller common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "fillBuyOrder", buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature, seller)
}

// FillBuyOrder is a paid mutator transaction binding the contract method 0x55e4a3fa.
//
// Solidity: function fillBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature, address seller) payable returns()
func (_Quixotic *QuixoticSession) FillBuyOrder(buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte, seller common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillBuyOrder(&_Quixotic.TransactOpts, buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature, seller)
}

// FillBuyOrder is a paid mutator transaction binding the contract method 0x55e4a3fa.
//
// Solidity: function fillBuyOrder(address buyer, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, bytes signature, address seller) payable returns()
func (_Quixotic *QuixoticTransactorSession) FillBuyOrder(buyer common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, signature []byte, seller common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillBuyOrder(&_Quixotic.TransactOpts, buyer, contractAddress, tokenId, startTime, expiration, price, quantity, signature, seller)
}

// FillDutchAuctionOrder is a paid mutator transaction binding the contract method 0xedf8301b.
//
// Solidity: function fillDutchAuctionOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticTransactor) FillDutchAuctionOrder(opts *bind.TransactOpts, seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "fillDutchAuctionOrder", seller, contractAddress, tokenId, startTime, endTime, startPrice, endPrice, quantity, createdAtBlockNumber, signature, buyer)
}

// FillDutchAuctionOrder is a paid mutator transaction binding the contract method 0xedf8301b.
//
// Solidity: function fillDutchAuctionOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticSession) FillDutchAuctionOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillDutchAuctionOrder(&_Quixotic.TransactOpts, seller, contractAddress, tokenId, startTime, endTime, startPrice, endPrice, quantity, createdAtBlockNumber, signature, buyer)
}

// FillDutchAuctionOrder is a paid mutator transaction binding the contract method 0xedf8301b.
//
// Solidity: function fillDutchAuctionOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 endTime, uint256 startPrice, uint256 endPrice, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticTransactorSession) FillDutchAuctionOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, endTime *big.Int, startPrice *big.Int, endPrice *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillDutchAuctionOrder(&_Quixotic.TransactOpts, seller, contractAddress, tokenId, startTime, endTime, startPrice, endPrice, quantity, createdAtBlockNumber, signature, buyer)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0xad6c8c5f.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticTransactor) FillSellOrder(opts *bind.TransactOpts, seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "fillSellOrder", seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, signature, buyer)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0xad6c8c5f.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticSession) FillSellOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillSellOrder(&_Quixotic.TransactOpts, seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, signature, buyer)
}

// FillSellOrder is a paid mutator transaction binding the contract method 0xad6c8c5f.
//
// Solidity: function fillSellOrder(address seller, address contractAddress, uint256 tokenId, uint256 startTime, uint256 expiration, uint256 price, uint256 quantity, uint256 createdAtBlockNumber, bytes signature, address buyer) payable returns()
func (_Quixotic *QuixoticTransactorSession) FillSellOrder(seller common.Address, contractAddress common.Address, tokenId *big.Int, startTime *big.Int, expiration *big.Int, price *big.Int, quantity *big.Int, createdAtBlockNumber *big.Int, signature []byte, buyer common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.FillSellOrder(&_Quixotic.TransactOpts, seller, contractAddress, tokenId, startTime, expiration, price, quantity, createdAtBlockNumber, signature, buyer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quixotic *QuixoticTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quixotic *QuixoticSession) Pause() (*types.Transaction, error) {
	return _Quixotic.Contract.Pause(&_Quixotic.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Quixotic *QuixoticTransactorSession) Pause() (*types.Transaction, error) {
	return _Quixotic.Contract.Pause(&_Quixotic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quixotic *QuixoticTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quixotic *QuixoticSession) RenounceOwnership() (*types.Transaction, error) {
	return _Quixotic.Contract.RenounceOwnership(&_Quixotic.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Quixotic *QuixoticTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Quixotic.Contract.RenounceOwnership(&_Quixotic.TransactOpts)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_Quixotic *QuixoticTransactor) SetMakerWallet(opts *bind.TransactOpts, _newMakerWallet common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "setMakerWallet", _newMakerWallet)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_Quixotic *QuixoticSession) SetMakerWallet(_newMakerWallet common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetMakerWallet(&_Quixotic.TransactOpts, _newMakerWallet)
}

// SetMakerWallet is a paid mutator transaction binding the contract method 0x2b812a34.
//
// Solidity: function setMakerWallet(address _newMakerWallet) returns()
func (_Quixotic *QuixoticTransactorSession) SetMakerWallet(_newMakerWallet common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetMakerWallet(&_Quixotic.TransactOpts, _newMakerWallet)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address royaltyRegistryAddr, address cancellationRegistryAddr) returns()
func (_Quixotic *QuixoticTransactor) SetRegistryContracts(opts *bind.TransactOpts, royaltyRegistryAddr common.Address, cancellationRegistryAddr common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "setRegistryContracts", royaltyRegistryAddr, cancellationRegistryAddr)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address royaltyRegistryAddr, address cancellationRegistryAddr) returns()
func (_Quixotic *QuixoticSession) SetRegistryContracts(royaltyRegistryAddr common.Address, cancellationRegistryAddr common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetRegistryContracts(&_Quixotic.TransactOpts, royaltyRegistryAddr, cancellationRegistryAddr)
}

// SetRegistryContracts is a paid mutator transaction binding the contract method 0xb1216b79.
//
// Solidity: function setRegistryContracts(address royaltyRegistryAddr, address cancellationRegistryAddr) returns()
func (_Quixotic *QuixoticTransactorSession) SetRegistryContracts(royaltyRegistryAddr common.Address, cancellationRegistryAddr common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetRegistryContracts(&_Quixotic.TransactOpts, royaltyRegistryAddr, cancellationRegistryAddr)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_Quixotic *QuixoticTransactor) SetRoyalty(opts *bind.TransactOpts, contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "setRoyalty", contractAddress, _payoutAddress, _payoutPerMille)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_Quixotic *QuixoticSession) SetRoyalty(contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _Quixotic.Contract.SetRoyalty(&_Quixotic.TransactOpts, contractAddress, _payoutAddress, _payoutPerMille)
}

// SetRoyalty is a paid mutator transaction binding the contract method 0x55c338aa.
//
// Solidity: function setRoyalty(address contractAddress, address _payoutAddress, uint256 _payoutPerMille) returns()
func (_Quixotic *QuixoticTransactorSession) SetRoyalty(contractAddress common.Address, _payoutAddress common.Address, _payoutPerMille *big.Int) (*types.Transaction, error) {
	return _Quixotic.Contract.SetRoyalty(&_Quixotic.TransactOpts, contractAddress, _payoutAddress, _payoutPerMille)
}

// SetWethContract is a paid mutator transaction binding the contract method 0xabe0906a.
//
// Solidity: function setWethContract(address wethAddress) returns()
func (_Quixotic *QuixoticTransactor) SetWethContract(opts *bind.TransactOpts, wethAddress common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "setWethContract", wethAddress)
}

// SetWethContract is a paid mutator transaction binding the contract method 0xabe0906a.
//
// Solidity: function setWethContract(address wethAddress) returns()
func (_Quixotic *QuixoticSession) SetWethContract(wethAddress common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetWethContract(&_Quixotic.TransactOpts, wethAddress)
}

// SetWethContract is a paid mutator transaction binding the contract method 0xabe0906a.
//
// Solidity: function setWethContract(address wethAddress) returns()
func (_Quixotic *QuixoticTransactorSession) SetWethContract(wethAddress common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.SetWethContract(&_Quixotic.TransactOpts, wethAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quixotic *QuixoticTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quixotic *QuixoticSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.TransferOwnership(&_Quixotic.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Quixotic *QuixoticTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Quixotic.Contract.TransferOwnership(&_Quixotic.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quixotic *QuixoticTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quixotic *QuixoticSession) Unpause() (*types.Transaction, error) {
	return _Quixotic.Contract.Unpause(&_Quixotic.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Quixotic *QuixoticTransactorSession) Unpause() (*types.Transaction, error) {
	return _Quixotic.Contract.Unpause(&_Quixotic.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Quixotic *QuixoticTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quixotic.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Quixotic *QuixoticSession) Withdraw() (*types.Transaction, error) {
	return _Quixotic.Contract.Withdraw(&_Quixotic.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Quixotic *QuixoticTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Quixotic.Contract.Withdraw(&_Quixotic.TransactOpts)
}

// QuixoticBuyOrderFilledIterator is returned from FilterBuyOrderFilled and is used to iterate over the raw logs and unpacked data for BuyOrderFilled events raised by the Quixotic contract.
type QuixoticBuyOrderFilledIterator struct {
	Event *QuixoticBuyOrderFilled // Event containing the contract specifics and raw log

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
func (it *QuixoticBuyOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticBuyOrderFilled)
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
		it.Event = new(QuixoticBuyOrderFilled)
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
func (it *QuixoticBuyOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticBuyOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticBuyOrderFilled represents a BuyOrderFilled event raised by the Quixotic contract.
type QuixoticBuyOrderFilled struct {
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBuyOrderFilled is a free log retrieval operation binding the contract event 0xcb3cd529428badade171c8b0ef6c2f25d13bc69785143c43ec869d386ae34141.
//
// Solidity: event BuyOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) FilterBuyOrderFilled(opts *bind.FilterOpts, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (*QuixoticBuyOrderFilledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "BuyOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuixoticBuyOrderFilledIterator{contract: _Quixotic.contract, event: "BuyOrderFilled", logs: logs, sub: sub}, nil
}

// WatchBuyOrderFilled is a free log subscription operation binding the contract event 0xcb3cd529428badade171c8b0ef6c2f25d13bc69785143c43ec869d386ae34141.
//
// Solidity: event BuyOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) WatchBuyOrderFilled(opts *bind.WatchOpts, sink chan<- *QuixoticBuyOrderFilled, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "BuyOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticBuyOrderFilled)
				if err := _Quixotic.contract.UnpackLog(event, "BuyOrderFilled", log); err != nil {
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

// ParseBuyOrderFilled is a log parse operation binding the contract event 0xcb3cd529428badade171c8b0ef6c2f25d13bc69785143c43ec869d386ae34141.
//
// Solidity: event BuyOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) ParseBuyOrderFilled(log types.Log) (*QuixoticBuyOrderFilled, error) {
	event := new(QuixoticBuyOrderFilled)
	if err := _Quixotic.contract.UnpackLog(event, "BuyOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuixoticDutchAuctionFilledIterator is returned from FilterDutchAuctionFilled and is used to iterate over the raw logs and unpacked data for DutchAuctionFilled events raised by the Quixotic contract.
type QuixoticDutchAuctionFilledIterator struct {
	Event *QuixoticDutchAuctionFilled // Event containing the contract specifics and raw log

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
func (it *QuixoticDutchAuctionFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticDutchAuctionFilled)
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
		it.Event = new(QuixoticDutchAuctionFilled)
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
func (it *QuixoticDutchAuctionFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticDutchAuctionFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticDutchAuctionFilled represents a DutchAuctionFilled event raised by the Quixotic contract.
type QuixoticDutchAuctionFilled struct {
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDutchAuctionFilled is a free log retrieval operation binding the contract event 0xcd2ac775eba6d68eb3bdbba5c500751ed102e5d2b7bf9c459b3e7b808d132cbd.
//
// Solidity: event DutchAuctionFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) FilterDutchAuctionFilled(opts *bind.FilterOpts, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (*QuixoticDutchAuctionFilledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "DutchAuctionFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuixoticDutchAuctionFilledIterator{contract: _Quixotic.contract, event: "DutchAuctionFilled", logs: logs, sub: sub}, nil
}

// WatchDutchAuctionFilled is a free log subscription operation binding the contract event 0xcd2ac775eba6d68eb3bdbba5c500751ed102e5d2b7bf9c459b3e7b808d132cbd.
//
// Solidity: event DutchAuctionFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) WatchDutchAuctionFilled(opts *bind.WatchOpts, sink chan<- *QuixoticDutchAuctionFilled, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "DutchAuctionFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticDutchAuctionFilled)
				if err := _Quixotic.contract.UnpackLog(event, "DutchAuctionFilled", log); err != nil {
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

// ParseDutchAuctionFilled is a log parse operation binding the contract event 0xcd2ac775eba6d68eb3bdbba5c500751ed102e5d2b7bf9c459b3e7b808d132cbd.
//
// Solidity: event DutchAuctionFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) ParseDutchAuctionFilled(log types.Log) (*QuixoticDutchAuctionFilled, error) {
	event := new(QuixoticDutchAuctionFilled)
	if err := _Quixotic.contract.UnpackLog(event, "DutchAuctionFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuixoticOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Quixotic contract.
type QuixoticOwnershipTransferredIterator struct {
	Event *QuixoticOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *QuixoticOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticOwnershipTransferred)
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
		it.Event = new(QuixoticOwnershipTransferred)
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
func (it *QuixoticOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticOwnershipTransferred represents a OwnershipTransferred event raised by the Quixotic contract.
type QuixoticOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Quixotic *QuixoticFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*QuixoticOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &QuixoticOwnershipTransferredIterator{contract: _Quixotic.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Quixotic *QuixoticFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *QuixoticOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticOwnershipTransferred)
				if err := _Quixotic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Quixotic *QuixoticFilterer) ParseOwnershipTransferred(log types.Log) (*QuixoticOwnershipTransferred, error) {
	event := new(QuixoticOwnershipTransferred)
	if err := _Quixotic.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuixoticPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Quixotic contract.
type QuixoticPausedIterator struct {
	Event *QuixoticPaused // Event containing the contract specifics and raw log

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
func (it *QuixoticPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticPaused)
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
		it.Event = new(QuixoticPaused)
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
func (it *QuixoticPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticPaused represents a Paused event raised by the Quixotic contract.
type QuixoticPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quixotic *QuixoticFilterer) FilterPaused(opts *bind.FilterOpts) (*QuixoticPausedIterator, error) {

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &QuixoticPausedIterator{contract: _Quixotic.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quixotic *QuixoticFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *QuixoticPaused) (event.Subscription, error) {

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticPaused)
				if err := _Quixotic.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Quixotic *QuixoticFilterer) ParsePaused(log types.Log) (*QuixoticPaused, error) {
	event := new(QuixoticPaused)
	if err := _Quixotic.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuixoticSellOrderFilledIterator is returned from FilterSellOrderFilled and is used to iterate over the raw logs and unpacked data for SellOrderFilled events raised by the Quixotic contract.
type QuixoticSellOrderFilledIterator struct {
	Event *QuixoticSellOrderFilled // Event containing the contract specifics and raw log

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
func (it *QuixoticSellOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticSellOrderFilled)
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
		it.Event = new(QuixoticSellOrderFilled)
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
func (it *QuixoticSellOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticSellOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticSellOrderFilled represents a SellOrderFilled event raised by the Quixotic contract.
type QuixoticSellOrderFilled struct {
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSellOrderFilled is a free log retrieval operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) FilterSellOrderFilled(opts *bind.FilterOpts, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (*QuixoticSellOrderFilledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "SellOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &QuixoticSellOrderFilledIterator{contract: _Quixotic.contract, event: "SellOrderFilled", logs: logs, sub: sub}, nil
}

// WatchSellOrderFilled is a free log subscription operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) WatchSellOrderFilled(opts *bind.WatchOpts, sink chan<- *QuixoticSellOrderFilled, seller []common.Address, contractAddress []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var contractAddressRule []interface{}
	for _, contractAddressItem := range contractAddress {
		contractAddressRule = append(contractAddressRule, contractAddressItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "SellOrderFilled", sellerRule, contractAddressRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticSellOrderFilled)
				if err := _Quixotic.contract.UnpackLog(event, "SellOrderFilled", log); err != nil {
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

// ParseSellOrderFilled is a log parse operation binding the contract event 0x70ba0d31158674eea8365d0f7b9ac70e552cc28b8bb848664e4feb939c6578f8.
//
// Solidity: event SellOrderFilled(address indexed seller, address buyer, address indexed contractAddress, uint256 indexed tokenId, uint256 price)
func (_Quixotic *QuixoticFilterer) ParseSellOrderFilled(log types.Log) (*QuixoticSellOrderFilled, error) {
	event := new(QuixoticSellOrderFilled)
	if err := _Quixotic.contract.UnpackLog(event, "SellOrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// QuixoticUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Quixotic contract.
type QuixoticUnpausedIterator struct {
	Event *QuixoticUnpaused // Event containing the contract specifics and raw log

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
func (it *QuixoticUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(QuixoticUnpaused)
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
		it.Event = new(QuixoticUnpaused)
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
func (it *QuixoticUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *QuixoticUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// QuixoticUnpaused represents a Unpaused event raised by the Quixotic contract.
type QuixoticUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quixotic *QuixoticFilterer) FilterUnpaused(opts *bind.FilterOpts) (*QuixoticUnpausedIterator, error) {

	logs, sub, err := _Quixotic.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &QuixoticUnpausedIterator{contract: _Quixotic.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quixotic *QuixoticFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *QuixoticUnpaused) (event.Subscription, error) {

	logs, sub, err := _Quixotic.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(QuixoticUnpaused)
				if err := _Quixotic.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Quixotic *QuixoticFilterer) ParseUnpaused(log types.Log) (*QuixoticUnpaused, error) {
	event := new(QuixoticUnpaused)
	if err := _Quixotic.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
