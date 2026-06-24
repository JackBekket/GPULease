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

// GPULeaseFrozenFundsInfo is an auto generated low-level Go binding around an user-defined struct.
type GPULeaseFrozenFundsInfo struct {
	LeaseId *big.Int
	Amount  *big.Int
}

// GPULeaseMetaData contains all meta data concerning the GPULease contract.
var GPULeaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wallet_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"treasury_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeaseCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeasePaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"name\":\"LeaseResumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LeaseStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFeePercentage\",\"type\":\"uint256\"}],\"name\":\"PlatformFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"ReferralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserFeeCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feePercentage\",\"type\":\"uint256\"}],\"name\":\"UserFeeUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"clearUserFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"completeLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"feePercentageForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"frozenFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserFrozenFunds\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structGPULease.FrozenFundsInfo[]\",\"name\":\"result\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasCustomUserFee\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaseCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leaseReferralInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralShareBps\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leases\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"storagePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"computePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leaseFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"completed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"pausedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pausedDuration\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"pauseLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"platformFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referralManager\",\"outputs\":[{\"internalType\":\"contractIGPULeaseReferral\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_leaseId\",\"type\":\"uint256\"}],\"name\":\"resumeLease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"setPlatformFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReferralManager\",\"type\":\"address\"}],\"name\":\"setReferralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTreasury\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feePercentage\",\"type\":\"uint256\"}],\"name\":\"setUserFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_storagePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_computePricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_provider\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"startLease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"leaseId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userActiveLeases\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"contractGPULeaseWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GPULeaseABI is the input ABI used to generate the binding from.
// Deprecated: Use GPULeaseMetaData.ABI instead.
var GPULeaseABI = GPULeaseMetaData.ABI

// GPULease is an auto generated Go binding around an Ethereum contract.
type GPULease struct {
	GPULeaseCaller     // Read-only binding to the contract
	GPULeaseTransactor // Write-only binding to the contract
	GPULeaseFilterer   // Log filterer for contract events
}

// GPULeaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type GPULeaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GPULeaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GPULeaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GPULeaseSession struct {
	Contract     *GPULease         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GPULeaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GPULeaseCallerSession struct {
	Contract *GPULeaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GPULeaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GPULeaseTransactorSession struct {
	Contract     *GPULeaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GPULeaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type GPULeaseRaw struct {
	Contract *GPULease // Generic contract binding to access the raw methods on
}

// GPULeaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GPULeaseCallerRaw struct {
	Contract *GPULeaseCaller // Generic read-only contract binding to access the raw methods on
}

// GPULeaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GPULeaseTransactorRaw struct {
	Contract *GPULeaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGPULease creates a new instance of GPULease, bound to a specific deployed contract.
func NewGPULease(address common.Address, backend bind.ContractBackend) (*GPULease, error) {
	contract, err := bindGPULease(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GPULease{GPULeaseCaller: GPULeaseCaller{contract: contract}, GPULeaseTransactor: GPULeaseTransactor{contract: contract}, GPULeaseFilterer: GPULeaseFilterer{contract: contract}}, nil
}

// NewGPULeaseCaller creates a new read-only instance of GPULease, bound to a specific deployed contract.
func NewGPULeaseCaller(address common.Address, caller bind.ContractCaller) (*GPULeaseCaller, error) {
	contract, err := bindGPULease(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseCaller{contract: contract}, nil
}

// NewGPULeaseTransactor creates a new write-only instance of GPULease, bound to a specific deployed contract.
func NewGPULeaseTransactor(address common.Address, transactor bind.ContractTransactor) (*GPULeaseTransactor, error) {
	contract, err := bindGPULease(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseTransactor{contract: contract}, nil
}

// NewGPULeaseFilterer creates a new log filterer instance of GPULease, bound to a specific deployed contract.
func NewGPULeaseFilterer(address common.Address, filterer bind.ContractFilterer) (*GPULeaseFilterer, error) {
	contract, err := bindGPULease(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GPULeaseFilterer{contract: contract}, nil
}

// bindGPULease binds a generic wrapper to an already deployed contract.
func bindGPULease(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GPULeaseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULease *GPULeaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULease.Contract.GPULeaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULease *GPULeaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULease.Contract.GPULeaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULease *GPULeaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULease.Contract.GPULeaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULease *GPULeaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULease.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULease *GPULeaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULease.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULease *GPULeaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULease.Contract.contract.Transact(opts, method, params...)
}

// FeePercentageForUser is a free data retrieval call binding the contract method 0xa8f4214e.
//
// Solidity: function feePercentageForUser(address user) view returns(uint256)
func (_GPULease *GPULeaseCaller) FeePercentageForUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "feePercentageForUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentageForUser is a free data retrieval call binding the contract method 0xa8f4214e.
//
// Solidity: function feePercentageForUser(address user) view returns(uint256)
func (_GPULease *GPULeaseSession) FeePercentageForUser(user common.Address) (*big.Int, error) {
	return _GPULease.Contract.FeePercentageForUser(&_GPULease.CallOpts, user)
}

// FeePercentageForUser is a free data retrieval call binding the contract method 0xa8f4214e.
//
// Solidity: function feePercentageForUser(address user) view returns(uint256)
func (_GPULease *GPULeaseCallerSession) FeePercentageForUser(user common.Address) (*big.Int, error) {
	return _GPULease.Contract.FeePercentageForUser(&_GPULease.CallOpts, user)
}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_GPULease *GPULeaseCaller) FrozenFunds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "frozenFunds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_GPULease *GPULeaseSession) FrozenFunds(arg0 *big.Int) (*big.Int, error) {
	return _GPULease.Contract.FrozenFunds(&_GPULease.CallOpts, arg0)
}

// FrozenFunds is a free data retrieval call binding the contract method 0xa641992a.
//
// Solidity: function frozenFunds(uint256 ) view returns(uint256)
func (_GPULease *GPULeaseCallerSession) FrozenFunds(arg0 *big.Int) (*big.Int, error) {
	return _GPULease.Contract.FrozenFunds(&_GPULease.CallOpts, arg0)
}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_GPULease *GPULeaseCaller) GetUserFrozenFunds(opts *bind.CallOpts, user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "getUserFrozenFunds", user)

	if err != nil {
		return *new([]GPULeaseFrozenFundsInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]GPULeaseFrozenFundsInfo)).(*[]GPULeaseFrozenFundsInfo)

	return out0, err

}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_GPULease *GPULeaseSession) GetUserFrozenFunds(user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	return _GPULease.Contract.GetUserFrozenFunds(&_GPULease.CallOpts, user)
}

