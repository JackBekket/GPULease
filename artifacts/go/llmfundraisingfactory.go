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

// LLMFundraisingFactoryMetaData contains all meta data concerning the LLMFundraisingFactory contract.
var LLMFundraisingFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gpuLeaseWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metadataRenderer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FailedDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"campaign\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"templateId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaignName\",\"type\":\"string\"}],\"name\":\"CampaignCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"campaign\",\"type\":\"address\"}],\"name\":\"CampaignParticipantRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaignById\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"campaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"campaignsByCreator\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"campaignsByParticipant\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"templateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"campaignName\",\"type\":\"string\"}],\"name\":\"createCampaign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"campaignId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"campaign\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gpuLeaseWallet\",\"outputs\":[{\"internalType\":\"contractIGPULeaseWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"campaign\",\"type\":\"address\"}],\"name\":\"hasParticipatedInCampaign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCampaign\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataRenderer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextCampaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"participantCampaignAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"participantCampaignsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"}],\"name\":\"registerParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LLMFundraisingFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use LLMFundraisingFactoryMetaData.ABI instead.
var LLMFundraisingFactoryABI = LLMFundraisingFactoryMetaData.ABI

// LLMFundraisingFactory is an auto generated Go binding around an Ethereum contract.
type LLMFundraisingFactory struct {
	LLMFundraisingFactoryCaller     // Read-only binding to the contract
	LLMFundraisingFactoryTransactor // Write-only binding to the contract
	LLMFundraisingFactoryFilterer   // Log filterer for contract events
}

// LLMFundraisingFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LLMFundraisingFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LLMFundraisingFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LLMFundraisingFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LLMFundraisingFactorySession struct {
	Contract     *LLMFundraisingFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// LLMFundraisingFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LLMFundraisingFactoryCallerSession struct {
	Contract *LLMFundraisingFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// LLMFundraisingFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LLMFundraisingFactoryTransactorSession struct {
	Contract     *LLMFundraisingFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// LLMFundraisingFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LLMFundraisingFactoryRaw struct {
	Contract *LLMFundraisingFactory // Generic contract binding to access the raw methods on
}

// LLMFundraisingFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LLMFundraisingFactoryCallerRaw struct {
	Contract *LLMFundraisingFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LLMFundraisingFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LLMFundraisingFactoryTransactorRaw struct {
	Contract *LLMFundraisingFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLLMFundraisingFactory creates a new instance of LLMFundraisingFactory, bound to a specific deployed contract.
func NewLLMFundraisingFactory(address common.Address, backend bind.ContractBackend) (*LLMFundraisingFactory, error) {
	contract, err := bindLLMFundraisingFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactory{LLMFundraisingFactoryCaller: LLMFundraisingFactoryCaller{contract: contract}, LLMFundraisingFactoryTransactor: LLMFundraisingFactoryTransactor{contract: contract}, LLMFundraisingFactoryFilterer: LLMFundraisingFactoryFilterer{contract: contract}}, nil
}

// NewLLMFundraisingFactoryCaller creates a new read-only instance of LLMFundraisingFactory, bound to a specific deployed contract.
func NewLLMFundraisingFactoryCaller(address common.Address, caller bind.ContractCaller) (*LLMFundraisingFactoryCaller, error) {
	contract, err := bindLLMFundraisingFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryCaller{contract: contract}, nil
}

// NewLLMFundraisingFactoryTransactor creates a new write-only instance of LLMFundraisingFactory, bound to a specific deployed contract.
func NewLLMFundraisingFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LLMFundraisingFactoryTransactor, error) {
	contract, err := bindLLMFundraisingFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryTransactor{contract: contract}, nil
}

// NewLLMFundraisingFactoryFilterer creates a new log filterer instance of LLMFundraisingFactory, bound to a specific deployed contract.
func NewLLMFundraisingFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LLMFundraisingFactoryFilterer, error) {
	contract, err := bindLLMFundraisingFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryFilterer{contract: contract}, nil
}

