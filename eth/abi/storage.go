// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storage

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
	_ = abi.ConvertType
)

// StorageTokenMetaData contains all meta data concerning the StorageToken contract.
var StorageTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StorageTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageTokenMetaData.ABI instead.
var StorageTokenABI = StorageTokenMetaData.ABI

// StorageToken is an auto generated Go binding around an Ethereum contract.
type StorageToken struct {
	StorageTokenCaller     // Read-only binding to the contract
	StorageTokenTransactor // Write-only binding to the contract
	StorageTokenFilterer   // Log filterer for contract events
}

// StorageTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageTokenSession struct {
	Contract     *StorageToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageTokenCallerSession struct {
	Contract *StorageTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StorageTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageTokenTransactorSession struct {
	Contract     *StorageTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StorageTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageTokenRaw struct {
	Contract *StorageToken // Generic contract binding to access the raw methods on
}

// StorageTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageTokenCallerRaw struct {
	Contract *StorageTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StorageTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageTokenTransactorRaw struct {
	Contract *StorageTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageToken creates a new instance of StorageToken, bound to a specific deployed contract.
func NewStorageToken(address common.Address, backend bind.ContractBackend) (*StorageToken, error) {
	contract, err := bindStorageToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageToken{StorageTokenCaller: StorageTokenCaller{contract: contract}, StorageTokenTransactor: StorageTokenTransactor{contract: contract}, StorageTokenFilterer: StorageTokenFilterer{contract: contract}}, nil
}

// NewStorageTokenCaller creates a new read-only instance of StorageToken, bound to a specific deployed contract.
func NewStorageTokenCaller(address common.Address, caller bind.ContractCaller) (*StorageTokenCaller, error) {
	contract, err := bindStorageToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTokenCaller{contract: contract}, nil
}

// NewStorageTokenTransactor creates a new write-only instance of StorageToken, bound to a specific deployed contract.
func NewStorageTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageTokenTransactor, error) {
	contract, err := bindStorageToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageTokenTransactor{contract: contract}, nil
}

// NewStorageTokenFilterer creates a new log filterer instance of StorageToken, bound to a specific deployed contract.
func NewStorageTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageTokenFilterer, error) {
	contract, err := bindStorageToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageTokenFilterer{contract: contract}, nil
}

// bindStorageToken binds a generic wrapper to an already deployed contract.
func bindStorageToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageToken *StorageTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageToken.Contract.StorageTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageToken *StorageTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageToken.Contract.StorageTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageToken *StorageTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageToken.Contract.StorageTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageToken *StorageTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageToken *StorageTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageToken *StorageTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageToken.Contract.contract.Transact(opts, method, params...)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageToken *StorageTokenCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StorageToken.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageToken *StorageTokenSession) Retrieve() (*big.Int, error) {
	return _StorageToken.Contract.Retrieve(&_StorageToken.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_StorageToken *StorageTokenCallerSession) Retrieve() (*big.Int, error) {
	return _StorageToken.Contract.Retrieve(&_StorageToken.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageToken *StorageTokenTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _StorageToken.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageToken *StorageTokenSession) Store(num *big.Int) (*types.Transaction, error) {
	return _StorageToken.Contract.Store(&_StorageToken.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_StorageToken *StorageTokenTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _StorageToken.Contract.Store(&_StorageToken.TransactOpts, num)
}
