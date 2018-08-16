// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapters

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AccountLevelsABI is the input ABI used to generate the binding from.
const AccountLevelsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"accountLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccountLevelsBin is the compiled bytecode used for deploying new contracts.
const AccountLevelsBin = `0x608060405234801561001057600080fd5b5060b28061001f6000396000f300608060405260043610603e5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631cbd051981146043575b600080fd5b348015604e57600080fd5b50606e73ffffffffffffffffffffffffffffffffffffffff600435166080565b60408051918252519081900360200190f35b506000905600a165627a7a7230582031714b34a2a8e73c43b2e8aafeae4a7ddd39b955fb2cdd7a595057442a4445620029`

// DeployAccountLevels deploys a new Ethereum contract, binding an instance of AccountLevels to it.
func DeployAccountLevels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountLevels, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountLevelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountLevels{AccountLevelsCaller: AccountLevelsCaller{contract: contract}, AccountLevelsTransactor: AccountLevelsTransactor{contract: contract}, AccountLevelsFilterer: AccountLevelsFilterer{contract: contract}}, nil
}

// AccountLevels is an auto generated Go binding around an Ethereum contract.
type AccountLevels struct {
	AccountLevelsCaller     // Read-only binding to the contract
	AccountLevelsTransactor // Write-only binding to the contract
	AccountLevelsFilterer   // Log filterer for contract events
}

// AccountLevelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountLevelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountLevelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountLevelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountLevelsSession struct {
	Contract     *AccountLevels    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountLevelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountLevelsCallerSession struct {
	Contract *AccountLevelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccountLevelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountLevelsTransactorSession struct {
	Contract     *AccountLevelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccountLevelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountLevelsRaw struct {
	Contract *AccountLevels // Generic contract binding to access the raw methods on
}

// AccountLevelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountLevelsCallerRaw struct {
	Contract *AccountLevelsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountLevelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountLevelsTransactorRaw struct {
	Contract *AccountLevelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountLevels creates a new instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevels(address common.Address, backend bind.ContractBackend) (*AccountLevels, error) {
	contract, err := bindAccountLevels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountLevels{AccountLevelsCaller: AccountLevelsCaller{contract: contract}, AccountLevelsTransactor: AccountLevelsTransactor{contract: contract}, AccountLevelsFilterer: AccountLevelsFilterer{contract: contract}}, nil
}

// NewAccountLevelsCaller creates a new read-only instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevelsCaller(address common.Address, caller bind.ContractCaller) (*AccountLevelsCaller, error) {
	contract, err := bindAccountLevels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsCaller{contract: contract}, nil
}

// NewAccountLevelsTransactor creates a new write-only instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevelsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountLevelsTransactor, error) {
	contract, err := bindAccountLevels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTransactor{contract: contract}, nil
}

// NewAccountLevelsFilterer creates a new log filterer instance of AccountLevels, bound to a specific deployed contract.
func NewAccountLevelsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountLevelsFilterer, error) {
	contract, err := bindAccountLevels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsFilterer{contract: contract}, nil
}

// bindAccountLevels binds a generic wrapper to an already deployed contract.
func bindAccountLevels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevels *AccountLevelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevels.Contract.AccountLevelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevels *AccountLevelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevels.Contract.AccountLevelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevels *AccountLevelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevels.Contract.AccountLevelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevels *AccountLevelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevels.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevels *AccountLevelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevels.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevels *AccountLevelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevels.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsCaller) AccountLevel(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevels.contract.Call(opts, out, "accountLevel", user)
	return *ret0, err
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevels.Contract.AccountLevel(&_AccountLevels.CallOpts, user)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevels *AccountLevelsCallerSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevels.Contract.AccountLevel(&_AccountLevels.CallOpts, user)
}

// AccountLevelsTestABI is the input ABI used to generate the binding from.
const AccountLevelsTestABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"accountLevels\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"accountLevel\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"level\",\"type\":\"uint256\"}],\"name\":\"setAccountLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccountLevelsTestBin is the compiled bytecode used for deploying new contracts.
const AccountLevelsTestBin = `0x608060405234801561001057600080fd5b5061018b806100206000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166314577c55811461005b5780631cbd05191461009b5780638abadb6b146100c9575b600080fd5b34801561006757600080fd5b5061008973ffffffffffffffffffffffffffffffffffffffff600435166100fc565b60408051918252519081900360200190f35b3480156100a757600080fd5b5061008973ffffffffffffffffffffffffffffffffffffffff6004351661010e565b3480156100d557600080fd5b506100fa73ffffffffffffffffffffffffffffffffffffffff60043516602435610136565b005b60006020819052908152604090205481565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b73ffffffffffffffffffffffffffffffffffffffff9091166000908152602081905260409020555600a165627a7a723058204be3977da9064547db07887ae7af847891964536156038531be11d3ece712e0c0029`

// DeployAccountLevelsTest deploys a new Ethereum contract, binding an instance of AccountLevelsTest to it.
func DeployAccountLevelsTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccountLevelsTest, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsTestABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccountLevelsTestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccountLevelsTest{AccountLevelsTestCaller: AccountLevelsTestCaller{contract: contract}, AccountLevelsTestTransactor: AccountLevelsTestTransactor{contract: contract}, AccountLevelsTestFilterer: AccountLevelsTestFilterer{contract: contract}}, nil
}

// AccountLevelsTest is an auto generated Go binding around an Ethereum contract.
type AccountLevelsTest struct {
	AccountLevelsTestCaller     // Read-only binding to the contract
	AccountLevelsTestTransactor // Write-only binding to the contract
	AccountLevelsTestFilterer   // Log filterer for contract events
}

// AccountLevelsTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountLevelsTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountLevelsTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountLevelsTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountLevelsTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountLevelsTestSession struct {
	Contract     *AccountLevelsTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AccountLevelsTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountLevelsTestCallerSession struct {
	Contract *AccountLevelsTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AccountLevelsTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountLevelsTestTransactorSession struct {
	Contract     *AccountLevelsTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AccountLevelsTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountLevelsTestRaw struct {
	Contract *AccountLevelsTest // Generic contract binding to access the raw methods on
}

// AccountLevelsTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountLevelsTestCallerRaw struct {
	Contract *AccountLevelsTestCaller // Generic read-only contract binding to access the raw methods on
}

// AccountLevelsTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountLevelsTestTransactorRaw struct {
	Contract *AccountLevelsTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccountLevelsTest creates a new instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTest(address common.Address, backend bind.ContractBackend) (*AccountLevelsTest, error) {
	contract, err := bindAccountLevelsTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTest{AccountLevelsTestCaller: AccountLevelsTestCaller{contract: contract}, AccountLevelsTestTransactor: AccountLevelsTestTransactor{contract: contract}, AccountLevelsTestFilterer: AccountLevelsTestFilterer{contract: contract}}, nil
}

// NewAccountLevelsTestCaller creates a new read-only instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTestCaller(address common.Address, caller bind.ContractCaller) (*AccountLevelsTestCaller, error) {
	contract, err := bindAccountLevelsTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTestCaller{contract: contract}, nil
}