// bindLLMFundraisingFactory binds a generic wrapper to an already deployed contract.
func bindLLMFundraisingFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LLMFundraisingFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMFundraisingFactory *LLMFundraisingFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMFundraisingFactory.Contract.LLMFundraisingFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMFundraisingFactory *LLMFundraisingFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.LLMFundraisingFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMFundraisingFactory *LLMFundraisingFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.LLMFundraisingFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMFundraisingFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.contract.Transact(opts, method, params...)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) CampaignById(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaignById", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CampaignById(arg0 *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignById(&_LLMFundraisingFactory.CallOpts, arg0)
}

// CampaignById is a free data retrieval call binding the contract method 0x98bf6de8.
//
// Solidity: function campaignById(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) CampaignById(arg0 *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignById(&_LLMFundraisingFactory.CallOpts, arg0)
}

// CampaignImplementation is a free data retrieval call binding the contract method 0x62a6e065.
//
// Solidity: function campaignImplementation() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) CampaignImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaignImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CampaignImplementation is a free data retrieval call binding the contract method 0x62a6e065.
//
// Solidity: function campaignImplementation() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CampaignImplementation() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignImplementation(&_LLMFundraisingFactory.CallOpts)
}

// CampaignImplementation is a free data retrieval call binding the contract method 0x62a6e065.
//
// Solidity: function campaignImplementation() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) CampaignImplementation() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignImplementation(&_LLMFundraisingFactory.CallOpts)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) Campaigns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaigns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) Campaigns(arg0 *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Campaigns(&_LLMFundraisingFactory.CallOpts, arg0)
}

// Campaigns is a free data retrieval call binding the contract method 0x141961bc.
//
// Solidity: function campaigns(uint256 ) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) Campaigns(arg0 *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Campaigns(&_LLMFundraisingFactory.CallOpts, arg0)
}

// CampaignsByCreator is a free data retrieval call binding the contract method 0x109b1611.
//
// Solidity: function campaignsByCreator(address creator) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) CampaignsByCreator(opts *bind.CallOpts, creator common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaignsByCreator", creator)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CampaignsByCreator is a free data retrieval call binding the contract method 0x109b1611.
//
// Solidity: function campaignsByCreator(address creator) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CampaignsByCreator(creator common.Address) ([]common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignsByCreator(&_LLMFundraisingFactory.CallOpts, creator)
}

// CampaignsByCreator is a free data retrieval call binding the contract method 0x109b1611.
//
// Solidity: function campaignsByCreator(address creator) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) CampaignsByCreator(creator common.Address) ([]common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignsByCreator(&_LLMFundraisingFactory.CallOpts, creator)
}

// CampaignsByParticipant is a free data retrieval call binding the contract method 0xcc63bac9.
//
// Solidity: function campaignsByParticipant(address participant) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) CampaignsByParticipant(opts *bind.CallOpts, participant common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaignsByParticipant", participant)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// CampaignsByParticipant is a free data retrieval call binding the contract method 0xcc63bac9.
//
// Solidity: function campaignsByParticipant(address participant) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CampaignsByParticipant(participant common.Address) ([]common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignsByParticipant(&_LLMFundraisingFactory.CallOpts, participant)
}

// CampaignsByParticipant is a free data retrieval call binding the contract method 0xcc63bac9.
//
// Solidity: function campaignsByParticipant(address participant) view returns(address[])
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) CampaignsByParticipant(participant common.Address) ([]common.Address, error) {
	return _LLMFundraisingFactory.Contract.CampaignsByParticipant(&_LLMFundraisingFactory.CallOpts, participant)
}

// CampaignsCount is a free data retrieval call binding the contract method 0x44c89b65.
//
// Solidity: function campaignsCount() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) CampaignsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "campaignsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignsCount is a free data retrieval call binding the contract method 0x44c89b65.
//
// Solidity: function campaignsCount() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CampaignsCount() (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.CampaignsCount(&_LLMFundraisingFactory.CallOpts)
}

