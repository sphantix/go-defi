// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package virtualfpair

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

// FPairMetaData contains all meta data concerning the FPair contract.
var FPairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceALast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceBLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FPairABI is the input ABI used to generate the binding from.
// Deprecated: Use FPairMetaData.ABI instead.
var FPairABI = FPairMetaData.ABI

// FPair is an auto generated Go binding around an Ethereum contract.
type FPair struct {
	FPairCaller     // Read-only binding to the contract
	FPairTransactor // Write-only binding to the contract
	FPairFilterer   // Log filterer for contract events
}

// FPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type FPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FPairSession struct {
	Contract     *FPair            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FPairCallerSession struct {
	Contract *FPairCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FPairTransactorSession struct {
	Contract     *FPairTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type FPairRaw struct {
	Contract *FPair // Generic contract binding to access the raw methods on
}

// FPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FPairCallerRaw struct {
	Contract *FPairCaller // Generic read-only contract binding to access the raw methods on
}

// FPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FPairTransactorRaw struct {
	Contract *FPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFPair creates a new instance of FPair, bound to a specific deployed contract.
func NewFPair(address common.Address, backend bind.ContractBackend) (*FPair, error) {
	contract, err := bindFPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FPair{FPairCaller: FPairCaller{contract: contract}, FPairTransactor: FPairTransactor{contract: contract}, FPairFilterer: FPairFilterer{contract: contract}}, nil
}

// NewFPairCaller creates a new read-only instance of FPair, bound to a specific deployed contract.
func NewFPairCaller(address common.Address, caller bind.ContractCaller) (*FPairCaller, error) {
	contract, err := bindFPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FPairCaller{contract: contract}, nil
}

// NewFPairTransactor creates a new write-only instance of FPair, bound to a specific deployed contract.
func NewFPairTransactor(address common.Address, transactor bind.ContractTransactor) (*FPairTransactor, error) {
	contract, err := bindFPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FPairTransactor{contract: contract}, nil
}

// NewFPairFilterer creates a new log filterer instance of FPair, bound to a specific deployed contract.
func NewFPairFilterer(address common.Address, filterer bind.ContractFilterer) (*FPairFilterer, error) {
	contract, err := bindFPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FPairFilterer{contract: contract}, nil
}

// bindFPair binds a generic wrapper to an already deployed contract.
func bindFPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FPairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FPair *FPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FPair.Contract.FPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FPair *FPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FPair.Contract.FPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FPair *FPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FPair.Contract.FPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FPair *FPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FPair *FPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FPair *FPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FPair.Contract.contract.Transact(opts, method, params...)
}

// AssetBalance is a free data retrieval call binding the contract method 0xc66f2455.
//
// Solidity: function assetBalance() view returns(uint256)
func (_FPair *FPairCaller) AssetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "assetBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetBalance is a free data retrieval call binding the contract method 0xc66f2455.
//
// Solidity: function assetBalance() view returns(uint256)
func (_FPair *FPairSession) AssetBalance() (*big.Int, error) {
	return _FPair.Contract.AssetBalance(&_FPair.CallOpts)
}

