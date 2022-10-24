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

// MarketFee is an auto generated low-level Go binding around an user-defined struct.
type MarketFee struct {
	Percentage *big.Int
	To         common.Address
}

// MarketOrder is an auto generated low-level Go binding around an user-defined struct.
type MarketOrder struct {
	Salt         *big.Int
	User         common.Address
	Network      *big.Int
	Intent       *big.Int
	DelegateType *big.Int
	Deadline     *big.Int
	Currency     common.Address
	DataMask     []byte
	Items        []MarketOrderItem
	R            [32]byte
	S            [32]byte
	V            uint8
	SignVersion  uint8
}

// MarketOrderItem is an auto generated low-level Go binding around an user-defined struct.
type MarketOrderItem struct {
	Price *big.Int
	Data  []byte
}

// MarketRunInput is an auto generated low-level Go binding around an user-defined struct.
type MarketRunInput struct {
	Orders  []MarketOrder
	Details []MarketSettleDetail
	Shared  MarketSettleShared
	R       [32]byte
	S       [32]byte
	V       uint8
}

// MarketSettleDetail is an auto generated low-level Go binding around an user-defined struct.
type MarketSettleDetail struct {
	Op                 uint8
	OrderIdx           *big.Int
	ItemIdx            *big.Int
	Price              *big.Int
	ItemHash           [32]byte
	ExecutionDelegate  common.Address
	DataReplacement    []byte
	BidIncentivePct    *big.Int
	AucMinIncrementPct *big.Int
	AucIncDurationSecs *big.Int
	Fees               []MarketFee
}

// MarketSettleShared is an auto generated low-level Go binding around an user-defined struct.
type MarketSettleShared struct {
	Salt         *big.Int
	Deadline     *big.Int
	AmountToEth  *big.Int
	AmountToWeth *big.Int
	User         common.Address
	CanFail      bool
}

// X2Y2R1MetaData contains all meta data concerning the X2Y2R1 contract.
var X2Y2R1MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"incentive\",\"type\":\"uint256\"}],\"name\":\"EvAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"}],\"name\":\"EvCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRemoval\",\"type\":\"bool\"}],\"name\":\"EvDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"error\",\"type\":\"bytes\"}],\"name\":\"EvFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"EvFeeCapUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"orderSalt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settleSalt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structMarket.OrderItem\",\"name\":\"item\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"indexed\":false,\"internalType\":\"structMarket.SettleDetail\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"EvInventory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EvProfit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isRemoval\",\"type\":\"bool\"}],\"name\":\"EvSigner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"RATE_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"itemHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCapPct\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"feeCapPct_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"weth_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"inventoryStatus\",\"outputs\":[{\"internalType\":\"enumMarket.InvStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"ongoingAuctions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"netPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMarket.OrderItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"signVersion\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order[]\",\"name\":\"orders\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"internalType\":\"structMarket.SettleDetail[]\",\"name\":\"details\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWeth\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"canFail\",\"type\":\"bool\"}],\"internalType\":\"structMarket.SettleShared\",\"name\":\"shared\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.RunInput\",\"name\":\"input\",\"type\":\"tuple\"}],\"name\":\"run\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"network\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"intent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"delegateType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"contractIERC20Upgradeable\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataMask\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMarket.OrderItem[]\",\"name\":\"items\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"signVersion\",\"type\":\"uint8\"}],\"internalType\":\"structMarket.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToEth\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToWeth\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"canFail\",\"type\":\"bool\"}],\"internalType\":\"structMarket.SettleShared\",\"name\":\"shared\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"enumMarket.Op\",\"name\":\"op\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"orderIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemIdx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itemHash\",\"type\":\"bytes32\"},{\"internalType\":\"contractIDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"dataReplacement\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"bidIncentivePct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucMinIncrementPct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aucIncDurationSecs\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"internalType\":\"structMarket.Fee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"}],\"internalType\":\"structMarket.SettleDetail\",\"name\":\"detail\",\"type\":\"tuple\"}],\"name\":\"run1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toRemove\",\"type\":\"address[]\"}],\"name\":\"updateDelegates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"updateFeeCap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toAdd\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"toRemove\",\"type\":\"address[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weth\",\"outputs\":[{\"internalType\":\"contractIWETHUpgradable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// X2Y2R1ABI is the input ABI used to generate the binding from.
// Deprecated: Use X2Y2R1MetaData.ABI instead.
var X2Y2R1ABI = X2Y2R1MetaData.ABI

