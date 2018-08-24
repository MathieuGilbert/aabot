// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dice

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

// DSSafeAddSubABI is the input ABI used to generate the binding from.
const DSSafeAddSubABI = "[]"

// DSSafeAddSubBin is the compiled bytecode used for deploying new contracts.
const DSSafeAddSubBin = `0x6080604052348015600f57600080fd5b50603580601d6000396000f3006080604052600080fd00a165627a7a723058209e35ea9a5a9f3478c318dbc1a728346aa56a0a3ea20b56262df6967debfa82810029`

// DeployDSSafeAddSub deploys a new Ethereum contract, binding an instance of DSSafeAddSub to it.
func DeployDSSafeAddSub(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DSSafeAddSub, error) {
	parsed, err := abi.JSON(strings.NewReader(DSSafeAddSubABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DSSafeAddSubBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DSSafeAddSub{DSSafeAddSubCaller: DSSafeAddSubCaller{contract: contract}, DSSafeAddSubTransactor: DSSafeAddSubTransactor{contract: contract}, DSSafeAddSubFilterer: DSSafeAddSubFilterer{contract: contract}}, nil
}

// DSSafeAddSub is an auto generated Go binding around an Ethereum contract.
type DSSafeAddSub struct {
	DSSafeAddSubCaller     // Read-only binding to the contract
	DSSafeAddSubTransactor // Write-only binding to the contract
	DSSafeAddSubFilterer   // Log filterer for contract events
}

// DSSafeAddSubCaller is an auto generated read-only Go binding around an Ethereum contract.
type DSSafeAddSubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSSafeAddSubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DSSafeAddSubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSSafeAddSubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DSSafeAddSubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DSSafeAddSubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DSSafeAddSubSession struct {
	Contract     *DSSafeAddSub     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DSSafeAddSubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DSSafeAddSubCallerSession struct {
	Contract *DSSafeAddSubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DSSafeAddSubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DSSafeAddSubTransactorSession struct {
	Contract     *DSSafeAddSubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DSSafeAddSubRaw is an auto generated low-level Go binding around an Ethereum contract.
type DSSafeAddSubRaw struct {
	Contract *DSSafeAddSub // Generic contract binding to access the raw methods on
}

// DSSafeAddSubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DSSafeAddSubCallerRaw struct {
	Contract *DSSafeAddSubCaller // Generic read-only contract binding to access the raw methods on
}

// DSSafeAddSubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DSSafeAddSubTransactorRaw struct {
	Contract *DSSafeAddSubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDSSafeAddSub creates a new instance of DSSafeAddSub, bound to a specific deployed contract.
func NewDSSafeAddSub(address common.Address, backend bind.ContractBackend) (*DSSafeAddSub, error) {
	contract, err := bindDSSafeAddSub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DSSafeAddSub{DSSafeAddSubCaller: DSSafeAddSubCaller{contract: contract}, DSSafeAddSubTransactor: DSSafeAddSubTransactor{contract: contract}, DSSafeAddSubFilterer: DSSafeAddSubFilterer{contract: contract}}, nil
}

// NewDSSafeAddSubCaller creates a new read-only instance of DSSafeAddSub, bound to a specific deployed contract.
func NewDSSafeAddSubCaller(address common.Address, caller bind.ContractCaller) (*DSSafeAddSubCaller, error) {
	contract, err := bindDSSafeAddSub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DSSafeAddSubCaller{contract: contract}, nil
}

// NewDSSafeAddSubTransactor creates a new write-only instance of DSSafeAddSub, bound to a specific deployed contract.
func NewDSSafeAddSubTransactor(address common.Address, transactor bind.ContractTransactor) (*DSSafeAddSubTransactor, error) {
	contract, err := bindDSSafeAddSub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DSSafeAddSubTransactor{contract: contract}, nil
}

// NewDSSafeAddSubFilterer creates a new log filterer instance of DSSafeAddSub, bound to a specific deployed contract.
func NewDSSafeAddSubFilterer(address common.Address, filterer bind.ContractFilterer) (*DSSafeAddSubFilterer, error) {
	contract, err := bindDSSafeAddSub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DSSafeAddSubFilterer{contract: contract}, nil
}

// bindDSSafeAddSub binds a generic wrapper to an already deployed contract.
func bindDSSafeAddSub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DSSafeAddSubABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSSafeAddSub *DSSafeAddSubRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSSafeAddSub.Contract.DSSafeAddSubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSSafeAddSub *DSSafeAddSubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSSafeAddSub.Contract.DSSafeAddSubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSSafeAddSub *DSSafeAddSubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSSafeAddSub.Contract.DSSafeAddSubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DSSafeAddSub *DSSafeAddSubCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DSSafeAddSub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DSSafeAddSub *DSSafeAddSubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DSSafeAddSub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DSSafeAddSub *DSSafeAddSubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DSSafeAddSub.Contract.contract.Transact(opts, method, params...)
}

// EtherollTokenABI is the input ABI used to generate the binding from.
const EtherollTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_newPriviledgedAddress\",\"type\":\"address\"}],\"name\":\"ownerSetPriviledgedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateTokenStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokensFrozen\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextThaw\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"priviledgedAddressBurnUnsoldCoins\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"crowdfundDeadline\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextFreeze\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"priviledgedAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"ownerTransferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"Frozen\",\"type\":\"bool\"}],\"name\":\"LogTokensFrozen\",\"type\":\"event\"}]"

// EtherollTokenBin is the compiled bytecode used for deploying new contracts.
const EtherollTokenBin = `0x60c0604052600960808190527f546f6b656e20312e30000000000000000000000000000000000000000000000060a090815261003e916001919061014c565b506040805180820190915260048082527f444943450000000000000000000000000000000000000000000000000000000060209092019182526100839160029161014c565b506040805180820190915260038082527f524f4c000000000000000000000000000000000000000000000000000000000060209092019182526100c6918161014c565b506004805460ff191660101790556934f086f3b33b684000006005556212750042908101600755626ebe0081016008556277f8800160095534801561010a57600080fd5b5060008054600160a060020a0319163390811782558152600a602052604090206934f086f3b33b6840000090556006805460a060020a60ff02191690556101e7565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061018d57805160ff19168380011785556101ba565b828001600101855582156101ba579182015b828111156101ba57825182559160200191906001019061019f565b506101c69291506101ca565b5090565b6101e491905b808211156101c657600081556001016101d0565b90565b610abe806101f66000396000f3006080604052600436106101115763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663040ea2f4811461011657806306fdde0314610139578063095ea7b3146101c357806318160ddd146101fb57806323b872dd1461022257806325ffba041461024c5780632efa79d314610261578063313ce5671461027657806358e879f3146102a15780635a3b7e42146102b65780636410c41a146102cb57806370a08231146102e05780638480021c146103015780638b64574b146103165780638da5cb5b1461032b578063937c0cdf1461035c57806395d89b4114610371578063a09cca9314610386578063a9059cbb146103a7578063dd62ed3e146103cb575b600080fd5b34801561012257600080fd5b50610137600160a060020a03600435166103f2565b005b34801561014557600080fd5b5061014e610438565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610188578181015183820152602001610170565b50505050905090810190601f1680156101b55780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101cf57600080fd5b506101e7600160a060020a03600435166024356104c3565b604080519115158252519081900360200190f35b34801561020757600080fd5b5061021061052a565b60408051918252519081900360200190f35b34801561022e57600080fd5b506101e7600160a060020a0360043581169060243516604435610530565b34801561025857600080fd5b50610137610674565b34801561026d57600080fd5b506101e76107b1565b34801561028257600080fd5b5061028b6107c1565b6040805160ff9092168252519081900360200190f35b3480156102ad57600080fd5b506102106107ca565b3480156102c257600080fd5b5061014e6107d0565b3480156102d757600080fd5b5061013761082a565b3480156102ec57600080fd5b50610210600160a060020a036004351661088b565b34801561030d57600080fd5b5061021061089d565b34801561032257600080fd5b506102106108a3565b34801561033757600080fd5b506103406108a9565b60408051600160a060020a039092168252519081900360200190f35b34801561036857600080fd5b506103406108b8565b34801561037d57600080fd5b5061014e6108c7565b34801561039257600080fd5b50610137600160a060020a0360043516610922565b3480156103b357600080fd5b506101e7600160a060020a0360043516602435610968565b3480156103d757600080fd5b50610210600160a060020a0360043581169060243516610a53565b600054600160a060020a0316331461040957600080fd5b6006805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156104bb5780601f10610490576101008083540402835291602001916104bb565b820191906000526020600020905b81548152906001019060200180831161049e57829003601f168201915b505050505081565b336000818152600b60209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a35060015b92915050565b60055481565b60065460009060a060020a900460ff1680156105575750600654600160a060020a03163314155b156105645750600061066d565b600160a060020a0384166000908152600a602052604090205482111561058c5750600061066d565b600160a060020a0383166000908152600a602052604090205482810110156105b65750600061066d565b600160a060020a0384166000908152600b602090815260408083203384529091529020548211156105e95750600061066d565b600160a060020a038085166000818152600a6020908152604080832080548890039055938716808352848320805488019055838352600b8252848320338452825291849020805487900390558351868152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35060015b9392505050565b6007544210156106d8576006805474ff0000000000000000000000000000000000000000191660a060020a90811791829055604051910460ff161515907fd3ae281451496b0a8d978ad222f5465aa683499b2ca2698ec1686a413e717f0790600090a25b600854421061073b576006805474ff0000000000000000000000000000000000000000191660a060020a90811791829055604051910460ff161515907fd3ae281451496b0a8d978ad222f5465aa683499b2ca2698ec1686a413e717f0790600090a25b60095442106107af576006805474ff0000000000000000000000000000000000000000191690819055626ebe00429081016008556277f8800160095560405160a060020a90910460ff161515907fd3ae281451496b0a8d978ad222f5465aa683499b2ca2698ec1686a413e717f0790600090a25b565b60065460a060020a900460ff1681565b60045460ff1681565b60095481565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104bb5780601f10610490576101008083540402835291602001916104bb565b600654600160a060020a031633811461084257600080fd5b600554600654600160a060020a03166000908152600a602052604090205461086a9190610a70565b60055550600654600160a060020a03166000908152600a6020526040812055565b600a6020526000908152604090205481565b60075481565b60085481565b600054600160a060020a031681565b600654600160a060020a031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156104bb5780601f10610490576101008083540402835291602001916104bb565b600054600160a060020a0316331461093957600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60065460009060a060020a900460ff16801561098f5750600654600160a060020a03163314155b1561099c57506000610524565b336000908152600a60205260409020548211156109bb57506000610524565b600160a060020a0383166000908152600a602052604090205482810110156109e557506000610524565b336000818152600a6020908152604080832080548790039055600160a060020a03871680845292819020805487019055805186815290519293927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a350600192915050565b600b60209081526000928352604080842090915290825290205481565b6000610a7c8383610a8d565b1515610a8757600080fd5b50900390565b1115905600a165627a7a72305820c10c273958cb125439bda1a92280ddbb42692783cfa8e1df6f6611bb651c65e40029`

// DeployEtherollToken deploys a new Ethereum contract, binding an instance of EtherollToken to it.
func DeployEtherollToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtherollToken, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherollTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EtherollTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherollToken{EtherollTokenCaller: EtherollTokenCaller{contract: contract}, EtherollTokenTransactor: EtherollTokenTransactor{contract: contract}, EtherollTokenFilterer: EtherollTokenFilterer{contract: contract}}, nil
}

// EtherollToken is an auto generated Go binding around an Ethereum contract.
type EtherollToken struct {
	EtherollTokenCaller     // Read-only binding to the contract
	EtherollTokenTransactor // Write-only binding to the contract
	EtherollTokenFilterer   // Log filterer for contract events
}

// EtherollTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherollTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherollTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherollTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherollTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherollTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherollTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherollTokenSession struct {
	Contract     *EtherollToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherollTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherollTokenCallerSession struct {
	Contract *EtherollTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EtherollTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherollTokenTransactorSession struct {
	Contract     *EtherollTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EtherollTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherollTokenRaw struct {
	Contract *EtherollToken // Generic contract binding to access the raw methods on
}

// EtherollTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherollTokenCallerRaw struct {
	Contract *EtherollTokenCaller // Generic read-only contract binding to access the raw methods on
}

// EtherollTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherollTokenTransactorRaw struct {
	Contract *EtherollTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherollToken creates a new instance of EtherollToken, bound to a specific deployed contract.
func NewEtherollToken(address common.Address, backend bind.ContractBackend) (*EtherollToken, error) {
	contract, err := bindEtherollToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherollToken{EtherollTokenCaller: EtherollTokenCaller{contract: contract}, EtherollTokenTransactor: EtherollTokenTransactor{contract: contract}, EtherollTokenFilterer: EtherollTokenFilterer{contract: contract}}, nil
}

// NewEtherollTokenCaller creates a new read-only instance of EtherollToken, bound to a specific deployed contract.
func NewEtherollTokenCaller(address common.Address, caller bind.ContractCaller) (*EtherollTokenCaller, error) {
	contract, err := bindEtherollToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenCaller{contract: contract}, nil
}

// NewEtherollTokenTransactor creates a new write-only instance of EtherollToken, bound to a specific deployed contract.
func NewEtherollTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherollTokenTransactor, error) {
	contract, err := bindEtherollToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenTransactor{contract: contract}, nil
}

// NewEtherollTokenFilterer creates a new log filterer instance of EtherollToken, bound to a specific deployed contract.
func NewEtherollTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherollTokenFilterer, error) {
	contract, err := bindEtherollToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenFilterer{contract: contract}, nil
}

// bindEtherollToken binds a generic wrapper to an already deployed contract.
func bindEtherollToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EtherollTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherollToken *EtherollTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherollToken.Contract.EtherollTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherollToken *EtherollTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherollToken.Contract.EtherollTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherollToken *EtherollTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherollToken.Contract.EtherollTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherollToken *EtherollTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EtherollToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherollToken *EtherollTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherollToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherollToken *EtherollTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherollToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherollToken.Contract.Allowance(&_EtherollToken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _EtherollToken.Contract.Allowance(&_EtherollToken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EtherollToken.Contract.BalanceOf(&_EtherollToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _EtherollToken.Contract.BalanceOf(&_EtherollToken.CallOpts, arg0)
}

// CrowdfundDeadline is a free data retrieval call binding the contract method 0x8480021c.
//
// Solidity: function crowdfundDeadline() constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) CrowdfundDeadline(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "crowdfundDeadline")
	return *ret0, err
}

// CrowdfundDeadline is a free data retrieval call binding the contract method 0x8480021c.
//
// Solidity: function crowdfundDeadline() constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) CrowdfundDeadline() (*big.Int, error) {
	return _EtherollToken.Contract.CrowdfundDeadline(&_EtherollToken.CallOpts)
}

// CrowdfundDeadline is a free data retrieval call binding the contract method 0x8480021c.
//
// Solidity: function crowdfundDeadline() constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) CrowdfundDeadline() (*big.Int, error) {
	return _EtherollToken.Contract.CrowdfundDeadline(&_EtherollToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EtherollToken *EtherollTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EtherollToken *EtherollTokenSession) Decimals() (uint8, error) {
	return _EtherollToken.Contract.Decimals(&_EtherollToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_EtherollToken *EtherollTokenCallerSession) Decimals() (uint8, error) {
	return _EtherollToken.Contract.Decimals(&_EtherollToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EtherollToken *EtherollTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EtherollToken *EtherollTokenSession) Name() (string, error) {
	return _EtherollToken.Contract.Name(&_EtherollToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_EtherollToken *EtherollTokenCallerSession) Name() (string, error) {
	return _EtherollToken.Contract.Name(&_EtherollToken.CallOpts)
}

// NextFreeze is a free data retrieval call binding the contract method 0x8b64574b.
//
// Solidity: function nextFreeze() constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) NextFreeze(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "nextFreeze")
	return *ret0, err
}

// NextFreeze is a free data retrieval call binding the contract method 0x8b64574b.
//
// Solidity: function nextFreeze() constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) NextFreeze() (*big.Int, error) {
	return _EtherollToken.Contract.NextFreeze(&_EtherollToken.CallOpts)
}

// NextFreeze is a free data retrieval call binding the contract method 0x8b64574b.
//
// Solidity: function nextFreeze() constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) NextFreeze() (*big.Int, error) {
	return _EtherollToken.Contract.NextFreeze(&_EtherollToken.CallOpts)
}

// NextThaw is a free data retrieval call binding the contract method 0x58e879f3.
//
// Solidity: function nextThaw() constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) NextThaw(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "nextThaw")
	return *ret0, err
}