// CampaignsCount is a free data retrieval call binding the contract method 0x44c89b65.
//
// Solidity: function campaignsCount() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) CampaignsCount() (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.CampaignsCount(&_LLMFundraisingFactory.CallOpts)
}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) GpuLeaseWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "gpuLeaseWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) GpuLeaseWallet() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.GpuLeaseWallet(&_LLMFundraisingFactory.CallOpts)
}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) GpuLeaseWallet() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.GpuLeaseWallet(&_LLMFundraisingFactory.CallOpts)
}

// HasParticipatedInCampaign is a free data retrieval call binding the contract method 0x9b287abb.
//
// Solidity: function hasParticipatedInCampaign(address participant, address campaign) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) HasParticipatedInCampaign(opts *bind.CallOpts, participant common.Address, campaign common.Address) (bool, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "hasParticipatedInCampaign", participant, campaign)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasParticipatedInCampaign is a free data retrieval call binding the contract method 0x9b287abb.
//
// Solidity: function hasParticipatedInCampaign(address participant, address campaign) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) HasParticipatedInCampaign(participant common.Address, campaign common.Address) (bool, error) {
	return _LLMFundraisingFactory.Contract.HasParticipatedInCampaign(&_LLMFundraisingFactory.CallOpts, participant, campaign)
}

// HasParticipatedInCampaign is a free data retrieval call binding the contract method 0x9b287abb.
//
// Solidity: function hasParticipatedInCampaign(address participant, address campaign) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) HasParticipatedInCampaign(participant common.Address, campaign common.Address) (bool, error) {
	return _LLMFundraisingFactory.Contract.HasParticipatedInCampaign(&_LLMFundraisingFactory.CallOpts, participant, campaign)
}

// IsCampaign is a free data retrieval call binding the contract method 0xe2714e1b.
//
// Solidity: function isCampaign(address ) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) IsCampaign(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "isCampaign", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCampaign is a free data retrieval call binding the contract method 0xe2714e1b.
//
// Solidity: function isCampaign(address ) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) IsCampaign(arg0 common.Address) (bool, error) {
	return _LLMFundraisingFactory.Contract.IsCampaign(&_LLMFundraisingFactory.CallOpts, arg0)
}

// IsCampaign is a free data retrieval call binding the contract method 0xe2714e1b.
//
// Solidity: function isCampaign(address ) view returns(bool)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) IsCampaign(arg0 common.Address) (bool, error) {
	return _LLMFundraisingFactory.Contract.IsCampaign(&_LLMFundraisingFactory.CallOpts, arg0)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) MetadataRenderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "metadataRenderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) MetadataRenderer() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.MetadataRenderer(&_LLMFundraisingFactory.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) MetadataRenderer() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.MetadataRenderer(&_LLMFundraisingFactory.CallOpts)
}

// NextCampaignId is a free data retrieval call binding the contract method 0x7903a756.
//
// Solidity: function nextCampaignId() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) NextCampaignId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "nextCampaignId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextCampaignId is a free data retrieval call binding the contract method 0x7903a756.
//
// Solidity: function nextCampaignId() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) NextCampaignId() (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.NextCampaignId(&_LLMFundraisingFactory.CallOpts)
}

// NextCampaignId is a free data retrieval call binding the contract method 0x7903a756.
//
// Solidity: function nextCampaignId() view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) NextCampaignId() (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.NextCampaignId(&_LLMFundraisingFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) Owner() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Owner(&_LLMFundraisingFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) Owner() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Owner(&_LLMFundraisingFactory.CallOpts)
}

// ParticipantCampaignAt is a free data retrieval call binding the contract method 0x36290f21.
//
// Solidity: function participantCampaignAt(address participant, uint256 index) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) ParticipantCampaignAt(opts *bind.CallOpts, participant common.Address, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "participantCampaignAt", participant, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantCampaignAt is a free data retrieval call binding the contract method 0x36290f21.
//
// Solidity: function participantCampaignAt(address participant, uint256 index) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) ParticipantCampaignAt(participant common.Address, index *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.ParticipantCampaignAt(&_LLMFundraisingFactory.CallOpts, participant, index)
}

