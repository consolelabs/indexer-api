// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package x2y2

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

// ERC721DelegatePair is an auto generated low-level Go binding around an user-defined struct.
type ERC721DelegatePair struct {
	Token   common.Address
	TokenId *big.Int
}

// X2Y2MetaData contains all meta data concerning the X2Y2 contract.
var X2Y2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELEGATION_CALLER\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delegateType\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeAuctionComplete\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeAuctionRefund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"previousBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeBid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeBuy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeSell\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"contractIERC721\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structERC721Delegate.Pair[]\",\"name\":\"pairs\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// X2Y2ABI is the input ABI used to generate the binding from.
// Deprecated: Use X2Y2MetaData.ABI instead.
var X2Y2ABI = X2Y2MetaData.ABI

// X2Y2 is an auto generated Go binding around an Ethereum contract.
type X2Y2 struct {
	X2Y2Caller     // Read-only binding to the contract
	X2Y2Transactor // Write-only binding to the contract
	X2Y2Filterer   // Log filterer for contract events
}

// X2Y2Caller is an auto generated read-only Go binding around an Ethereum contract.
type X2Y2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type X2Y2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type X2Y2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type X2Y2Session struct {
	Contract     *X2Y2             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// X2Y2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type X2Y2CallerSession struct {
	Contract *X2Y2Caller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// X2Y2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type X2Y2TransactorSession struct {
	Contract     *X2Y2Transactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// X2Y2Raw is an auto generated low-level Go binding around an Ethereum contract.
type X2Y2Raw struct {
	Contract *X2Y2 // Generic contract binding to access the raw methods on
}

// X2Y2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type X2Y2CallerRaw struct {
	Contract *X2Y2Caller // Generic read-only contract binding to access the raw methods on
}

// X2Y2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type X2Y2TransactorRaw struct {
	Contract *X2Y2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewX2Y2 creates a new instance of X2Y2, bound to a specific deployed contract.
func NewX2Y2(address common.Address, backend bind.ContractBackend) (*X2Y2, error) {
	contract, err := bindX2Y2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &X2Y2{X2Y2Caller: X2Y2Caller{contract: contract}, X2Y2Transactor: X2Y2Transactor{contract: contract}, X2Y2Filterer: X2Y2Filterer{contract: contract}}, nil
}

// NewX2Y2Caller creates a new read-only instance of X2Y2, bound to a specific deployed contract.
func NewX2Y2Caller(address common.Address, caller bind.ContractCaller) (*X2Y2Caller, error) {
	contract, err := bindX2Y2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &X2Y2Caller{contract: contract}, nil
}

// NewX2Y2Transactor creates a new write-only instance of X2Y2, bound to a specific deployed contract.
func NewX2Y2Transactor(address common.Address, transactor bind.ContractTransactor) (*X2Y2Transactor, error) {
	contract, err := bindX2Y2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &X2Y2Transactor{contract: contract}, nil
}

// NewX2Y2Filterer creates a new log filterer instance of X2Y2, bound to a specific deployed contract.
func NewX2Y2Filterer(address common.Address, filterer bind.ContractFilterer) (*X2Y2Filterer, error) {
	contract, err := bindX2Y2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &X2Y2Filterer{contract: contract}, nil
}

