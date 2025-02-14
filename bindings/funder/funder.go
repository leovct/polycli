// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package funder

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

// FunderMetaData contains all meta data concerning the Funder contract.
var FunderMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"amount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkFund\",\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"fund\",\"inputs\":[{\"name\":\"_address\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506040516108eb3803806108eb833981810160405281019061003191906100b6565b5f8111610073576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161006a90610161565b60405180910390fd5b805f819055505061017f565b5f80fd5b5f819050919050565b61009581610083565b811461009f575f80fd5b50565b5f815190506100b08161008c565b92915050565b5f602082840312156100cb576100ca61007f565b5b5f6100d8848285016100a2565b91505092915050565b5f82825260208201905092915050565b7f5468652066756e64696e6720616d6f756e742073686f756c64206265206772655f8201527f61746572207468616e207a65726f000000000000000000000000000000000000602082015250565b5f61014b602e836100e1565b9150610156826100f1565b604082019050919050565b5f6020820190508181035f8301526101788161013f565b9050919050565b61075f8061018c5f395ff3fe608060405260043610610037575f3560e01c80632302440814610042578063a4626b851461006a578063aa8c217c146100925761003e565b3661003e57005b5f80fd5b34801561004d575f80fd5b5061006860048036038101906100639190610323565b6100bc565b005b348015610075575f80fd5b50610090600480360381019061008b91906103af565b61021b565b005b34801561009d575f80fd5b506100a66102bc565b6040516100b39190610412565b60405180910390f35b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361012a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610121906104ab565b60405180910390fd5b5f5447101561016e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016590610539565b60405180910390fd5b5f8173ffffffffffffffffffffffffffffffffffffffff165f5460405161019490610584565b5f6040518083038185875af1925050503d805f81146101ce576040519150601f19603f3d011682016040523d82523d5f602084013e6101d3565b606091505b5050905080610217576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020e906105e2565b60405180910390fd5b5050565b818190505f5461022b919061062d565b47101561026d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610264906106de565b60405180910390fd5b5f5b828290508110156102b7576102aa8383838181106102905761028f6106fc565b5b90506020020160208101906102a59190610323565b6100bc565b808060010191505061026f565b505050565b5f5481565b5f80fd5b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6102f2826102c9565b9050919050565b610302816102e8565b811461030c575f80fd5b50565b5f8135905061031d816102f9565b92915050565b5f60208284031215610338576103376102c1565b5b5f6103458482850161030f565b91505092915050565b5f80fd5b5f80fd5b5f80fd5b5f8083601f84011261036f5761036e61034e565b5b8235905067ffffffffffffffff81111561038c5761038b610352565b5b6020830191508360208202830111156103a8576103a7610356565b5b9250929050565b5f80602083850312156103c5576103c46102c1565b5b5f83013567ffffffffffffffff8111156103e2576103e16102c5565b5b6103ee8582860161035a565b92509250509250929050565b5f819050919050565b61040c816103fa565b82525050565b5f6020820190506104255f830184610403565b92915050565b5f82825260208201905092915050565b7f5468652066756e64656420616464726573732073686f756c64206265206469665f8201527f666572656e74207468616e20746865207a65726f206164647265737300000000602082015250565b5f610495603c8361042b565b91506104a08261043b565b604082019050919050565b5f6020820190508181035f8301526104c281610489565b9050919050565b7f496e73756666696369656e7420636f6e74726163742062616c616e636520666f5f8201527f722066756e64696e670000000000000000000000000000000000000000000000602082015250565b5f61052360298361042b565b915061052e826104c9565b604082019050919050565b5f6020820190508181035f83015261055081610517565b9050919050565b5f81905092915050565b50565b5f61056f5f83610557565b915061057a82610561565b5f82019050919050565b5f61058e82610564565b9150819050919050565b7f46756e64696e67206661696c65640000000000000000000000000000000000005f82015250565b5f6105cc600e8361042b565b91506105d782610598565b602082019050919050565b5f6020820190508181035f8301526105f9816105c0565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610637826103fa565b9150610642836103fa565b9250828202610650816103fa565b9150828204841483151761066757610666610600565b5b5092915050565b7f496e73756666696369656e7420636f6e74726163742062616c616e636520666f5f8201527f722062617463682066756e64696e670000000000000000000000000000000000602082015250565b5f6106c8602f8361042b565b91506106d38261066e565b604082019050919050565b5f6020820190508181035f8301526106f5816106bc565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea264697066735822122030232f3adc0ffd6f6839ea5d84cd5d5bda77ec7657c8f80d81d8f6139a9bdf6664736f6c63430008170033",
}

// FunderABI is the input ABI used to generate the binding from.
// Deprecated: Use FunderMetaData.ABI instead.
var FunderABI = FunderMetaData.ABI

// FunderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FunderMetaData.Bin instead.
var FunderBin = FunderMetaData.Bin

// DeployFunder deploys a new Ethereum contract, binding an instance of Funder to it.
func DeployFunder(auth *bind.TransactOpts, backend bind.ContractBackend, _amount *big.Int) (common.Address, *types.Transaction, *Funder, error) {
	parsed, err := FunderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FunderBin), backend, _amount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Funder{FunderCaller: FunderCaller{contract: contract}, FunderTransactor: FunderTransactor{contract: contract}, FunderFilterer: FunderFilterer{contract: contract}}, nil
}