// NewAccountLevelsTestTransactor creates a new write-only instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTestTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountLevelsTestTransactor, error) {
	contract, err := bindAccountLevelsTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTestTransactor{contract: contract}, nil
}

// NewAccountLevelsTestFilterer creates a new log filterer instance of AccountLevelsTest, bound to a specific deployed contract.
func NewAccountLevelsTestFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountLevelsTestFilterer, error) {
	contract, err := bindAccountLevelsTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountLevelsTestFilterer{contract: contract}, nil
}

// bindAccountLevelsTest binds a generic wrapper to an already deployed contract.
func bindAccountLevelsTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountLevelsTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevelsTest *AccountLevelsTestRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevelsTest.Contract.AccountLevelsTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevelsTest *AccountLevelsTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.AccountLevelsTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevelsTest *AccountLevelsTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.AccountLevelsTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccountLevelsTest *AccountLevelsTestCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccountLevelsTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccountLevelsTest *AccountLevelsTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccountLevelsTest *AccountLevelsTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.contract.Transact(opts, method, params...)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCaller) AccountLevel(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevelsTest.contract.Call(opts, out, "accountLevel", user)
	return *ret0, err
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevel(&_AccountLevelsTest.CallOpts, user)
}

// AccountLevel is a free data retrieval call binding the contract method 0x1cbd0519.
//
// Solidity: function accountLevel(user address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCallerSession) AccountLevel(user common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevel(&_AccountLevelsTest.CallOpts, user)
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCaller) AccountLevels(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AccountLevelsTest.contract.Call(opts, out, "accountLevels", arg0)
	return *ret0, err
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestSession) AccountLevels(arg0 common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevels(&_AccountLevelsTest.CallOpts, arg0)
}

// AccountLevels is a free data retrieval call binding the contract method 0x14577c55.
//
// Solidity: function accountLevels( address) constant returns(uint256)
func (_AccountLevelsTest *AccountLevelsTestCallerSession) AccountLevels(arg0 common.Address) (*big.Int, error) {
	return _AccountLevelsTest.Contract.AccountLevels(&_AccountLevelsTest.CallOpts, arg0)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestTransactor) SetAccountLevel(opts *bind.TransactOpts, user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.contract.Transact(opts, "setAccountLevel", user, level)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestSession) SetAccountLevel(user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.SetAccountLevel(&_AccountLevelsTest.TransactOpts, user, level)
}

// SetAccountLevel is a paid mutator transaction binding the contract method 0x8abadb6b.
//
// Solidity: function setAccountLevel(user address, level uint256) returns()
func (_AccountLevelsTest *AccountLevelsTestTransactorSession) SetAccountLevel(user common.Address, level *big.Int) (*types.Transaction, error) {
	return _AccountLevelsTest.Contract.SetAccountLevel(&_AccountLevelsTest.TransactOpts, user, level)
}