// ParticipantCampaignAt is a free data retrieval call binding the contract method 0x36290f21.
//
// Solidity: function participantCampaignAt(address participant, uint256 index) view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) ParticipantCampaignAt(participant common.Address, index *big.Int) (common.Address, error) {
	return _LLMFundraisingFactory.Contract.ParticipantCampaignAt(&_LLMFundraisingFactory.CallOpts, participant, index)
}

// ParticipantCampaignsCount is a free data retrieval call binding the contract method 0xe5371aa3.
//
// Solidity: function participantCampaignsCount(address participant) view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) ParticipantCampaignsCount(opts *bind.CallOpts, participant common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "participantCampaignsCount", participant)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ParticipantCampaignsCount is a free data retrieval call binding the contract method 0xe5371aa3.
//
// Solidity: function participantCampaignsCount(address participant) view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) ParticipantCampaignsCount(participant common.Address) (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.ParticipantCampaignsCount(&_LLMFundraisingFactory.CallOpts, participant)
}

// ParticipantCampaignsCount is a free data retrieval call binding the contract method 0xe5371aa3.
//
// Solidity: function participantCampaignsCount(address participant) view returns(uint256)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) ParticipantCampaignsCount(participant common.Address) (*big.Int, error) {
	return _LLMFundraisingFactory.Contract.ParticipantCampaignsCount(&_LLMFundraisingFactory.CallOpts, participant)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraisingFactory.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) Usdc() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Usdc(&_LLMFundraisingFactory.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraisingFactory *LLMFundraisingFactoryCallerSession) Usdc() (common.Address, error) {
	return _LLMFundraisingFactory.Contract.Usdc(&_LLMFundraisingFactory.CallOpts)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x2761ef42.
//
// Solidity: function createCampaign(uint256 targetAmount, uint256 duration, uint256 startTimestamp, uint256 templateId, string campaignName) returns(uint256 campaignId, address campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactor) CreateCampaign(opts *bind.TransactOpts, targetAmount *big.Int, duration *big.Int, startTimestamp *big.Int, templateId *big.Int, campaignName string) (*types.Transaction, error) {
	return _LLMFundraisingFactory.contract.Transact(opts, "createCampaign", targetAmount, duration, startTimestamp, templateId, campaignName)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x2761ef42.
//
// Solidity: function createCampaign(uint256 targetAmount, uint256 duration, uint256 startTimestamp, uint256 templateId, string campaignName) returns(uint256 campaignId, address campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) CreateCampaign(targetAmount *big.Int, duration *big.Int, startTimestamp *big.Int, templateId *big.Int, campaignName string) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.CreateCampaign(&_LLMFundraisingFactory.TransactOpts, targetAmount, duration, startTimestamp, templateId, campaignName)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0x2761ef42.
//
// Solidity: function createCampaign(uint256 targetAmount, uint256 duration, uint256 startTimestamp, uint256 templateId, string campaignName) returns(uint256 campaignId, address campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorSession) CreateCampaign(targetAmount *big.Int, duration *big.Int, startTimestamp *big.Int, templateId *big.Int, campaignName string) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.CreateCampaign(&_LLMFundraisingFactory.TransactOpts, targetAmount, duration, startTimestamp, templateId, campaignName)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x7cc9e4f0.
//
// Solidity: function registerParticipant(address participant) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactor) RegisterParticipant(opts *bind.TransactOpts, participant common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.contract.Transact(opts, "registerParticipant", participant)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x7cc9e4f0.
//
// Solidity: function registerParticipant(address participant) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) RegisterParticipant(participant common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.RegisterParticipant(&_LLMFundraisingFactory.TransactOpts, participant)
}

