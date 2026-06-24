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

// GPULeaseWalletMetaData contains all meta data concerning the GPULeaseWallet contract.
var GPULeaseWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"credit_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newManager\",\"type\":\"address\"}],\"name\":\"LeaseManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credit\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"creditBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"debitBalance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leaseManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"moveBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLeaseManager\",\"type\":\"address\"}],\"name\":\"setLeaseManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GPULeaseWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use GPULeaseWalletMetaData.ABI instead.
var GPULeaseWalletABI = GPULeaseWalletMetaData.ABI

// GPULeaseWallet is an auto generated Go binding around an Ethereum contract.
type GPULeaseWallet struct {
	GPULeaseWalletCaller     // Read-only binding to the contract
	GPULeaseWalletTransactor // Write-only binding to the contract
	GPULeaseWalletFilterer   // Log filterer for contract events
}

// GPULeaseWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type GPULeaseWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GPULeaseWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GPULeaseWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GPULeaseWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GPULeaseWalletSession struct {
	Contract     *GPULeaseWallet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GPULeaseWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GPULeaseWalletCallerSession struct {
	Contract *GPULeaseWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GPULeaseWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GPULeaseWalletTransactorSession struct {
	Contract     *GPULeaseWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GPULeaseWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type GPULeaseWalletRaw struct {
	Contract *GPULeaseWallet // Generic contract binding to access the raw methods on
}

// GPULeaseWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GPULeaseWalletCallerRaw struct {
	Contract *GPULeaseWalletCaller // Generic read-only contract binding to access the raw methods on
}

// GPULeaseWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GPULeaseWalletTransactorRaw struct {
	Contract *GPULeaseWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGPULeaseWallet creates a new instance of GPULeaseWallet, bound to a specific deployed contract.
func NewGPULeaseWallet(address common.Address, backend bind.ContractBackend) (*GPULeaseWallet, error) {
	contract, err := bindGPULeaseWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWallet{GPULeaseWalletCaller: GPULeaseWalletCaller{contract: contract}, GPULeaseWalletTransactor: GPULeaseWalletTransactor{contract: contract}, GPULeaseWalletFilterer: GPULeaseWalletFilterer{contract: contract}}, nil
}

// NewGPULeaseWalletCaller creates a new read-only instance of GPULeaseWallet, bound to a specific deployed contract.
func NewGPULeaseWalletCaller(address common.Address, caller bind.ContractCaller) (*GPULeaseWalletCaller, error) {
	contract, err := bindGPULeaseWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletCaller{contract: contract}, nil
}

// NewGPULeaseWalletTransactor creates a new write-only instance of GPULeaseWallet, bound to a specific deployed contract.
func NewGPULeaseWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*GPULeaseWalletTransactor, error) {
	contract, err := bindGPULeaseWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletTransactor{contract: contract}, nil
}

// NewGPULeaseWalletFilterer creates a new log filterer instance of GPULeaseWallet, bound to a specific deployed contract.
func NewGPULeaseWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*GPULeaseWalletFilterer, error) {
	contract, err := bindGPULeaseWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletFilterer{contract: contract}, nil
}

// bindGPULeaseWallet binds a generic wrapper to an already deployed contract.
func bindGPULeaseWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GPULeaseWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULeaseWallet *GPULeaseWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULeaseWallet.Contract.GPULeaseWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULeaseWallet *GPULeaseWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.GPULeaseWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULeaseWallet *GPULeaseWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.GPULeaseWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GPULeaseWallet *GPULeaseWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GPULeaseWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GPULeaseWallet *GPULeaseWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GPULeaseWallet *GPULeaseWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.contract.Transact(opts, method, params...)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GPULeaseWallet.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GPULeaseWallet.Contract.Balances(&_GPULeaseWallet.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _GPULeaseWallet.Contract.Balances(&_GPULeaseWallet.CallOpts, arg0)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCaller) Credit(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULeaseWallet.contract.Call(opts, &out, "credit")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletSession) Credit() (common.Address, error) {
	return _GPULeaseWallet.Contract.Credit(&_GPULeaseWallet.CallOpts)
}

