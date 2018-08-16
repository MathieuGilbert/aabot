// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapters

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// DexBotABI is the input ABI used to generate the binding from.
const DexBotABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newNumber\",\"type\":\"uint256\"}],\"name\":\"setNumber\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"number\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// DexBotBin is the compiled bytecode used for deploying new contracts.
const DexBotBin = `0x608060405234801561001057600080fd5b506040516020806100f2833981016040525160005560bf806100336000396000f30060806040526004361060485763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633fb5c1cb8114604d5780638381f58a146064575b600080fd5b348015605857600080fd5b5060626004356088565b005b348015606f57600080fd5b506076608d565b60408051918252519081900360200190f35b600055565b600054815600a165627a7a72305820b9d5fc1a3ac562520cfeed3254963fb18b7f913c183c7c6b52c9cde0ce1edd930029`

// DeployDexBot deploys a new Ethereum contract, binding an instance of DexBot to it.
func DeployDexBot(auth *bind.TransactOpts, backend bind.ContractBackend, initialNumber *big.Int) (common.Address, *types.Transaction, *DexBot, error) {
	parsed, err := abi.JSON(strings.NewReader(DexBotABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DexBotBin), backend, initialNumber)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DexBot{DexBotCaller: DexBotCaller{contract: contract}, DexBotTransactor: DexBotTransactor{contract: contract}, DexBotFilterer: DexBotFilterer{contract: contract}}, nil
}

// DexBot is an auto generated Go binding around an Ethereum contract.
type DexBot struct {
	DexBotCaller     // Read-only binding to the contract
	DexBotTransactor // Write-only binding to the contract
	DexBotFilterer   // Log filterer for contract events
}

// DexBotCaller is an auto generated read-only Go binding around an Ethereum contract.
type DexBotCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexBotTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DexBotTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexBotFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DexBotFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DexBotSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DexBotSession struct {
	Contract     *DexBot           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexBotCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DexBotCallerSession struct {
	Contract *DexBotCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DexBotTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DexBotTransactorSession struct {
	Contract     *DexBotTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DexBotRaw is an auto generated low-level Go binding around an Ethereum contract.
type DexBotRaw struct {
	Contract *DexBot // Generic contract binding to access the raw methods on
}

// DexBotCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DexBotCallerRaw struct {
	Contract *DexBotCaller // Generic read-only contract binding to access the raw methods on
}

// DexBotTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DexBotTransactorRaw struct {
	Contract *DexBotTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDexBot creates a new instance of DexBot, bound to a specific deployed contract.
func NewDexBot(address common.Address, backend bind.ContractBackend) (*DexBot, error) {
	contract, err := bindDexBot(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DexBot{DexBotCaller: DexBotCaller{contract: contract}, DexBotTransactor: DexBotTransactor{contract: contract}, DexBotFilterer: DexBotFilterer{contract: contract}}, nil
}

// NewDexBotCaller creates a new read-only instance of DexBot, bound to a specific deployed contract.
func NewDexBotCaller(address common.Address, caller bind.ContractCaller) (*DexBotCaller, error) {
	contract, err := bindDexBot(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DexBotCaller{contract: contract}, nil
}

// NewDexBotTransactor creates a new write-only instance of DexBot, bound to a specific deployed contract.
func NewDexBotTransactor(address common.Address, transactor bind.ContractTransactor) (*DexBotTransactor, error) {
	contract, err := bindDexBot(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DexBotTransactor{contract: contract}, nil
}

// NewDexBotFilterer creates a new log filterer instance of DexBot, bound to a specific deployed contract.
func NewDexBotFilterer(address common.Address, filterer bind.ContractFilterer) (*DexBotFilterer, error) {
	contract, err := bindDexBot(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DexBotFilterer{contract: contract}, nil
}

// bindDexBot binds a generic wrapper to an already deployed contract.
func bindDexBot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DexBotABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexBot *DexBotRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DexBot.Contract.DexBotCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexBot *DexBotRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexBot.Contract.DexBotTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexBot *DexBotRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexBot.Contract.DexBotTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DexBot *DexBotCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DexBot.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DexBot *DexBotTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DexBot.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DexBot *DexBotTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DexBot.Contract.contract.Transact(opts, method, params...)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint256)
func (_DexBot *DexBotCaller) Number(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DexBot.contract.Call(opts, out, "number")
	return *ret0, err
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint256)
func (_DexBot *DexBotSession) Number() (*big.Int, error) {
	return _DexBot.Contract.Number(&_DexBot.CallOpts)
}

// Number is a free data retrieval call binding the contract method 0x8381f58a.
//
// Solidity: function number() constant returns(uint256)
func (_DexBot *DexBotCallerSession) Number() (*big.Int, error) {
	return _DexBot.Contract.Number(&_DexBot.CallOpts)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(newNumber uint256) returns()
func (_DexBot *DexBotTransactor) SetNumber(opts *bind.TransactOpts, newNumber *big.Int) (*types.Transaction, error) {
	return _DexBot.contract.Transact(opts, "setNumber", newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(newNumber uint256) returns()
func (_DexBot *DexBotSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _DexBot.Contract.SetNumber(&_DexBot.TransactOpts, newNumber)
}

// SetNumber is a paid mutator transaction binding the contract method 0x3fb5c1cb.
//
// Solidity: function setNumber(newNumber uint256) returns()
func (_DexBot *DexBotTransactorSession) SetNumber(newNumber *big.Int) (*types.Transaction, error) {
	return _DexBot.Contract.SetNumber(&_DexBot.TransactOpts, newNumber)
}