// X2Y2R1 is an auto generated Go binding around an Ethereum contract.
type X2Y2R1 struct {
	X2Y2R1Caller     // Read-only binding to the contract
	X2Y2R1Transactor // Write-only binding to the contract
	X2Y2R1Filterer   // Log filterer for contract events
}

// X2Y2R1Caller is an auto generated read-only Go binding around an Ethereum contract.
type X2Y2R1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2R1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type X2Y2R1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2R1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type X2Y2R1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// X2Y2R1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type X2Y2R1Session struct {
	Contract     *X2Y2R1           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// X2Y2R1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type X2Y2R1CallerSession struct {
	Contract *X2Y2R1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// X2Y2R1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type X2Y2R1TransactorSession struct {
	Contract     *X2Y2R1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// X2Y2R1Raw is an auto generated low-level Go binding around an Ethereum contract.
type X2Y2R1Raw struct {
	Contract *X2Y2R1 // Generic contract binding to access the raw methods on
}

// X2Y2R1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type X2Y2R1CallerRaw struct {
	Contract *X2Y2R1Caller // Generic read-only contract binding to access the raw methods on
}

// X2Y2R1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type X2Y2R1TransactorRaw struct {
	Contract *X2Y2R1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewX2Y2R1 creates a new instance of X2Y2R1, bound to a specific deployed contract.
func NewX2Y2R1(address common.Address, backend bind.ContractBackend) (*X2Y2R1, error) {
	contract, err := bindX2Y2R1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1{X2Y2R1Caller: X2Y2R1Caller{contract: contract}, X2Y2R1Transactor: X2Y2R1Transactor{contract: contract}, X2Y2R1Filterer: X2Y2R1Filterer{contract: contract}}, nil
}

// NewX2Y2R1Caller creates a new read-only instance of X2Y2R1, bound to a specific deployed contract.
func NewX2Y2R1Caller(address common.Address, caller bind.ContractCaller) (*X2Y2R1Caller, error) {
	contract, err := bindX2Y2R1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1Caller{contract: contract}, nil
}

// NewX2Y2R1Transactor creates a new write-only instance of X2Y2R1, bound to a specific deployed contract.
func NewX2Y2R1Transactor(address common.Address, transactor bind.ContractTransactor) (*X2Y2R1Transactor, error) {
	contract, err := bindX2Y2R1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1Transactor{contract: contract}, nil
}

// NewX2Y2R1Filterer creates a new log filterer instance of X2Y2R1, bound to a specific deployed contract.
func NewX2Y2R1Filterer(address common.Address, filterer bind.ContractFilterer) (*X2Y2R1Filterer, error) {
	contract, err := bindX2Y2R1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1Filterer{contract: contract}, nil
}

// bindX2Y2R1 binds a generic wrapper to an already deployed contract.
func bindX2Y2R1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(X2Y2R1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X2Y2R1 *X2Y2R1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X2Y2R1.Contract.X2Y2R1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X2Y2R1 *X2Y2R1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.Contract.X2Y2R1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X2Y2R1 *X2Y2R1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X2Y2R1.Contract.X2Y2R1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_X2Y2R1 *X2Y2R1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _X2Y2R1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_X2Y2R1 *X2Y2R1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_X2Y2R1 *X2Y2R1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _X2Y2R1.Contract.contract.Transact(opts, method, params...)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_X2Y2R1 *X2Y2R1Caller) RATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_X2Y2R1 *X2Y2R1Session) RATEBASE() (*big.Int, error) {
	return _X2Y2R1.Contract.RATEBASE(&_X2Y2R1.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_X2Y2R1 *X2Y2R1CallerSession) RATEBASE() (*big.Int, error) {
	return _X2Y2R1.Contract.RATEBASE(&_X2Y2R1.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1Caller) Delegates(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "delegates", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1Session) Delegates(arg0 common.Address) (bool, error) {
	return _X2Y2R1.Contract.Delegates(&_X2Y2R1.CallOpts, arg0)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1CallerSession) Delegates(arg0 common.Address) (bool, error) {
	return _X2Y2R1.Contract.Delegates(&_X2Y2R1.CallOpts, arg0)
}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_X2Y2R1 *X2Y2R1Caller) FeeCapPct(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "feeCapPct")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_X2Y2R1 *X2Y2R1Session) FeeCapPct() (*big.Int, error) {
	return _X2Y2R1.Contract.FeeCapPct(&_X2Y2R1.CallOpts)
}