// NextThaw is a free data retrieval call binding the contract method 0x58e879f3.
//
// Solidity: function nextThaw() constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) NextThaw() (*big.Int, error) {
	return _EtherollToken.Contract.NextThaw(&_EtherollToken.CallOpts)
}

// NextThaw is a free data retrieval call binding the contract method 0x58e879f3.
//
// Solidity: function nextThaw() constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) NextThaw() (*big.Int, error) {
	return _EtherollToken.Contract.NextThaw(&_EtherollToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EtherollToken *EtherollTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EtherollToken *EtherollTokenSession) Owner() (common.Address, error) {
	return _EtherollToken.Contract.Owner(&_EtherollToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_EtherollToken *EtherollTokenCallerSession) Owner() (common.Address, error) {
	return _EtherollToken.Contract.Owner(&_EtherollToken.CallOpts)
}

// PriviledgedAddress is a free data retrieval call binding the contract method 0x937c0cdf.
//
// Solidity: function priviledgedAddress() constant returns(address)
func (_EtherollToken *EtherollTokenCaller) PriviledgedAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "priviledgedAddress")
	return *ret0, err
}

// PriviledgedAddress is a free data retrieval call binding the contract method 0x937c0cdf.
//
// Solidity: function priviledgedAddress() constant returns(address)
func (_EtherollToken *EtherollTokenSession) PriviledgedAddress() (common.Address, error) {
	return _EtherollToken.Contract.PriviledgedAddress(&_EtherollToken.CallOpts)
}

