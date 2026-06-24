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

// LLMFundraisingMetaData contains all meta data concerning the LLMFundraising contract.
var LLMFundraisingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"previousGrade\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"newGrade\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDonated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetShareBps\",\"type\":\"uint256\"}],\"name\":\"BackerGradeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"campaignName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"grade\",\"type\":\"uint8\"}],\"name\":\"BackerRewardMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRaised\",\"type\":\"uint256\"}],\"name\":\"CampaignFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalRaised\",\"type\":\"uint256\"}],\"name\":\"CampaignSucceeded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalDonated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"grade\",\"type\":\"uint8\"}],\"name\":\"Donated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wallet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"backerGrades\",\"outputs\":[{\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"campaignName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkConditions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"expired\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"reached\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"checkState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimBackerReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"donorAt\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"donorInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"participated\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"donatedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetShareBps\",\"type\":\"uint256\"},{\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"grade\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"wasRefunded\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"rewardTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"donorShareBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donors\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"donorsSlice\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"donors_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gpuLeaseWallet\",\"outputs\":[{\"internalType\":\"contractIGPULeaseWallet\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"gradeForDonation\",\"outputs\":[{\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"hasDonated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_campaignId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_templateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_campaignName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_usdc\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_gpuLeaseWallet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_participantRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_metadataRenderer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_campaignOwner\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isTargetReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataRenderer\",\"outputs\":[{\"internalType\":\"contractICampaignMetadataRenderer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"mintBackerRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minted\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextRewardTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"participantRegistry\",\"outputs\":[{\"internalType\":\"contractICampaignParticipantRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refunded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardTokenByDonor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardTokenGrades\",\"outputs\":[{\"internalType\":\"enumLLMFundraising.BackerGrade\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumLLMFundraising.CampaignState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"templateId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdc\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LLMFundraisingABI is the input ABI used to generate the binding from.
// Deprecated: Use LLMFundraisingMetaData.ABI instead.
var LLMFundraisingABI = LLMFundraisingMetaData.ABI

// LLMFundraising is an auto generated Go binding around an Ethereum contract.
type LLMFundraising struct {
	LLMFundraisingCaller     // Read-only binding to the contract
	LLMFundraisingTransactor // Write-only binding to the contract
	LLMFundraisingFilterer   // Log filterer for contract events
}