// FeeCapPct is a free data retrieval call binding the contract method 0x9fb51467.
//
// Solidity: function feeCapPct() view returns(uint256)
func (_X2Y2R1 *X2Y2R1CallerSession) FeeCapPct() (*big.Int, error) {
	return _X2Y2R1.Contract.FeeCapPct(&_X2Y2R1.CallOpts)
}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_X2Y2R1 *X2Y2R1Caller) InventoryStatus(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "inventoryStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_X2Y2R1 *X2Y2R1Session) InventoryStatus(arg0 [32]byte) (uint8, error) {
	return _X2Y2R1.Contract.InventoryStatus(&_X2Y2R1.CallOpts, arg0)
}

// InventoryStatus is a free data retrieval call binding the contract method 0x912c860c.
//
// Solidity: function inventoryStatus(bytes32 ) view returns(uint8)
func (_X2Y2R1 *X2Y2R1CallerSession) InventoryStatus(arg0 [32]byte) (uint8, error) {
	return _X2Y2R1.Contract.InventoryStatus(&_X2Y2R1.CallOpts, arg0)
}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_X2Y2R1 *X2Y2R1Caller) OngoingAuctions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "ongoingAuctions", arg0)

	outstruct := new(struct {
		Price    *big.Int
		NetPrice *big.Int
		EndAt    *big.Int
		Bidder   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NetPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_X2Y2R1 *X2Y2R1Session) OngoingAuctions(arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	return _X2Y2R1.Contract.OngoingAuctions(&_X2Y2R1.CallOpts, arg0)
}

// OngoingAuctions is a free data retrieval call binding the contract method 0xea805917.
//
// Solidity: function ongoingAuctions(bytes32 ) view returns(uint256 price, uint256 netPrice, uint256 endAt, address bidder)
func (_X2Y2R1 *X2Y2R1CallerSession) OngoingAuctions(arg0 [32]byte) (struct {
	Price    *big.Int
	NetPrice *big.Int
	EndAt    *big.Int
	Bidder   common.Address
}, error) {
	return _X2Y2R1.Contract.OngoingAuctions(&_X2Y2R1.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_X2Y2R1 *X2Y2R1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_X2Y2R1 *X2Y2R1Session) Owner() (common.Address, error) {
	return _X2Y2R1.Contract.Owner(&_X2Y2R1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_X2Y2R1 *X2Y2R1CallerSession) Owner() (common.Address, error) {
	return _X2Y2R1.Contract.Owner(&_X2Y2R1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_X2Y2R1 *X2Y2R1Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_X2Y2R1 *X2Y2R1Session) Paused() (bool, error) {
	return _X2Y2R1.Contract.Paused(&_X2Y2R1.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_X2Y2R1 *X2Y2R1CallerSession) Paused() (bool, error) {
	return _X2Y2R1.Contract.Paused(&_X2Y2R1.CallOpts)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1Caller) Signers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "signers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1Session) Signers(arg0 common.Address) (bool, error) {
	return _X2Y2R1.Contract.Signers(&_X2Y2R1.CallOpts, arg0)
}

// Signers is a free data retrieval call binding the contract method 0x736c0d5b.
//
// Solidity: function signers(address ) view returns(bool)
func (_X2Y2R1 *X2Y2R1CallerSession) Signers(arg0 common.Address) (bool, error) {
	return _X2Y2R1.Contract.Signers(&_X2Y2R1.CallOpts, arg0)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_X2Y2R1 *X2Y2R1Caller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _X2Y2R1.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_X2Y2R1 *X2Y2R1Session) Weth() (common.Address, error) {
	return _X2Y2R1.Contract.Weth(&_X2Y2R1.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_X2Y2R1 *X2Y2R1CallerSession) Weth() (common.Address, error) {
	return _X2Y2R1.Contract.Weth(&_X2Y2R1.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_X2Y2R1 *X2Y2R1Transactor) Cancel(opts *bind.TransactOpts, itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "cancel", itemHashes, deadline, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_X2Y2R1 *X2Y2R1Session) Cancel(itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Cancel(&_X2Y2R1.TransactOpts, itemHashes, deadline, v, r, s)
}

// Cancel is a paid mutator transaction binding the contract method 0x2295f9bf.
//
// Solidity: function cancel(bytes32[] itemHashes, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Cancel(itemHashes [][32]byte, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Cancel(&_X2Y2R1.TransactOpts, itemHashes, deadline, v, r, s)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_X2Y2R1 *X2Y2R1Transactor) Initialize(opts *bind.TransactOpts, feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "initialize", feeCapPct_, weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_X2Y2R1 *X2Y2R1Session) Initialize(feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Initialize(&_X2Y2R1.TransactOpts, feeCapPct_, weth_)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 feeCapPct_, address weth_) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Initialize(feeCapPct_ *big.Int, weth_ common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Initialize(&_X2Y2R1.TransactOpts, feeCapPct_, weth_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_X2Y2R1 *X2Y2R1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_X2Y2R1 *X2Y2R1Session) Pause() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Pause(&_X2Y2R1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Pause() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Pause(&_X2Y2R1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_X2Y2R1 *X2Y2R1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_X2Y2R1 *X2Y2R1Session) RenounceOwnership() (*types.Transaction, error) {
	return _X2Y2R1.Contract.RenounceOwnership(&_X2Y2R1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _X2Y2R1.Contract.RenounceOwnership(&_X2Y2R1.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_X2Y2R1 *X2Y2R1Transactor) Run(opts *bind.TransactOpts, input MarketRunInput) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "run", input)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_X2Y2R1 *X2Y2R1Session) Run(input MarketRunInput) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Run(&_X2Y2R1.TransactOpts, input)
}

// Run is a paid mutator transaction binding the contract method 0x357a150b.
//
// Solidity: function run(((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8)[],(uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[])[],(uint256,uint256,uint256,uint256,address,bool),bytes32,bytes32,uint8) input) payable returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Run(input MarketRunInput) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Run(&_X2Y2R1.TransactOpts, input)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_X2Y2R1 *X2Y2R1Transactor) Run1(opts *bind.TransactOpts, order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "run1", order, shared, detail)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_X2Y2R1 *X2Y2R1Session) Run1(order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Run1(&_X2Y2R1.TransactOpts, order, shared, detail)
}

// Run1 is a paid mutator transaction binding the contract method 0xd95e3c54.
//
// Solidity: function run1((uint256,address,uint256,uint256,uint256,uint256,address,bytes,(uint256,bytes)[],bytes32,bytes32,uint8,uint8) order, (uint256,uint256,uint256,uint256,address,bool) shared, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail) returns(uint256)
func (_X2Y2R1 *X2Y2R1TransactorSession) Run1(order MarketOrder, shared MarketSettleShared, detail MarketSettleDetail) (*types.Transaction, error) {
	return _X2Y2R1.Contract.Run1(&_X2Y2R1.TransactOpts, order, shared, detail)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_X2Y2R1 *X2Y2R1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_X2Y2R1 *X2Y2R1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.TransferOwnership(&_X2Y2R1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.TransferOwnership(&_X2Y2R1.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_X2Y2R1 *X2Y2R1Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_X2Y2R1 *X2Y2R1Session) Unpause() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Unpause(&_X2Y2R1.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Unpause() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Unpause(&_X2Y2R1.TransactOpts)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1Transactor) UpdateDelegates(opts *bind.TransactOpts, toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "updateDelegates", toAdd, toRemove)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1Session) UpdateDelegates(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateDelegates(&_X2Y2R1.TransactOpts, toAdd, toRemove)
}

// UpdateDelegates is a paid mutator transaction binding the contract method 0xe59f739a.
//
// Solidity: function updateDelegates(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) UpdateDelegates(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateDelegates(&_X2Y2R1.TransactOpts, toAdd, toRemove)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_X2Y2R1 *X2Y2R1Transactor) UpdateFeeCap(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "updateFeeCap", val)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_X2Y2R1 *X2Y2R1Session) UpdateFeeCap(val *big.Int) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateFeeCap(&_X2Y2R1.TransactOpts, val)
}

// UpdateFeeCap is a paid mutator transaction binding the contract method 0x95835fea.
//
// Solidity: function updateFeeCap(uint256 val) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) UpdateFeeCap(val *big.Int) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateFeeCap(&_X2Y2R1.TransactOpts, val)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1Transactor) UpdateSigners(opts *bind.TransactOpts, toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.contract.Transact(opts, "updateSigners", toAdd, toRemove)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1Session) UpdateSigners(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateSigners(&_X2Y2R1.TransactOpts, toAdd, toRemove)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0x350b2369.
//
// Solidity: function updateSigners(address[] toAdd, address[] toRemove) returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) UpdateSigners(toAdd []common.Address, toRemove []common.Address) (*types.Transaction, error) {
	return _X2Y2R1.Contract.UpdateSigners(&_X2Y2R1.TransactOpts, toAdd, toRemove)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_X2Y2R1 *X2Y2R1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _X2Y2R1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_X2Y2R1 *X2Y2R1Session) Receive() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Receive(&_X2Y2R1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_X2Y2R1 *X2Y2R1TransactorSession) Receive() (*types.Transaction, error) {
	return _X2Y2R1.Contract.Receive(&_X2Y2R1.TransactOpts)
}

// X2Y2R1EvAuctionRefundIterator is returned from FilterEvAuctionRefund and is used to iterate over the raw logs and unpacked data for EvAuctionRefund events raised by the X2Y2R1 contract.
type X2Y2R1EvAuctionRefundIterator struct {
	Event *X2Y2R1EvAuctionRefund // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvAuctionRefund)
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
		it.Event = new(X2Y2R1EvAuctionRefund)
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
func (it *X2Y2R1EvAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvAuctionRefund represents a EvAuctionRefund event raised by the X2Y2R1 contract.
type X2Y2R1EvAuctionRefund struct {
	ItemHash  [32]byte
	Currency  common.Address
	To        common.Address
	Amount    *big.Int
	Incentive *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvAuctionRefund is a free log retrieval operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvAuctionRefund(opts *bind.FilterOpts, itemHash [][32]byte) (*X2Y2R1EvAuctionRefundIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvAuctionRefund", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvAuctionRefundIterator{contract: _X2Y2R1.contract, event: "EvAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchEvAuctionRefund is a free log subscription operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvAuctionRefund(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvAuctionRefund, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvAuctionRefund", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvAuctionRefund)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
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

// ParseEvAuctionRefund is a log parse operation binding the contract event 0x681e2055b67e23ce693a446bd0567fb9df559ce6f82da4397482bad968551ac2.
//
// Solidity: event EvAuctionRefund(bytes32 indexed itemHash, address currency, address to, uint256 amount, uint256 incentive)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvAuctionRefund(log types.Log) (*X2Y2R1EvAuctionRefund, error) {
	event := new(X2Y2R1EvAuctionRefund)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvCancelIterator is returned from FilterEvCancel and is used to iterate over the raw logs and unpacked data for EvCancel events raised by the X2Y2R1 contract.
type X2Y2R1EvCancelIterator struct {
	Event *X2Y2R1EvCancel // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvCancel)
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
		it.Event = new(X2Y2R1EvCancel)
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
func (it *X2Y2R1EvCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvCancel represents a EvCancel event raised by the X2Y2R1 contract.
type X2Y2R1EvCancel struct {
	ItemHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvCancel is a free log retrieval operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvCancel(opts *bind.FilterOpts, itemHash [][32]byte) (*X2Y2R1EvCancelIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvCancel", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvCancelIterator{contract: _X2Y2R1.contract, event: "EvCancel", logs: logs, sub: sub}, nil
}

// WatchEvCancel is a free log subscription operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvCancel(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvCancel, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvCancel", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvCancel)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvCancel", log); err != nil {
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

// ParseEvCancel is a log parse operation binding the contract event 0x5b0b06d07e20243724d90e17a20034972f339eb28bd1c9437a71999bd15a1e7a.
//
// Solidity: event EvCancel(bytes32 indexed itemHash)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvCancel(log types.Log) (*X2Y2R1EvCancel, error) {
	event := new(X2Y2R1EvCancel)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvCancel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvDelegateIterator is returned from FilterEvDelegate and is used to iterate over the raw logs and unpacked data for EvDelegate events raised by the X2Y2R1 contract.
type X2Y2R1EvDelegateIterator struct {
	Event *X2Y2R1EvDelegate // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvDelegate)
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
		it.Event = new(X2Y2R1EvDelegate)
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
func (it *X2Y2R1EvDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvDelegate represents a EvDelegate event raised by the X2Y2R1 contract.
type X2Y2R1EvDelegate struct {
	Delegate  common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvDelegate is a free log retrieval operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvDelegate(opts *bind.FilterOpts) (*X2Y2R1EvDelegateIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvDelegate")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvDelegateIterator{contract: _X2Y2R1.contract, event: "EvDelegate", logs: logs, sub: sub}, nil
}

// WatchEvDelegate is a free log subscription operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvDelegate(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvDelegate) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvDelegate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvDelegate)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvDelegate", log); err != nil {
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

// ParseEvDelegate is a log parse operation binding the contract event 0x4a31a64b928a0e8aff42ef84d144ffe82d08cb41c8027060593135e2026899b2.
//
// Solidity: event EvDelegate(address delegate, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvDelegate(log types.Log) (*X2Y2R1EvDelegate, error) {
	event := new(X2Y2R1EvDelegate)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvFailureIterator is returned from FilterEvFailure and is used to iterate over the raw logs and unpacked data for EvFailure events raised by the X2Y2R1 contract.
type X2Y2R1EvFailureIterator struct {
	Event *X2Y2R1EvFailure // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvFailure)
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
		it.Event = new(X2Y2R1EvFailure)
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
func (it *X2Y2R1EvFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvFailure represents a EvFailure event raised by the X2Y2R1 contract.
type X2Y2R1EvFailure struct {
	Index *big.Int
	Error []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEvFailure is a free log retrieval operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvFailure(opts *bind.FilterOpts) (*X2Y2R1EvFailureIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvFailure")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvFailureIterator{contract: _X2Y2R1.contract, event: "EvFailure", logs: logs, sub: sub}, nil
}

// WatchEvFailure is a free log subscription operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvFailure(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvFailure) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvFailure)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvFailure", log); err != nil {
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

// ParseEvFailure is a log parse operation binding the contract event 0x97c789f43a3e7ac27906b5fbdac832f54441771021fba06f71207d9be6d4b623.
//
// Solidity: event EvFailure(uint256 index, bytes error)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvFailure(log types.Log) (*X2Y2R1EvFailure, error) {
	event := new(X2Y2R1EvFailure)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvFeeCapUpdateIterator is returned from FilterEvFeeCapUpdate and is used to iterate over the raw logs and unpacked data for EvFeeCapUpdate events raised by the X2Y2R1 contract.
type X2Y2R1EvFeeCapUpdateIterator struct {
	Event *X2Y2R1EvFeeCapUpdate // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvFeeCapUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvFeeCapUpdate)
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
		it.Event = new(X2Y2R1EvFeeCapUpdate)
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
func (it *X2Y2R1EvFeeCapUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvFeeCapUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvFeeCapUpdate represents a EvFeeCapUpdate event raised by the X2Y2R1 contract.
type X2Y2R1EvFeeCapUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvFeeCapUpdate is a free log retrieval operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvFeeCapUpdate(opts *bind.FilterOpts) (*X2Y2R1EvFeeCapUpdateIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvFeeCapUpdate")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvFeeCapUpdateIterator{contract: _X2Y2R1.contract, event: "EvFeeCapUpdate", logs: logs, sub: sub}, nil
}

// WatchEvFeeCapUpdate is a free log subscription operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvFeeCapUpdate(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvFeeCapUpdate) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvFeeCapUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvFeeCapUpdate)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvFeeCapUpdate", log); err != nil {
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

// ParseEvFeeCapUpdate is a log parse operation binding the contract event 0x19fc3beddeea399f0966d5f8664ad94006f16a10fb28c4e2fe6fae62626b7128.
//
// Solidity: event EvFeeCapUpdate(uint256 newValue)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvFeeCapUpdate(log types.Log) (*X2Y2R1EvFeeCapUpdate, error) {
	event := new(X2Y2R1EvFeeCapUpdate)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvFeeCapUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvInventoryIterator is returned from FilterEvInventory and is used to iterate over the raw logs and unpacked data for EvInventory events raised by the X2Y2R1 contract.
type X2Y2R1EvInventoryIterator struct {
	Event *X2Y2R1EvInventory // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvInventoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvInventory)
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
		it.Event = new(X2Y2R1EvInventory)
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
func (it *X2Y2R1EvInventoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvInventoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvInventory represents a EvInventory event raised by the X2Y2R1 contract.
type X2Y2R1EvInventory struct {
	ItemHash     [32]byte
	Maker        common.Address
	Taker        common.Address
	OrderSalt    *big.Int
	SettleSalt   *big.Int
	Intent       *big.Int
	DelegateType *big.Int
	Deadline     *big.Int
	Currency     common.Address
	DataMask     []byte
	Item         MarketOrderItem
	Detail       MarketSettleDetail
	Raw          types.Log // Blockchain specific contextual infos
	
}

// FilterEvInventory is a free log retrieval operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvInventory(opts *bind.FilterOpts, itemHash [][32]byte) (*X2Y2R1EvInventoryIterator, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvInventory", itemHashRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvInventoryIterator{contract: _X2Y2R1.contract, event: "EvInventory", logs: logs, sub: sub}, nil
}

// WatchEvInventory is a free log subscription operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvInventory(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvInventory, itemHash [][32]byte) (event.Subscription, error) {

	var itemHashRule []interface{}
	for _, itemHashItem := range itemHash {
		itemHashRule = append(itemHashRule, itemHashItem)
	}

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvInventory", itemHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvInventory)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvInventory", log); err != nil {
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

// ParseEvInventory is a log parse operation binding the contract event 0x3cbb63f144840e5b1b0a38a7c19211d2e89de4d7c5faf8b2d3c1776c302d1d33.
//
// Solidity: event EvInventory(bytes32 indexed itemHash, address maker, address taker, uint256 orderSalt, uint256 settleSalt, uint256 intent, uint256 delegateType, uint256 deadline, address currency, bytes dataMask, (uint256,bytes) item, (uint8,uint256,uint256,uint256,bytes32,address,bytes,uint256,uint256,uint256,(uint256,address)[]) detail)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvInventory(log types.Log) (*X2Y2R1EvInventory, error) {
	event := new(X2Y2R1EvInventory)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvInventory", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvProfitIterator is returned from FilterEvProfit and is used to iterate over the raw logs and unpacked data for EvProfit events raised by the X2Y2R1 contract.
type X2Y2R1EvProfitIterator struct {
	Event *X2Y2R1EvProfit // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvProfitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvProfit)
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
		it.Event = new(X2Y2R1EvProfit)
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
func (it *X2Y2R1EvProfitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvProfitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvProfit represents a EvProfit event raised by the X2Y2R1 contract.
type X2Y2R1EvProfit struct {
	ItemHash [32]byte
	Currency common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvProfit is a free log retrieval operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvProfit(opts *bind.FilterOpts) (*X2Y2R1EvProfitIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvProfit")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvProfitIterator{contract: _X2Y2R1.contract, event: "EvProfit", logs: logs, sub: sub}, nil
}

// WatchEvProfit is a free log subscription operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvProfit(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvProfit) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvProfit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvProfit)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvProfit", log); err != nil {
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

// ParseEvProfit is a log parse operation binding the contract event 0xe2c49856b032c255ae7e325d18109bc4e22a2804e2e49a017ec0f59f19cd447b.
//
// Solidity: event EvProfit(bytes32 itemHash, address currency, address to, uint256 amount)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvProfit(log types.Log) (*X2Y2R1EvProfit, error) {
	event := new(X2Y2R1EvProfit)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvProfit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1EvSignerIterator is returned from FilterEvSigner and is used to iterate over the raw logs and unpacked data for EvSigner events raised by the X2Y2R1 contract.
type X2Y2R1EvSignerIterator struct {
	Event *X2Y2R1EvSigner // Event containing the contract specifics and raw log

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
func (it *X2Y2R1EvSignerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1EvSigner)
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
		it.Event = new(X2Y2R1EvSigner)
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
func (it *X2Y2R1EvSignerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1EvSignerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1EvSigner represents a EvSigner event raised by the X2Y2R1 contract.
type X2Y2R1EvSigner struct {
	Signer    common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvSigner is a free log retrieval operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) FilterEvSigner(opts *bind.FilterOpts) (*X2Y2R1EvSignerIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "EvSigner")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1EvSignerIterator{contract: _X2Y2R1.contract, event: "EvSigner", logs: logs, sub: sub}, nil
}

// WatchEvSigner is a free log subscription operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) WatchEvSigner(opts *bind.WatchOpts, sink chan<- *X2Y2R1EvSigner) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "EvSigner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1EvSigner)
				if err := _X2Y2R1.contract.UnpackLog(event, "EvSigner", log); err != nil {
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

// ParseEvSigner is a log parse operation binding the contract event 0x0127aee741cbb6bc48b5475b8eb3eb2e5d053809d551dedd517a0b5b52b80fd5.
//
// Solidity: event EvSigner(address signer, bool isRemoval)
func (_X2Y2R1 *X2Y2R1Filterer) ParseEvSigner(log types.Log) (*X2Y2R1EvSigner, error) {
	event := new(X2Y2R1EvSigner)
	if err := _X2Y2R1.contract.UnpackLog(event, "EvSigner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the X2Y2R1 contract.
type X2Y2R1OwnershipTransferredIterator struct {
	Event *X2Y2R1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *X2Y2R1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1OwnershipTransferred)
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
		it.Event = new(X2Y2R1OwnershipTransferred)
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
func (it *X2Y2R1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1OwnershipTransferred represents a OwnershipTransferred event raised by the X2Y2R1 contract.
type X2Y2R1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_X2Y2R1 *X2Y2R1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*X2Y2R1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &X2Y2R1OwnershipTransferredIterator{contract: _X2Y2R1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_X2Y2R1 *X2Y2R1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *X2Y2R1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1OwnershipTransferred)
				if err := _X2Y2R1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_X2Y2R1 *X2Y2R1Filterer) ParseOwnershipTransferred(log types.Log) (*X2Y2R1OwnershipTransferred, error) {
	event := new(X2Y2R1OwnershipTransferred)
	if err := _X2Y2R1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the X2Y2R1 contract.
type X2Y2R1PausedIterator struct {
	Event *X2Y2R1Paused // Event containing the contract specifics and raw log

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
func (it *X2Y2R1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1Paused)
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
		it.Event = new(X2Y2R1Paused)
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
func (it *X2Y2R1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1Paused represents a Paused event raised by the X2Y2R1 contract.
type X2Y2R1Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_X2Y2R1 *X2Y2R1Filterer) FilterPaused(opts *bind.FilterOpts) (*X2Y2R1PausedIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1PausedIterator{contract: _X2Y2R1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_X2Y2R1 *X2Y2R1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *X2Y2R1Paused) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1Paused)
				if err := _X2Y2R1.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_X2Y2R1 *X2Y2R1Filterer) ParsePaused(log types.Log) (*X2Y2R1Paused, error) {
	event := new(X2Y2R1Paused)
	if err := _X2Y2R1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// X2Y2R1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the X2Y2R1 contract.
type X2Y2R1UnpausedIterator struct {
	Event *X2Y2R1Unpaused // Event containing the contract specifics and raw log

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
func (it *X2Y2R1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(X2Y2R1Unpaused)
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
		it.Event = new(X2Y2R1Unpaused)
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
func (it *X2Y2R1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *X2Y2R1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// X2Y2R1Unpaused represents a Unpaused event raised by the X2Y2R1 contract.
type X2Y2R1Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_X2Y2R1 *X2Y2R1Filterer) FilterUnpaused(opts *bind.FilterOpts) (*X2Y2R1UnpausedIterator, error) {

	logs, sub, err := _X2Y2R1.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &X2Y2R1UnpausedIterator{contract: _X2Y2R1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_X2Y2R1 *X2Y2R1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *X2Y2R1Unpaused) (event.Subscription, error) {

	logs, sub, err := _X2Y2R1.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(X2Y2R1Unpaused)
				if err := _X2Y2R1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_X2Y2R1 *X2Y2R1Filterer) ParseUnpaused(log types.Log) (*X2Y2R1Unpaused, error) {
	event := new(X2Y2R1Unpaused)
	if err := _X2Y2R1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}