// bindX2Y2 binds a generic wrapper to an already deployed contract.
func bindX2Y2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(X2Y2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X2Y2 *X2Y2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X2Y2.Contract.X2Y2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X2Y2 *X2Y2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2.Contract.X2Y2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X2Y2 *X2Y2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X2Y2.Contract.X2Y2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X2Y2 *X2Y2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X2Y2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X2Y2 *X2Y2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X2Y2 *X2Y2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X2Y2.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_X2Y2 *X2Y2Caller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_X2Y2 *X2Y2Session) DEFAULTADMINROLE() ([32]byte, error) {
	return _X2Y2.Contract.DEFAULTADMINROLE(&_X2Y2.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_X2Y2 *X2Y2CallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _X2Y2.Contract.DEFAULTADMINROLE(&_X2Y2.CallOpts)
}

// DELEGATIONCALLER is a free data retrieval call binding the contract method 0x8e325979.
//
// Solidity: function DELEGATION_CALLER() view returns(bytes32)
func (_X2Y2 *X2Y2Caller) DELEGATIONCALLER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "DELEGATION_CALLER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONCALLER is a free data retrieval call binding the contract method 0x8e325979.
//
// Solidity: function DELEGATION_CALLER() view returns(bytes32)
func (_X2Y2 *X2Y2Session) DELEGATIONCALLER() ([32]byte, error) {
	return _X2Y2.Contract.DELEGATIONCALLER(&_X2Y2.CallOpts)
}

// DELEGATIONCALLER is a free data retrieval call binding the contract method 0x8e325979.
//
// Solidity: function DELEGATION_CALLER() view returns(bytes32)
func (_X2Y2 *X2Y2CallerSession) DELEGATIONCALLER() ([32]byte, error) {
	return _X2Y2.Contract.DELEGATIONCALLER(&_X2Y2.CallOpts)
}

// DelegateType is a free data retrieval call binding the contract method 0x2c436e5b.
//
// Solidity: function delegateType() view returns(uint256)
func (_X2Y2 *X2Y2Caller) DelegateType(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "delegateType")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelegateType is a free data retrieval call binding the contract method 0x2c436e5b.
//
// Solidity: function delegateType() view returns(uint256)
func (_X2Y2 *X2Y2Session) DelegateType() (*big.Int, error) {
	return _X2Y2.Contract.DelegateType(&_X2Y2.CallOpts)
}

// DelegateType is a free data retrieval call binding the contract method 0x2c436e5b.
//
// Solidity: function delegateType() view returns(uint256)
func (_X2Y2 *X2Y2CallerSession) DelegateType() (*big.Int, error) {
	return _X2Y2.Contract.DelegateType(&_X2Y2.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_X2Y2 *X2Y2Caller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_X2Y2 *X2Y2Session) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _X2Y2.Contract.GetRoleAdmin(&_X2Y2.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_X2Y2 *X2Y2CallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _X2Y2.Contract.GetRoleAdmin(&_X2Y2.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_X2Y2 *X2Y2Caller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_X2Y2 *X2Y2Session) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _X2Y2.Contract.HasRole(&_X2Y2.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_X2Y2 *X2Y2CallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _X2Y2.Contract.HasRole(&_X2Y2.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_X2Y2 *X2Y2Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _X2Y2.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_X2Y2 *X2Y2Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _X2Y2.Contract.SupportsInterface(&_X2Y2.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_X2Y2 *X2Y2CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _X2Y2.Contract.SupportsInterface(&_X2Y2.CallOpts, interfaceId)
}

// ExecuteAuctionComplete is a paid mutator transaction binding the contract method 0x3672c911.
//
// Solidity: function executeAuctionComplete(address , address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Transactor) ExecuteAuctionComplete(opts *bind.TransactOpts, arg0 common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "executeAuctionComplete", arg0, buyer, data)
}

// ExecuteAuctionComplete is a paid mutator transaction binding the contract method 0x3672c911.
//
// Solidity: function executeAuctionComplete(address , address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Session) ExecuteAuctionComplete(arg0 common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteAuctionComplete(&_X2Y2.TransactOpts, arg0, buyer, data)
}

// ExecuteAuctionComplete is a paid mutator transaction binding the contract method 0x3672c911.
//
// Solidity: function executeAuctionComplete(address , address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2TransactorSession) ExecuteAuctionComplete(arg0 common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteAuctionComplete(&_X2Y2.TransactOpts, arg0, buyer, data)
}

// ExecuteAuctionRefund is a paid mutator transaction binding the contract method 0xf477e4fd.
//
// Solidity: function executeAuctionRefund(address seller, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2Transactor) ExecuteAuctionRefund(opts *bind.TransactOpts, seller common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "executeAuctionRefund", seller, arg1, data)
}

// ExecuteAuctionRefund is a paid mutator transaction binding the contract method 0xf477e4fd.
//
// Solidity: function executeAuctionRefund(address seller, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2Session) ExecuteAuctionRefund(seller common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteAuctionRefund(&_X2Y2.TransactOpts, seller, arg1, data)
}