// Credit is a free data retrieval call binding the contract method 0xa06d083c.
//
// Solidity: function credit() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCallerSession) Credit() (common.Address, error) {
	return _GPULeaseWallet.Contract.Credit(&_GPULeaseWallet.CallOpts)
}

// LeaseManager is a free data retrieval call binding the contract method 0x51ab95c6.
//
// Solidity: function leaseManager() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCaller) LeaseManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULeaseWallet.contract.Call(opts, &out, "leaseManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LeaseManager is a free data retrieval call binding the contract method 0x51ab95c6.
//
// Solidity: function leaseManager() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletSession) LeaseManager() (common.Address, error) {
	return _GPULeaseWallet.Contract.LeaseManager(&_GPULeaseWallet.CallOpts)
}

// LeaseManager is a free data retrieval call binding the contract method 0x51ab95c6.
//
// Solidity: function leaseManager() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCallerSession) LeaseManager() (common.Address, error) {
	return _GPULeaseWallet.Contract.LeaseManager(&_GPULeaseWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GPULeaseWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletSession) Owner() (common.Address, error) {
	return _GPULeaseWallet.Contract.Owner(&_GPULeaseWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GPULeaseWallet *GPULeaseWalletCallerSession) Owner() (common.Address, error) {
	return _GPULeaseWallet.Contract.Owner(&_GPULeaseWallet.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletCaller) UserBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GPULeaseWallet.contract.Call(opts, &out, "userBalance", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletSession) UserBalance(user common.Address) (*big.Int, error) {
	return _GPULeaseWallet.Contract.UserBalance(&_GPULeaseWallet.CallOpts, user)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(address user) view returns(uint256)
func (_GPULeaseWallet *GPULeaseWalletCallerSession) UserBalance(user common.Address) (*big.Int, error) {
	return _GPULeaseWallet.Contract.UserBalance(&_GPULeaseWallet.CallOpts, user)
}

// CreditBalance is a paid mutator transaction binding the contract method 0x5afa3f23.
//
// Solidity: function creditBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) CreditBalance(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "creditBalance", user, amount)
}

// CreditBalance is a paid mutator transaction binding the contract method 0x5afa3f23.
//
// Solidity: function creditBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) CreditBalance(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.CreditBalance(&_GPULeaseWallet.TransactOpts, user, amount)
}

// CreditBalance is a paid mutator transaction binding the contract method 0x5afa3f23.
//
// Solidity: function creditBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) CreditBalance(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.CreditBalance(&_GPULeaseWallet.TransactOpts, user, amount)
}

// DebitBalance is a paid mutator transaction binding the contract method 0x43992a14.
//
// Solidity: function debitBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) DebitBalance(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "debitBalance", user, amount)
}

// DebitBalance is a paid mutator transaction binding the contract method 0x43992a14.
//
// Solidity: function debitBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) DebitBalance(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.DebitBalance(&_GPULeaseWallet.TransactOpts, user, amount)
}

// DebitBalance is a paid mutator transaction binding the contract method 0x43992a14.
//
// Solidity: function debitBalance(address user, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) DebitBalance(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.DebitBalance(&_GPULeaseWallet.TransactOpts, user, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.Deposit(&_GPULeaseWallet.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.Deposit(&_GPULeaseWallet.TransactOpts, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address beneficiary, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) DepositFor(opts *bind.TransactOpts, beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "depositFor", beneficiary, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address beneficiary, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) DepositFor(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.DepositFor(&_GPULeaseWallet.TransactOpts, beneficiary, amount)
}

// DepositFor is a paid mutator transaction binding the contract method 0x2f4f21e2.
//
// Solidity: function depositFor(address beneficiary, uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) DepositFor(beneficiary common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.DepositFor(&_GPULeaseWallet.TransactOpts, beneficiary, amount)
}

// MoveBalance is a paid mutator transaction binding the contract method 0xd1b8631d.
//
// Solidity: function moveBalance(address from, address to) returns(uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletTransactor) MoveBalance(opts *bind.TransactOpts, from common.Address, to common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "moveBalance", from, to)
}