// LLMFundraisingCaller is an auto generated read-only Go binding around an Ethereum contract.
type LLMFundraisingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LLMFundraisingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LLMFundraisingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LLMFundraisingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LLMFundraisingSession struct {
	Contract     *LLMFundraising   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LLMFundraisingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LLMFundraisingCallerSession struct {
	Contract *LLMFundraisingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// LLMFundraisingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LLMFundraisingTransactorSession struct {
	Contract     *LLMFundraisingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LLMFundraisingRaw is an auto generated low-level Go binding around an Ethereum contract.
type LLMFundraisingRaw struct {
	Contract *LLMFundraising // Generic contract binding to access the raw methods on
}

// LLMFundraisingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LLMFundraisingCallerRaw struct {
	Contract *LLMFundraisingCaller // Generic read-only contract binding to access the raw methods on
}

// LLMFundraisingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LLMFundraisingTransactorRaw struct {
	Contract *LLMFundraisingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLLMFundraising creates a new instance of LLMFundraising, bound to a specific deployed contract.
func NewLLMFundraising(address common.Address, backend bind.ContractBackend) (*LLMFundraising, error) {
	contract, err := bindLLMFundraising(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LLMFundraising{LLMFundraisingCaller: LLMFundraisingCaller{contract: contract}, LLMFundraisingTransactor: LLMFundraisingTransactor{contract: contract}, LLMFundraisingFilterer: LLMFundraisingFilterer{contract: contract}}, nil
}

// NewLLMFundraisingCaller creates a new read-only instance of LLMFundraising, bound to a specific deployed contract.
func NewLLMFundraisingCaller(address common.Address, caller bind.ContractCaller) (*LLMFundraisingCaller, error) {
	contract, err := bindLLMFundraising(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingCaller{contract: contract}, nil
}

// NewLLMFundraisingTransactor creates a new write-only instance of LLMFundraising, bound to a specific deployed contract.
func NewLLMFundraisingTransactor(address common.Address, transactor bind.ContractTransactor) (*LLMFundraisingTransactor, error) {
	contract, err := bindLLMFundraising(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingTransactor{contract: contract}, nil
}

// NewLLMFundraisingFilterer creates a new log filterer instance of LLMFundraising, bound to a specific deployed contract.
func NewLLMFundraisingFilterer(address common.Address, filterer bind.ContractFilterer) (*LLMFundraisingFilterer, error) {
	contract, err := bindLLMFundraising(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFilterer{contract: contract}, nil
}

// bindLLMFundraising binds a generic wrapper to an already deployed contract.
func bindLLMFundraising(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LLMFundraisingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMFundraising *LLMFundraisingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMFundraising.Contract.LLMFundraisingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMFundraising *LLMFundraisingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.Contract.LLMFundraisingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMFundraising *LLMFundraisingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMFundraising.Contract.LLMFundraisingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LLMFundraising *LLMFundraisingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LLMFundraising.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LLMFundraising *LLMFundraisingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LLMFundraising *LLMFundraisingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LLMFundraising.Contract.contract.Transact(opts, method, params...)
}

// BackerGrades is a free data retrieval call binding the contract method 0x0a19f69c.
//
// Solidity: function backerGrades(address ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCaller) BackerGrades(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "backerGrades", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BackerGrades is a free data retrieval call binding the contract method 0x0a19f69c.
//
// Solidity: function backerGrades(address ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingSession) BackerGrades(arg0 common.Address) (uint8, error) {
	return _LLMFundraising.Contract.BackerGrades(&_LLMFundraising.CallOpts, arg0)
}

// BackerGrades is a free data retrieval call binding the contract method 0x0a19f69c.
//
// Solidity: function backerGrades(address ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCallerSession) BackerGrades(arg0 common.Address) (uint8, error) {
	return _LLMFundraising.Contract.BackerGrades(&_LLMFundraising.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.BalanceOf(&_LLMFundraising.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.BalanceOf(&_LLMFundraising.CallOpts, owner)
}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) CampaignId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "campaignId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) CampaignId() (*big.Int, error) {
	return _LLMFundraising.Contract.CampaignId(&_LLMFundraising.CallOpts)
}

// CampaignId is a free data retrieval call binding the contract method 0x8ed5b0fc.
//
// Solidity: function campaignId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) CampaignId() (*big.Int, error) {
	return _LLMFundraising.Contract.CampaignId(&_LLMFundraising.CallOpts)
}

// CampaignName is a free data retrieval call binding the contract method 0x0724fda9.
//
// Solidity: function campaignName() view returns(string)
func (_LLMFundraising *LLMFundraisingCaller) CampaignName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "campaignName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CampaignName is a free data retrieval call binding the contract method 0x0724fda9.
//
// Solidity: function campaignName() view returns(string)
func (_LLMFundraising *LLMFundraisingSession) CampaignName() (string, error) {
	return _LLMFundraising.Contract.CampaignName(&_LLMFundraising.CallOpts)
}

// CampaignName is a free data retrieval call binding the contract method 0x0724fda9.
//
// Solidity: function campaignName() view returns(string)
func (_LLMFundraising *LLMFundraisingCallerSession) CampaignName() (string, error) {
	return _LLMFundraising.Contract.CampaignName(&_LLMFundraising.CallOpts)
}

// CheckConditions is a free data retrieval call binding the contract method 0xdfd09fa5.
//
// Solidity: function checkConditions() view returns(bool expired, bool reached)
func (_LLMFundraising *LLMFundraisingCaller) CheckConditions(opts *bind.CallOpts) (struct {
	Expired bool
	Reached bool
}, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "checkConditions")

	outstruct := new(struct {
		Expired bool
		Reached bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Expired = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Reached = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// CheckConditions is a free data retrieval call binding the contract method 0xdfd09fa5.
//
// Solidity: function checkConditions() view returns(bool expired, bool reached)
func (_LLMFundraising *LLMFundraisingSession) CheckConditions() (struct {
	Expired bool
	Reached bool
}, error) {
	return _LLMFundraising.Contract.CheckConditions(&_LLMFundraising.CallOpts)
}

// CheckConditions is a free data retrieval call binding the contract method 0xdfd09fa5.
//
// Solidity: function checkConditions() view returns(bool expired, bool reached)
func (_LLMFundraising *LLMFundraisingCallerSession) CheckConditions() (struct {
	Expired bool
	Reached bool
}, error) {
	return _LLMFundraising.Contract.CheckConditions(&_LLMFundraising.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) Deadline(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "deadline")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) Deadline() (*big.Int, error) {
	return _LLMFundraising.Contract.Deadline(&_LLMFundraising.CallOpts)
}

// Deadline is a free data retrieval call binding the contract method 0x29dcb0cf.
//
// Solidity: function deadline() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) Deadline() (*big.Int, error) {
	return _LLMFundraising.Contract.Deadline(&_LLMFundraising.CallOpts)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.Donations(&_LLMFundraising.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.Donations(&_LLMFundraising.CallOpts, arg0)
}

// DonorAt is a free data retrieval call binding the contract method 0xaad7c90c.
//
// Solidity: function donorAt(uint256 index) view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) DonorAt(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donorAt", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonorAt is a free data retrieval call binding the contract method 0xaad7c90c.
//
// Solidity: function donorAt(uint256 index) view returns(address)
func (_LLMFundraising *LLMFundraisingSession) DonorAt(index *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.DonorAt(&_LLMFundraising.CallOpts, index)
}

// DonorAt is a free data retrieval call binding the contract method 0xaad7c90c.
//
// Solidity: function donorAt(uint256 index) view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) DonorAt(index *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.DonorAt(&_LLMFundraising.CallOpts, index)
}

// DonorInfo is a free data retrieval call binding the contract method 0xa729dd7c.
//
// Solidity: function donorInfo(address donor) view returns(bool participated, uint256 donatedAmount, uint256 targetShareBps, uint8 grade, bool wasRefunded, uint256 rewardTokenId)
func (_LLMFundraising *LLMFundraisingCaller) DonorInfo(opts *bind.CallOpts, donor common.Address) (struct {
	Participated   bool
	DonatedAmount  *big.Int
	TargetShareBps *big.Int
	Grade          uint8
	WasRefunded    bool
	RewardTokenId  *big.Int
}, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donorInfo", donor)

	outstruct := new(struct {
		Participated   bool
		DonatedAmount  *big.Int
		TargetShareBps *big.Int
		Grade          uint8
		WasRefunded    bool
		RewardTokenId  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Participated = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.DonatedAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TargetShareBps = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Grade = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.WasRefunded = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.RewardTokenId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DonorInfo is a free data retrieval call binding the contract method 0xa729dd7c.
//
// Solidity: function donorInfo(address donor) view returns(bool participated, uint256 donatedAmount, uint256 targetShareBps, uint8 grade, bool wasRefunded, uint256 rewardTokenId)
func (_LLMFundraising *LLMFundraisingSession) DonorInfo(donor common.Address) (struct {
	Participated   bool
	DonatedAmount  *big.Int
	TargetShareBps *big.Int
	Grade          uint8
	WasRefunded    bool
	RewardTokenId  *big.Int
}, error) {
	return _LLMFundraising.Contract.DonorInfo(&_LLMFundraising.CallOpts, donor)
}

// DonorInfo is a free data retrieval call binding the contract method 0xa729dd7c.
//
// Solidity: function donorInfo(address donor) view returns(bool participated, uint256 donatedAmount, uint256 targetShareBps, uint8 grade, bool wasRefunded, uint256 rewardTokenId)
func (_LLMFundraising *LLMFundraisingCallerSession) DonorInfo(donor common.Address) (struct {
	Participated   bool
	DonatedAmount  *big.Int
	TargetShareBps *big.Int
	Grade          uint8
	WasRefunded    bool
	RewardTokenId  *big.Int
}, error) {
	return _LLMFundraising.Contract.DonorInfo(&_LLMFundraising.CallOpts, donor)
}

// DonorShareBps is a free data retrieval call binding the contract method 0x77d6444c.
//
// Solidity: function donorShareBps(address donor) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) DonorShareBps(opts *bind.CallOpts, donor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donorShareBps", donor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonorShareBps is a free data retrieval call binding the contract method 0x77d6444c.
//
// Solidity: function donorShareBps(address donor) view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) DonorShareBps(donor common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.DonorShareBps(&_LLMFundraising.CallOpts, donor)
}

// DonorShareBps is a free data retrieval call binding the contract method 0x77d6444c.
//
// Solidity: function donorShareBps(address donor) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) DonorShareBps(donor common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.DonorShareBps(&_LLMFundraising.CallOpts, donor)
}

// Donors is a free data retrieval call binding the contract method 0x3087110a.
//
// Solidity: function donors() view returns(address[])
func (_LLMFundraising *LLMFundraisingCaller) Donors(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donors")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// Donors is a free data retrieval call binding the contract method 0x3087110a.
//
// Solidity: function donors() view returns(address[])
func (_LLMFundraising *LLMFundraisingSession) Donors() ([]common.Address, error) {
	return _LLMFundraising.Contract.Donors(&_LLMFundraising.CallOpts)
}

// Donors is a free data retrieval call binding the contract method 0x3087110a.
//
// Solidity: function donors() view returns(address[])
func (_LLMFundraising *LLMFundraisingCallerSession) Donors() ([]common.Address, error) {
	return _LLMFundraising.Contract.Donors(&_LLMFundraising.CallOpts)
}

// DonorsCount is a free data retrieval call binding the contract method 0x9ff534dc.
//
// Solidity: function donorsCount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) DonorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DonorsCount is a free data retrieval call binding the contract method 0x9ff534dc.
//
// Solidity: function donorsCount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) DonorsCount() (*big.Int, error) {
	return _LLMFundraising.Contract.DonorsCount(&_LLMFundraising.CallOpts)
}

// DonorsCount is a free data retrieval call binding the contract method 0x9ff534dc.
//
// Solidity: function donorsCount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) DonorsCount() (*big.Int, error) {
	return _LLMFundraising.Contract.DonorsCount(&_LLMFundraising.CallOpts)
}

// DonorsSlice is a free data retrieval call binding the contract method 0xa263d772.
//
// Solidity: function donorsSlice(uint256 offset, uint256 limit) view returns(address[] donors_)
func (_LLMFundraising *LLMFundraisingCaller) DonorsSlice(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "donorsSlice", offset, limit)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DonorsSlice is a free data retrieval call binding the contract method 0xa263d772.
//
// Solidity: function donorsSlice(uint256 offset, uint256 limit) view returns(address[] donors_)
func (_LLMFundraising *LLMFundraisingSession) DonorsSlice(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _LLMFundraising.Contract.DonorsSlice(&_LLMFundraising.CallOpts, offset, limit)
}

// DonorsSlice is a free data retrieval call binding the contract method 0xa263d772.
//
// Solidity: function donorsSlice(uint256 offset, uint256 limit) view returns(address[] donors_)
func (_LLMFundraising *LLMFundraisingCallerSession) DonorsSlice(offset *big.Int, limit *big.Int) ([]common.Address, error) {
	return _LLMFundraising.Contract.DonorsSlice(&_LLMFundraising.CallOpts, offset, limit)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) Duration() (*big.Int, error) {
	return _LLMFundraising.Contract.Duration(&_LLMFundraising.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) Duration() (*big.Int, error) {
	return _LLMFundraising.Contract.Duration(&_LLMFundraising.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.GetApproved(&_LLMFundraising.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.GetApproved(&_LLMFundraising.CallOpts, tokenId)
}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) GpuLeaseWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "gpuLeaseWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraising *LLMFundraisingSession) GpuLeaseWallet() (common.Address, error) {
	return _LLMFundraising.Contract.GpuLeaseWallet(&_LLMFundraising.CallOpts)
}

// GpuLeaseWallet is a free data retrieval call binding the contract method 0x2b350856.
//
// Solidity: function gpuLeaseWallet() view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) GpuLeaseWallet() (common.Address, error) {
	return _LLMFundraising.Contract.GpuLeaseWallet(&_LLMFundraising.CallOpts)
}

// GradeForDonation is a free data retrieval call binding the contract method 0xbffd11bb.
//
// Solidity: function gradeForDonation(address donor) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCaller) GradeForDonation(opts *bind.CallOpts, donor common.Address) (uint8, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "gradeForDonation", donor)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GradeForDonation is a free data retrieval call binding the contract method 0xbffd11bb.
//
// Solidity: function gradeForDonation(address donor) view returns(uint8)
func (_LLMFundraising *LLMFundraisingSession) GradeForDonation(donor common.Address) (uint8, error) {
	return _LLMFundraising.Contract.GradeForDonation(&_LLMFundraising.CallOpts, donor)
}

// GradeForDonation is a free data retrieval call binding the contract method 0xbffd11bb.
//
// Solidity: function gradeForDonation(address donor) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCallerSession) GradeForDonation(donor common.Address) (uint8, error) {
	return _LLMFundraising.Contract.GradeForDonation(&_LLMFundraising.CallOpts, donor)
}

// HasDonated is a free data retrieval call binding the contract method 0x17294a11.
//
// Solidity: function hasDonated(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) HasDonated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "hasDonated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasDonated is a free data retrieval call binding the contract method 0x17294a11.
//
// Solidity: function hasDonated(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) HasDonated(arg0 common.Address) (bool, error) {
	return _LLMFundraising.Contract.HasDonated(&_LLMFundraising.CallOpts, arg0)
}

// HasDonated is a free data retrieval call binding the contract method 0x17294a11.
//
// Solidity: function hasDonated(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) HasDonated(arg0 common.Address) (bool, error) {
	return _LLMFundraising.Contract.HasDonated(&_LLMFundraising.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) Initialized() (bool, error) {
	return _LLMFundraising.Contract.Initialized(&_LLMFundraising.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) Initialized() (bool, error) {
	return _LLMFundraising.Contract.Initialized(&_LLMFundraising.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LLMFundraising.Contract.IsApprovedForAll(&_LLMFundraising.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _LLMFundraising.Contract.IsApprovedForAll(&_LLMFundraising.CallOpts, owner, operator)
}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) IsExpired(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "isExpired")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) IsExpired() (bool, error) {
	return _LLMFundraising.Contract.IsExpired(&_LLMFundraising.CallOpts)
}

// IsExpired is a free data retrieval call binding the contract method 0x2f13b60c.
//
// Solidity: function isExpired() view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) IsExpired() (bool, error) {
	return _LLMFundraising.Contract.IsExpired(&_LLMFundraising.CallOpts)
}

// IsTargetReached is a free data retrieval call binding the contract method 0x5e098f05.
//
// Solidity: function isTargetReached() view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) IsTargetReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "isTargetReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTargetReached is a free data retrieval call binding the contract method 0x5e098f05.
//
// Solidity: function isTargetReached() view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) IsTargetReached() (bool, error) {
	return _LLMFundraising.Contract.IsTargetReached(&_LLMFundraising.CallOpts)
}