// PriviledgedAddress is a free data retrieval call binding the contract method 0x937c0cdf.
//
// Solidity: function priviledgedAddress() constant returns(address)
func (_EtherollToken *EtherollTokenCallerSession) PriviledgedAddress() (common.Address, error) {
	return _EtherollToken.Contract.PriviledgedAddress(&_EtherollToken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_EtherollToken *EtherollTokenCaller) Standard(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "standard")
	return *ret0, err
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_EtherollToken *EtherollTokenSession) Standard() (string, error) {
	return _EtherollToken.Contract.Standard(&_EtherollToken.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() constant returns(string)
func (_EtherollToken *EtherollTokenCallerSession) Standard() (string, error) {
	return _EtherollToken.Contract.Standard(&_EtherollToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EtherollToken *EtherollTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EtherollToken *EtherollTokenSession) Symbol() (string, error) {
	return _EtherollToken.Contract.Symbol(&_EtherollToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_EtherollToken *EtherollTokenCallerSession) Symbol() (string, error) {
	return _EtherollToken.Contract.Symbol(&_EtherollToken.CallOpts)
}

// TokensFrozen is a free data retrieval call binding the contract method 0x2efa79d3.
//
// Solidity: function tokensFrozen() constant returns(bool)
func (_EtherollToken *EtherollTokenCaller) TokensFrozen(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "tokensFrozen")
	return *ret0, err
}

// TokensFrozen is a free data retrieval call binding the contract method 0x2efa79d3.
//
// Solidity: function tokensFrozen() constant returns(bool)
func (_EtherollToken *EtherollTokenSession) TokensFrozen() (bool, error) {
	return _EtherollToken.Contract.TokensFrozen(&_EtherollToken.CallOpts)
}

// TokensFrozen is a free data retrieval call binding the contract method 0x2efa79d3.
//
// Solidity: function tokensFrozen() constant returns(bool)
func (_EtherollToken *EtherollTokenCallerSession) TokensFrozen() (bool, error) {
	return _EtherollToken.Contract.TokensFrozen(&_EtherollToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherollToken *EtherollTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EtherollToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherollToken *EtherollTokenSession) TotalSupply() (*big.Int, error) {
	return _EtherollToken.Contract.TotalSupply(&_EtherollToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_EtherollToken *EtherollTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _EtherollToken.Contract.TotalSupply(&_EtherollToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.Approve(&_EtherollToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.Approve(&_EtherollToken.TransactOpts, _spender, _value)
}

// OwnerSetPriviledgedAddress is a paid mutator transaction binding the contract method 0x040ea2f4.
//
// Solidity: function ownerSetPriviledgedAddress(_newPriviledgedAddress address) returns()
func (_EtherollToken *EtherollTokenTransactor) OwnerSetPriviledgedAddress(opts *bind.TransactOpts, _newPriviledgedAddress common.Address) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "ownerSetPriviledgedAddress", _newPriviledgedAddress)
}

// OwnerSetPriviledgedAddress is a paid mutator transaction binding the contract method 0x040ea2f4.
//
// Solidity: function ownerSetPriviledgedAddress(_newPriviledgedAddress address) returns()
func (_EtherollToken *EtherollTokenSession) OwnerSetPriviledgedAddress(_newPriviledgedAddress common.Address) (*types.Transaction, error) {
	return _EtherollToken.Contract.OwnerSetPriviledgedAddress(&_EtherollToken.TransactOpts, _newPriviledgedAddress)
}

// OwnerSetPriviledgedAddress is a paid mutator transaction binding the contract method 0x040ea2f4.
//
// Solidity: function ownerSetPriviledgedAddress(_newPriviledgedAddress address) returns()
func (_EtherollToken *EtherollTokenTransactorSession) OwnerSetPriviledgedAddress(_newPriviledgedAddress common.Address) (*types.Transaction, error) {
	return _EtherollToken.Contract.OwnerSetPriviledgedAddress(&_EtherollToken.TransactOpts, _newPriviledgedAddress)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_EtherollToken *EtherollTokenTransactor) OwnerTransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "ownerTransferOwnership", newOwner)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_EtherollToken *EtherollTokenSession) OwnerTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherollToken.Contract.OwnerTransferOwnership(&_EtherollToken.TransactOpts, newOwner)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_EtherollToken *EtherollTokenTransactorSession) OwnerTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EtherollToken.Contract.OwnerTransferOwnership(&_EtherollToken.TransactOpts, newOwner)
}

// PriviledgedAddressBurnUnsoldCoins is a paid mutator transaction binding the contract method 0x6410c41a.
//
// Solidity: function priviledgedAddressBurnUnsoldCoins() returns()
func (_EtherollToken *EtherollTokenTransactor) PriviledgedAddressBurnUnsoldCoins(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "priviledgedAddressBurnUnsoldCoins")
}

