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

// GPULeaseReferralMetaData contains all meta data concerning the GPULeaseReferral contract.
var GPULeaseReferralMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousShareBps\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShareBps\",\"type\":\"uint256\"}],\"name\":\"ReferralShareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousReferrer\",\"type\":\"address\"}],\"name\":\"ReferrerCleared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"ReferrerUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"clearReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"referralShareBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"referrerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_shareBps\",\"type\":\"uint256\"}],\"name\":\"setReferralShareBps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"setReferrer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GPULeaseReferralABI is the input ABI used to generate the binding from.
// Deprecated: Use GPULeaseReferralMetaData.ABI instead.
var GPULeaseReferralABI = GPULeaseReferralMetaData.ABI

// GPULeaseReferral is an auto generated Go binding around an Ethereum contract.
type GPULeaseReferral struct {
	GPULeaseReferralCaller     // Read-only binding to the contract
	GPULeaseReferralTransactor // Write-only binding to the contract
	GPULeaseReferralFilterer   // Log filterer for contract events
}

// GPULeaseReferralCaller is an auto generated read-only Go binding around an Ethereum contract.
type GPULeaseReferralCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseReferralTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GPULeaseReferralTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseReferralFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GPULeaseReferralFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseReferralSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GPULeaseReferralSession struct {
	Contract     *GPULeaseReferral // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GPULeaseReferralCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GPULeaseReferralCallerSession struct {
	Contract *GPULeaseReferralCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// GPULeaseReferralTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GPULeaseReferralTransactorSession struct {
	Contract     *GPULeaseReferralTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GPULeaseReferralRaw is an auto generated low-level Go binding around an Ethereum contract.
type GPULeaseReferralRaw struct {
	Contract *GPULeaseReferral // Generic contract binding to access the raw methods on
}

// GPULeaseReferralCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GPULeaseReferralCallerRaw struct {
	Contract *GPULeaseReferralCaller // Generic read-only contract binding to access the raw methods on
}

// GPULeaseReferralTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GPULeaseReferralTransactorRaw struct {
	Contract *GPULeaseReferralTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGPULeaseReferral creates a new instance of GPULeaseReferral, bound to a specific deployed contract.