// GetUserFrozenFunds is a free data retrieval call binding the contract method 0x97c4d284.
//
// Solidity: function getUserFrozenFunds(address user) view returns((uint256,uint256)[] result)
func (_GPULease *GPULeaseCallerSession) GetUserFrozenFunds(user common.Address) ([]GPULeaseFrozenFundsInfo, error) {
	return _GPULease.Contract.GetUserFrozenFunds(&_GPULease.CallOpts, user)
}

// HasCustomUserFee is a free data retrieval call binding the contract method 0xe67d3732.
//
// Solidity: function hasCustomUserFee(address ) view returns(bool)
func (_GPULease *GPULeaseCaller) HasCustomUserFee(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "hasCustomUserFee", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasCustomUserFee is a free data retrieval call binding the contract method 0xe67d3732.
//
// Solidity: function hasCustomUserFee(address ) view returns(bool)
func (_GPULease *GPULeaseSession) HasCustomUserFee(arg0 common.Address) (bool, error) {
	return _GPULease.Contract.HasCustomUserFee(&_GPULease.CallOpts, arg0)
}

// HasCustomUserFee is a free data retrieval call binding the contract method 0xe67d3732.
//
// Solidity: function hasCustomUserFee(address ) view returns(bool)
func (_GPULease *GPULeaseCallerSession) HasCustomUserFee(arg0 common.Address) (bool, error) {
	return _GPULease.Contract.HasCustomUserFee(&_GPULease.CallOpts, arg0)
}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_GPULease *GPULeaseCaller) LeaseCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "leaseCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_GPULease *GPULeaseSession) LeaseCount() (*big.Int, error) {
	return _GPULease.Contract.LeaseCount(&_GPULease.CallOpts)
}

// LeaseCount is a free data retrieval call binding the contract method 0xb4c0498b.
//
// Solidity: function leaseCount() view returns(uint256)
func (_GPULease *GPULeaseCallerSession) LeaseCount() (*big.Int, error) {
	return _GPULease.Contract.LeaseCount(&_GPULease.CallOpts)
}