// EtherDeltaABI is the input ABI used to generate the binding from.
const EtherDeltaABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"order\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderFills\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"amountFilled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeMake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeMake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeMake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"name\":\"changeFeeRebate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"testTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeAccount_\",\"type\":\"address\"}],\"name\":\"changeFeeAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeRebate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"feeTake_\",\"type\":\"uint256\"}],\"name\":\"changeFeeTake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"changeAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeTake\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"}],\"name\":\"changeAccountLevelsAddr\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accountLevelsAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenGet\",\"type\":\"address\"},{\"name\":\"amountGet\",\"type\":\"uint256\"},{\"name\":\"tokenGive\",\"type\":\"address\"},{\"name\":\"amountGive\",\"type\":\"uint256\"},{\"name\":\"expires\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"availableVolume\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\"},{\"name\":\"feeAccount_\",\"type\":\"address\"},{\"name\":\"accountLevelsAddr_\",\"type\":\"address\"},{\"name\":\"feeMake_\",\"type\":\"uint256\"},{\"name\":\"feeTake_\",\"type\":\"uint256\"},{\"name\":\"feeRebate_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"Order\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expires\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"v\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"r\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"Cancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenGet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokenGive\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountGive\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"get\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"give\",\"type\":\"address\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// EtherDeltaBin is the compiled bytecode used for deploying new contracts.
const EtherDeltaBin = `0x608060405234801561001057600080fd5b5060405160c08061198883398101604090815281516020830151918301516060840151608085015160a09095015160008054600160a060020a03958616600160a060020a03199182161790915560018054968616968216969096179095556002805494909316939094169290921790556003556004919091556005556118ed8061009b6000396000f3006080604052600436106101535763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630a19b14a81146101655780630b927666146101b657806319774d43146101ea578063278b8c0e146102205780632e1a7d4d14610261578063338b5dea1461027957806346be96c31461029d578063508493bc146102e857806354d03b5c1461030f57806357786394146103275780635e1d7ae41461033c57806365e17c9d146103545780636c86888b1461038557806371ffcb16146103f3578063731c2f81146104145780638823a9c0146104295780638f283970146104415780639e281a9814610462578063bb5f462914610486578063c281309e146104aa578063d0e30db0146104bf578063e8f6bc2e146104c7578063f3412942146104e8578063f7888aec146104fd578063f851a44014610524578063fb6e155f14610539575b34801561015f57600080fd5b50600080fd5b34801561017157600080fd5b506101b4600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e43516610104356101243561014435610584565b005b3480156101c257600080fd5b506101b4600160a060020a03600435811690602435906044351660643560843560a435610822565b3480156101f657600080fd5b5061020e600160a060020a036004351660243561094a565b60408051918252519081900360200190f35b34801561022c57600080fd5b506101b4600160a060020a03600435811690602435906044351660643560843560a43560ff60c4351660e43561010435610967565b34801561026d57600080fd5b506101b4600435610b77565b34801561028557600080fd5b506101b4600160a060020a0360043516602435610c6b565b3480156102a957600080fd5b5061020e600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e435166101043561012435610dc6565b3480156102f457600080fd5b5061020e600160a060020a0360043581169060243516610e91565b34801561031b57600080fd5b506101b4600435610eae565b34801561033357600080fd5b5061020e610ed9565b34801561034857600080fd5b506101b4600435610edf565b34801561036057600080fd5b50610369610f16565b60408051600160a060020a039092168252519081900360200190f35b34801561039157600080fd5b506103df600160a060020a0360043581169060243590604435811690606435906084359060a4359060c43581169060ff60e43516906101043590610124359061014435906101643516610f25565b604080519115158252519081900360200190f35b3480156103ff57600080fd5b506101b4600160a060020a0360043516610f8f565b34801561042057600080fd5b5061020e610fd5565b34801561043557600080fd5b506101b4600435610fdb565b34801561044d57600080fd5b506101b4600160a060020a0360043516611012565b34801561046e57600080fd5b506101b4600160a060020a0360043516602435611058565b34801561049257600080fd5b506103df600160a060020a03600435166024356111f3565b3480156104b657600080fd5b5061020e611213565b6101b4611219565b3480156104d357600080fd5b506101b4600160a060020a03600435166112a8565b3480156104f457600080fd5b506103696112ee565b34801561050957600080fd5b5061020e600160a060020a03600435811690602435166112fd565b34801561053057600080fd5b50610369611328565b34801561054557600080fd5b5061020e600160a060020a0360043581169060243590604435811690606435906084359060a4359060c4351660ff60e435166101043561012435611337565b604080516c010000000000000000000000003081028252600160a060020a03808f1682026014840152602883018e90528c16026048820152605c81018a9052607c8101899052609c8101889052905160009160029160bc808301926020929190829003018186865af11580156105fe573d6000803e3d6000fd5b5050506040513d602081101561061357600080fd5b5051600160a060020a038716600090815260076020908152604080832084845290915290205490915060ff16806106f35750604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052815190819003603c018120600080835260208381018086529290925260ff89168385015260608301889052608083018790529251600160a060020a038a169360019360a08082019493601f198101939281900390910191865af11580156106de573d6000803e3d6000fd5b50505060206040510351600160a060020a0316145b80156106ff5750874311155b80156107395750600160a060020a03861660009081526008602090815260408083208484529091529020548b90610736908461155d565b11155b151561074457600080fd5b6107528c8c8c8c8a87611581565b600160a060020a0386166000908152600860209081526040808320848452909152902054610780908361155d565b600160a060020a03871660009081526008602090815260408083208584529091529020557f6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d8c838c8e8d83028115156107d557fe5b60408051600160a060020a03968716815260208101959095529285168484015204606083015291891660808201523360a082015290519081900360c00190a1505050505050505050505050565b604080516c010000000000000000000000003081028252600160a060020a03808a1682026014840152602883018990528716026048820152605c8101859052607c8101849052609c8101839052905160009160029160bc808301926020929190829003018186865af115801561089c573d6000803e3d6000fd5b5050506040513d60208110156108b157600080fd5b5051336000818152600760209081526040808320858452825291829020805460ff191660011790558151600160a060020a038c811682529181018b905290891681830152606081018890526080810187905260a0810186905260c0810192909252519192507f3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e85919081900360e00190a150505050505050565b600860209081526000928352604080842090915290825290205481565b604080516c010000000000000000000000003081028252600160a060020a03808d1682026014840152602883018c90528a16026048820152605c8101889052607c8101879052609c8101869052905160009160029160bc808301926020929190829003018186865af11580156109e1573d6000803e3d6000fd5b5050506040513d60208110156109f657600080fd5b505133600090815260076020908152604080832084845290915290205490915060ff1680610ac45750604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101839052815190819003603c018120600080835260208381018086529290925260ff88168385015260608301879052608083018690529251339360019360a08082019493601f198101939281900390910191865af1158015610aaf573d6000803e3d6000fd5b50505060206040510351600160a060020a0316145b1515610acf57600080fd5b3360008181526008602090815260408083208584528252918290208c90558151600160a060020a038e811682529181018d9052908b1681830152606081018a90526080810189905260a0810188905260c081019290925260ff861660e083015261010082018590526101208201849052517f1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0918190036101400190a150505050505050505050565b3360009081526000805160206118a28339815191526020526040902054811115610ba057600080fd5b3360009081526000805160206118a28339815191526020526040902054610bc7908261185f565b3360008181526000805160206118a28339815191526020526040808220939093559151909183919081818185875af1925050501515610c0557600080fd5b3360008181526000805160206118a28339815191526020908152604080832054815193845291830193909352818301849052606082015290517ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb5679181900360800190a150565b600160a060020a0382161515610c8057600080fd5b604080517f23b872dd000000000000000000000000000000000000000000000000000000008152336004820152306024820152604481018390529051600160a060020a038416916323b872dd9160648083019260209291908290030181600087803b158015610cee57600080fd5b505af1158015610d02573d6000803e3d6000fd5b505050506040513d6020811015610d1857600080fd5b50511515610d2557600080fd5b600160a060020a0382166000908152600660209081526040808320338452909152902054610d53908261155d565b600160a060020a03831660008181526006602090815260408083203380855290835292819020859055805193845290830191909152818101849052606082019290925290517fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d79181900360800190a15050565b604080516c010000000000000000000000003081028252600160a060020a03808e1682026014840152602883018d90528b16026048820152605c8101899052607c8101889052609c81018790529051600091829160029160bc80820192602092909190829003018186865af1158015610e43573d6000803e3d6000fd5b5050506040513d6020811015610e5857600080fd5b5051600160a060020a038716600090815260086020908152604080832084845290915290205492509050509a9950505050505050505050565b600660209081526000928352604080842090915290825290205481565b600054600160a060020a03163314610ec557600080fd5b600354811115610ed457600080fd5b600355565b60035481565b600054600160a060020a03163314610ef657600080fd5b600554811080610f07575060045481115b15610f1157600080fd5b600555565b600154600160a060020a031681565b600160a060020a03808d1660009081526006602090815260408083209385168352929052908120548311801590610f6d575082610f6a8e8e8e8e8e8e8e8e8e8e611337565b10155b1515610f7b57506000610f7f565b5060015b9c9b505050505050505050505050565b600054600160a060020a03163314610fa657600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60055481565b600054600160a060020a03163314610ff257600080fd5b600454811180611003575060055481105b1561100d57600080fd5b600455565b600054600160a060020a0316331461102957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a038216151561106d57600080fd5b600160a060020a038216600090815260066020908152604080832033845290915290205481111561109d57600080fd5b600160a060020a03821660009081526006602090815260408083203384529091529020546110cb908261185f565b600160a060020a0383166000818152600660209081526040808320338085529083528184209590955580517fa9059cbb00000000000000000000000000000000000000000000000000000000815260048101959095526024850186905251929363a9059cbb9360448083019491928390030190829087803b15801561114f57600080fd5b505af1158015611163573d6000803e3d6000fd5b505050506040513d602081101561117957600080fd5b5051151561118657600080fd5b600160a060020a03821660008181526006602090815260408083203380855290835292819020548151948552918401929092528282018490526060830152517ff341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb5679181900360800190a15050565b600760209081526000928352604080842090915290825290205460ff1681565b60045481565b3360009081526000805160206118a28339815191526020526040902054611240903461155d565b3360008181526000805160206118a28339815191526020908152604080832085905580519283529082019290925234818301526060810192909252517fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d79181900360800190a1565b600054600160a060020a031633146112bf57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600254600160a060020a031681565b600160a060020a03918216600090815260066020908152604080832093909416825291909152205490565b600054600160a060020a031681565b604080516c010000000000000000000000003081028252600160a060020a03808e1682026014840152602883018d90528b16026048820152605c8101899052607c8101889052609c8101879052905160009182918291829160029160bc80820192602092909190829003018186865af11580156113b8573d6000803e3d6000fd5b5050506040513d60208110156113cd57600080fd5b5051600160a060020a038916600090815260076020908152604080832084845290915290205490935060ff16806114ad5750604080517f19457468657265756d205369676e6564204d6573736167653a0a3332000000008152601c8101859052815190819003603c018120600080835260208381018086529290925260ff8b1683850152606083018a9052608083018990529251600160a060020a038c169360019360a08082019493601f198101939281900390910191865af1158015611498573d6000803e3d6000fd5b50505060206040510351600160a060020a0316145b80156114b95750894311155b15156114c8576000935061154c565b600160a060020a03881660009081526008602090815260408083208684529091529020546114f7908e9061185f565b600160a060020a03808e166000908152600660209081526040808320938d16835292905220549092508b9061152c908f611873565b81151561153557fe5b049050808210156115485781935061154c565b8093505b5050509a9950505050505050505050565b600082820161157a8482108015906115755750838210155b611892565b9392505050565b600080600080670de0b6b3a764000061159c86600354611873565b8115156115a557fe5b049350670de0b6b3a76400006115bd86600454611873565b8115156115c657fe5b600254919004935060009250600160a060020a0316156116ae57600254604080517f1cbd0519000000000000000000000000000000000000000000000000000000008152600160a060020a03898116600483015291519190921691631cbd05199160248083019260209291908290030181600087803b15801561164857600080fd5b505af115801561165c573d6000803e3d6000fd5b505050506040513d602081101561167257600080fd5b5051905060018114156116a157670de0b6b3a764000061169486600554611873565b81151561169d57fe5b0491505b80600214156116ae578291505b600160a060020a038a1660009081526006602090815260408083203384529091529020546116e5906116e0878661155d565b61185f565b600160a060020a038b8116600090815260066020908152604080832033845290915280822093909355908816815220546117319061172c611726888661155d565b8761185f565b61155d565b600160a060020a038b811660009081526006602090815260408083208b85168452909152808220939093556001549091168152205461177d9061172c611777878761155d565b8561185f565b600160a060020a03808c166000908152600660208181526040808420600154861685528252808420959095558c84168352908152838220928a1682529190915220546117dd908a6117ce8a89611873565b8115156117d757fe5b0461185f565b600160a060020a038981166000908152600660209081526040808320938b1683529290528181209290925533825290205461182c908a61181d8a89611873565b81151561182657fe5b0461155d565b600160a060020a039098166000908152600660209081526040808320338452909152902097909755505050505050505050565b600061186d83831115611892565b50900390565b600082820261157a841580611575575083858381151561188f57fe5b04145b80151561189e57600080fd5b50560054cdd369e4e8a8515e52ca72ec816c2101831ad1f18bf44102ed171459c9b4f8a165627a7a7230582091cea87986247442b9657a39905275f8460ff3e4fc70fe2cb4118eccbc52d5950029`