func NewGPULeaseReferral(address common.Address, backend bind.ContractBackend) (*GPULeaseReferral, error) {
	contract, err := bindGPULeaseReferral(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferral{GPULeaseReferralCaller: GPULeaseReferralCaller{contract: contract}, GPULeaseReferralTransactor: GPULeaseReferralTransactor{contract: contract}, GPULeaseReferralFilterer: GPULeaseReferralFilterer{contract: contract}}, nil
}

// NewGPULeaseReferralCaller creates a new read-only instance of GPULeaseReferral, bound to a specific deployed contract.
func NewGPULeaseReferralCaller(address common.Address, caller bind.ContractCaller) (*GPULeaseReferralCaller, error) {
	contract, err := bindGPULeaseReferral(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralCaller{contract: contract}, nil
}

// NewGPULeaseReferralTransactor creates a new write-only instance of GPULeaseReferral, bound to a specific deployed contract.
func NewGPULeaseReferralTransactor(address common.Address, transactor bind.ContractTransactor) (*GPULeaseReferralTransactor, error) {
	contract, err := bindGPULeaseReferral(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralTransactor{contract: contract}, nil
}

// NewGPULeaseReferralFilterer creates a new log filterer instance of GPULeaseReferral, bound to a specific deployed contract.
func NewGPULeaseReferralFilterer(address common.Address, filterer bind.ContractFilterer) (*GPULeaseReferralFilterer, error) {
	contract, err := bindGPULeaseReferral(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralFilterer{contract: contract}, nil
}

// bindGPULeaseReferral binds a generic wrapper to an already deployed contract.
func bindGPULeaseReferral(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GPULeaseReferralMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULeaseReferral *GPULeaseReferralRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULeaseReferral.Contract.GPULeaseReferralCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULeaseReferral *GPULeaseReferralRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.GPULeaseReferralTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULeaseReferral *GPULeaseReferralRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.GPULeaseReferralTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULeaseReferral *GPULeaseReferralCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULeaseReferral.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULeaseReferral *GPULeaseReferralTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULeaseReferral *GPULeaseReferralTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseReferral *GPULeaseReferralCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULeaseReferral.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseReferral *GPULeaseReferralSession) Owner() (common.Address, error) {
	return _GPULeaseReferral.Contract.Owner(&_GPULeaseReferral.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseReferral *GPULeaseReferralCallerSession) Owner() (common.Address, error) {
	return _GPULeaseReferral.Contract.Owner(&_GPULeaseReferral.CallOpts)
}

// ReferralShareBps is a free data retrieval call binding the contract method 0x47c9bc2d.
//
// Solidity: function referralShareBps() view returns(uint256)
func (_GPULeaseReferral *GPULeaseReferralCaller) ReferralShareBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GPULeaseReferral.contract.Call(opts, &out, "referralShareBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReferralShareBps is a free data retrieval call binding the contract method 0x47c9bc2d.
//
// Solidity: function referralShareBps() view returns(uint256)
func (_GPULeaseReferral *GPULeaseReferralSession) ReferralShareBps() (*big.Int, error) {
	return _GPULeaseReferral.Contract.ReferralShareBps(&_GPULeaseReferral.CallOpts)
}

// ReferralShareBps is a free data retrieval call binding the contract method 0x47c9bc2d.
//
// Solidity: function referralShareBps() view returns(uint256)
func (_GPULeaseReferral *GPULeaseReferralCallerSession) ReferralShareBps() (*big.Int, error) {
	return _GPULeaseReferral.Contract.ReferralShareBps(&_GPULeaseReferral.CallOpts)
}

// ReferrerOf is a free data retrieval call binding the contract method 0xd21cacdf.
//
// Solidity: function referrerOf(address ) view returns(address)
func (_GPULeaseReferral *GPULeaseReferralCaller) ReferrerOf(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _GPULeaseReferral.contract.Call(opts, &out, "referrerOf", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReferrerOf is a free data retrieval call binding the contract method 0xd21cacdf.
//
// Solidity: function referrerOf(address ) view returns(address)
func (_GPULeaseReferral *GPULeaseReferralSession) ReferrerOf(arg0 common.Address) (common.Address, error) {
	return _GPULeaseReferral.Contract.ReferrerOf(&_GPULeaseReferral.CallOpts, arg0)
}

// ReferrerOf is a free data retrieval call binding the contract method 0xd21cacdf.
//
// Solidity: function referrerOf(address ) view returns(address)
func (_GPULeaseReferral *GPULeaseReferralCallerSession) ReferrerOf(arg0 common.Address) (common.Address, error) {
	return _GPULeaseReferral.Contract.ReferrerOf(&_GPULeaseReferral.CallOpts, arg0)
}

// ClearReferrer is a paid mutator transaction binding the contract method 0x45f97f1d.
//
// Solidity: function clearReferrer(address user) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactor) ClearReferrer(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.contract.Transact(opts, "clearReferrer", user)
}

// ClearReferrer is a paid mutator transaction binding the contract method 0x45f97f1d.
//
// Solidity: function clearReferrer(address user) returns()
func (_GPULeaseReferral *GPULeaseReferralSession) ClearReferrer(user common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.ClearReferrer(&_GPULeaseReferral.TransactOpts, user)
}

// ClearReferrer is a paid mutator transaction binding the contract method 0x45f97f1d.
//
// Solidity: function clearReferrer(address user) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactorSession) ClearReferrer(user common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.ClearReferrer(&_GPULeaseReferral.TransactOpts, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseReferral *GPULeaseReferralTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseReferral.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseReferral *GPULeaseReferralSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.RenounceOwnership(&_GPULeaseReferral.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseReferral *GPULeaseReferralTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.RenounceOwnership(&_GPULeaseReferral.TransactOpts)
}

// SetReferralShareBps is a paid mutator transaction binding the contract method 0xd07e995b.
//
// Solidity: function setReferralShareBps(uint256 _shareBps) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactor) SetReferralShareBps(opts *bind.TransactOpts, _shareBps *big.Int) (*types.Transaction, error) {
	return _GPULeaseReferral.contract.Transact(opts, "setReferralShareBps", _shareBps)
}

// SetReferralShareBps is a paid mutator transaction binding the contract method 0xd07e995b.
//
// Solidity: function setReferralShareBps(uint256 _shareBps) returns()
func (_GPULeaseReferral *GPULeaseReferralSession) SetReferralShareBps(_shareBps *big.Int) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.SetReferralShareBps(&_GPULeaseReferral.TransactOpts, _shareBps)
}

// SetReferralShareBps is a paid mutator transaction binding the contract method 0xd07e995b.
//
// Solidity: function setReferralShareBps(uint256 _shareBps) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactorSession) SetReferralShareBps(_shareBps *big.Int) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.SetReferralShareBps(&_GPULeaseReferral.TransactOpts, _shareBps)
}

// SetReferrer is a paid mutator transaction binding the contract method 0xbbddaca3.
//
// Solidity: function setReferrer(address user, address referrer) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactor) SetReferrer(opts *bind.TransactOpts, user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.contract.Transact(opts, "setReferrer", user, referrer)
}

// SetReferrer is a paid mutator transaction binding the contract method 0xbbddaca3.
//
// Solidity: function setReferrer(address user, address referrer) returns()
func (_GPULeaseReferral *GPULeaseReferralSession) SetReferrer(user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.SetReferrer(&_GPULeaseReferral.TransactOpts, user, referrer)
}

// SetReferrer is a paid mutator transaction binding the contract method 0xbbddaca3.
//
// Solidity: function setReferrer(address user, address referrer) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactorSession) SetReferrer(user common.Address, referrer common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.SetReferrer(&_GPULeaseReferral.TransactOpts, user, referrer)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseReferral *GPULeaseReferralSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.TransferOwnership(&_GPULeaseReferral.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseReferral *GPULeaseReferralTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseReferral.Contract.TransferOwnership(&_GPULeaseReferral.TransactOpts, newOwner)
}

// GPULeaseReferralOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GPULeaseReferral contract.
type GPULeaseReferralOwnershipTransferredIterator struct {
	Event *GPULeaseReferralOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GPULeaseReferralOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseReferralOwnershipTransferred)
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
		it.Event = new(GPULeaseReferralOwnershipTransferred)
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
func (it *GPULeaseReferralOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseReferralOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseReferralOwnershipTransferred represents a OwnershipTransferred event raised by the GPULeaseReferral contract.
type GPULeaseReferralOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULeaseReferral *GPULeaseReferralFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GPULeaseReferralOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralOwnershipTransferredIterator{contract: _GPULeaseReferral.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULeaseReferral *GPULeaseReferralFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GPULeaseReferralOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseReferralOwnershipTransferred)
				if err := _GPULeaseReferral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GPULeaseReferral *GPULeaseReferralFilterer) ParseOwnershipTransferred(log types.Log) (*GPULeaseReferralOwnershipTransferred, error) {
	event := new(GPULeaseReferralOwnershipTransferred)
	if err := _GPULeaseReferral.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseReferralReferralShareUpdatedIterator is returned from FilterReferralShareUpdated and is used to iterate over the raw logs and unpacked data for ReferralShareUpdated events raised by the GPULeaseReferral contract.
type GPULeaseReferralReferralShareUpdatedIterator struct {
	Event *GPULeaseReferralReferralShareUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeaseReferralReferralShareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseReferralReferralShareUpdated)
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
		it.Event = new(GPULeaseReferralReferralShareUpdated)
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
func (it *GPULeaseReferralReferralShareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseReferralReferralShareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseReferralReferralShareUpdated represents a ReferralShareUpdated event raised by the GPULeaseReferral contract.
type GPULeaseReferralReferralShareUpdated struct {
	PreviousShareBps *big.Int
	NewShareBps      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReferralShareUpdated is a free log retrieval operation binding the contract event 0xf4da55ee58d2a4fa2cf3eb9963234feccae3e2080231395b2bd5ff8c41c4ebf2.
//
// Solidity: event ReferralShareUpdated(uint256 previousShareBps, uint256 newShareBps)
func (_GPULeaseReferral *GPULeaseReferralFilterer) FilterReferralShareUpdated(opts *bind.FilterOpts) (*GPULeaseReferralReferralShareUpdatedIterator, error) {

	logs, sub, err := _GPULeaseReferral.contract.FilterLogs(opts, "ReferralShareUpdated")
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralReferralShareUpdatedIterator{contract: _GPULeaseReferral.contract, event: "ReferralShareUpdated", logs: logs, sub: sub}, nil
}

// WatchReferralShareUpdated is a free log subscription operation binding the contract event 0xf4da55ee58d2a4fa2cf3eb9963234feccae3e2080231395b2bd5ff8c41c4ebf2.
//
// Solidity: event ReferralShareUpdated(uint256 previousShareBps, uint256 newShareBps)
func (_GPULeaseReferral *GPULeaseReferralFilterer) WatchReferralShareUpdated(opts *bind.WatchOpts, sink chan<- *GPULeaseReferralReferralShareUpdated) (event.Subscription, error) {

	logs, sub, err := _GPULeaseReferral.contract.WatchLogs(opts, "ReferralShareUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseReferralReferralShareUpdated)
				if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferralShareUpdated", log); err != nil {
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

// ParseReferralShareUpdated is a log parse operation binding the contract event 0xf4da55ee58d2a4fa2cf3eb9963234feccae3e2080231395b2bd5ff8c41c4ebf2.
//
// Solidity: event ReferralShareUpdated(uint256 previousShareBps, uint256 newShareBps)
func (_GPULeaseReferral *GPULeaseReferralFilterer) ParseReferralShareUpdated(log types.Log) (*GPULeaseReferralReferralShareUpdated, error) {
	event := new(GPULeaseReferralReferralShareUpdated)
	if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferralShareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseReferralReferrerClearedIterator is returned from FilterReferrerCleared and is used to iterate over the raw logs and unpacked data for ReferrerCleared events raised by the GPULeaseReferral contract.
type GPULeaseReferralReferrerClearedIterator struct {
	Event *GPULeaseReferralReferrerCleared // Event containing the contract specifics and raw log

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
func (it *GPULeaseReferralReferrerClearedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseReferralReferrerCleared)
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
		it.Event = new(GPULeaseReferralReferrerCleared)
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
func (it *GPULeaseReferralReferrerClearedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseReferralReferrerClearedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseReferralReferrerCleared represents a ReferrerCleared event raised by the GPULeaseReferral contract.
type GPULeaseReferralReferrerCleared struct {
	User             common.Address
	PreviousReferrer common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReferrerCleared is a free log retrieval operation binding the contract event 0x37f0927e0f56b1cc3358bea3ddc70d4d2faad25a2764df9628203d6918525956.
//
// Solidity: event ReferrerCleared(address indexed user, address indexed previousReferrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) FilterReferrerCleared(opts *bind.FilterOpts, user []common.Address, previousReferrer []common.Address) (*GPULeaseReferralReferrerClearedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var previousReferrerRule []interface{}
	for _, previousReferrerItem := range previousReferrer {
		previousReferrerRule = append(previousReferrerRule, previousReferrerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.FilterLogs(opts, "ReferrerCleared", userRule, previousReferrerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralReferrerClearedIterator{contract: _GPULeaseReferral.contract, event: "ReferrerCleared", logs: logs, sub: sub}, nil
}

// WatchReferrerCleared is a free log subscription operation binding the contract event 0x37f0927e0f56b1cc3358bea3ddc70d4d2faad25a2764df9628203d6918525956.
//
// Solidity: event ReferrerCleared(address indexed user, address indexed previousReferrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) WatchReferrerCleared(opts *bind.WatchOpts, sink chan<- *GPULeaseReferralReferrerCleared, user []common.Address, previousReferrer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var previousReferrerRule []interface{}
	for _, previousReferrerItem := range previousReferrer {
		previousReferrerRule = append(previousReferrerRule, previousReferrerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.WatchLogs(opts, "ReferrerCleared", userRule, previousReferrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseReferralReferrerCleared)
				if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferrerCleared", log); err != nil {
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

// ParseReferrerCleared is a log parse operation binding the contract event 0x37f0927e0f56b1cc3358bea3ddc70d4d2faad25a2764df9628203d6918525956.
//
// Solidity: event ReferrerCleared(address indexed user, address indexed previousReferrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) ParseReferrerCleared(log types.Log) (*GPULeaseReferralReferrerCleared, error) {
	event := new(GPULeaseReferralReferrerCleared)
	if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferrerCleared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseReferralReferrerUpdatedIterator is returned from FilterReferrerUpdated and is used to iterate over the raw logs and unpacked data for ReferrerUpdated events raised by the GPULeaseReferral contract.
type GPULeaseReferralReferrerUpdatedIterator struct {
	Event *GPULeaseReferralReferrerUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeaseReferralReferrerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseReferralReferrerUpdated)
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
		it.Event = new(GPULeaseReferralReferrerUpdated)
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
func (it *GPULeaseReferralReferrerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseReferralReferrerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseReferralReferrerUpdated represents a ReferrerUpdated event raised by the GPULeaseReferral contract.
type GPULeaseReferralReferrerUpdated struct {
	User     common.Address
	Referrer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterReferrerUpdated is a free log retrieval operation binding the contract event 0xc40302e3b5897f6966b131753cb09f65aa712ae82e3f49b189d089d5694256e3.
//
// Solidity: event ReferrerUpdated(address indexed user, address indexed referrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) FilterReferrerUpdated(opts *bind.FilterOpts, user []common.Address, referrer []common.Address) (*GPULeaseReferralReferrerUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.FilterLogs(opts, "ReferrerUpdated", userRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseReferralReferrerUpdatedIterator{contract: _GPULeaseReferral.contract, event: "ReferrerUpdated", logs: logs, sub: sub}, nil
}

// WatchReferrerUpdated is a free log subscription operation binding the contract event 0xc40302e3b5897f6966b131753cb09f65aa712ae82e3f49b189d089d5694256e3.
//
// Solidity: event ReferrerUpdated(address indexed user, address indexed referrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) WatchReferrerUpdated(opts *bind.WatchOpts, sink chan<- *GPULeaseReferralReferrerUpdated, user []common.Address, referrer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _GPULeaseReferral.contract.WatchLogs(opts, "ReferrerUpdated", userRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseReferralReferrerUpdated)
				if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferrerUpdated", log); err != nil {
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

// ParseReferrerUpdated is a log parse operation binding the contract event 0xc40302e3b5897f6966b131753cb09f65aa712ae82e3f49b189d089d5694256e3.
//
// Solidity: event ReferrerUpdated(address indexed user, address indexed referrer)
func (_GPULeaseReferral *GPULeaseReferralFilterer) ParseReferrerUpdated(log types.Log) (*GPULeaseReferralReferrerUpdated, error) {
	event := new(GPULeaseReferralReferrerUpdated)
	if err := _GPULeaseReferral.contract.UnpackLog(event, "ReferrerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