// LeaseReferralInfo is a free data retrieval call binding the contract method 0x4f9479a2.
//
// Solidity: function leaseReferralInfo(uint256 ) view returns(address referrer, uint256 referralShareBps)
func (_GPULease *GPULeaseCaller) LeaseReferralInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Referrer         common.Address
	ReferralShareBps *big.Int
}, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "leaseReferralInfo", arg0)

	outstruct := new(struct {
		Referrer         common.Address
		ReferralShareBps *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Referrer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ReferralShareBps = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LeaseReferralInfo is a free data retrieval call binding the contract method 0x4f9479a2.
//
// Solidity: function leaseReferralInfo(uint256 ) view returns(address referrer, uint256 referralShareBps)
func (_GPULease *GPULeaseSession) LeaseReferralInfo(arg0 *big.Int) (struct {
	Referrer         common.Address
	ReferralShareBps *big.Int
}, error) {
	return _GPULease.Contract.LeaseReferralInfo(&_GPULease.CallOpts, arg0)
}

// LeaseReferralInfo is a free data retrieval call binding the contract method 0x4f9479a2.
//
// Solidity: function leaseReferralInfo(uint256 ) view returns(address referrer, uint256 referralShareBps)
func (_GPULease *GPULeaseCallerSession) LeaseReferralInfo(arg0 *big.Int) (struct {
	Referrer         common.Address
	ReferralShareBps *big.Int
}, error) {
	return _GPULease.Contract.LeaseReferralInfo(&_GPULease.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 duration, uint256 storagePricePerSecond, uint256 computePricePerSecond, uint256 leaseFeePercentage, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_GPULease *GPULeaseCaller) Leases(opts *bind.CallOpts, arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	Duration              *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	LeaseFeePercentage    *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "leases", arg0)

	outstruct := new(struct {
		User                  common.Address
		Provider              common.Address
		StartTime             *big.Int
		Duration              *big.Int
		StoragePricePerSecond *big.Int
		ComputePricePerSecond *big.Int
		LeaseFeePercentage    *big.Int
		Active                bool
		Completed             bool
		Paused                bool
		PausedAt              *big.Int
		PausedDuration        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Provider = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.StartTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StoragePricePerSecond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ComputePricePerSecond = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LeaseFeePercentage = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Completed = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.Paused = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.PausedAt = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.PausedDuration = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 duration, uint256 storagePricePerSecond, uint256 computePricePerSecond, uint256 leaseFeePercentage, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_GPULease *GPULeaseSession) Leases(arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	Duration              *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	LeaseFeePercentage    *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	return _GPULease.Contract.Leases(&_GPULease.CallOpts, arg0)
}

// Leases is a free data retrieval call binding the contract method 0x8927a106.
//
// Solidity: function leases(uint256 ) view returns(address user, address provider, uint256 startTime, uint256 duration, uint256 storagePricePerSecond, uint256 computePricePerSecond, uint256 leaseFeePercentage, bool active, bool completed, bool paused, uint256 pausedAt, uint256 pausedDuration)
func (_GPULease *GPULeaseCallerSession) Leases(arg0 *big.Int) (struct {
	User                  common.Address
	Provider              common.Address
	StartTime             *big.Int
	Duration              *big.Int
	StoragePricePerSecond *big.Int
	ComputePricePerSecond *big.Int
	LeaseFeePercentage    *big.Int
	Active                bool
	Completed             bool
	Paused                bool
	PausedAt              *big.Int
	PausedDuration        *big.Int
}, error) {
	return _GPULease.Contract.Leases(&_GPULease.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULease *GPULeaseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULease *GPULeaseSession) Owner() (common.Address, error) {
	return _GPULease.Contract.Owner(&_GPULease.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULease *GPULeaseCallerSession) Owner() (common.Address, error) {
	return _GPULease.Contract.Owner(&_GPULease.CallOpts)
}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_GPULease *GPULeaseCaller) PlatformFeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "platformFeePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_GPULease *GPULeaseSession) PlatformFeePercentage() (*big.Int, error) {
	return _GPULease.Contract.PlatformFeePercentage(&_GPULease.CallOpts)
}

// PlatformFeePercentage is a free data retrieval call binding the contract method 0xcdd78cfc.
//
// Solidity: function platformFeePercentage() view returns(uint256)
func (_GPULease *GPULeaseCallerSession) PlatformFeePercentage() (*big.Int, error) {
	return _GPULease.Contract.PlatformFeePercentage(&_GPULease.CallOpts)
}

// ReferralManager is a free data retrieval call binding the contract method 0x2ccd9623.
//
// Solidity: function referralManager() view returns(address)
func (_GPULease *GPULeaseCaller) ReferralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "referralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferralManager is a free data retrieval call binding the contract method 0x2ccd9623.
//
// Solidity: function referralManager() view returns(address)
func (_GPULease *GPULeaseSession) ReferralManager() (common.Address, error) {
	return _GPULease.Contract.ReferralManager(&_GPULease.CallOpts)
}

// ReferralManager is a free data retrieval call binding the contract method 0x2ccd9623.
//
// Solidity: function referralManager() view returns(address)
func (_GPULease *GPULeaseCallerSession) ReferralManager() (common.Address, error) {
	return _GPULease.Contract.ReferralManager(&_GPULease.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_GPULease *GPULeaseCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_GPULease *GPULeaseSession) Treasury() (common.Address, error) {
	return _GPULease.Contract.Treasury(&_GPULease.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_GPULease *GPULeaseCallerSession) Treasury() (common.Address, error) {
	return _GPULease.Contract.Treasury(&_GPULease.CallOpts)
}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_GPULease *GPULeaseCaller) UserActiveLeases(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "userActiveLeases", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_GPULease *GPULeaseSession) UserActiveLeases(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _GPULease.Contract.UserActiveLeases(&_GPULease.CallOpts, arg0, arg1)
}

// UserActiveLeases is a free data retrieval call binding the contract method 0xe869a060.
//
// Solidity: function userActiveLeases(address , uint256 ) view returns(uint256)
func (_GPULease *GPULeaseCallerSession) UserActiveLeases(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _GPULease.Contract.UserActiveLeases(&_GPULease.CallOpts, arg0, arg1)
}

// UserFeePercentage is a free data retrieval call binding the contract method 0xd36824b0.
//
// Solidity: function userFeePercentage(address ) view returns(uint256)
func (_GPULease *GPULeaseCaller) UserFeePercentage(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "userFeePercentage", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserFeePercentage is a free data retrieval call binding the contract method 0xd36824b0.
//
// Solidity: function userFeePercentage(address ) view returns(uint256)
func (_GPULease *GPULeaseSession) UserFeePercentage(arg0 common.Address) (*big.Int, error) {
	return _GPULease.Contract.UserFeePercentage(&_GPULease.CallOpts, arg0)
}

// UserFeePercentage is a free data retrieval call binding the contract method 0xd36824b0.
//
// Solidity: function userFeePercentage(address ) view returns(uint256)
func (_GPULease *GPULeaseCallerSession) UserFeePercentage(arg0 common.Address) (*big.Int, error) {
	return _GPULease.Contract.UserFeePercentage(&_GPULease.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_GPULease *GPULeaseCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULease.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_GPULease *GPULeaseSession) Wallet() (common.Address, error) {
	return _GPULease.Contract.Wallet(&_GPULease.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_GPULease *GPULeaseCallerSession) Wallet() (common.Address, error) {
	return _GPULease.Contract.Wallet(&_GPULease.CallOpts)
}

// ClearUserFee is a paid mutator transaction binding the contract method 0xfbea0c1a.
//
// Solidity: function clearUserFee(address user) returns()
func (_GPULease *GPULeaseTransactor) ClearUserFee(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "clearUserFee", user)
}

// ClearUserFee is a paid mutator transaction binding the contract method 0xfbea0c1a.
//
// Solidity: function clearUserFee(address user) returns()
func (_GPULease *GPULeaseSession) ClearUserFee(user common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.ClearUserFee(&_GPULease.TransactOpts, user)
}

// ClearUserFee is a paid mutator transaction binding the contract method 0xfbea0c1a.
//
// Solidity: function clearUserFee(address user) returns()
func (_GPULease *GPULeaseTransactorSession) ClearUserFee(user common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.ClearUserFee(&_GPULease.TransactOpts, user)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactor) CompleteLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "completeLease", _leaseId)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseSession) CompleteLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.CompleteLease(&_GPULease.TransactOpts, _leaseId)
}

// CompleteLease is a paid mutator transaction binding the contract method 0x95e6e242.
//
// Solidity: function completeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactorSession) CompleteLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.CompleteLease(&_GPULease.TransactOpts, _leaseId)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactor) PauseLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "pauseLease", _leaseId)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseSession) PauseLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.PauseLease(&_GPULease.TransactOpts, _leaseId)
}

// PauseLease is a paid mutator transaction binding the contract method 0xb8eeaa78.
//
// Solidity: function pauseLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactorSession) PauseLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.PauseLease(&_GPULease.TransactOpts, _leaseId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULease *GPULeaseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULease *GPULeaseSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULease.Contract.RenounceOwnership(&_GPULease.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULease *GPULeaseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULease.Contract.RenounceOwnership(&_GPULease.TransactOpts)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactor) ResumeLease(opts *bind.TransactOpts, _leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "resumeLease", _leaseId)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseSession) ResumeLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.ResumeLease(&_GPULease.TransactOpts, _leaseId)
}

// ResumeLease is a paid mutator transaction binding the contract method 0x18ba11c9.
//
// Solidity: function resumeLease(uint256 _leaseId) returns()
func (_GPULease *GPULeaseTransactorSession) ResumeLease(_leaseId *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.ResumeLease(&_GPULease.TransactOpts, _leaseId)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_GPULease *GPULeaseTransactor) SetPlatformFee(opts *bind.TransactOpts, _feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "setPlatformFee", _feePercentage)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_GPULease *GPULeaseSession) SetPlatformFee(_feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.SetPlatformFee(&_GPULease.TransactOpts, _feePercentage)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feePercentage) returns()
func (_GPULease *GPULeaseTransactorSession) SetPlatformFee(_feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.SetPlatformFee(&_GPULease.TransactOpts, _feePercentage)
}

// SetReferralManager is a paid mutator transaction binding the contract method 0xc5f3a4c4.
//
// Solidity: function setReferralManager(address newReferralManager) returns()
func (_GPULease *GPULeaseTransactor) SetReferralManager(opts *bind.TransactOpts, newReferralManager common.Address) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "setReferralManager", newReferralManager)
}