// RegisterParticipant is a paid mutator transaction binding the contract method 0x7cc9e4f0.
//
// Solidity: function registerParticipant(address participant) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorSession) RegisterParticipant(participant common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.RegisterParticipant(&_LLMFundraisingFactory.TransactOpts, participant)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraisingFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.RenounceOwnership(&_LLMFundraisingFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.RenounceOwnership(&_LLMFundraisingFactory.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.TransferOwnership(&_LLMFundraisingFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraisingFactory *LLMFundraisingFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraisingFactory.Contract.TransferOwnership(&_LLMFundraisingFactory.TransactOpts, newOwner)
}

// LLMFundraisingFactoryCampaignCreatedIterator is returned from FilterCampaignCreated and is used to iterate over the raw logs and unpacked data for CampaignCreated events raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryCampaignCreatedIterator struct {
	Event *LLMFundraisingFactoryCampaignCreated // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingFactoryCampaignCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingFactoryCampaignCreated)
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
		it.Event = new(LLMFundraisingFactoryCampaignCreated)
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
func (it *LLMFundraisingFactoryCampaignCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingFactoryCampaignCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingFactoryCampaignCreated represents a CampaignCreated event raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryCampaignCreated struct {
	CampaignId     *big.Int
	Campaign       common.Address
	Creator        common.Address
	TargetAmount   *big.Int
	StartTimestamp *big.Int
	Duration       *big.Int
	TemplateId     *big.Int
	CampaignName   string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCampaignCreated is a free log retrieval operation binding the contract event 0x586eb302feca65b7cd23eca1855aeac790d55a9bbd4e42168873521f5d5bca47.
//
// Solidity: event CampaignCreated(uint256 indexed campaignId, address indexed campaign, address indexed creator, uint256 targetAmount, uint256 startTimestamp, uint256 duration, uint256 templateId, string campaignName)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) FilterCampaignCreated(opts *bind.FilterOpts, campaignId []*big.Int, campaign []common.Address, creator []common.Address) (*LLMFundraisingFactoryCampaignCreatedIterator, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var campaignRule []interface{}
	for _, campaignItem := range campaign {
		campaignRule = append(campaignRule, campaignItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.FilterLogs(opts, "CampaignCreated", campaignIdRule, campaignRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryCampaignCreatedIterator{contract: _LLMFundraisingFactory.contract, event: "CampaignCreated", logs: logs, sub: sub}, nil
}

// WatchCampaignCreated is a free log subscription operation binding the contract event 0x586eb302feca65b7cd23eca1855aeac790d55a9bbd4e42168873521f5d5bca47.
//
// Solidity: event CampaignCreated(uint256 indexed campaignId, address indexed campaign, address indexed creator, uint256 targetAmount, uint256 startTimestamp, uint256 duration, uint256 templateId, string campaignName)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) WatchCampaignCreated(opts *bind.WatchOpts, sink chan<- *LLMFundraisingFactoryCampaignCreated, campaignId []*big.Int, campaign []common.Address, creator []common.Address) (event.Subscription, error) {

	var campaignIdRule []interface{}
	for _, campaignIdItem := range campaignId {
		campaignIdRule = append(campaignIdRule, campaignIdItem)
	}
	var campaignRule []interface{}
	for _, campaignItem := range campaign {
		campaignRule = append(campaignRule, campaignItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.WatchLogs(opts, "CampaignCreated", campaignIdRule, campaignRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingFactoryCampaignCreated)
				if err := _LLMFundraisingFactory.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
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

// ParseCampaignCreated is a log parse operation binding the contract event 0x586eb302feca65b7cd23eca1855aeac790d55a9bbd4e42168873521f5d5bca47.
//
// Solidity: event CampaignCreated(uint256 indexed campaignId, address indexed campaign, address indexed creator, uint256 targetAmount, uint256 startTimestamp, uint256 duration, uint256 templateId, string campaignName)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) ParseCampaignCreated(log types.Log) (*LLMFundraisingFactoryCampaignCreated, error) {
	event := new(LLMFundraisingFactoryCampaignCreated)
	if err := _LLMFundraisingFactory.contract.UnpackLog(event, "CampaignCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingFactoryCampaignParticipantRegisteredIterator is returned from FilterCampaignParticipantRegistered and is used to iterate over the raw logs and unpacked data for CampaignParticipantRegistered events raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryCampaignParticipantRegisteredIterator struct {
	Event *LLMFundraisingFactoryCampaignParticipantRegistered // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingFactoryCampaignParticipantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingFactoryCampaignParticipantRegistered)
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
		it.Event = new(LLMFundraisingFactoryCampaignParticipantRegistered)
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
func (it *LLMFundraisingFactoryCampaignParticipantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingFactoryCampaignParticipantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingFactoryCampaignParticipantRegistered represents a CampaignParticipantRegistered event raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryCampaignParticipantRegistered struct {
	Participant common.Address
	Campaign    common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCampaignParticipantRegistered is a free log retrieval operation binding the contract event 0x8300376baa4bb1123d5bd8ac536d00577f4b5ae198e7a07b213634f837c25cb5.
//
// Solidity: event CampaignParticipantRegistered(address indexed participant, address indexed campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) FilterCampaignParticipantRegistered(opts *bind.FilterOpts, participant []common.Address, campaign []common.Address) (*LLMFundraisingFactoryCampaignParticipantRegisteredIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var campaignRule []interface{}
	for _, campaignItem := range campaign {
		campaignRule = append(campaignRule, campaignItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.FilterLogs(opts, "CampaignParticipantRegistered", participantRule, campaignRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryCampaignParticipantRegisteredIterator{contract: _LLMFundraisingFactory.contract, event: "CampaignParticipantRegistered", logs: logs, sub: sub}, nil
}

// WatchCampaignParticipantRegistered is a free log subscription operation binding the contract event 0x8300376baa4bb1123d5bd8ac536d00577f4b5ae198e7a07b213634f837c25cb5.
//
// Solidity: event CampaignParticipantRegistered(address indexed participant, address indexed campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) WatchCampaignParticipantRegistered(opts *bind.WatchOpts, sink chan<- *LLMFundraisingFactoryCampaignParticipantRegistered, participant []common.Address, campaign []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}
	var campaignRule []interface{}
	for _, campaignItem := range campaign {
		campaignRule = append(campaignRule, campaignItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.WatchLogs(opts, "CampaignParticipantRegistered", participantRule, campaignRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingFactoryCampaignParticipantRegistered)
				if err := _LLMFundraisingFactory.contract.UnpackLog(event, "CampaignParticipantRegistered", log); err != nil {
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

// ParseCampaignParticipantRegistered is a log parse operation binding the contract event 0x8300376baa4bb1123d5bd8ac536d00577f4b5ae198e7a07b213634f837c25cb5.
//
// Solidity: event CampaignParticipantRegistered(address indexed participant, address indexed campaign)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) ParseCampaignParticipantRegistered(log types.Log) (*LLMFundraisingFactoryCampaignParticipantRegistered, error) {
	event := new(LLMFundraisingFactoryCampaignParticipantRegistered)
	if err := _LLMFundraisingFactory.contract.UnpackLog(event, "CampaignParticipantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryOwnershipTransferredIterator struct {
	Event *LLMFundraisingFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingFactoryOwnershipTransferred)
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
		it.Event = new(LLMFundraisingFactoryOwnershipTransferred)
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
func (it *LLMFundraisingFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the LLMFundraisingFactory contract.
type LLMFundraisingFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LLMFundraisingFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFactoryOwnershipTransferredIterator{contract: _LLMFundraisingFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LLMFundraisingFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LLMFundraisingFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingFactoryOwnershipTransferred)
				if err := _LLMFundraisingFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LLMFundraisingFactory *LLMFundraisingFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*LLMFundraisingFactoryOwnershipTransferred, error) {
	event := new(LLMFundraisingFactoryOwnershipTransferred)
	if err := _LLMFundraisingFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
