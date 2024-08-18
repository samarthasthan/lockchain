// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// SecretStorageSecretStruct is an auto generated low-level Go binding around an user-defined struct.
type SecretStorageSecretStruct struct {
	Id        *big.Int
	Secret    string
	CreatedAt *big.Int
}

// EthereumMetaData contains all meta data concerning the Ethereum contract.
var EthereumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_secret\",\"type\":\"string\"}],\"name\":\"setSecret\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSecrets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"secret\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structSecretStorage.SecretStruct[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getSecretByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"secret\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"internalType\":\"structSecretStorage.SecretStruct\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]",
}

// EthereumABI is the input ABI used to generate the binding from.
// Deprecated: Use EthereumMetaData.ABI instead.
var EthereumABI = EthereumMetaData.ABI

// Ethereum is an auto generated Go binding around an Ethereum contract.
type Ethereum struct {
	EthereumCaller     // Read-only binding to the contract
	EthereumTransactor // Write-only binding to the contract
	EthereumFilterer   // Log filterer for contract events
}

// EthereumCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthereumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthereumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthereumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthereumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthereumSession struct {
	Contract     *Ethereum         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthereumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthereumCallerSession struct {
	Contract *EthereumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthereumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthereumTransactorSession struct {
	Contract     *EthereumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthereumRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthereumRaw struct {
	Contract *Ethereum // Generic contract binding to access the raw methods on
}

// EthereumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthereumCallerRaw struct {
	Contract *EthereumCaller // Generic read-only contract binding to access the raw methods on
}

// EthereumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthereumTransactorRaw struct {
	Contract *EthereumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthereum creates a new instance of Ethereum, bound to a specific deployed contract.
func NewEthereum(address common.Address, backend bind.ContractBackend) (*Ethereum, error) {
	contract, err := bindEthereum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ethereum{EthereumCaller: EthereumCaller{contract: contract}, EthereumTransactor: EthereumTransactor{contract: contract}, EthereumFilterer: EthereumFilterer{contract: contract}}, nil
}

// NewEthereumCaller creates a new read-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumCaller(address common.Address, caller bind.ContractCaller) (*EthereumCaller, error) {
	contract, err := bindEthereum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumCaller{contract: contract}, nil
}

// NewEthereumTransactor creates a new write-only instance of Ethereum, bound to a specific deployed contract.
func NewEthereumTransactor(address common.Address, transactor bind.ContractTransactor) (*EthereumTransactor, error) {
	contract, err := bindEthereum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthereumTransactor{contract: contract}, nil
}

// NewEthereumFilterer creates a new log filterer instance of Ethereum, bound to a specific deployed contract.
func NewEthereumFilterer(address common.Address, filterer bind.ContractFilterer) (*EthereumFilterer, error) {
	contract, err := bindEthereum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthereumFilterer{contract: contract}, nil
}

// bindEthereum binds a generic wrapper to an already deployed contract.
func bindEthereum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthereumMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.EthereumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.EthereumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ethereum *EthereumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ethereum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ethereum *EthereumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ethereum *EthereumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ethereum.Contract.contract.Transact(opts, method, params...)
}

// GetSecretByIndex is a free data retrieval call binding the contract method 0x35c600e6.
//
// Solidity: function getSecretByIndex(uint256 _index) view returns((uint256,string,uint256))
func (_Ethereum *EthereumCaller) GetSecretByIndex(opts *bind.CallOpts, _index *big.Int) (SecretStorageSecretStruct, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getSecretByIndex", _index)

	if err != nil {
		return *new(SecretStorageSecretStruct), err
	}

	out0 := *abi.ConvertType(out[0], new(SecretStorageSecretStruct)).(*SecretStorageSecretStruct)

	return out0, err

}

// GetSecretByIndex is a free data retrieval call binding the contract method 0x35c600e6.
//
// Solidity: function getSecretByIndex(uint256 _index) view returns((uint256,string,uint256))
func (_Ethereum *EthereumSession) GetSecretByIndex(_index *big.Int) (SecretStorageSecretStruct, error) {
	return _Ethereum.Contract.GetSecretByIndex(&_Ethereum.CallOpts, _index)
}

// GetSecretByIndex is a free data retrieval call binding the contract method 0x35c600e6.
//
// Solidity: function getSecretByIndex(uint256 _index) view returns((uint256,string,uint256))
func (_Ethereum *EthereumCallerSession) GetSecretByIndex(_index *big.Int) (SecretStorageSecretStruct, error) {
	return _Ethereum.Contract.GetSecretByIndex(&_Ethereum.CallOpts, _index)
}

// GetSecrets is a free data retrieval call binding the contract method 0x996e2bee.
//
// Solidity: function getSecrets() view returns((uint256,string,uint256)[])
func (_Ethereum *EthereumCaller) GetSecrets(opts *bind.CallOpts) ([]SecretStorageSecretStruct, error) {
	var out []interface{}
	err := _Ethereum.contract.Call(opts, &out, "getSecrets")

	if err != nil {
		return *new([]SecretStorageSecretStruct), err
	}

	out0 := *abi.ConvertType(out[0], new([]SecretStorageSecretStruct)).(*[]SecretStorageSecretStruct)

	return out0, err

}

// GetSecrets is a free data retrieval call binding the contract method 0x996e2bee.
//
// Solidity: function getSecrets() view returns((uint256,string,uint256)[])
func (_Ethereum *EthereumSession) GetSecrets() ([]SecretStorageSecretStruct, error) {
	return _Ethereum.Contract.GetSecrets(&_Ethereum.CallOpts)
}

// GetSecrets is a free data retrieval call binding the contract method 0x996e2bee.
//
// Solidity: function getSecrets() view returns((uint256,string,uint256)[])
func (_Ethereum *EthereumCallerSession) GetSecrets() ([]SecretStorageSecretStruct, error) {
	return _Ethereum.Contract.GetSecrets(&_Ethereum.CallOpts)
}

// SetSecret is a paid mutator transaction binding the contract method 0x7ed6c926.
//
// Solidity: function setSecret(string _secret) returns()
func (_Ethereum *EthereumTransactor) SetSecret(opts *bind.TransactOpts, _secret string) (*types.Transaction, error) {
	return _Ethereum.contract.Transact(opts, "setSecret", _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x7ed6c926.
//
// Solidity: function setSecret(string _secret) returns()
func (_Ethereum *EthereumSession) SetSecret(_secret string) (*types.Transaction, error) {
	return _Ethereum.Contract.SetSecret(&_Ethereum.TransactOpts, _secret)
}

// SetSecret is a paid mutator transaction binding the contract method 0x7ed6c926.
//
// Solidity: function setSecret(string _secret) returns()
func (_Ethereum *EthereumTransactorSession) SetSecret(_secret string) (*types.Transaction, error) {
	return _Ethereum.Contract.SetSecret(&_Ethereum.TransactOpts, _secret)
}