// SetReferralManager is a paid mutator transaction binding the contract method 0xc5f3a4c4.
//
// Solidity: function setReferralManager(address newReferralManager) returns()
func (_GPULease *GPULeaseSession) SetReferralManager(newReferralManager common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.SetReferralManager(&_GPULease.TransactOpts, newReferralManager)
}

// SetReferralManager is a paid mutator transaction binding the contract method 0xc5f3a4c4.
//
// Solidity: function setReferralManager(address newReferralManager) returns()
func (_GPULease *GPULeaseTransactorSession) SetReferralManager(newReferralManager common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.SetReferralManager(&_GPULease.TransactOpts, newReferralManager)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_GPULease *GPULeaseTransactor) SetTreasury(opts *bind.TransactOpts, newTreasury common.Address) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "setTreasury", newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_GPULease *GPULeaseSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.SetTreasury(&_GPULease.TransactOpts, newTreasury)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(address newTreasury) returns()
func (_GPULease *GPULeaseTransactorSession) SetTreasury(newTreasury common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.SetTreasury(&_GPULease.TransactOpts, newTreasury)
}

// SetUserFee is a paid mutator transaction binding the contract method 0x809ad1f3.
//
// Solidity: function setUserFee(address user, uint256 _feePercentage) returns()
func (_GPULease *GPULeaseTransactor) SetUserFee(opts *bind.TransactOpts, user common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "setUserFee", user, _feePercentage)
}