// MoveBalance is a paid mutator transaction binding the contract method 0xd1b8631d.
//
// Solidity: function moveBalance(address from, address to) returns(uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletSession) MoveBalance(from common.Address, to common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.MoveBalance(&_GPULeaseWallet.TransactOpts, from, to)
}

// MoveBalance is a paid mutator transaction binding the contract method 0xd1b8631d.
//
// Solidity: function moveBalance(address from, address to) returns(uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) MoveBalance(from common.Address, to common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.MoveBalance(&_GPULeaseWallet.TransactOpts, from, to)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseWallet *GPULeaseWalletSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.RenounceOwnership(&_GPULeaseWallet.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.RenounceOwnership(&_GPULeaseWallet.TransactOpts)
}

// SetLeaseManager is a paid mutator transaction binding the contract method 0x1dedaca1.
//
// Solidity: function setLeaseManager(address newLeaseManager) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) SetLeaseManager(opts *bind.TransactOpts, newLeaseManager common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "setLeaseManager", newLeaseManager)
}

// SetLeaseManager is a paid mutator transaction binding the contract method 0x1dedaca1.
//
// Solidity: function setLeaseManager(address newLeaseManager) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) SetLeaseManager(newLeaseManager common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.SetLeaseManager(&_GPULeaseWallet.TransactOpts, newLeaseManager)
}

// SetLeaseManager is a paid mutator transaction binding the contract method 0x1dedaca1.
//
// Solidity: function setLeaseManager(address newLeaseManager) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) SetLeaseManager(newLeaseManager common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.SetLeaseManager(&_GPULeaseWallet.TransactOpts, newLeaseManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.TransferOwnership(&_GPULeaseWallet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.TransferOwnership(&_GPULeaseWallet.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.Withdraw(&_GPULeaseWallet.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_GPULeaseWallet *GPULeaseWalletTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _GPULeaseWallet.Contract.Withdraw(&_GPULeaseWallet.TransactOpts, amount)
}

// GPULeaseWalletDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GPULeaseWallet contract.
type GPULeaseWalletDepositIterator struct {
	Event *GPULeaseWalletDeposit // Event containing the contract specifics and raw log

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
func (it *GPULeaseWalletDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseWalletDeposit)
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
		it.Event = new(GPULeaseWalletDeposit)
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
func (it *GPULeaseWalletDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseWalletDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseWalletDeposit represents a Deposit event raised by the GPULeaseWallet contract.
type GPULeaseWalletDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*GPULeaseWalletDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletDepositIterator{contract: _GPULeaseWallet.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GPULeaseWalletDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseWalletDeposit)
				if err := _GPULeaseWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) ParseDeposit(log types.Log) (*GPULeaseWalletDeposit, error) {
	event := new(GPULeaseWalletDeposit)
	if err := _GPULeaseWallet.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseWalletLeaseManagerUpdatedIterator is returned from FilterLeaseManagerUpdated and is used to iterate over the raw logs and unpacked data for LeaseManagerUpdated events raised by the GPULeaseWallet contract.
type GPULeaseWalletLeaseManagerUpdatedIterator struct {
	Event *GPULeaseWalletLeaseManagerUpdated // Event containing the contract specifics and raw log

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
func (it *GPULeaseWalletLeaseManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseWalletLeaseManagerUpdated)
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
		it.Event = new(GPULeaseWalletLeaseManagerUpdated)
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
func (it *GPULeaseWalletLeaseManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseWalletLeaseManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseWalletLeaseManagerUpdated represents a LeaseManagerUpdated event raised by the GPULeaseWallet contract.
type GPULeaseWalletLeaseManagerUpdated struct {
	PreviousManager common.Address
	NewManager      common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLeaseManagerUpdated is a free log retrieval operation binding the contract event 0xad83250657202693cac734fbb1a2fdb15e818fd2017094d63bb8ffde9247049c.
//
// Solidity: event LeaseManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULeaseWallet *GPULeaseWalletFilterer) FilterLeaseManagerUpdated(opts *bind.FilterOpts, previousManager []common.Address, newManager []common.Address) (*GPULeaseWalletLeaseManagerUpdatedIterator, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.FilterLogs(opts, "LeaseManagerUpdated", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletLeaseManagerUpdatedIterator{contract: _GPULeaseWallet.contract, event: "LeaseManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLeaseManagerUpdated is a free log subscription operation binding the contract event 0xad83250657202693cac734fbb1a2fdb15e818fd2017094d63bb8ffde9247049c.
//
// Solidity: event LeaseManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULeaseWallet *GPULeaseWalletFilterer) WatchLeaseManagerUpdated(opts *bind.WatchOpts, sink chan<- *GPULeaseWalletLeaseManagerUpdated, previousManager []common.Address, newManager []common.Address) (event.Subscription, error) {

	var previousManagerRule []interface{}
	for _, previousManagerItem := range previousManager {
		previousManagerRule = append(previousManagerRule, previousManagerItem)
	}
	var newManagerRule []interface{}
	for _, newManagerItem := range newManager {
		newManagerRule = append(newManagerRule, newManagerItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.WatchLogs(opts, "LeaseManagerUpdated", previousManagerRule, newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseWalletLeaseManagerUpdated)
				if err := _GPULeaseWallet.contract.UnpackLog(event, "LeaseManagerUpdated", log); err != nil {
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

// ParseLeaseManagerUpdated is a log parse operation binding the contract event 0xad83250657202693cac734fbb1a2fdb15e818fd2017094d63bb8ffde9247049c.
//
// Solidity: event LeaseManagerUpdated(address indexed previousManager, address indexed newManager)
func (_GPULeaseWallet *GPULeaseWalletFilterer) ParseLeaseManagerUpdated(log types.Log) (*GPULeaseWalletLeaseManagerUpdated, error) {
	event := new(GPULeaseWalletLeaseManagerUpdated)
	if err := _GPULeaseWallet.contract.UnpackLog(event, "LeaseManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseWalletOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GPULeaseWallet contract.
type GPULeaseWalletOwnershipTransferredIterator struct {
	Event *GPULeaseWalletOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GPULeaseWalletOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseWalletOwnershipTransferred)
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
		it.Event = new(GPULeaseWalletOwnershipTransferred)
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
func (it *GPULeaseWalletOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseWalletOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseWalletOwnershipTransferred represents a OwnershipTransferred event raised by the GPULeaseWallet contract.
type GPULeaseWalletOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULeaseWallet *GPULeaseWalletFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GPULeaseWalletOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletOwnershipTransferredIterator{contract: _GPULeaseWallet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GPULeaseWallet *GPULeaseWalletFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GPULeaseWalletOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseWalletOwnershipTransferred)
				if err := _GPULeaseWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GPULeaseWallet *GPULeaseWalletFilterer) ParseOwnershipTransferred(log types.Log) (*GPULeaseWalletOwnershipTransferred, error) {
	event := new(GPULeaseWalletOwnershipTransferred)
	if err := _GPULeaseWallet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GPULeaseWalletWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the GPULeaseWallet contract.
type GPULeaseWalletWithdrawIterator struct {
	Event *GPULeaseWalletWithdraw // Event containing the contract specifics and raw log

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
func (it *GPULeaseWalletWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GPULeaseWalletWithdraw)
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
		it.Event = new(GPULeaseWalletWithdraw)
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
func (it *GPULeaseWalletWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GPULeaseWalletWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GPULeaseWalletWithdraw represents a Withdraw event raised by the GPULeaseWallet contract.
type GPULeaseWalletWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*GPULeaseWalletWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &GPULeaseWalletWithdrawIterator{contract: _GPULeaseWallet.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *GPULeaseWalletWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _GPULeaseWallet.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GPULeaseWalletWithdraw)
				if err := _GPULeaseWallet.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_GPULeaseWallet *GPULeaseWalletFilterer) ParseWithdraw(log types.Log) (*GPULeaseWalletWithdraw, error) {
	event := new(GPULeaseWalletWithdraw)
	if err := _GPULeaseWallet.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