// IsTargetReached is a free data retrieval call binding the contract method 0x5e098f05.
//
// Solidity: function isTargetReached() view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) IsTargetReached() (bool, error) {
	return _LLMFundraising.Contract.IsTargetReached(&_LLMFundraising.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) MetadataRenderer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "metadataRenderer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraising *LLMFundraisingSession) MetadataRenderer() (common.Address, error) {
	return _LLMFundraising.Contract.MetadataRenderer(&_LLMFundraising.CallOpts)
}

// MetadataRenderer is a free data retrieval call binding the contract method 0x70319970.
//
// Solidity: function metadataRenderer() view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) MetadataRenderer() (common.Address, error) {
	return _LLMFundraising.Contract.MetadataRenderer(&_LLMFundraising.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LLMFundraising *LLMFundraisingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LLMFundraising *LLMFundraisingSession) Name() (string, error) {
	return _LLMFundraising.Contract.Name(&_LLMFundraising.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_LLMFundraising *LLMFundraisingCallerSession) Name() (string, error) {
	return _LLMFundraising.Contract.Name(&_LLMFundraising.CallOpts)
}

// NextRewardTokenId is a free data retrieval call binding the contract method 0x35909603.
//
// Solidity: function nextRewardTokenId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) NextRewardTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "nextRewardTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRewardTokenId is a free data retrieval call binding the contract method 0x35909603.
//
// Solidity: function nextRewardTokenId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) NextRewardTokenId() (*big.Int, error) {
	return _LLMFundraising.Contract.NextRewardTokenId(&_LLMFundraising.CallOpts)
}