// SetUserFee is a paid mutator transaction binding the contract method 0x809ad1f3.
//
// Solidity: function setUserFee(address user, uint256 _feePercentage) returns()
func (_GPULease *GPULeaseSession) SetUserFee(user common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.SetUserFee(&_GPULease.TransactOpts, user, _feePercentage)
}

// SetUserFee is a paid mutator transaction binding the contract method 0x809ad1f3.
//
// Solidity: function setUserFee(address user, uint256 _feePercentage) returns()
func (_GPULease *GPULeaseTransactorSession) SetUserFee(user common.Address, _feePercentage *big.Int) (*types.Transaction, error) {
	return _GPULease.Contract.SetUserFee(&_GPULease.TransactOpts, user, _feePercentage)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_GPULease *GPULeaseTransactor) StartLease(opts *bind.TransactOpts, _duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "startLease", _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_GPULease *GPULeaseSession) StartLease(_duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.StartLease(&_GPULease.TransactOpts, _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// StartLease is a paid mutator transaction binding the contract method 0xf648da80.
//
// Solidity: function startLease(uint256 _duration, uint256 _storagePricePerSecond, uint256 _computePricePerSecond, address _provider, address _user) returns(uint256 leaseId)
func (_GPULease *GPULeaseTransactorSession) StartLease(_duration *big.Int, _storagePricePerSecond *big.Int, _computePricePerSecond *big.Int, _provider common.Address, _user common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.StartLease(&_GPULease.TransactOpts, _duration, _storagePricePerSecond, _computePricePerSecond, _provider, _user)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULease *GPULeaseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GPULease.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULease *GPULeaseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.TransferOwnership(&_GPULease.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULease *GPULeaseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULease.Contract.TransferOwnership(&_GPULease.TransactOpts, newOwner)
}

// GPULeaseLeaseCompletedIterator is returned from FilterLeaseCompleted and is used to iterate over the raw logs and unpacked data for LeaseCompleted events raised by the GPULease contract.
type GPULeaseLeaseCompletedIterator struct {
	Event *GPULeaseLeaseCompleted // Event containing the contract specifics and raw log

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
func (it *GPULeaseLeaseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseLeaseCompleted)
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
		it.Event = new(GPULeaseLeaseCompleted)
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
func (it *GPULeaseLeaseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseLeaseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseLeaseCompleted represents a LeaseCompleted event raised by the GPULease contract.
type GPULeaseLeaseCompleted struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeaseCompleted is a free log retrieval operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) FilterLeaseCompleted(opts *bind.FilterOpts) (*GPULeaseLeaseCompletedIterator, error) {

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "LeaseCompleted")
	if err != nil {
		return nil, err
	}
	return &GPULeaseLeaseCompletedIterator{contract: _GPULease.contract, event: "LeaseCompleted", logs: logs, sub: sub}, nil
}

// WatchLeaseCompleted is a free log subscription operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) WatchLeaseCompleted(opts *bind.WatchOpts, sink chan<- *GPULeaseLeaseCompleted) (event.Subscription, error) {

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "LeaseCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseLeaseCompleted)
				if err := _GPULease.contract.UnpackLog(event, "LeaseCompleted", log); err != nil {
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

// ParseLeaseCompleted is a log parse operation binding the contract event 0x5a1241b1f059ce7cc7acee21bf522f44360638a58f17c58bd390a0f0af7e1937.
//
// Solidity: event LeaseCompleted(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) ParseLeaseCompleted(log types.Log) (*GPULeaseLeaseCompleted, error) {
	event := new(GPULeaseLeaseCompleted)
	if err := _GPULease.contract.UnpackLog(event, "LeaseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseLeasePausedIterator is returned from FilterLeasePaused and is used to iterate over the raw logs and unpacked data for LeasePaused events raised by the GPULease contract.
type GPULeaseLeasePausedIterator struct {
	Event *GPULeaseLeasePaused // Event containing the contract specifics and raw log

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
func (it *GPULeaseLeasePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseLeasePaused)
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
		it.Event = new(GPULeaseLeasePaused)
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
func (it *GPULeaseLeasePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseLeasePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseLeasePaused represents a LeasePaused event raised by the GPULease contract.
type GPULeaseLeasePaused struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeasePaused is a free log retrieval operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) FilterLeasePaused(opts *bind.FilterOpts) (*GPULeaseLeasePausedIterator, error) {

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "LeasePaused")
	if err != nil {
		return nil, err
	}
	return &GPULeaseLeasePausedIterator{contract: _GPULease.contract, event: "LeasePaused", logs: logs, sub: sub}, nil
}

// WatchLeasePaused is a free log subscription operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) WatchLeasePaused(opts *bind.WatchOpts, sink chan<- *GPULeaseLeasePaused) (event.Subscription, error) {

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "LeasePaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseLeasePaused)
				if err := _GPULease.contract.UnpackLog(event, "LeasePaused", log); err != nil {
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

// ParseLeasePaused is a log parse operation binding the contract event 0xb04140f2785e8aa46842b88fdbee148f5bcd51a8e9c4b367b7f431e80a6ed5a1.
//
// Solidity: event LeasePaused(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) ParseLeasePaused(log types.Log) (*GPULeaseLeasePaused, error) {
	event := new(GPULeaseLeasePaused)
	if err := _GPULease.contract.UnpackLog(event, "LeasePaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseLeaseResumedIterator is returned from FilterLeaseResumed and is used to iterate over the raw logs and unpacked data for LeaseResumed events raised by the GPULease contract.
type GPULeaseLeaseResumedIterator struct {
	Event *GPULeaseLeaseResumed // Event containing the contract specifics and raw log

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
func (it *GPULeaseLeaseResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseLeaseResumed)
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
		it.Event = new(GPULeaseLeaseResumed)
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
func (it *GPULeaseLeaseResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseLeaseResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseLeaseResumed represents a LeaseResumed event raised by the GPULease contract.
type GPULeaseLeaseResumed struct {
	LeaseId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLeaseResumed is a free log retrieval operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) FilterLeaseResumed(opts *bind.FilterOpts) (*GPULeaseLeaseResumedIterator, error) {

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "LeaseResumed")
	if err != nil {
		return nil, err
	}
	return &GPULeaseLeaseResumedIterator{contract: _GPULease.contract, event: "LeaseResumed", logs: logs, sub: sub}, nil
}

// WatchLeaseResumed is a free log subscription operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) WatchLeaseResumed(opts *bind.WatchOpts, sink chan<- *GPULeaseLeaseResumed) (event.Subscription, error) {

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "LeaseResumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseLeaseResumed)
				if err := _GPULease.contract.UnpackLog(event, "LeaseResumed", log); err != nil {
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

// ParseLeaseResumed is a log parse operation binding the contract event 0x0ec1e889672cca90b2287dde9161b83295c31ef8ed2a8bd943a7f04b183f9167.
//
// Solidity: event LeaseResumed(uint256 leaseId)
func (_GPULease *GPULeaseFilterer) ParseLeaseResumed(log types.Log) (*GPULeaseLeaseResumed, error) {
	event := new(GPULeaseLeaseResumed)
	if err := _GPULease.contract.UnpackLog(event, "LeaseResumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseLeaseStartedIterator is returned from FilterLeaseStarted and is used to iterate over the raw logs and unpacked data for LeaseStarted events raised by the GPULease contract.
type GPULeaseLeaseStartedIterator struct {
	Event *GPULeaseLeaseStarted // Event containing the contract specifics and raw log

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
func (it *GPULeaseLeaseStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseLeaseStarted)
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
		it.Event = new(GPULeaseLeaseStarted)
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
func (it *GPULeaseLeaseStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseLeaseStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseLeaseStarted represents a LeaseStarted event raised by the GPULease contract.
type GPULeaseLeaseStarted struct {
	LeaseId  *big.Int
	User     common.Address
	Provider common.Address
	Duration *big.Int
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLeaseStarted is a free log retrieval operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_GPULease *GPULeaseFilterer) FilterLeaseStarted(opts *bind.FilterOpts) (*GPULeaseLeaseStartedIterator, error) {

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "LeaseStarted")
	if err != nil {
		return nil, err
	}
	return &GPULeaseLeaseStartedIterator{contract: _GPULease.contract, event: "LeaseStarted", logs: logs, sub: sub}, nil
}

// WatchLeaseStarted is a free log subscription operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_GPULease *GPULeaseFilterer) WatchLeaseStarted(opts *bind.WatchOpts, sink chan<- *GPULeaseLeaseStarted) (event.Subscription, error) {

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "LeaseStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseLeaseStarted)
				if err := _GPULease.contract.UnpackLog(event, "LeaseStarted", log); err != nil {
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

// ParseLeaseStarted is a log parse operation binding the contract event 0x4501059f1dbad8e151132c2515ee7ea78d4b2540b9e56e941a03fdb9888efe28.
//
// Solidity: event LeaseStarted(uint256 leaseId, address user, address provider, uint256 duration, uint256 amount)
func (_GPULease *GPULeaseFilterer) ParseLeaseStarted(log types.Log) (*GPULeaseLeaseStarted, error) {
	event := new(GPULeaseLeaseStarted)
	if err := _GPULease.contract.UnpackLog(event, "LeaseStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GPULease contract.
type GPULeaseOwnershipTransferredIterator struct {
	Event *GPULeaseOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GPULeaseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseOwnershipTransferred)
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
		it.Event = new(GPULeaseOwnershipTransferred)
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
func (it *GPULeaseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseOwnershipTransferred represents a OwnershipTransferred event raised by the GPULease contract.
type GPULeaseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULease *GPULeaseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GPULeaseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseOwnershipTransferredIterator{contract: _GPULease.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULease *GPULeaseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GPULeaseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseOwnershipTransferred)
				if err := _GPULease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GPULease *GPULeaseFilterer) ParseOwnershipTransferred(log types.Log) (*GPULeaseOwnershipTransferred, error) {
	event := new(GPULeaseOwnershipTransferred)
	if err := _GPULease.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeasePlatformFeeUpdatedIterator is returned from FilterPlatformFeeUpdated and is used to iterate over the raw logs and unpacked data for PlatformFeeUpdated events raised by the GPULease contract.
type GPULeasePlatformFeeUpdatedIterator struct {
	Event *GPULeasePlatformFeeUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeasePlatformFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeasePlatformFeeUpdated)
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
		it.Event = new(GPULeasePlatformFeeUpdated)
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
func (it *GPULeasePlatformFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeasePlatformFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeasePlatformFeeUpdated represents a PlatformFeeUpdated event raised by the GPULease contract.
type GPULeasePlatformFeeUpdated struct {
	PreviousFeePercentage *big.Int
	NewFeePercentage      *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterPlatformFeeUpdated is a free log retrieval operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 previousFeePercentage, uint256 newFeePercentage)
func (_GPULease *GPULeaseFilterer) FilterPlatformFeeUpdated(opts *bind.FilterOpts) (*GPULeasePlatformFeeUpdatedIterator, error) {

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &GPULeasePlatformFeeUpdatedIterator{contract: _GPULease.contract, event: "PlatformFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchPlatformFeeUpdated is a free log subscription operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 previousFeePercentage, uint256 newFeePercentage)
func (_GPULease *GPULeaseFilterer) WatchPlatformFeeUpdated(opts *bind.WatchOpts, sink chan<- *GPULeasePlatformFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "PlatformFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeasePlatformFeeUpdated)
				if err := _GPULease.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
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

// ParsePlatformFeeUpdated is a log parse operation binding the contract event 0xd347e206f25a89b917fc9482f1a2d294d749baa4dc9bde7fb495ee11fe491643.
//
// Solidity: event PlatformFeeUpdated(uint256 previousFeePercentage, uint256 newFeePercentage)
func (_GPULease *GPULeaseFilterer) ParsePlatformFeeUpdated(log types.Log) (*GPULeasePlatformFeeUpdated, error) {
	event := new(GPULeasePlatformFeeUpdated)
	if err := _GPULease.contract.UnpackLog(event, "PlatformFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseReferralManagerUpdatedIterator is returned from FilterReferralManagerUpdated and is used to iterate over the raw logs and unpacked data for ReferralManagerUpdated events raised by the GPULease contract.
type GPULeaseReferralManagerUpdatedIterator struct {
	Event *GPULeaseReferralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeaseReferralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseReferralManagerUpdated)
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
		it.Event = new(GPULeaseReferralManagerUpdated)
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
func (it *GPULeaseReferralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseReferralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseReferralManagerUpdated represents a ReferralManagerUpdated event raised by the GPULease contract.
type GPULeaseReferralManagerUpdated struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReferralManagerUpdated is a free log retrieval operation binding the contract event 0x8fb6763bbb7019b7a1b61a1fcf72799033d91373edd63d44a6304d464fe903e6.
//
// Solidity: event ReferralManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULease *GPULeaseFilterer) FilterReferralManagerUpdated(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*GPULeaseReferralManagerUpdatedIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "ReferralManagerUpdated", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralManagerUpdatedIterator{contract: _GPULease.contract, event: "ReferralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchReferralManagerUpdated is a free log subscription operation binding the contract event 0x8fb6763bbb7019b7a1b61a1fcf72799033d91373edd63d44a6304d464fe903e6.
//
// Solidity: event ReferralManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULease *GPULeaseFilterer) WatchReferralManagerUpdated(opts *bind.WatchOpts, sink chan<- *GPULeaseReferralManagerUpdated, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "ReferralManagerUpdated", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseReferralManagerUpdated)
				if err := _GPULease.contract.UnpackLog(event, "ReferralManagerUpdated", log); err != nil {
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

// ParseReferralManagerUpdated is a log parse operation binding the contract event 0x8fb6763bbb7019b7a1b61a1fcf72799033d91373edd63d44a6304d464fe903e6.
//
// Solidity: event ReferralManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULease *GPULeaseFilterer) ParseReferralManagerUpdated(log types.Log) (*GPULeaseReferralManagerUpdated, error) {
	event := new(GPULeaseReferralManagerUpdated)
	if err := _GPULease.contract.UnpackLog(event, "ReferralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseUserFeeClearedIterator is returned from FilterUserFeeCleared and is used to iterate over the raw logs and unpacked data for UserFeeCleared events raised by the GPULease contract.
type GPULeaseUserFeeClearedIterator struct {
	Event *GPULeaseUserFeeCleared // Event containing the contract specifics and raw log

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
func (it *GPULeaseUserFeeClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseUserFeeCleared)
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
		it.Event = new(GPULeaseUserFeeCleared)
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
func (it *GPULeaseUserFeeClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseUserFeeClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseUserFeeCleared represents a UserFeeCleared event raised by the GPULease contract.
type GPULeaseUserFeeCleared struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserFeeCleared is a free log retrieval operation binding the contract event 0xdd35a637cf2e753ff8db76de0e69ba771854c070874289c1d4d0698178144177.
//
// Solidity: event UserFeeCleared(address indexed user)
func (_GPULease *GPULeaseFilterer) FilterUserFeeCleared(opts *bind.FilterOpts, user []common.Address) (*GPULeaseUserFeeClearedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "UserFeeCleared", userRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseUserFeeClearedIterator{contract: _GPULease.contract, event: "UserFeeCleared", logs: logs, sub: sub}, nil
}

// WatchUserFeeCleared is a free log subscription operation binding the contract event 0xdd35a637cf2e753ff8db76de0e69ba771854c070874289c1d4d0698178144177.
//
// Solidity: event UserFeeCleared(address indexed user)
func (_GPULease *GPULeaseFilterer) WatchUserFeeCleared(opts *bind.WatchOpts, sink chan<- *GPULeaseUserFeeCleared, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "UserFeeCleared", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseUserFeeCleared)
				if err := _GPULease.contract.UnpackLog(event, "UserFeeCleared", log); err != nil {
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

// ParseUserFeeCleared is a log parse operation binding the contract event 0xdd35a637cf2e753ff8db76de0e69ba771854c070874289c1d4d0698178144177.
//
// Solidity: event UserFeeCleared(address indexed user)
func (_GPULease *GPULeaseFilterer) ParseUserFeeCleared(log types.Log) (*GPULeaseUserFeeCleared, error) {
	event := new(GPULeaseUserFeeCleared)
	if err := _GPULease.contract.UnpackLog(event, "UserFeeCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseUserFeeUpdatedIterator is returned from FilterUserFeeUpdated and is used to iterate over the raw logs and unpacked data for UserFeeUpdated events raised by the GPULease contract.
type GPULeaseUserFeeUpdatedIterator struct {
	Event *GPULeaseUserFeeUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeaseUserFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseUserFeeUpdated)
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
		it.Event = new(GPULeaseUserFeeUpdated)
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
func (it *GPULeaseUserFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseUserFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseUserFeeUpdated represents a UserFeeUpdated event raised by the GPULease contract.
type GPULeaseUserFeeUpdated struct {
	User          common.Address
	FeePercentage *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterUserFeeUpdated is a free log retrieval operation binding the contract event 0x4a8c85f9ead8a36e5bfa09e1d7b5f83286089f8bb79e4cb0668049659c8cc919.
//
// Solidity: event UserFeeUpdated(address indexed user, uint256 feePercentage)
func (_GPULease *GPULeaseFilterer) FilterUserFeeUpdated(opts *bind.FilterOpts, user []common.Address) (*GPULeaseUserFeeUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULease.contract.FilterLogs(opts, "UserFeeUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseUserFeeUpdatedIterator{contract: _GPULease.contract, event: "UserFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchUserFeeUpdated is a free log subscription operation binding the contract event 0x4a8c85f9ead8a36e5bfa09e1d7b5f83286089f8bb79e4cb0668049659c8cc919.
//
// Solidity: event UserFeeUpdated(address indexed user, uint256 feePercentage)
func (_GPULease *GPULeaseFilterer) WatchUserFeeUpdated(opts *bind.WatchOpts, sink chan<- *GPULeaseUserFeeUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULease.contract.WatchLogs(opts, "UserFeeUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseUserFeeUpdated)
				if err := _GPULease.contract.UnpackLog(event, "UserFeeUpdated", log); err != nil {
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

// ParseUserFeeUpdated is a log parse operation binding the contract event 0x4a8c85f9ead8a36e5bfa09e1d7b5f83286089f8bb79e4cb0668049659c8cc919.
//
// Solidity: event UserFeeUpdated(address indexed user, uint256 feePercentage)
func (_GPULease *GPULeaseFilterer) ParseUserFeeUpdated(log types.Log) (*GPULeaseUserFeeUpdated, error) {
	event := new(GPULeaseUserFeeUpdated)
	if err := _GPULease.contract.UnpackLog(event, "UserFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
