// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gpulease

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

// CampaignMetadataRendererMetaData contains all meta data concerning the CampaignMetadataRenderer contract.
var CampaignMetadataRendererMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"campaignName\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"grade\",\"type\":\"uint8\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// CampaignMetadataRendererABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignMetadataRendererMetaData.ABI instead.
var CampaignMetadataRendererABI = CampaignMetadataRendererMetaData.ABI

// CampaignMetadataRenderer is an auto generated Go binding around an Ethereum contract.
type CampaignMetadataRenderer struct {
	CampaignMetadataRendererCaller     // Read-only binding to the contract
	CampaignMetadataRendererTransactor // Write-only binding to the contract
	CampaignMetadataRendererFilterer   // Log filterer for contract events
}

// CampaignMetadataRendererCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignMetadataRendererCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignMetadataRendererTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignMetadataRendererTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignMetadataRendererFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignMetadataRendererFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignMetadataRendererSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignMetadataRendererSession struct {
	Contract     *CampaignMetadataRenderer // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// CampaignMetadataRendererCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignMetadataRendererCallerSession struct {
	Contract *CampaignMetadataRendererCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// CampaignMetadataRendererTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignMetadataRendererTransactorSession struct {
	Contract     *CampaignMetadataRendererTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// CampaignMetadataRendererRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignMetadataRendererRaw struct {
	Contract *CampaignMetadataRenderer // Generic contract binding to access the raw methods on
}

// CampaignMetadataRendererCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignMetadataRendererCallerRaw struct {
	Contract *CampaignMetadataRendererCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignMetadataRendererTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignMetadataRendererTransactorRaw struct {
	Contract *CampaignMetadataRendererTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaignMetadataRenderer creates a new instance of CampaignMetadataRenderer, bound to a specific deployed contract.
func NewCampaignMetadataRenderer(address common.Address, backend bind.ContractBackend) (*CampaignMetadataRenderer, error) {
	contract, err := bindCampaignMetadataRenderer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CampaignMetadataRenderer{CampaignMetadataRendererCaller: CampaignMetadataRendererCaller{contract: contract}, CampaignMetadataRendererTransactor: CampaignMetadataRendererTransactor{contract: contract}, CampaignMetadataRendererFilterer: CampaignMetadataRendererFilterer{contract: contract}}, nil
}

// NewCampaignMetadataRendererCaller creates a new read-only instance of CampaignMetadataRenderer, bound to a specific deployed contract.
func NewCampaignMetadataRendererCaller(address common.Address, caller bind.ContractCaller) (*CampaignMetadataRendererCaller, error) {
	contract, err := bindCampaignMetadataRenderer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignMetadataRendererCaller{contract: contract}, nil
}

// NewCampaignMetadataRendererTransactor creates a new write-only instance of CampaignMetadataRenderer, bound to a specific deployed contract.
func NewCampaignMetadataRendererTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignMetadataRendererTransactor, error) {
	contract, err := bindCampaignMetadataRenderer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignMetadataRendererTransactor{contract: contract}, nil
}

// NewCampaignMetadataRendererFilterer creates a new log filterer instance of CampaignMetadataRenderer, bound to a specific deployed contract.
func NewCampaignMetadataRendererFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignMetadataRendererFilterer, error) {
	contract, err := bindCampaignMetadataRenderer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignMetadataRendererFilterer{contract: contract}, nil
}

// bindCampaignMetadataRenderer binds a generic wrapper to an already deployed contract.
func bindCampaignMetadataRenderer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CampaignMetadataRendererMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignMetadataRenderer *CampaignMetadataRendererRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignMetadataRenderer.Contract.CampaignMetadataRendererCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignMetadataRenderer *CampaignMetadataRendererRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignMetadataRenderer.Contract.CampaignMetadataRendererTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignMetadataRenderer *CampaignMetadataRendererRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignMetadataRenderer.Contract.CampaignMetadataRendererTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignMetadataRenderer *CampaignMetadataRendererCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignMetadataRenderer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignMetadataRenderer *CampaignMetadataRendererTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignMetadataRenderer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignMetadataRenderer *CampaignMetadataRendererTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignMetadataRenderer.Contract.contract.Transact(opts, method, params...)
}

// TokenURI is a free data retrieval call binding the contract method 0xd5a893fc.
//
// Solidity: function tokenURI(string campaignName, uint256 campaignId, uint8 grade) pure returns(string)
func (_CampaignMetadataRenderer *CampaignMetadataRendererCaller) TokenURI(opts *bind.CallOpts, campaignName string, campaignId *big.Int, grade uint8) (string, error) {
	var out []interface{}
	err := _CampaignMetadataRenderer.contract.Call(opts, &out, "tokenURI", campaignName, campaignId, grade)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xd5a893fc.
//
// Solidity: function tokenURI(string campaignName, uint256 campaignId, uint8 grade) pure returns(string)
func (_CampaignMetadataRenderer *CampaignMetadataRendererSession) TokenURI(campaignName string, campaignId *big.Int, grade uint8) (string, error) {
	return _CampaignMetadataRenderer.Contract.TokenURI(&_CampaignMetadataRenderer.CallOpts, campaignName, campaignId, grade)
}

// TokenURI is a free data retrieval call binding the contract method 0xd5a893fc.
//
// Solidity: function tokenURI(string campaignName, uint256 campaignId, uint8 grade) pure returns(string)
func (_CampaignMetadataRenderer *CampaignMetadataRendererCallerSession) TokenURI(campaignName string, campaignId *big.Int, grade uint8) (string, error) {
	return _CampaignMetadataRenderer.Contract.TokenURI(&_CampaignMetadataRenderer.CallOpts, campaignName, campaignId, grade)
}