// NextRewardTokenId is a free data retrieval call binding the contract method 0x35909603.
//
// Solidity: function nextRewardTokenId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) NextRewardTokenId() (*big.Int, error) {
	return _LLMFundraising.Contract.NextRewardTokenId(&_LLMFundraising.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraising *LLMFundraisingSession) Owner() (common.Address, error) {
	return _LLMFundraising.Contract.Owner(&_LLMFundraising.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) Owner() (common.Address, error) {
	return _LLMFundraising.Contract.Owner(&_LLMFundraising.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.OwnerOf(&_LLMFundraising.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _LLMFundraising.Contract.OwnerOf(&_LLMFundraising.CallOpts, tokenId)
}

// ParticipantRegistry is a free data retrieval call binding the contract method 0x9e9ca591.
//
// Solidity: function participantRegistry() view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) ParticipantRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "participantRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ParticipantRegistry is a free data retrieval call binding the contract method 0x9e9ca591.
//
// Solidity: function participantRegistry() view returns(address)
func (_LLMFundraising *LLMFundraisingSession) ParticipantRegistry() (common.Address, error) {
	return _LLMFundraising.Contract.ParticipantRegistry(&_LLMFundraising.CallOpts)
}

// ParticipantRegistry is a free data retrieval call binding the contract method 0x9e9ca591.
//
// Solidity: function participantRegistry() view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) ParticipantRegistry() (common.Address, error) {
	return _LLMFundraising.Contract.ParticipantRegistry(&_LLMFundraising.CallOpts)
}

// Refunded is a free data retrieval call binding the contract method 0xc033a490.
//
// Solidity: function refunded(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) Refunded(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "refunded", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Refunded is a free data retrieval call binding the contract method 0xc033a490.
//
// Solidity: function refunded(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) Refunded(arg0 common.Address) (bool, error) {
	return _LLMFundraising.Contract.Refunded(&_LLMFundraising.CallOpts, arg0)
}

// Refunded is a free data retrieval call binding the contract method 0xc033a490.
//
// Solidity: function refunded(address ) view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) Refunded(arg0 common.Address) (bool, error) {
	return _LLMFundraising.Contract.Refunded(&_LLMFundraising.CallOpts, arg0)
}

// RewardTokenByDonor is a free data retrieval call binding the contract method 0x8cc05bec.
//
// Solidity: function rewardTokenByDonor(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) RewardTokenByDonor(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "rewardTokenByDonor", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardTokenByDonor is a free data retrieval call binding the contract method 0x8cc05bec.
//
// Solidity: function rewardTokenByDonor(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) RewardTokenByDonor(arg0 common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.RewardTokenByDonor(&_LLMFundraising.CallOpts, arg0)
}

// RewardTokenByDonor is a free data retrieval call binding the contract method 0x8cc05bec.
//
// Solidity: function rewardTokenByDonor(address ) view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) RewardTokenByDonor(arg0 common.Address) (*big.Int, error) {
	return _LLMFundraising.Contract.RewardTokenByDonor(&_LLMFundraising.CallOpts, arg0)
}

// RewardTokenGrades is a free data retrieval call binding the contract method 0x7592e888.
//
// Solidity: function rewardTokenGrades(uint256 ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCaller) RewardTokenGrades(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "rewardTokenGrades", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RewardTokenGrades is a free data retrieval call binding the contract method 0x7592e888.
//
// Solidity: function rewardTokenGrades(uint256 ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingSession) RewardTokenGrades(arg0 *big.Int) (uint8, error) {
	return _LLMFundraising.Contract.RewardTokenGrades(&_LLMFundraising.CallOpts, arg0)
}

// RewardTokenGrades is a free data retrieval call binding the contract method 0x7592e888.
//
// Solidity: function rewardTokenGrades(uint256 ) view returns(uint8)
func (_LLMFundraising *LLMFundraisingCallerSession) RewardTokenGrades(arg0 *big.Int) (uint8, error) {
	return _LLMFundraising.Contract.RewardTokenGrades(&_LLMFundraising.CallOpts, arg0)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) StartTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "startTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) StartTimestamp() (*big.Int, error) {
	return _LLMFundraising.Contract.StartTimestamp(&_LLMFundraising.CallOpts)
}