// ExecuteAuctionRefund is a paid mutator transaction binding the contract method 0xf477e4fd.
//
// Solidity: function executeAuctionRefund(address seller, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2TransactorSession) ExecuteAuctionRefund(seller common.Address, arg1 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteAuctionRefund(&_X2Y2.TransactOpts, seller, arg1, data)
}

// ExecuteBid is a paid mutator transaction binding the contract method 0xc23725f9.
//
// Solidity: function executeBid(address seller, address previousBidder, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2Transactor) ExecuteBid(opts *bind.TransactOpts, seller common.Address, previousBidder common.Address, arg2 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "executeBid", seller, previousBidder, arg2, data)
}

// ExecuteBid is a paid mutator transaction binding the contract method 0xc23725f9.
//
// Solidity: function executeBid(address seller, address previousBidder, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2Session) ExecuteBid(seller common.Address, previousBidder common.Address, arg2 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteBid(&_X2Y2.TransactOpts, seller, previousBidder, arg2, data)
}

// ExecuteBid is a paid mutator transaction binding the contract method 0xc23725f9.
//
// Solidity: function executeBid(address seller, address previousBidder, address , bytes data) returns(bool)
func (_X2Y2 *X2Y2TransactorSession) ExecuteBid(seller common.Address, previousBidder common.Address, arg2 common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteBid(&_X2Y2.TransactOpts, seller, previousBidder, arg2, data)
}

// ExecuteBuy is a paid mutator transaction binding the contract method 0x16721626.
//
// Solidity: function executeBuy(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Transactor) ExecuteBuy(opts *bind.TransactOpts, seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "executeBuy", seller, buyer, data)
}

// ExecuteBuy is a paid mutator transaction binding the contract method 0x16721626.
//
// Solidity: function executeBuy(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Session) ExecuteBuy(seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteBuy(&_X2Y2.TransactOpts, seller, buyer, data)
}

// ExecuteBuy is a paid mutator transaction binding the contract method 0x16721626.
//
// Solidity: function executeBuy(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2TransactorSession) ExecuteBuy(seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteBuy(&_X2Y2.TransactOpts, seller, buyer, data)
}

// ExecuteSell is a paid mutator transaction binding the contract method 0xbc553f0f.
//
// Solidity: function executeSell(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Transactor) ExecuteSell(opts *bind.TransactOpts, seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "executeSell", seller, buyer, data)
}

// ExecuteSell is a paid mutator transaction binding the contract method 0xbc553f0f.
//
// Solidity: function executeSell(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2Session) ExecuteSell(seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteSell(&_X2Y2.TransactOpts, seller, buyer, data)
}