// DeployEtherDelta deploys a new Ethereum contract, binding an instance of EtherDelta to it.
func DeployEtherDelta(auth *bind.TransactOpts, backend bind.ContractBackend, admin_ common.Address, feeAccount_ common.Address, accountLevelsAddr_ common.Address, feeMake_ *big.Int, feeTake_ *big.Int, feeRebate_ *big.Int) (common.Address, *types.Transaction, *EtherDelta, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherDeltaABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherDeltaBin), backend, admin_, feeAccount_, accountLevelsAddr_, feeMake_, feeTake_, feeRebate_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherDelta{EtherDeltaCaller: EtherDeltaCaller{contract: contract}, EtherDeltaTransactor: EtherDeltaTransactor{contract: contract}, EtherDeltaFilterer: EtherDeltaFilterer{contract: contract}}, nil
}

// EtherDelta is an auto generated Go binding around an Ethereum contract.
type EtherDelta struct {
	EtherDeltaCaller     // Read-only binding to the contract
	EtherDeltaTransactor // Write-only binding to the contract
	EtherDeltaFilterer   // Log filterer for contract events
}

// EtherDeltaCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherDeltaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherDeltaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherDeltaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherDeltaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherDeltaSession struct {
	Contract     *EtherDelta       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherDeltaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherDeltaCallerSession struct {
	Contract *EtherDeltaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EtherDeltaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherDeltaTransactorSession struct {
	Contract     *EtherDeltaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EtherDeltaRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherDeltaRaw struct {
	Contract *EtherDelta // Generic contract binding to access the raw methods on
}

// EtherDeltaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherDeltaCallerRaw struct {
	Contract *EtherDeltaCaller // Generic read-only contract binding to access the raw methods on
}

// EtherDeltaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherDeltaTransactorRaw struct {
	Contract *EtherDeltaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherDelta creates a new instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDelta(address common.Address, backend bind.ContractBackend) (*EtherDelta, error) {
	contract, err := bindEtherDelta(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherDelta{EtherDeltaCaller: EtherDeltaCaller{contract: contract}, EtherDeltaTransactor: EtherDeltaTransactor{contract: contract}, EtherDeltaFilterer: EtherDeltaFilterer{contract: contract}}, nil
}

// NewEtherDeltaCaller creates a new read-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaCaller(address common.Address, caller bind.ContractCaller) (*EtherDeltaCaller, error) {
	contract, err := bindEtherDelta(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaCaller{contract: contract}, nil
}

// NewEtherDeltaTransactor creates a new write-only instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherDeltaTransactor, error) {
	contract, err := bindEtherDelta(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaTransactor{contract: contract}, nil
}

// NewEtherDeltaFilterer creates a new log filterer instance of EtherDelta, bound to a specific deployed contract.
func NewEtherDeltaFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherDeltaFilterer, error) {
	contract, err := bindEtherDelta(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherDeltaFilterer{contract: contract}, nil
}

// bindEtherDelta binds a generic wrapper to an already deployed contract.
func bindEtherDelta(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherDeltaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.EtherDeltaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.EtherDeltaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherDelta *EtherDeltaCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherDelta.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherDelta *EtherDeltaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherDelta *EtherDeltaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherDelta.Contract.contract.Transact(opts, method, params...)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) AccountLevelsAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "accountLevelsAddr")
	return *ret0, err
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// AccountLevelsAddr is a free data retrieval call binding the contract method 0xf3412942.
//
// Solidity: function accountLevelsAddr() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) AccountLevelsAddr() (common.Address, error) {
	return _EtherDelta.Contract.AccountLevelsAddr(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) Admin() (common.Address, error) {
	return _EtherDelta.Contract.Admin(&_EtherDelta.CallOpts)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AmountFilled(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "amountFilled", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AmountFilled is a free data retrieval call binding the contract method 0x46be96c3.
//
// Solidity: function amountFilled(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AmountFilled(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AmountFilled(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) AvailableVolume(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "availableVolume", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
	return *ret0, err
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// AvailableVolume is a free data retrieval call binding the contract method 0xfb6e155f.
//
// Solidity: function availableVolume(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) AvailableVolume(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.AvailableVolume(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) BalanceOf(opts *bind.CallOpts, token common.Address, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "balanceOf", token, user)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(token address, user address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) BalanceOf(token common.Address, user common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.BalanceOf(&_EtherDelta.CallOpts, token, user)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCaller) FeeAccount(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeAccount")
	return *ret0, err
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeAccount is a free data retrieval call binding the contract method 0x65e17c9d.
//
// Solidity: function feeAccount() constant returns(address)
func (_EtherDelta *EtherDeltaCallerSession) FeeAccount() (common.Address, error) {
	return _EtherDelta.Contract.FeeAccount(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeMake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeMake")
	return *ret0, err
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeMake is a free data retrieval call binding the contract method 0x57786394.
//
// Solidity: function feeMake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeMake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeMake(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeRebate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeRebate")
	return *ret0, err
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeRebate is a free data retrieval call binding the contract method 0x731c2f81.
//
// Solidity: function feeRebate() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeRebate() (*big.Int, error) {
	return _EtherDelta.Contract.FeeRebate(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) FeeTake(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "feeTake")
	return *ret0, err
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// FeeTake is a free data retrieval call binding the contract method 0xc281309e.
//
// Solidity: function feeTake() constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) FeeTake() (*big.Int, error) {
	return _EtherDelta.Contract.FeeTake(&_EtherDelta.CallOpts)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) OrderFills(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orderFills", arg0, arg1)
	return *ret0, err
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// OrderFills is a free data retrieval call binding the contract method 0x19774d43.
//
// Solidity: function orderFills( address,  bytes32) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) OrderFills(arg0 common.Address, arg1 [32]byte) (*big.Int, error) {
	return _EtherDelta.Contract.OrderFills(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) Orders(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "orders", arg0, arg1)
	return *ret0, err
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xbb5f4629.
//
// Solidity: function orders( address,  bytes32) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) Orders(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _EtherDelta.Contract.Orders(&_EtherDelta.CallOpts, arg0, arg1)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCaller) TestTrade(opts *bind.CallOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "testTrade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
	return *ret0, err
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// TestTrade is a free data retrieval call binding the contract method 0x6c86888b.
//
// Solidity: function testTrade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256, sender address) constant returns(bool)
func (_EtherDelta *EtherDeltaCallerSession) TestTrade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int, sender common.Address) (bool, error) {
	return _EtherDelta.Contract.TestTrade(&_EtherDelta.CallOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount, sender)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCaller) Tokens(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherDelta.contract.Call(opts, out, "tokens", arg0, arg1)
	return *ret0, err
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// Tokens is a free data retrieval call binding the contract method 0x508493bc.
//
// Solidity: function tokens( address,  address) constant returns(uint256)
func (_EtherDelta *EtherDeltaCallerSession) Tokens(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherDelta.Contract.Tokens(&_EtherDelta.CallOpts, arg0, arg1)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactor) CancelOrder(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "cancelOrder", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x278b8c0e.
//
// Solidity: function cancelOrder(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, v uint8, r bytes32, s bytes32) returns()
func (_EtherDelta *EtherDeltaTransactorSession) CancelOrder(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _EtherDelta.Contract.CancelOrder(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, v, r, s)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAccountLevelsAddr(opts *bind.TransactOpts, accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAccountLevelsAddr", accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAccountLevelsAddr is a paid mutator transaction binding the contract method 0xe8f6bc2e.
//
// Solidity: function changeAccountLevelsAddr(accountLevelsAddr_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAccountLevelsAddr(accountLevelsAddr_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAccountLevelsAddr(&_EtherDelta.TransactOpts, accountLevelsAddr_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeAdmin", admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeAdmin is a paid mutator transaction binding the contract method 0x8f283970.
//
// Solidity: function changeAdmin(admin_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeAdmin(&_EtherDelta.TransactOpts, admin_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeAccount(opts *bind.TransactOpts, feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeAccount", feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeAccount is a paid mutator transaction binding the contract method 0x71ffcb16.
//
// Solidity: function changeFeeAccount(feeAccount_ address) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeAccount(feeAccount_ common.Address) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeAccount(&_EtherDelta.TransactOpts, feeAccount_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeMake(opts *bind.TransactOpts, feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeMake", feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeMake is a paid mutator transaction binding the contract method 0x54d03b5c.
//
// Solidity: function changeFeeMake(feeMake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeMake(feeMake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeMake(&_EtherDelta.TransactOpts, feeMake_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeRebate(opts *bind.TransactOpts, feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeRebate", feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeRebate is a paid mutator transaction binding the contract method 0x5e1d7ae4.
//
// Solidity: function changeFeeRebate(feeRebate_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeRebate(feeRebate_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeRebate(&_EtherDelta.TransactOpts, feeRebate_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) ChangeFeeTake(opts *bind.TransactOpts, feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "changeFeeTake", feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// ChangeFeeTake is a paid mutator transaction binding the contract method 0x8823a9c0.
//
// Solidity: function changeFeeTake(feeTake_ uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) ChangeFeeTake(feeTake_ *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.ChangeFeeTake(&_EtherDelta.TransactOpts, feeTake_)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_EtherDelta *EtherDeltaTransactorSession) Deposit() (*types.Transaction, error) {
	return _EtherDelta.Contract.Deposit(&_EtherDelta.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) DepositToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "depositToken", token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.DepositToken(&_EtherDelta.TransactOpts, token, amount)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Order(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "order", tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Order is a paid mutator transaction binding the contract method 0x0b927666.
//
// Solidity: function order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Order(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Order(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Trade(opts *bind.TransactOpts, tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "trade", tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Trade is a paid mutator transaction binding the contract method 0x0a19b14a.
//
// Solidity: function trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Trade(tokenGet common.Address, amountGet *big.Int, tokenGive common.Address, amountGive *big.Int, expires *big.Int, nonce *big.Int, user common.Address, v uint8, r [32]byte, s [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Trade(&_EtherDelta.TransactOpts, tokenGet, amountGet, tokenGive, amountGive, expires, nonce, user, v, r, s, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.Withdraw(&_EtherDelta.TransactOpts, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.contract.Transact(opts, "withdrawToken", token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x9e281a98.
//
// Solidity: function withdrawToken(token address, amount uint256) returns()
func (_EtherDelta *EtherDeltaTransactorSession) WithdrawToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EtherDelta.Contract.WithdrawToken(&_EtherDelta.TransactOpts, token, amount)
}

// EtherDeltaCancelIterator is returned from FilterCancel and is used to iterate over the raw logs and unpacked data for Cancel events raised by the EtherDelta contract.
type EtherDeltaCancelIterator struct {
	Event *EtherDeltaCancel // Event containing the contract specifics and raw log

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
func (it *EtherDeltaCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaCancel)
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
		it.Event = new(EtherDeltaCancel)
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
func (it *EtherDeltaCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaCancel represents a Cancel event raised by the EtherDelta contract.
type EtherDeltaCancel struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Expires    *big.Int
	Nonce      *big.Int
	User       common.Address
	V          uint8
	R          [32]byte
	S          [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCancel is a free log retrieval operation binding the contract event 0x1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0.
//
// Solidity: e Cancel(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32)
func (_EtherDelta *EtherDeltaFilterer) FilterCancel(opts *bind.FilterOpts) (*EtherDeltaCancelIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaCancelIterator{contract: _EtherDelta.contract, event: "Cancel", logs: logs, sub: sub}, nil
}

// WatchCancel is a free log subscription operation binding the contract event 0x1e0b760c386003e9cb9bcf4fcf3997886042859d9b6ed6320e804597fcdb28b0.
//
// Solidity: e Cancel(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address, v uint8, r bytes32, s bytes32)
func (_EtherDelta *EtherDeltaFilterer) WatchCancel(opts *bind.WatchOpts, sink chan<- *EtherDeltaCancel) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Cancel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaCancel)
				if err := _EtherDelta.contract.UnpackLog(event, "Cancel", log); err != nil {
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

// EtherDeltaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the EtherDelta contract.
type EtherDeltaDepositIterator struct {
	Event *EtherDeltaDeposit // Event containing the contract specifics and raw log

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
func (it *EtherDeltaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaDeposit)
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
		it.Event = new(EtherDeltaDeposit)
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
func (it *EtherDeltaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaDeposit represents a Deposit event raised by the EtherDelta contract.
type EtherDeltaDeposit struct {
	Token   common.Address
	User    common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: e Deposit(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) FilterDeposit(opts *bind.FilterOpts) (*EtherDeltaDepositIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaDepositIterator{contract: _EtherDelta.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: e Deposit(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *EtherDeltaDeposit) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaDeposit)
				if err := _EtherDelta.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// EtherDeltaOrderIterator is returned from FilterOrder and is used to iterate over the raw logs and unpacked data for Order events raised by the EtherDelta contract.
type EtherDeltaOrderIterator struct {
	Event *EtherDeltaOrder // Event containing the contract specifics and raw log

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
func (it *EtherDeltaOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaOrder)
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
		it.Event = new(EtherDeltaOrder)
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
func (it *EtherDeltaOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaOrder represents a Order event raised by the EtherDelta contract.
type EtherDeltaOrder struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Expires    *big.Int
	Nonce      *big.Int
	User       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrder is a free log retrieval operation binding the contract event 0x3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e85.
//
// Solidity: e Order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address)
func (_EtherDelta *EtherDeltaFilterer) FilterOrder(opts *bind.FilterOpts) (*EtherDeltaOrderIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Order")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaOrderIterator{contract: _EtherDelta.contract, event: "Order", logs: logs, sub: sub}, nil
}

// WatchOrder is a free log subscription operation binding the contract event 0x3f7f2eda73683c21a15f9435af1028c93185b5f1fa38270762dc32be606b3e85.
//
// Solidity: e Order(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, expires uint256, nonce uint256, user address)
func (_EtherDelta *EtherDeltaFilterer) WatchOrder(opts *bind.WatchOpts, sink chan<- *EtherDeltaOrder) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Order")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaOrder)
				if err := _EtherDelta.contract.UnpackLog(event, "Order", log); err != nil {
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

// EtherDeltaTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the EtherDelta contract.
type EtherDeltaTradeIterator struct {
	Event *EtherDeltaTrade // Event containing the contract specifics and raw log

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
func (it *EtherDeltaTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaTrade)
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
		it.Event = new(EtherDeltaTrade)
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
func (it *EtherDeltaTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaTrade represents a Trade event raised by the EtherDelta contract.
type EtherDeltaTrade struct {
	TokenGet   common.Address
	AmountGet  *big.Int
	TokenGive  common.Address
	AmountGive *big.Int
	Get        common.Address
	Give       common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0x6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d.
//
// Solidity: e Trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, get address, give address)
func (_EtherDelta *EtherDeltaFilterer) FilterTrade(opts *bind.FilterOpts) (*EtherDeltaTradeIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaTradeIterator{contract: _EtherDelta.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0x6effdda786735d5033bfad5f53e5131abcced9e52be6c507b62d639685fbed6d.
//
// Solidity: e Trade(tokenGet address, amountGet uint256, tokenGive address, amountGive uint256, get address, give address)
func (_EtherDelta *EtherDeltaFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *EtherDeltaTrade) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Trade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaTrade)
				if err := _EtherDelta.contract.UnpackLog(event, "Trade", log); err != nil {
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

// EtherDeltaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the EtherDelta contract.
type EtherDeltaWithdrawIterator struct {
	Event *EtherDeltaWithdraw // Event containing the contract specifics and raw log

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
func (it *EtherDeltaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherDeltaWithdraw)
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
		it.Event = new(EtherDeltaWithdraw)
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
func (it *EtherDeltaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherDeltaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherDeltaWithdraw represents a Withdraw event raised by the EtherDelta contract.
type EtherDeltaWithdraw struct {
	Token   common.Address
	User    common.Address
	Amount  *big.Int
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: e Withdraw(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) FilterWithdraw(opts *bind.FilterOpts) (*EtherDeltaWithdrawIterator, error) {

	logs, sub, err := _EtherDelta.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &EtherDeltaWithdrawIterator{contract: _EtherDelta.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xf341246adaac6f497bc2a656f546ab9e182111d630394f0c57c710a59a2cb567.
//
// Solidity: e Withdraw(token address, user address, amount uint256, balance uint256)
func (_EtherDelta *EtherDeltaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *EtherDeltaWithdraw) (event.Subscription, error) {

	logs, sub, err := _EtherDelta.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherDeltaWithdraw)
				if err := _EtherDelta.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ReserveTokenABI is the input ABI used to generate the binding from.
const ReserveTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ReserveTokenBin is the compiled bytecode used for deploying new contracts.
const ReserveTokenBin = `0x608060405234801561001057600080fd5b5060058054600160a060020a03191633179055610727806100326000396000f3006080604052600436106100ae5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100b3578063075461721461013d578063095ea7b31461016e5780630ecaea73146101a657806318160ddd146101cc57806323b872dd146101f3578063313ce5671461021d57806370a0823114610232578063a24835d114610253578063a9059cbb14610277578063dd62ed3e1461029b575b600080fd5b3480156100bf57600080fd5b506100c86102c2565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101025781810151838201526020016100ea565b50505050905090810190601f16801561012f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561014957600080fd5b5061015261034f565b60408051600160a060020a039092168252519081900360200190f35b34801561017a57600080fd5b50610192600160a060020a036004351660243561035e565b604080519115158252519081900360200190f35b3480156101b257600080fd5b506101ca600160a060020a03600435166024356103c5565b005b3480156101d857600080fd5b506101e161042c565b60408051918252519081900360200190f35b3480156101ff57600080fd5b50610192600160a060020a0360043581169060243516604435610432565b34801561022957600080fd5b506101e1610538565b34801561023e57600080fd5b506101e1600160a060020a036004351661053e565b34801561025f57600080fd5b506101ca600160a060020a0360043516602435610559565b34801561028357600080fd5b50610192600160a060020a03600435166024356105de565b3480156102a757600080fd5b506101e1600160a060020a0360043581169060243516610690565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103475780601f1061031c57610100808354040283529160200191610347565b820191906000526020600020905b81548152906001019060200180831161032a57829003601f168201915b505050505081565b600554600160a060020a031681565b336000818152600360209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b600554600160a060020a031633146103dc57600080fd5b600160a060020a0382166000908152600260205260409020546103ff90826106bb565b600160a060020a03831660009081526002602052604090205560045461042590826106bb565b6004555050565b60045481565b600160a060020a038316600090815260026020526040812054821180159061047d5750600160a060020a03841660009081526003602090815260408083203384529091529020548211155b80156104a25750600160a060020a038316600090815260026020526040902054828101115b1561052d57600160a060020a03808416600081815260026020908152604080832080548801905593881680835284832080548890039055600382528483203384528252918490208054879003905583518681529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a3506001610531565b5060005b9392505050565b60005481565b600160a060020a031660009081526002602052604090205490565b600554600160a060020a0316331461057057600080fd5b600160a060020a03821660009081526002602052604090205481111561059557600080fd5b600160a060020a0382166000908152600260205260409020546105b890826106d8565b600160a060020a03831660009081526002602052604090205560045461042590826106d8565b3360009081526002602052604081205482118015906106165750600160a060020a038316600090815260026020526040902054828101115b156106885733600081815260026020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a35060016103bf565b5060006103bf565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b60008282016105318482108015906106d35750838210155b6106ec565b60006106e6838311156106ec565b50900390565b8015156106f857600080fd5b505600a165627a7a7230582094bf5a65f2ce7ac17c6d727ad5b5363685eae6eddc86c977a8b6c17761fa0c690029`

// DeployReserveToken deploys a new Ethereum contract, binding an instance of ReserveToken to it.
func DeployReserveToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReserveToken, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReserveTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReserveToken{ReserveTokenCaller: ReserveTokenCaller{contract: contract}, ReserveTokenTransactor: ReserveTokenTransactor{contract: contract}, ReserveTokenFilterer: ReserveTokenFilterer{contract: contract}}, nil
}

// ReserveToken is an auto generated Go binding around an Ethereum contract.
type ReserveToken struct {
	ReserveTokenCaller     // Read-only binding to the contract
	ReserveTokenTransactor // Write-only binding to the contract
	ReserveTokenFilterer   // Log filterer for contract events
}

// ReserveTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveTokenSession struct {
	Contract     *ReserveToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveTokenCallerSession struct {
	Contract *ReserveTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ReserveTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTokenTransactorSession struct {
	Contract     *ReserveTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReserveTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveTokenRaw struct {
	Contract *ReserveToken // Generic contract binding to access the raw methods on
}

// ReserveTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveTokenCallerRaw struct {
	Contract *ReserveTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTokenTransactorRaw struct {
	Contract *ReserveTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveToken creates a new instance of ReserveToken, bound to a specific deployed contract.
func NewReserveToken(address common.Address, backend bind.ContractBackend) (*ReserveToken, error) {
	contract, err := bindReserveToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveToken{ReserveTokenCaller: ReserveTokenCaller{contract: contract}, ReserveTokenTransactor: ReserveTokenTransactor{contract: contract}, ReserveTokenFilterer: ReserveTokenFilterer{contract: contract}}, nil
}

// NewReserveTokenCaller creates a new read-only instance of ReserveToken, bound to a specific deployed contract.
func NewReserveTokenCaller(address common.Address, caller bind.ContractCaller) (*ReserveTokenCaller, error) {
	contract, err := bindReserveToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenCaller{contract: contract}, nil
}

// NewReserveTokenTransactor creates a new write-only instance of ReserveToken, bound to a specific deployed contract.
func NewReserveTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTokenTransactor, error) {
	contract, err := bindReserveToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenTransactor{contract: contract}, nil
}

// NewReserveTokenFilterer creates a new log filterer instance of ReserveToken, bound to a specific deployed contract.
func NewReserveTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveTokenFilterer, error) {
	contract, err := bindReserveToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenFilterer{contract: contract}, nil
}

// bindReserveToken binds a generic wrapper to an already deployed contract.
func bindReserveToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveToken *ReserveTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveToken.Contract.ReserveTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveToken *ReserveTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveToken.Contract.ReserveTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveToken *ReserveTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveToken.Contract.ReserveTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveToken *ReserveTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveToken *ReserveTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveToken *ReserveTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.Allowance(&_ReserveToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_ReserveToken *ReserveTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.Allowance(&_ReserveToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.BalanceOf(&_ReserveToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_ReserveToken *ReserveTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ReserveToken.Contract.BalanceOf(&_ReserveToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenSession) Decimals() (*big.Int, error) {
	return _ReserveToken.Contract.Decimals(&_ReserveToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_ReserveToken *ReserveTokenCallerSession) Decimals() (*big.Int, error) {
	return _ReserveToken.Contract.Decimals(&_ReserveToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "minter")
	return *ret0, err
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenSession) Minter() (common.Address, error) {
	return _ReserveToken.Contract.Minter(&_ReserveToken.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_ReserveToken *ReserveTokenCallerSession) Minter() (common.Address, error) {
	return _ReserveToken.Contract.Minter(&_ReserveToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenSession) Name() (string, error) {
	return _ReserveToken.Contract.Name(&_ReserveToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_ReserveToken *ReserveTokenCallerSession) Name() (string, error) {
	return _ReserveToken.Contract.Name(&_ReserveToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenSession) TotalSupply() (*big.Int, error) {
	return _ReserveToken.Contract.TotalSupply(&_ReserveToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ReserveToken *ReserveTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _ReserveToken.Contract.TotalSupply(&_ReserveToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Approve(&_ReserveToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Approve(&_ReserveToken.TransactOpts, _spender, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactor) Create(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "create", account, amount)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenSession) Create(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Create(&_ReserveToken.TransactOpts, account, amount)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactorSession) Create(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Create(&_ReserveToken.TransactOpts, account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactor) Destroy(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "destroy", account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Destroy(&_ReserveToken.TransactOpts, account, amount)
}

// Destroy is a paid mutator transaction binding the contract method 0xa24835d1.
//
// Solidity: function destroy(account address, amount uint256) returns()
func (_ReserveToken *ReserveTokenTransactorSession) Destroy(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Destroy(&_ReserveToken.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Transfer(&_ReserveToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.Transfer(&_ReserveToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.TransferFrom(&_ReserveToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_ReserveToken *ReserveTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ReserveToken.Contract.TransferFrom(&_ReserveToken.TransactOpts, _from, _to, _value)
}

// ReserveTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ReserveToken contract.
type ReserveTokenApprovalIterator struct {
	Event *ReserveTokenApproval // Event containing the contract specifics and raw log

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
func (it *ReserveTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTokenApproval)
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
		it.Event = new(ReserveTokenApproval)
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
func (it *ReserveTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTokenApproval represents a Approval event raised by the ReserveToken contract.
type ReserveTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ReserveToken *ReserveTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*ReserveTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ReserveToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenApprovalIterator{contract: _ReserveToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_ReserveToken *ReserveTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ReserveTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _ReserveToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTokenApproval)
				if err := _ReserveToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ReserveTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ReserveToken contract.
type ReserveTokenTransferIterator struct {
	Event *ReserveTokenTransfer // Event containing the contract specifics and raw log

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
func (it *ReserveTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTokenTransfer)
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
		it.Event = new(ReserveTokenTransfer)
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
func (it *ReserveTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTokenTransfer represents a Transfer event raised by the ReserveToken contract.
type ReserveTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ReserveToken *ReserveTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*ReserveTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ReserveToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenTransferIterator{contract: _ReserveToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_ReserveToken *ReserveTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ReserveTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _ReserveToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTokenTransfer)
				if err := _ReserveToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a7230582053f02e6179492ffb473cf02705dd4371fc0349629c9818cf90b3815d7a02d09e0029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405234801561001057600080fd5b50610550806100206000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610092578063095ea7b31461011c57806318160ddd1461015457806323b872dd1461017b578063313ce567146101a557806370a08231146101ba578063a9059cbb146101db578063dd62ed3e146101ff575b600080fd5b34801561009e57600080fd5b506100a7610226565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e15781810151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012857600080fd5b50610140600160a060020a03600435166024356102b3565b604080519115158252519081900360200190f35b34801561016057600080fd5b5061016961031a565b60408051918252519081900360200190f35b34801561018757600080fd5b50610140600160a060020a0360043581169060243516604435610320565b3480156101b157600080fd5b50610169610426565b3480156101c657600080fd5b50610169600160a060020a036004351661042c565b3480156101e757600080fd5b50610140600160a060020a0360043516602435610447565b34801561020b57600080fd5b50610169600160a060020a03600435811690602435166104f9565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102ab5780601f10610280576101008083540402835291602001916102ab565b820191906000526020600020905b81548152906001019060200180831161028e57829003601f168201915b505050505081565b336000818152600360209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60045481565b600160a060020a038316600090815260026020526040812054821180159061036b5750600160a060020a03841660009081526003602090815260408083203384529091529020548211155b80156103905750600160a060020a038316600090815260026020526040902054828101115b1561041b57600160a060020a03808416600081815260026020908152604080832080548801905593881680835284832080548890039055600382528483203384528252918490208054879003905583518681529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600161041f565b5060005b9392505050565b60005481565b600160a060020a031660009081526002602052604090205490565b33600090815260026020526040812054821180159061047f5750600160a060020a038316600090815260026020526040902054828101115b156104f15733600081815260026020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3506001610314565b506000610314565b600160a060020a039182166000908152600360209081526040808320939094168252919091522054905600a165627a7a723058201fc4bbfb4a0dd82a7c2e44de57bd815e3a213b98cc99277c04299f563b269d940029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenSession) Decimals() (*big.Int, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Decimals() (*big.Int, error) {
	return _StandardToken.Contract.Decimals(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StandardToken *StandardTokenCallerSession) Name() (string, error) {
	return _StandardToken.Contract.Name(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StandardTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x608060405234801561001057600080fd5b50610311806100206000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610092578063095ea7b31461011c57806318160ddd1461016157806323b872dd14610188578063313ce567146101bf57806370a08231146101d4578063a9059cbb1461011c578063dd62ed3e14610202575b600080fd5b34801561009e57600080fd5b506100a7610236565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e15781810151838201526020016100c9565b50505050905090810190601f16801561010e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561012857600080fd5b5061014d73ffffffffffffffffffffffffffffffffffffffff600435166024356102c3565b604080519115158252519081900360200190f35b34801561016d57600080fd5b506101766102cb565b60408051918252519081900360200190f35b34801561019457600080fd5b5061014d73ffffffffffffffffffffffffffffffffffffffff600435811690602435166044356102d0565b3480156101cb57600080fd5b506101766102d9565b3480156101e057600080fd5b5061017673ffffffffffffffffffffffffffffffffffffffff600435166102df565b34801561020e57600080fd5b5061017673ffffffffffffffffffffffffffffffffffffffff600435811690602435166102c3565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156102bb5780601f10610290576101008083540402835291602001916102bb565b820191906000526020600020905b81548152906001019060200180831161029e57829003601f168201915b505050505081565b600092915050565b600090565b60009392505050565b60005481565b506000905600a165627a7a7230582022c05692e16156e989acbd7e0c328f48adb6f4694c3e4a2c1c59bf3d76d102570029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(remaining uint256)
func (_Token *TokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_Token *TokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCallerSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(supply uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*TokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*TokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