// AssetBalance is a free data retrieval call binding the contract method 0xc66f2455.
//
// Solidity: function assetBalance() view returns(uint256)
func (_FPair *FPairCallerSession) AssetBalance() (*big.Int, error) {
	return _FPair.Contract.AssetBalance(&_FPair.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_FPair *FPairCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_FPair *FPairSession) Balance() (*big.Int, error) {
	return _FPair.Contract.Balance(&_FPair.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_FPair *FPairCallerSession) Balance() (*big.Int, error) {
	return _FPair.Contract.Balance(&_FPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_FPair *FPairCaller) GetReserves(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "getReserves")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_FPair *FPairSession) GetReserves() (*big.Int, *big.Int, error) {
	return _FPair.Contract.GetReserves(&_FPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256, uint256)
func (_FPair *FPairCallerSession) GetReserves() (*big.Int, *big.Int, error) {
	return _FPair.Contract.GetReserves(&_FPair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FPair *FPairCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FPair *FPairSession) KLast() (*big.Int, error) {
	return _FPair.Contract.KLast(&_FPair.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_FPair *FPairCallerSession) KLast() (*big.Int, error) {
	return _FPair.Contract.KLast(&_FPair.CallOpts)
}

// PriceALast is a free data retrieval call binding the contract method 0x5c9d6938.
//
// Solidity: function priceALast() view returns(uint256)
func (_FPair *FPairCaller) PriceALast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "priceALast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceALast is a free data retrieval call binding the contract method 0x5c9d6938.
//
// Solidity: function priceALast() view returns(uint256)
func (_FPair *FPairSession) PriceALast() (*big.Int, error) {
	return _FPair.Contract.PriceALast(&_FPair.CallOpts)
}

// PriceALast is a free data retrieval call binding the contract method 0x5c9d6938.
//
// Solidity: function priceALast() view returns(uint256)
func (_FPair *FPairCallerSession) PriceALast() (*big.Int, error) {
	return _FPair.Contract.PriceALast(&_FPair.CallOpts)
}

// PriceBLast is a free data retrieval call binding the contract method 0x0e06dfc9.
//
// Solidity: function priceBLast() view returns(uint256)
func (_FPair *FPairCaller) PriceBLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "priceBLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceBLast is a free data retrieval call binding the contract method 0x0e06dfc9.
//
// Solidity: function priceBLast() view returns(uint256)
func (_FPair *FPairSession) PriceBLast() (*big.Int, error) {
	return _FPair.Contract.PriceBLast(&_FPair.CallOpts)
}

// PriceBLast is a free data retrieval call binding the contract method 0x0e06dfc9.
//
// Solidity: function priceBLast() view returns(uint256)
func (_FPair *FPairCallerSession) PriceBLast() (*big.Int, error) {
	return _FPair.Contract.PriceBLast(&_FPair.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FPair *FPairCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FPair *FPairSession) Router() (common.Address, error) {
	return _FPair.Contract.Router(&_FPair.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_FPair *FPairCallerSession) Router() (common.Address, error) {
	return _FPair.Contract.Router(&_FPair.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_FPair *FPairCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_FPair *FPairSession) TokenA() (common.Address, error) {
	return _FPair.Contract.TokenA(&_FPair.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_FPair *FPairCallerSession) TokenA() (common.Address, error) {
	return _FPair.Contract.TokenA(&_FPair.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_FPair *FPairCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FPair.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_FPair *FPairSession) TokenB() (common.Address, error) {
	return _FPair.Contract.TokenB(&_FPair.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_FPair *FPairCallerSession) TokenB() (common.Address, error) {
	return _FPair.Contract.TokenB(&_FPair.CallOpts)
}

// Approval is a paid mutator transaction binding the contract method 0x5c52a5f2.
//
// Solidity: function approval(address _user, address _token, uint256 amount) returns(bool)
func (_FPair *FPairTransactor) Approval(opts *bind.TransactOpts, _user common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.contract.Transact(opts, "approval", _user, _token, amount)
}

// Approval is a paid mutator transaction binding the contract method 0x5c52a5f2.
//
// Solidity: function approval(address _user, address _token, uint256 amount) returns(bool)
func (_FPair *FPairSession) Approval(_user common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Approval(&_FPair.TransactOpts, _user, _token, amount)
}

// Approval is a paid mutator transaction binding the contract method 0x5c52a5f2.
//
// Solidity: function approval(address _user, address _token, uint256 amount) returns(bool)
func (_FPair *FPairTransactorSession) Approval(_user common.Address, _token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Approval(&_FPair.TransactOpts, _user, _token, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 reserve0, uint256 reserve1) returns(bool)
func (_FPair *FPairTransactor) Mint(opts *bind.TransactOpts, reserve0 *big.Int, reserve1 *big.Int) (*types.Transaction, error) {
	return _FPair.contract.Transact(opts, "mint", reserve0, reserve1)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 reserve0, uint256 reserve1) returns(bool)
func (_FPair *FPairSession) Mint(reserve0 *big.Int, reserve1 *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Mint(&_FPair.TransactOpts, reserve0, reserve1)
}

// Mint is a paid mutator transaction binding the contract method 0x1b2ef1ca.
//
// Solidity: function mint(uint256 reserve0, uint256 reserve1) returns(bool)
func (_FPair *FPairTransactorSession) Mint(reserve0 *big.Int, reserve1 *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Mint(&_FPair.TransactOpts, reserve0, reserve1)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out) returns(bool)
func (_FPair *FPairTransactor) Swap(opts *bind.TransactOpts, amount0In *big.Int, amount0Out *big.Int, amount1In *big.Int, amount1Out *big.Int) (*types.Transaction, error) {
	return _FPair.contract.Transact(opts, "swap", amount0In, amount0Out, amount1In, amount1Out)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out) returns(bool)
func (_FPair *FPairSession) Swap(amount0In *big.Int, amount0Out *big.Int, amount1In *big.Int, amount1Out *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Swap(&_FPair.TransactOpts, amount0In, amount0Out, amount1In, amount1Out)
}

// Swap is a paid mutator transaction binding the contract method 0x5673b02d.
//
// Solidity: function swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out) returns(bool)
func (_FPair *FPairTransactorSession) Swap(amount0In *big.Int, amount0Out *big.Int, amount1In *big.Int, amount1Out *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.Swap(&_FPair.TransactOpts, amount0In, amount0Out, amount1In, amount1Out)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5c921eb9.
//
// Solidity: function transferAsset(address recipient, uint256 amount) returns()
func (_FPair *FPairTransactor) TransferAsset(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.contract.Transact(opts, "transferAsset", recipient, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5c921eb9.
//
// Solidity: function transferAsset(address recipient, uint256 amount) returns()
func (_FPair *FPairSession) TransferAsset(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.TransferAsset(&_FPair.TransactOpts, recipient, amount)
}

// TransferAsset is a paid mutator transaction binding the contract method 0x5c921eb9.
//
// Solidity: function transferAsset(address recipient, uint256 amount) returns()
func (_FPair *FPairTransactorSession) TransferAsset(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.TransferAsset(&_FPair.TransactOpts, recipient, amount)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address recipient, uint256 amount) returns()
func (_FPair *FPairTransactor) TransferTo(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.contract.Transact(opts, "transferTo", recipient, amount)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address recipient, uint256 amount) returns()
func (_FPair *FPairSession) TransferTo(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.TransferTo(&_FPair.TransactOpts, recipient, amount)
}

// TransferTo is a paid mutator transaction binding the contract method 0x2ccb1b30.
//
// Solidity: function transferTo(address recipient, uint256 amount) returns()
func (_FPair *FPairTransactorSession) TransferTo(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FPair.Contract.TransferTo(&_FPair.TransactOpts, recipient, amount)
}

// FPairMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the FPair contract.
type FPairMintIterator struct {
	Event *FPairMint // Event containing the contract specifics and raw log

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
func (it *FPairMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FPairMint)
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
		it.Event = new(FPairMint)
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
func (it *FPairMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FPairMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FPairMint represents a Mint event raised by the FPair contract.
type FPairMint struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 reserve0, uint256 reserve1)
func (_FPair *FPairFilterer) FilterMint(opts *bind.FilterOpts) (*FPairMintIterator, error) {

	logs, sub, err := _FPair.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &FPairMintIterator{contract: _FPair.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 reserve0, uint256 reserve1)
func (_FPair *FPairFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *FPairMint) (event.Subscription, error) {

	logs, sub, err := _FPair.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FPairMint)
				if err := _FPair.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xcc9c58b575eabd3f6a1ee653e91fcea3ff546867ffc3782a3bbca1f9b6dbb8df.
//
// Solidity: event Mint(uint256 reserve0, uint256 reserve1)
func (_FPair *FPairFilterer) ParseMint(log types.Log) (*FPairMint, error) {
	event := new(FPairMint)
	if err := _FPair.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the FPair contract.
type FPairSwapIterator struct {
	Event *FPairSwap // Event containing the contract specifics and raw log

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
func (it *FPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FPairSwap)
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
		it.Event = new(FPairSwap)
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
func (it *FPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FPairSwap represents a Swap event raised by the FPair contract.
type FPairSwap struct {
	Amount0In  *big.Int
	Amount0Out *big.Int
	Amount1In  *big.Int
	Amount1Out *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x298c349c742327269dc8de6ad66687767310c948ea309df826f5bd103e19d207.
//
// Solidity: event Swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out)
func (_FPair *FPairFilterer) FilterSwap(opts *bind.FilterOpts) (*FPairSwapIterator, error) {

	logs, sub, err := _FPair.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &FPairSwapIterator{contract: _FPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x298c349c742327269dc8de6ad66687767310c948ea309df826f5bd103e19d207.
//
// Solidity: event Swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out)
func (_FPair *FPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *FPairSwap) (event.Subscription, error) {

	logs, sub, err := _FPair.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FPairSwap)
				if err := _FPair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x298c349c742327269dc8de6ad66687767310c948ea309df826f5bd103e19d207.
//
// Solidity: event Swap(uint256 amount0In, uint256 amount0Out, uint256 amount1In, uint256 amount1Out)
func (_FPair *FPairFilterer) ParseSwap(log types.Log) (*FPairSwap, error) {
	event := new(FPairSwap)
	if err := _FPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