// StartTimestamp is a free data retrieval call binding the contract method 0xe6fd48bc.
//
// Solidity: function startTimestamp() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) StartTimestamp() (*big.Int, error) {
	return _LLMFundraising.Contract.StartTimestamp(&_LLMFundraising.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_LLMFundraising *LLMFundraisingCaller) State(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "state")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_LLMFundraising *LLMFundraisingSession) State() (uint8, error) {
	return _LLMFundraising.Contract.State(&_LLMFundraising.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_LLMFundraising *LLMFundraisingCallerSession) State() (uint8, error) {
	return _LLMFundraising.Contract.State(&_LLMFundraising.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMFundraising *LLMFundraisingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMFundraising *LLMFundraisingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LLMFundraising.Contract.SupportsInterface(&_LLMFundraising.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LLMFundraising *LLMFundraisingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LLMFundraising.Contract.SupportsInterface(&_LLMFundraising.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LLMFundraising *LLMFundraisingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LLMFundraising *LLMFundraisingSession) Symbol() (string, error) {
	return _LLMFundraising.Contract.Symbol(&_LLMFundraising.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_LLMFundraising *LLMFundraisingCallerSession) Symbol() (string, error) {
	return _LLMFundraising.Contract.Symbol(&_LLMFundraising.CallOpts)
}

// TargetAmount is a free data retrieval call binding the contract method 0x953b8fb8.
//
// Solidity: function targetAmount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) TargetAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "targetAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetAmount is a free data retrieval call binding the contract method 0x953b8fb8.
//
// Solidity: function targetAmount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) TargetAmount() (*big.Int, error) {
	return _LLMFundraising.Contract.TargetAmount(&_LLMFundraising.CallOpts)
}

// TargetAmount is a free data retrieval call binding the contract method 0x953b8fb8.
//
// Solidity: function targetAmount() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) TargetAmount() (*big.Int, error) {
	return _LLMFundraising.Contract.TargetAmount(&_LLMFundraising.CallOpts)
}

// TemplateId is a free data retrieval call binding the contract method 0x7aa77f29.
//
// Solidity: function templateId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) TemplateId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "templateId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TemplateId is a free data retrieval call binding the contract method 0x7aa77f29.
//
// Solidity: function templateId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) TemplateId() (*big.Int, error) {
	return _LLMFundraising.Contract.TemplateId(&_LLMFundraising.CallOpts)
}

// TemplateId is a free data retrieval call binding the contract method 0x7aa77f29.
//
// Solidity: function templateId() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) TemplateId() (*big.Int, error) {
	return _LLMFundraising.Contract.TemplateId(&_LLMFundraising.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMFundraising *LLMFundraisingCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMFundraising *LLMFundraisingSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LLMFundraising.Contract.TokenURI(&_LLMFundraising.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_LLMFundraising *LLMFundraisingCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _LLMFundraising.Contract.TokenURI(&_LLMFundraising.CallOpts, tokenId)
}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCaller) TotalRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "totalRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_LLMFundraising *LLMFundraisingSession) TotalRaised() (*big.Int, error) {
	return _LLMFundraising.Contract.TotalRaised(&_LLMFundraising.CallOpts)
}

// TotalRaised is a free data retrieval call binding the contract method 0xc5c4744c.
//
// Solidity: function totalRaised() view returns(uint256)
func (_LLMFundraising *LLMFundraisingCallerSession) TotalRaised() (*big.Int, error) {
	return _LLMFundraising.Contract.TotalRaised(&_LLMFundraising.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraising *LLMFundraisingCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LLMFundraising.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraising *LLMFundraisingSession) Usdc() (common.Address, error) {
	return _LLMFundraising.Contract.Usdc(&_LLMFundraising.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_LLMFundraising *LLMFundraisingCallerSession) Usdc() (common.Address, error) {
	return _LLMFundraising.Contract.Usdc(&_LLMFundraising.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Approve(&_LLMFundraising.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Approve(&_LLMFundraising.TransactOpts, to, tokenId)
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_LLMFundraising *LLMFundraisingTransactor) CheckState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "checkState")
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_LLMFundraising *LLMFundraisingSession) CheckState() (*types.Transaction, error) {
	return _LLMFundraising.Contract.CheckState(&_LLMFundraising.TransactOpts)
}

// CheckState is a paid mutator transaction binding the contract method 0x96dfcbea.
//
// Solidity: function checkState() returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) CheckState() (*types.Transaction, error) {
	return _LLMFundraising.Contract.CheckState(&_LLMFundraising.TransactOpts)
}

// ClaimBackerReward is a paid mutator transaction binding the contract method 0x7b338e14.
//
// Solidity: function claimBackerReward() returns(uint256 tokenId)
func (_LLMFundraising *LLMFundraisingTransactor) ClaimBackerReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "claimBackerReward")
}

// ClaimBackerReward is a paid mutator transaction binding the contract method 0x7b338e14.
//
// Solidity: function claimBackerReward() returns(uint256 tokenId)
func (_LLMFundraising *LLMFundraisingSession) ClaimBackerReward() (*types.Transaction, error) {
	return _LLMFundraising.Contract.ClaimBackerReward(&_LLMFundraising.TransactOpts)
}