// ExecuteSell is a paid mutator transaction binding the contract method 0xbc553f0f.
//
// Solidity: function executeSell(address seller, address buyer, bytes data) returns(bool)
func (_X2Y2 *X2Y2TransactorSession) ExecuteSell(seller common.Address, buyer common.Address, data []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.ExecuteSell(&_X2Y2.TransactOpts, seller, buyer, data)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Transactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Session) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.GrantRole(&_X2Y2.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2TransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.GrantRole(&_X2Y2.TransactOpts, role, account)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_X2Y2 *X2Y2Transactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_X2Y2 *X2Y2Session) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.OnERC721Received(&_X2Y2.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_X2Y2 *X2Y2TransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _X2Y2.Contract.OnERC721Received(&_X2Y2.TransactOpts, arg0, arg1, arg2, arg3)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Transactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Session) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.RenounceRole(&_X2Y2.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2TransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.RenounceRole(&_X2Y2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Transactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2Session) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.RevokeRole(&_X2Y2.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_X2Y2 *X2Y2TransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.RevokeRole(&_X2Y2.TransactOpts, role, account)
}

// TransferBatch is a paid mutator transaction binding the contract method 0x3cbf4f8a.
//
// Solidity: function transferBatch((address,uint256)[] pairs, address to) returns()
func (_X2Y2 *X2Y2Transactor) TransferBatch(opts *bind.TransactOpts, pairs []ERC721DelegatePair, to common.Address) (*types.Transaction, error) {
	return _X2Y2.contract.Transact(opts, "transferBatch", pairs, to)
}

// TransferBatch is a paid mutator transaction binding the contract method 0x3cbf4f8a.
//
// Solidity: function transferBatch((address,uint256)[] pairs, address to) returns()
func (_X2Y2 *X2Y2Session) TransferBatch(pairs []ERC721DelegatePair, to common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.TransferBatch(&_X2Y2.TransactOpts, pairs, to)
}

// TransferBatch is a paid mutator transaction binding the contract method 0x3cbf4f8a.
//
// Solidity: function transferBatch((address,uint256)[] pairs, address to) returns()
func (_X2Y2 *X2Y2TransactorSession) TransferBatch(pairs []ERC721DelegatePair, to common.Address) (*types.Transaction, error) {
	return _X2Y2.Contract.TransferBatch(&_X2Y2.TransactOpts, pairs, to)
}

// X2Y2RoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the X2Y2 contract.
type X2Y2RoleAdminChangedIterator struct {
	Event *X2Y2RoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *X2Y2RoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2RoleAdminChanged)
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
		it.Event = new(X2Y2RoleAdminChanged)
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
func (it *X2Y2RoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2RoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2RoleAdminChanged represents a RoleAdminChanged event raised by the X2Y2 contract.
type X2Y2RoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_X2Y2 *X2Y2Filterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*X2Y2RoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _X2Y2.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2RoleAdminChangedIterator{contract: _X2Y2.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_X2Y2 *X2Y2Filterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *X2Y2RoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _X2Y2.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2RoleAdminChanged)
				if err := _X2Y2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_X2Y2 *X2Y2Filterer) ParseRoleAdminChanged(log types.Log) (*X2Y2RoleAdminChanged, error) {
	event := new(X2Y2RoleAdminChanged)
	if err := _X2Y2.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2RoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the X2Y2 contract.
type X2Y2RoleGrantedIterator struct {
	Event *X2Y2RoleGranted // Event containing the contract specifics and raw log

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
func (it *X2Y2RoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2RoleGranted)
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
		it.Event = new(X2Y2RoleGranted)
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
func (it *X2Y2RoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2RoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2RoleGranted represents a RoleGranted event raised by the X2Y2 contract.
type X2Y2RoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*X2Y2RoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _X2Y2.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2RoleGrantedIterator{contract: _X2Y2.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *X2Y2RoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _X2Y2.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2RoleGranted)
				if err := _X2Y2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) ParseRoleGranted(log types.Log) (*X2Y2RoleGranted, error) {
	event := new(X2Y2RoleGranted)
	if err := _X2Y2.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2RoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the X2Y2 contract.
type X2Y2RoleRevokedIterator struct {
	Event *X2Y2RoleRevoked // Event containing the contract specifics and raw log

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
func (it *X2Y2RoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2RoleRevoked)
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
		it.Event = new(X2Y2RoleRevoked)
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
func (it *X2Y2RoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2RoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2RoleRevoked represents a RoleRevoked event raised by the X2Y2 contract.
type X2Y2RoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*X2Y2RoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _X2Y2.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2RoleRevokedIterator{contract: _X2Y2.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *X2Y2RoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _X2Y2.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2RoleRevoked)
				if err := _X2Y2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_X2Y2 *X2Y2Filterer) ParseRoleRevoked(log types.Log) (*X2Y2RoleRevoked, error) {
	event := new(X2Y2RoleRevoked)
	if err := _X2Y2.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