// Funder is an auto generated Go binding around an Ethereum contract.
type Funder struct {
	FunderCaller     // Read-only binding to the contract
	FunderTransactor // Write-only binding to the contract
	FunderFilterer   // Log filterer for contract events
}

// FunderCaller is an auto generated read-only Go binding around an Ethereum contract.
type FunderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FunderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FunderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FunderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FunderSession struct {
	Contract     *Funder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FunderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FunderCallerSession struct {
	Contract *FunderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FunderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FunderTransactorSession struct {
	Contract     *FunderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FunderRaw is an auto generated low-level Go binding around an Ethereum contract.
type FunderRaw struct {
	Contract *Funder // Generic contract binding to access the raw methods on
}

// FunderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FunderCallerRaw struct {
	Contract *FunderCaller // Generic read-only contract binding to access the raw methods on
}

// FunderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FunderTransactorRaw struct {
	Contract *FunderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFunder creates a new instance of Funder, bound to a specific deployed contract.
func NewFunder(address common.Address, backend bind.ContractBackend) (*Funder, error) {
	contract, err := bindFunder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Funder{FunderCaller: FunderCaller{contract: contract}, FunderTransactor: FunderTransactor{contract: contract}, FunderFilterer: FunderFilterer{contract: contract}}, nil
}

// NewFunderCaller creates a new read-only instance of Funder, bound to a specific deployed contract.
func NewFunderCaller(address common.Address, caller bind.ContractCaller) (*FunderCaller, error) {
	contract, err := bindFunder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FunderCaller{contract: contract}, nil
}

// NewFunderTransactor creates a new write-only instance of Funder, bound to a specific deployed contract.
func NewFunderTransactor(address common.Address, transactor bind.ContractTransactor) (*FunderTransactor, error) {
	contract, err := bindFunder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FunderTransactor{contract: contract}, nil
}

// NewFunderFilterer creates a new log filterer instance of Funder, bound to a specific deployed contract.
func NewFunderFilterer(address common.Address, filterer bind.ContractFilterer) (*FunderFilterer, error) {
	contract, err := bindFunder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FunderFilterer{contract: contract}, nil
}

// bindFunder binds a generic wrapper to an already deployed contract.
func bindFunder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FunderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Funder *FunderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Funder.Contract.FunderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Funder *FunderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Funder.Contract.FunderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Funder *FunderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Funder.Contract.FunderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Funder *FunderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Funder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Funder *FunderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Funder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Funder *FunderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Funder.Contract.contract.Transact(opts, method, params...)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Funder *FunderCaller) Amount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Funder.contract.Call(opts, &out, "amount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Funder *FunderSession) Amount() (*big.Int, error) {
	return _Funder.Contract.Amount(&_Funder.CallOpts)
}

// Amount is a free data retrieval call binding the contract method 0xaa8c217c.
//
// Solidity: function amount() view returns(uint256)
func (_Funder *FunderCallerSession) Amount() (*big.Int, error) {
	return _Funder.Contract.Amount(&_Funder.CallOpts)
}

// BulkFund is a paid mutator transaction binding the contract method 0xa4626b85.
//
// Solidity: function bulkFund(address[] _addresses) returns()
func (_Funder *FunderTransactor) BulkFund(opts *bind.TransactOpts, _addresses []common.Address) (*types.Transaction, error) {
	return _Funder.contract.Transact(opts, "bulkFund", _addresses)
}

// BulkFund is a paid mutator transaction binding the contract method 0xa4626b85.
//
// Solidity: function bulkFund(address[] _addresses) returns()
func (_Funder *FunderSession) BulkFund(_addresses []common.Address) (*types.Transaction, error) {
	return _Funder.Contract.BulkFund(&_Funder.TransactOpts, _addresses)
}

// BulkFund is a paid mutator transaction binding the contract method 0xa4626b85.
//
// Solidity: function bulkFund(address[] _addresses) returns()
func (_Funder *FunderTransactorSession) BulkFund(_addresses []common.Address) (*types.Transaction, error) {
	return _Funder.Contract.BulkFund(&_Funder.TransactOpts, _addresses)
}

// Fund is a paid mutator transaction binding the contract method 0x23024408.
//
// Solidity: function fund(address _address) returns()
func (_Funder *FunderTransactor) Fund(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Funder.contract.Transact(opts, "fund", _address)
}

// Fund is a paid mutator transaction binding the contract method 0x23024408.
//
// Solidity: function fund(address _address) returns()
func (_Funder *FunderSession) Fund(_address common.Address) (*types.Transaction, error) {
	return _Funder.Contract.Fund(&_Funder.TransactOpts, _address)
}

// Fund is a paid mutator transaction binding the contract method 0x23024408.
//
// Solidity: function fund(address _address) returns()
func (_Funder *FunderTransactorSession) Fund(_address common.Address) (*types.Transaction, error) {
	return _Funder.Contract.Fund(&_Funder.TransactOpts, _address)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Funder *FunderTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Funder.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Funder *FunderSession) Receive() (*types.Transaction, error) {
	return _Funder.Contract.Receive(&_Funder.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Funder *FunderTransactorSession) Receive() (*types.Transaction, error) {
	return _Funder.Contract.Receive(&_Funder.TransactOpts)
}