// PriviledgedAddressBurnUnsoldCoins is a paid mutator transaction binding the contract method 0x6410c41a.
//
// Solidity: function priviledgedAddressBurnUnsoldCoins() returns()
func (_EtherollToken *EtherollTokenSession) PriviledgedAddressBurnUnsoldCoins() (*types.Transaction, error) {
	return _EtherollToken.Contract.PriviledgedAddressBurnUnsoldCoins(&_EtherollToken.TransactOpts)
}

// PriviledgedAddressBurnUnsoldCoins is a paid mutator transaction binding the contract method 0x6410c41a.
//
// Solidity: function priviledgedAddressBurnUnsoldCoins() returns()
func (_EtherollToken *EtherollTokenTransactorSession) PriviledgedAddressBurnUnsoldCoins() (*types.Transaction, error) {
	return _EtherollToken.Contract.PriviledgedAddressBurnUnsoldCoins(&_EtherollToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.Transfer(&_EtherollToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.Transfer(&_EtherollToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.TransferFrom(&_EtherollToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_EtherollToken *EtherollTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EtherollToken.Contract.TransferFrom(&_EtherollToken.TransactOpts, _from, _to, _value)
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0x25ffba04.
//
// Solidity: function updateTokenStatus() returns()
func (_EtherollToken *EtherollTokenTransactor) UpdateTokenStatus(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherollToken.contract.Transact(opts, "updateTokenStatus")
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0x25ffba04.
//
// Solidity: function updateTokenStatus() returns()
func (_EtherollToken *EtherollTokenSession) UpdateTokenStatus() (*types.Transaction, error) {
	return _EtherollToken.Contract.UpdateTokenStatus(&_EtherollToken.TransactOpts)
}

// UpdateTokenStatus is a paid mutator transaction binding the contract method 0x25ffba04.
//
// Solidity: function updateTokenStatus() returns()
func (_EtherollToken *EtherollTokenTransactorSession) UpdateTokenStatus() (*types.Transaction, error) {
	return _EtherollToken.Contract.UpdateTokenStatus(&_EtherollToken.TransactOpts)
}

// EtherollTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EtherollToken contract.
type EtherollTokenApprovalIterator struct {
	Event *EtherollTokenApproval // Event containing the contract specifics and raw log

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
func (it *EtherollTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherollTokenApproval)
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
		it.Event = new(EtherollTokenApproval)
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
func (it *EtherollTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherollTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherollTokenApproval represents a Approval event raised by the EtherollToken contract.
type EtherollTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_EtherollToken *EtherollTokenFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*EtherollTokenApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _EtherollToken.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenApprovalIterator{contract: _EtherollToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _spender indexed address, _value uint256)
func (_EtherollToken *EtherollTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EtherollTokenApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _EtherollToken.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherollTokenApproval)
				if err := _EtherollToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// EtherollTokenLogTokensFrozenIterator is returned from FilterLogTokensFrozen and is used to iterate over the raw logs and unpacked data for LogTokensFrozen events raised by the EtherollToken contract.
type EtherollTokenLogTokensFrozenIterator struct {
	Event *EtherollTokenLogTokensFrozen // Event containing the contract specifics and raw log

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
func (it *EtherollTokenLogTokensFrozenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherollTokenLogTokensFrozen)
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
		it.Event = new(EtherollTokenLogTokensFrozen)
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
func (it *EtherollTokenLogTokensFrozenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherollTokenLogTokensFrozenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherollTokenLogTokensFrozen represents a LogTokensFrozen event raised by the EtherollToken contract.
type EtherollTokenLogTokensFrozen struct {
	Frozen bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogTokensFrozen is a free log retrieval operation binding the contract event 0xd3ae281451496b0a8d978ad222f5465aa683499b2ca2698ec1686a413e717f07.
//
// Solidity: e LogTokensFrozen(Frozen indexed bool)
func (_EtherollToken *EtherollTokenFilterer) FilterLogTokensFrozen(opts *bind.FilterOpts, Frozen []bool) (*EtherollTokenLogTokensFrozenIterator, error) {

	var FrozenRule []interface{}
	for _, FrozenItem := range Frozen {
		FrozenRule = append(FrozenRule, FrozenItem)
	}

	logs, sub, err := _EtherollToken.contract.FilterLogs(opts, "LogTokensFrozen", FrozenRule)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenLogTokensFrozenIterator{contract: _EtherollToken.contract, event: "LogTokensFrozen", logs: logs, sub: sub}, nil
}

// WatchLogTokensFrozen is a free log subscription operation binding the contract event 0xd3ae281451496b0a8d978ad222f5465aa683499b2ca2698ec1686a413e717f07.
//
// Solidity: e LogTokensFrozen(Frozen indexed bool)
func (_EtherollToken *EtherollTokenFilterer) WatchLogTokensFrozen(opts *bind.WatchOpts, sink chan<- *EtherollTokenLogTokensFrozen, Frozen []bool) (event.Subscription, error) {

	var FrozenRule []interface{}
	for _, FrozenItem := range Frozen {
		FrozenRule = append(FrozenRule, FrozenItem)
	}

	logs, sub, err := _EtherollToken.contract.WatchLogs(opts, "LogTokensFrozen", FrozenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherollTokenLogTokensFrozen)
				if err := _EtherollToken.contract.UnpackLog(event, "LogTokensFrozen", log); err != nil {
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

// EtherollTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EtherollToken contract.
type EtherollTokenTransferIterator struct {
	Event *EtherollTokenTransfer // Event containing the contract specifics and raw log

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
func (it *EtherollTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EtherollTokenTransfer)
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
		it.Event = new(EtherollTokenTransfer)
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
func (it *EtherollTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EtherollTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EtherollTokenTransfer represents a Transfer event raised by the EtherollToken contract.
type EtherollTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_EtherollToken *EtherollTokenFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*EtherollTokenTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _EtherollToken.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &EtherollTokenTransferIterator{contract: _EtherollToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _value uint256)
func (_EtherollToken *EtherollTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EtherollTokenTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _EtherollToken.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EtherollTokenTransfer)
				if err := _EtherollToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// OwnedABI is the input ABI used to generate the binding from.
const OwnedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"ownerTransferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnedBin is the compiled bytecode used for deploying new contracts.
const OwnedBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610166806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063a09cca931461008e575b600080fd5b34801561005c57600080fd5b506100656100be565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b34801561009a57600080fd5b506100bc73ffffffffffffffffffffffffffffffffffffffff600435166100da565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b60005473ffffffffffffffffffffffffffffffffffffffff1633146100fe57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff929092169190911790555600a165627a7a72305820fab53500ade5bdaa309f06111fdb87c3b59144c939878365faf95ff8d7d0f1750029`

// DeployOwned deploys a new Ethereum contract, binding an instance of Owned to it.
func DeployOwned(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Owned, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// Owned is an auto generated Go binding around an Ethereum contract.
type Owned struct {
	OwnedCaller     // Read-only binding to the contract
	OwnedTransactor // Write-only binding to the contract
	OwnedFilterer   // Log filterer for contract events
}

// OwnedCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnedSession struct {
	Contract     *Owned            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnedCallerSession struct {
	Contract *OwnedCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OwnedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnedTransactorSession struct {
	Contract     *OwnedTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnedRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnedRaw struct {
	Contract *Owned // Generic contract binding to access the raw methods on
}

// OwnedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnedCallerRaw struct {
	Contract *OwnedCaller // Generic read-only contract binding to access the raw methods on
}

// OwnedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnedTransactorRaw struct {
	Contract *OwnedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwned creates a new instance of Owned, bound to a specific deployed contract.
func NewOwned(address common.Address, backend bind.ContractBackend) (*Owned, error) {
	contract, err := bindOwned(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Owned{OwnedCaller: OwnedCaller{contract: contract}, OwnedTransactor: OwnedTransactor{contract: contract}, OwnedFilterer: OwnedFilterer{contract: contract}}, nil
}

// NewOwnedCaller creates a new read-only instance of Owned, bound to a specific deployed contract.
func NewOwnedCaller(address common.Address, caller bind.ContractCaller) (*OwnedCaller, error) {
	contract, err := bindOwned(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedCaller{contract: contract}, nil
}

// NewOwnedTransactor creates a new write-only instance of Owned, bound to a specific deployed contract.
func NewOwnedTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnedTransactor, error) {
	contract, err := bindOwned(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnedTransactor{contract: contract}, nil
}

// NewOwnedFilterer creates a new log filterer instance of Owned, bound to a specific deployed contract.
func NewOwnedFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnedFilterer, error) {
	contract, err := bindOwned(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnedFilterer{contract: contract}, nil
}

// bindOwned binds a generic wrapper to an already deployed contract.
func bindOwned(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.OwnedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.OwnedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Owned *OwnedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Owned.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Owned *OwnedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Owned *OwnedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Owned.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Owned.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Owned *OwnedCallerSession) Owner() (common.Address, error) {
	return _Owned.Contract.Owner(&_Owned.CallOpts)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactor) OwnerTransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Owned.contract.Transact(opts, "ownerTransferOwnership", newOwner)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_Owned *OwnedSession) OwnerTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.OwnerTransferOwnership(&_Owned.TransactOpts, newOwner)
}

// OwnerTransferOwnership is a paid mutator transaction binding the contract method 0xa09cca93.
//
// Solidity: function ownerTransferOwnership(newOwner address) returns()
func (_Owned *OwnedTransactorSession) OwnerTransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Owned.Contract.OwnerTransferOwnership(&_Owned.TransactOpts, newOwner)
}