// ClaimBackerReward is a paid mutator transaction binding the contract method 0x7b338e14.
//
// Solidity: function claimBackerReward() returns(uint256 tokenId)
func (_LLMFundraising *LLMFundraisingTransactorSession) ClaimBackerReward() (*types.Transaction, error) {
	return _LLMFundraising.Contract.ClaimBackerReward(&_LLMFundraising.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 amount) returns()
func (_LLMFundraising *LLMFundraisingTransactor) Donate(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "donate", amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 amount) returns()
func (_LLMFundraising *LLMFundraisingSession) Donate(amount *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Donate(&_LLMFundraising.TransactOpts, amount)
}

// Donate is a paid mutator transaction binding the contract method 0xf14faf6f.
//
// Solidity: function donate(uint256 amount) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) Donate(amount *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Donate(&_LLMFundraising.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x19ae4f0d.
//
// Solidity: function initialize(uint256 _campaignId, uint256 _targetAmount, uint256 _duration, uint256 _startTimestamp, uint256 _templateId, string _campaignName, address _usdc, address _gpuLeaseWallet, address _participantRegistry, address _metadataRenderer, address _campaignOwner) returns()
func (_LLMFundraising *LLMFundraisingTransactor) Initialize(opts *bind.TransactOpts, _campaignId *big.Int, _targetAmount *big.Int, _duration *big.Int, _startTimestamp *big.Int, _templateId *big.Int, _campaignName string, _usdc common.Address, _gpuLeaseWallet common.Address, _participantRegistry common.Address, _metadataRenderer common.Address, _campaignOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "initialize", _campaignId, _targetAmount, _duration, _startTimestamp, _templateId, _campaignName, _usdc, _gpuLeaseWallet, _participantRegistry, _metadataRenderer, _campaignOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x19ae4f0d.
//
// Solidity: function initialize(uint256 _campaignId, uint256 _targetAmount, uint256 _duration, uint256 _startTimestamp, uint256 _templateId, string _campaignName, address _usdc, address _gpuLeaseWallet, address _participantRegistry, address _metadataRenderer, address _campaignOwner) returns()
func (_LLMFundraising *LLMFundraisingSession) Initialize(_campaignId *big.Int, _targetAmount *big.Int, _duration *big.Int, _startTimestamp *big.Int, _templateId *big.Int, _campaignName string, _usdc common.Address, _gpuLeaseWallet common.Address, _participantRegistry common.Address, _metadataRenderer common.Address, _campaignOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Initialize(&_LLMFundraising.TransactOpts, _campaignId, _targetAmount, _duration, _startTimestamp, _templateId, _campaignName, _usdc, _gpuLeaseWallet, _participantRegistry, _metadataRenderer, _campaignOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0x19ae4f0d.
//
// Solidity: function initialize(uint256 _campaignId, uint256 _targetAmount, uint256 _duration, uint256 _startTimestamp, uint256 _templateId, string _campaignName, address _usdc, address _gpuLeaseWallet, address _participantRegistry, address _metadataRenderer, address _campaignOwner) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) Initialize(_campaignId *big.Int, _targetAmount *big.Int, _duration *big.Int, _startTimestamp *big.Int, _templateId *big.Int, _campaignName string, _usdc common.Address, _gpuLeaseWallet common.Address, _participantRegistry common.Address, _metadataRenderer common.Address, _campaignOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.Contract.Initialize(&_LLMFundraising.TransactOpts, _campaignId, _targetAmount, _duration, _startTimestamp, _templateId, _campaignName, _usdc, _gpuLeaseWallet, _participantRegistry, _metadataRenderer, _campaignOwner)
}

// MintBackerRewards is a paid mutator transaction binding the contract method 0x10641782.
//
// Solidity: function mintBackerRewards(uint256 offset, uint256 limit) returns(uint256 minted)
func (_LLMFundraising *LLMFundraisingTransactor) MintBackerRewards(opts *bind.TransactOpts, offset *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "mintBackerRewards", offset, limit)
}

// MintBackerRewards is a paid mutator transaction binding the contract method 0x10641782.
//
// Solidity: function mintBackerRewards(uint256 offset, uint256 limit) returns(uint256 minted)
func (_LLMFundraising *LLMFundraisingSession) MintBackerRewards(offset *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.MintBackerRewards(&_LLMFundraising.TransactOpts, offset, limit)
}

// MintBackerRewards is a paid mutator transaction binding the contract method 0x10641782.
//
// Solidity: function mintBackerRewards(uint256 offset, uint256 limit) returns(uint256 minted)
func (_LLMFundraising *LLMFundraisingTransactorSession) MintBackerRewards(offset *big.Int, limit *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.MintBackerRewards(&_LLMFundraising.TransactOpts, offset, limit)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_LLMFundraising *LLMFundraisingTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_LLMFundraising *LLMFundraisingSession) Refund() (*types.Transaction, error) {
	return _LLMFundraising.Contract.Refund(&_LLMFundraising.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) Refund() (*types.Transaction, error) {
	return _LLMFundraising.Contract.Refund(&_LLMFundraising.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraising *LLMFundraisingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraising *LLMFundraisingSession) RenounceOwnership() (*types.Transaction, error) {
	return _LLMFundraising.Contract.RenounceOwnership(&_LLMFundraising.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LLMFundraising.Contract.RenounceOwnership(&_LLMFundraising.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SafeTransferFrom(&_LLMFundraising.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SafeTransferFrom(&_LLMFundraising.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMFundraising *LLMFundraisingTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMFundraising *LLMFundraisingSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SafeTransferFrom0(&_LLMFundraising.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SafeTransferFrom0(&_LLMFundraising.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMFundraising *LLMFundraisingTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMFundraising *LLMFundraisingSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SetApprovalForAll(&_LLMFundraising.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _LLMFundraising.Contract.SetApprovalForAll(&_LLMFundraising.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.TransferFrom(&_LLMFundraising.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _LLMFundraising.Contract.TransferFrom(&_LLMFundraising.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraising *LLMFundraisingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraising *LLMFundraisingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.Contract.TransferOwnership(&_LLMFundraising.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LLMFundraising *LLMFundraisingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LLMFundraising.Contract.TransferOwnership(&_LLMFundraising.TransactOpts, newOwner)
}

// LLMFundraisingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LLMFundraising contract.
type LLMFundraisingApprovalIterator struct {
	Event *LLMFundraisingApproval // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingApproval)
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
		it.Event = new(LLMFundraisingApproval)
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
func (it *LLMFundraisingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingApproval represents a Approval event raised by the LLMFundraising contract.
type LLMFundraisingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*LLMFundraisingApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingApprovalIterator{contract: _LLMFundraising.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LLMFundraisingApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingApproval)
				if err := _LLMFundraising.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) ParseApproval(log types.Log) (*LLMFundraisingApproval, error) {
	event := new(LLMFundraisingApproval)
	if err := _LLMFundraising.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the LLMFundraising contract.
type LLMFundraisingApprovalForAllIterator struct {
	Event *LLMFundraisingApprovalForAll // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingApprovalForAll)
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
		it.Event = new(LLMFundraisingApprovalForAll)
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
func (it *LLMFundraisingApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingApprovalForAll represents a ApprovalForAll event raised by the LLMFundraising contract.
type LLMFundraisingApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMFundraising *LLMFundraisingFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*LLMFundraisingApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingApprovalForAllIterator{contract: _LLMFundraising.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMFundraising *LLMFundraisingFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LLMFundraisingApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingApprovalForAll)
				if err := _LLMFundraising.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_LLMFundraising *LLMFundraisingFilterer) ParseApprovalForAll(log types.Log) (*LLMFundraisingApprovalForAll, error) {
	event := new(LLMFundraisingApprovalForAll)
	if err := _LLMFundraising.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingBackerGradeUpdatedIterator is returned from FilterBackerGradeUpdated and is used to iterate over the raw logs and unpacked data for BackerGradeUpdated events raised by the LLMFundraising contract.
type LLMFundraisingBackerGradeUpdatedIterator struct {
	Event *LLMFundraisingBackerGradeUpdated // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingBackerGradeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingBackerGradeUpdated)
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
		it.Event = new(LLMFundraisingBackerGradeUpdated)
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
func (it *LLMFundraisingBackerGradeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingBackerGradeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingBackerGradeUpdated represents a BackerGradeUpdated event raised by the LLMFundraising contract.
type LLMFundraisingBackerGradeUpdated struct {
	Donor          common.Address
	PreviousGrade  uint8
	NewGrade       uint8
	TotalDonated   *big.Int
	TargetShareBps *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBackerGradeUpdated is a free log retrieval operation binding the contract event 0x9edbeab148e39aabe0c088d94c9704b598cce01e40da413e8da79013a4849efb.
//
// Solidity: event BackerGradeUpdated(address indexed donor, uint8 previousGrade, uint8 newGrade, uint256 totalDonated, uint256 targetShareBps)
func (_LLMFundraising *LLMFundraisingFilterer) FilterBackerGradeUpdated(opts *bind.FilterOpts, donor []common.Address) (*LLMFundraisingBackerGradeUpdatedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "BackerGradeUpdated", donorRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingBackerGradeUpdatedIterator{contract: _LLMFundraising.contract, event: "BackerGradeUpdated", logs: logs, sub: sub}, nil
}

// WatchBackerGradeUpdated is a free log subscription operation binding the contract event 0x9edbeab148e39aabe0c088d94c9704b598cce01e40da413e8da79013a4849efb.
//
// Solidity: event BackerGradeUpdated(address indexed donor, uint8 previousGrade, uint8 newGrade, uint256 totalDonated, uint256 targetShareBps)
func (_LLMFundraising *LLMFundraisingFilterer) WatchBackerGradeUpdated(opts *bind.WatchOpts, sink chan<- *LLMFundraisingBackerGradeUpdated, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "BackerGradeUpdated", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingBackerGradeUpdated)
				if err := _LLMFundraising.contract.UnpackLog(event, "BackerGradeUpdated", log); err != nil {
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

// ParseBackerGradeUpdated is a log parse operation binding the contract event 0x9edbeab148e39aabe0c088d94c9704b598cce01e40da413e8da79013a4849efb.
//
// Solidity: event BackerGradeUpdated(address indexed donor, uint8 previousGrade, uint8 newGrade, uint256 totalDonated, uint256 targetShareBps)
func (_LLMFundraising *LLMFundraisingFilterer) ParseBackerGradeUpdated(log types.Log) (*LLMFundraisingBackerGradeUpdated, error) {
	event := new(LLMFundraisingBackerGradeUpdated)
	if err := _LLMFundraising.contract.UnpackLog(event, "BackerGradeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingBackerRewardMintedIterator is returned from FilterBackerRewardMinted and is used to iterate over the raw logs and unpacked data for BackerRewardMinted events raised by the LLMFundraising contract.
type LLMFundraisingBackerRewardMintedIterator struct {
	Event *LLMFundraisingBackerRewardMinted // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingBackerRewardMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingBackerRewardMinted)
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
		it.Event = new(LLMFundraisingBackerRewardMinted)
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
func (it *LLMFundraisingBackerRewardMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingBackerRewardMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingBackerRewardMinted represents a BackerRewardMinted event raised by the LLMFundraising contract.
type LLMFundraisingBackerRewardMinted struct {
	Donor        common.Address
	TokenId      *big.Int
	CampaignName string
	Grade        uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBackerRewardMinted is a free log retrieval operation binding the contract event 0xacc61df05a1322ee40c5459cddfcee8a7dfe0ec8d93272de2f0d978ad481f03e.
//
// Solidity: event BackerRewardMinted(address indexed donor, uint256 indexed tokenId, string campaignName, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) FilterBackerRewardMinted(opts *bind.FilterOpts, donor []common.Address, tokenId []*big.Int) (*LLMFundraisingBackerRewardMintedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "BackerRewardMinted", donorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingBackerRewardMintedIterator{contract: _LLMFundraising.contract, event: "BackerRewardMinted", logs: logs, sub: sub}, nil
}

// WatchBackerRewardMinted is a free log subscription operation binding the contract event 0xacc61df05a1322ee40c5459cddfcee8a7dfe0ec8d93272de2f0d978ad481f03e.
//
// Solidity: event BackerRewardMinted(address indexed donor, uint256 indexed tokenId, string campaignName, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) WatchBackerRewardMinted(opts *bind.WatchOpts, sink chan<- *LLMFundraisingBackerRewardMinted, donor []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "BackerRewardMinted", donorRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingBackerRewardMinted)
				if err := _LLMFundraising.contract.UnpackLog(event, "BackerRewardMinted", log); err != nil {
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

// ParseBackerRewardMinted is a log parse operation binding the contract event 0xacc61df05a1322ee40c5459cddfcee8a7dfe0ec8d93272de2f0d978ad481f03e.
//
// Solidity: event BackerRewardMinted(address indexed donor, uint256 indexed tokenId, string campaignName, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) ParseBackerRewardMinted(log types.Log) (*LLMFundraisingBackerRewardMinted, error) {
	event := new(LLMFundraisingBackerRewardMinted)
	if err := _LLMFundraising.contract.UnpackLog(event, "BackerRewardMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingCampaignFailedIterator is returned from FilterCampaignFailed and is used to iterate over the raw logs and unpacked data for CampaignFailed events raised by the LLMFundraising contract.
type LLMFundraisingCampaignFailedIterator struct {
	Event *LLMFundraisingCampaignFailed // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingCampaignFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingCampaignFailed)
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
		it.Event = new(LLMFundraisingCampaignFailed)
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
func (it *LLMFundraisingCampaignFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingCampaignFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingCampaignFailed represents a CampaignFailed event raised by the LLMFundraising contract.
type LLMFundraisingCampaignFailed struct {
	TotalRaised *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCampaignFailed is a free log retrieval operation binding the contract event 0x7b1418b19a03b176f641639ee79855da661fa88a08377ecd32ad6d38152b9536.
//
// Solidity: event CampaignFailed(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) FilterCampaignFailed(opts *bind.FilterOpts) (*LLMFundraisingCampaignFailedIterator, error) {

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "CampaignFailed")
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingCampaignFailedIterator{contract: _LLMFundraising.contract, event: "CampaignFailed", logs: logs, sub: sub}, nil
}

// WatchCampaignFailed is a free log subscription operation binding the contract event 0x7b1418b19a03b176f641639ee79855da661fa88a08377ecd32ad6d38152b9536.
//
// Solidity: event CampaignFailed(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) WatchCampaignFailed(opts *bind.WatchOpts, sink chan<- *LLMFundraisingCampaignFailed) (event.Subscription, error) {

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "CampaignFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingCampaignFailed)
				if err := _LLMFundraising.contract.UnpackLog(event, "CampaignFailed", log); err != nil {
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

// ParseCampaignFailed is a log parse operation binding the contract event 0x7b1418b19a03b176f641639ee79855da661fa88a08377ecd32ad6d38152b9536.
//
// Solidity: event CampaignFailed(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) ParseCampaignFailed(log types.Log) (*LLMFundraisingCampaignFailed, error) {
	event := new(LLMFundraisingCampaignFailed)
	if err := _LLMFundraising.contract.UnpackLog(event, "CampaignFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingCampaignSucceededIterator is returned from FilterCampaignSucceeded and is used to iterate over the raw logs and unpacked data for CampaignSucceeded events raised by the LLMFundraising contract.
type LLMFundraisingCampaignSucceededIterator struct {
	Event *LLMFundraisingCampaignSucceeded // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingCampaignSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingCampaignSucceeded)
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
		it.Event = new(LLMFundraisingCampaignSucceeded)
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
func (it *LLMFundraisingCampaignSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingCampaignSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingCampaignSucceeded represents a CampaignSucceeded event raised by the LLMFundraising contract.
type LLMFundraisingCampaignSucceeded struct {
	TotalRaised *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCampaignSucceeded is a free log retrieval operation binding the contract event 0x2b3f27cb35b5581f9baeb414c2592d8fc9c9067287476ca5624392320e9bf68e.
//
// Solidity: event CampaignSucceeded(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) FilterCampaignSucceeded(opts *bind.FilterOpts) (*LLMFundraisingCampaignSucceededIterator, error) {

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "CampaignSucceeded")
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingCampaignSucceededIterator{contract: _LLMFundraising.contract, event: "CampaignSucceeded", logs: logs, sub: sub}, nil
}

// WatchCampaignSucceeded is a free log subscription operation binding the contract event 0x2b3f27cb35b5581f9baeb414c2592d8fc9c9067287476ca5624392320e9bf68e.
//
// Solidity: event CampaignSucceeded(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) WatchCampaignSucceeded(opts *bind.WatchOpts, sink chan<- *LLMFundraisingCampaignSucceeded) (event.Subscription, error) {

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "CampaignSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingCampaignSucceeded)
				if err := _LLMFundraising.contract.UnpackLog(event, "CampaignSucceeded", log); err != nil {
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

// ParseCampaignSucceeded is a log parse operation binding the contract event 0x2b3f27cb35b5581f9baeb414c2592d8fc9c9067287476ca5624392320e9bf68e.
//
// Solidity: event CampaignSucceeded(uint256 totalRaised)
func (_LLMFundraising *LLMFundraisingFilterer) ParseCampaignSucceeded(log types.Log) (*LLMFundraisingCampaignSucceeded, error) {
	event := new(LLMFundraisingCampaignSucceeded)
	if err := _LLMFundraising.contract.UnpackLog(event, "CampaignSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingDonatedIterator is returned from FilterDonated and is used to iterate over the raw logs and unpacked data for Donated events raised by the LLMFundraising contract.
type LLMFundraisingDonatedIterator struct {
	Event *LLMFundraisingDonated // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingDonated)
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
		it.Event = new(LLMFundraisingDonated)
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
func (it *LLMFundraisingDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingDonated represents a Donated event raised by the LLMFundraising contract.
type LLMFundraisingDonated struct {
	Donor        common.Address
	Amount       *big.Int
	TotalDonated *big.Int
	Grade        uint8
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDonated is a free log retrieval operation binding the contract event 0xf04d1237cbd0bf21dc045d2160a825f225935347027ba04bc8645d8138708d19.
//
// Solidity: event Donated(address indexed donor, uint256 amount, uint256 totalDonated, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) FilterDonated(opts *bind.FilterOpts, donor []common.Address) (*LLMFundraisingDonatedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingDonatedIterator{contract: _LLMFundraising.contract, event: "Donated", logs: logs, sub: sub}, nil
}

// WatchDonated is a free log subscription operation binding the contract event 0xf04d1237cbd0bf21dc045d2160a825f225935347027ba04bc8645d8138708d19.
//
// Solidity: event Donated(address indexed donor, uint256 amount, uint256 totalDonated, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) WatchDonated(opts *bind.WatchOpts, sink chan<- *LLMFundraisingDonated, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingDonated)
				if err := _LLMFundraising.contract.UnpackLog(event, "Donated", log); err != nil {
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

// ParseDonated is a log parse operation binding the contract event 0xf04d1237cbd0bf21dc045d2160a825f225935347027ba04bc8645d8138708d19.
//
// Solidity: event Donated(address indexed donor, uint256 amount, uint256 totalDonated, uint8 grade)
func (_LLMFundraising *LLMFundraisingFilterer) ParseDonated(log types.Log) (*LLMFundraisingDonated, error) {
	event := new(LLMFundraisingDonated)
	if err := _LLMFundraising.contract.UnpackLog(event, "Donated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingFundsTransferredIterator is returned from FilterFundsTransferred and is used to iterate over the raw logs and unpacked data for FundsTransferred events raised by the LLMFundraising contract.
type LLMFundraisingFundsTransferredIterator struct {
	Event *LLMFundraisingFundsTransferred // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingFundsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingFundsTransferred)
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
		it.Event = new(LLMFundraisingFundsTransferred)
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
func (it *LLMFundraisingFundsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingFundsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingFundsTransferred represents a FundsTransferred event raised by the LLMFundraising contract.
type LLMFundraisingFundsTransferred struct {
	Wallet common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsTransferred is a free log retrieval operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed wallet, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) FilterFundsTransferred(opts *bind.FilterOpts, wallet []common.Address) (*LLMFundraisingFundsTransferredIterator, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "FundsTransferred", walletRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingFundsTransferredIterator{contract: _LLMFundraising.contract, event: "FundsTransferred", logs: logs, sub: sub}, nil
}

// WatchFundsTransferred is a free log subscription operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed wallet, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) WatchFundsTransferred(opts *bind.WatchOpts, sink chan<- *LLMFundraisingFundsTransferred, wallet []common.Address) (event.Subscription, error) {

	var walletRule []interface{}
	for _, walletItem := range wallet {
		walletRule = append(walletRule, walletItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "FundsTransferred", walletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingFundsTransferred)
				if err := _LLMFundraising.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
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

// ParseFundsTransferred is a log parse operation binding the contract event 0x8c9a4f13b67cb64d7c6aa1ae0c9bf07694af521a28b93e7060020810ab4bc59f.
//
// Solidity: event FundsTransferred(address indexed wallet, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) ParseFundsTransferred(log types.Log) (*LLMFundraisingFundsTransferred, error) {
	event := new(LLMFundraisingFundsTransferred)
	if err := _LLMFundraising.contract.UnpackLog(event, "FundsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LLMFundraising contract.
type LLMFundraisingOwnershipTransferredIterator struct {
	Event *LLMFundraisingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingOwnershipTransferred)
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
		it.Event = new(LLMFundraisingOwnershipTransferred)
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
func (it *LLMFundraisingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingOwnershipTransferred represents a OwnershipTransferred event raised by the LLMFundraising contract.
type LLMFundraisingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LLMFundraising *LLMFundraisingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LLMFundraisingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingOwnershipTransferredIterator{contract: _LLMFundraising.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LLMFundraising *LLMFundraisingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LLMFundraisingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingOwnershipTransferred)
				if err := _LLMFundraising.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LLMFundraising *LLMFundraisingFilterer) ParseOwnershipTransferred(log types.Log) (*LLMFundraisingOwnershipTransferred, error) {
	event := new(LLMFundraisingOwnershipTransferred)
	if err := _LLMFundraising.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the LLMFundraising contract.
type LLMFundraisingRefundedIterator struct {
	Event *LLMFundraisingRefunded // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingRefunded)
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
		it.Event = new(LLMFundraisingRefunded)
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
func (it *LLMFundraisingRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingRefunded represents a Refunded event raised by the LLMFundraising contract.
type LLMFundraisingRefunded struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed donor, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) FilterRefunded(opts *bind.FilterOpts, donor []common.Address) (*LLMFundraisingRefundedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "Refunded", donorRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingRefundedIterator{contract: _LLMFundraising.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed donor, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *LLMFundraisingRefunded, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "Refunded", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingRefunded)
				if err := _LLMFundraising.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xd7dee2702d63ad89917b6a4da9981c90c4d24f8c2bdfd64c604ecae57d8d0651.
//
// Solidity: event Refunded(address indexed donor, uint256 amount)
func (_LLMFundraising *LLMFundraisingFilterer) ParseRefunded(log types.Log) (*LLMFundraisingRefunded, error) {
	event := new(LLMFundraisingRefunded)
	if err := _LLMFundraising.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LLMFundraisingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LLMFundraising contract.
type LLMFundraisingTransferIterator struct {
	Event *LLMFundraisingTransfer // Event containing the contract specifics and raw log

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
func (it *LLMFundraisingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LLMFundraisingTransfer)
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
		it.Event = new(LLMFundraisingTransfer)
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
func (it *LLMFundraisingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LLMFundraisingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LLMFundraisingTransfer represents a Transfer event raised by the LLMFundraising contract.
type LLMFundraisingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*LLMFundraisingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &LLMFundraisingTransferIterator{contract: _LLMFundraising.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LLMFundraisingTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _LLMFundraising.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LLMFundraisingTransfer)
				if err := _LLMFundraising.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_LLMFundraising *LLMFundraisingFilterer) ParseTransfer(log types.Log) (*LLMFundraisingTransfer, error) {
	event := new(LLMFundraisingTransfer)
	if err := _LLMFundraising.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
