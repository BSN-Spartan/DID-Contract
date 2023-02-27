// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DidContract

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
)

// DidContractMetaData contains all meta data concerning the DidContract contract.
var DidContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"createDidRetLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msg\",\"type\":\"bytes\"}],\"name\":\"getInitializeDataLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"updateDidAuthRetLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"didDocument\",\"type\":\"string\"}],\"name\":\"createDid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDocument\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"didDocument\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"isDidExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"didDocument\",\"type\":\"string\"}],\"name\":\"updateDocument\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"msgcode\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060006100216100c460201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506100cc565b600033905090565b6119ec806100db6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146104cf578063b0d8e36914610503578063e51af990146105d4578063f2fde38b1461080b5761007d565b80634d300e0c14610082578063715018a6146102b95780637ccb6a64146102c3575b600080fd5b6101d26004803603604081101561009857600080fd5b81019080803590602001906401000000008111156100b557600080fd5b8201836020820111156100c757600080fd5b803590602001918460018302840111640100000000831117156100e957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561014c57600080fd5b82018360208201111561015e57600080fd5b8035906020019184600183028401116401000000008311171561018057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061084f565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156102165780820151818401526020810190506101fb565b50505050905090810190601f1680156102435780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b8381101561027c578082015181840152602081019050610261565b50505050905090810190601f1680156102a95780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6102c1610c9f565b005b61037c600480360360208110156102d957600080fd5b81019080803590602001906401000000008111156102f657600080fd5b82018360208201111561030857600080fd5b8035906020019184600183028401116401000000008311171561032a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610e0c565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b838110156103c45780820151818401526020810190506103a9565b50505050905090810190601f1680156103f15780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b8381101561042a57808201518184015260208101905061040f565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b83811015610490578082015181840152602081019050610475565b50505050905090810190601f1680156104bd5780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b6104d7610f47565b604051808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105bc6004803603602081101561051957600080fd5b810190808035906020019064010000000081111561053657600080fd5b82018360208201111561054857600080fd5b8035906020019184600183028401116401000000008311171561056a57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610f70565b60405180821515815260200191505060405180910390f35b610724600480360360408110156105ea57600080fd5b810190808035906020019064010000000081111561060757600080fd5b82018360208201111561061957600080fd5b8035906020019184600183028401116401000000008311171561063b57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192908035906020019064010000000081111561069e57600080fd5b8201836020820111156106b057600080fd5b803590602001918460018302840111640100000000831117156106d257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611008565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561076857808201518184015260208101905061074d565b50505050905090810190601f1680156107955780820380516001836020036101000a031916815260200191505b50838103825284818151815260200191508051906020019080838360005b838110156107ce5780820151818401526020810190506107b3565b50505050905090810190601f1680156107fb5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b61084d6004803603602081101561082157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506113f2565b005b606080606084905060608490506108646115e4565b73ffffffffffffffffffffffffffffffffffffffff166004836040518082805190602001908083835b602083106108b0578051825260208201915060208101905060208303925061088d565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610970576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602881526020018061198f6028913960400191505060405180910390fd5b600061097b87610f70565b905080610b04577f9af1b91de8cb6c51fb6a8211f6b04fc6c6c3642f77f3e417b9dd5f82f9e860a56040518060400160405280600481526020017f3230313100000000000000000000000000000000000000000000000000000000815250604051808060200180602001838103835284818151815260200191508051906020019080838360005b83811015610a1d578082015181840152602081019050610a02565b50505050905090810190601f168015610a4a5780820380516001836020036101000a031916815260200191505b508381038252600d8152602001807f646964206e6f7420657869737400000000000000000000000000000000000000815250602001935050505060405180910390a16040518060400160405280600481526020017f32303131000000000000000000000000000000000000000000000000000000008152506040518060400160405280600d81526020017f646964206e6f742065786973740000000000000000000000000000000000000081525094509450505050610c98565b610b1a838360016115ec9092919063ffffffff16565b7f9af1b91de8cb6c51fb6a8211f6b04fc6c6c3642f77f3e417b9dd5f82f9e860a56040518060400160405280600481526020017f3030303000000000000000000000000000000000000000000000000000000000815250604051808060200180602001838103835284818151815260200191508051906020019080838360005b83811015610bb5578082015181840152602081019050610b9a565b50505050905090810190601f168015610be25780820380516001836020036101000a031916815260200191505b50838103825260078152602001807f7375636365737300000000000000000000000000000000000000000000000000815250602001935050505060405180910390a16040518060400160405280600481526020017f30303030000000000000000000000000000000000000000000000000000000008152506040518060400160405280600781526020017f7375636365737300000000000000000000000000000000000000000000000000815250945094505050505b9250929050565b610ca76115e4565b73ffffffffffffffffffffffffffffffffffffffff16610cc5610f47565b73ffffffffffffffffffffffffffffffffffffffff1614610d4e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60608060606000610e1c85610f70565b905080610eab576040518060400160405280600481526020017f32303131000000000000000000000000000000000000000000000000000000008152506040518060400160405280600d81526020017f646964206e6f74206578697374000000000000000000000000000000000000008152506040518060200160405280600081525093509350935050610f40565b60608590506060610ec682600161179590919063ffffffff16565b90506040518060400160405280600481526020017f3030303000000000000000000000000000000000000000000000000000000000815250816040518060400160405280600781526020017f7375636365737300000000000000000000000000000000000000000000000000815250909550955095505050505b9193909250565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000606082905060006001600001826040518082805190602001908083835b60208310610fb25780518252602082019150602081019050602083039250610f8f565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205490506000811415610ffc57600092505050611003565b6001925050505b919050565b606080600061101685610f70565b9050801561119e577f12bc3c9c8a99ef08acd8f1191e57dc6246466d47e920eba9634c46bf19c9cbc96040518060400160405280600481526020017f3230313200000000000000000000000000000000000000000000000000000000815250604051808060200180602001838103835284818151815260200191508051906020019080838360005b838110156110b957808201518184015260208101905061109e565b50505050905090810190601f1680156110e65780820380516001836020036101000a031916815260200191505b50838103825260118152602001807f64696420616c7265616479206578697374000000000000000000000000000000815250602001935050505060405180910390a16040518060400160405280600481526020017f32303132000000000000000000000000000000000000000000000000000000008152506040518060400160405280601181526020017f64696420616c726561647920657869737400000000000000000000000000000081525092509250506113eb565b606085905060608590506111be828260016115ec9092919063ffffffff16565b6111c66115e4565b6004836040518082805190602001908083835b602083106111fc57805182526020820191506020810190506020830392506111d9565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f12bc3c9c8a99ef08acd8f1191e57dc6246466d47e920eba9634c46bf19c9cbc96040518060400160405280600481526020017f3030303000000000000000000000000000000000000000000000000000000000815250604051808060200180602001838103835284818151815260200191508051906020019080838360005b838110156113085780820151818401526020810190506112ed565b50505050905090810190601f1680156113355780820380516001836020036101000a031916815260200191505b50838103825260078152602001807f7375636365737300000000000000000000000000000000000000000000000000815250602001935050505060405180910390a16040518060400160405280600481526020017f30303030000000000000000000000000000000000000000000000000000000008152506040518060400160405280600781526020017f7375636365737300000000000000000000000000000000000000000000000000815250945094505050505b9250929050565b6113fa6115e4565b73ffffffffffffffffffffffffffffffffffffffff16611418610f47565b73ffffffffffffffffffffffffffffffffffffffff16146114a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611527576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806119696026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b600083600001836040518082805190602001908083835b602083106116265780518252602082019150602081019050602083039250611603565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050600081141561175d5783600101839080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906116a29291906118cb565b5083600201829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906116e09291906118cb565b50836001018054905084600001846040518082805190602001908083835b6020831061172157805182526020820191506020810190506020830392506116fe565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208190555061178f565b8184600201600183038154811061177057fe5b90600052602060002001908051906020019061178d9291906118cb565b505b50505050565b6060600083600001836040518082805190602001908083835b602083106117d157805182526020820191506020810190506020830392506117ae565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020549050606084600201600183038154811061181b57fe5b906000526020600020018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156118b95780601f1061188e576101008083540402835291602001916118b9565b820191906000526020600020905b81548152906001019060200180831161189c57829003601f168201915b50505050509050809250505092915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061190c57805160ff191683800117855561193a565b8280016001018555821561193a579182015b8281111561193957825182559160200191906001019061191e565b5b509050611947919061194b565b5090565b5b8082111561196457600081600090555060010161194c565b509056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373446964436f6e74726163743a2063616c6c6572206973206e6f742074686520506f73736573736f72a2646970667358221220148aa6d0d4091265fc4265dee89b6ac27df55768d852b8b9f9db877903d414c164736f6c634300060c0033",
}

// DidContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DidContractMetaData.ABI instead.
var DidContractABI = DidContractMetaData.ABI

// DidContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DidContractMetaData.Bin instead.
var DidContractBin = DidContractMetaData.Bin

// DeployDidContract deploys a new Ethereum contract, binding an instance of DidContract to it.
func DeployDidContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DidContract, error) {
	parsed, err := DidContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DidContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DidContract{DidContractCaller: DidContractCaller{contract: contract}, DidContractTransactor: DidContractTransactor{contract: contract}, DidContractFilterer: DidContractFilterer{contract: contract}}, nil
}

// DidContract is an auto generated Go binding around an Ethereum contract.
type DidContract struct {
	DidContractCaller     // Read-only binding to the contract
	DidContractTransactor // Write-only binding to the contract
	DidContractFilterer   // Log filterer for contract events
}

// DidContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DidContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DidContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DidContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DidContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DidContractSession struct {
	Contract     *DidContract      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DidContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DidContractCallerSession struct {
	Contract *DidContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DidContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DidContractTransactorSession struct {
	Contract     *DidContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DidContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DidContractRaw struct {
	Contract *DidContract // Generic contract binding to access the raw methods on
}

// DidContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DidContractCallerRaw struct {
	Contract *DidContractCaller // Generic read-only contract binding to access the raw methods on
}

// DidContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DidContractTransactorRaw struct {
	Contract *DidContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDidContract creates a new instance of DidContract, bound to a specific deployed contract.
func NewDidContract(address common.Address, backend bind.ContractBackend) (*DidContract, error) {
	contract, err := bindDidContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DidContract{DidContractCaller: DidContractCaller{contract: contract}, DidContractTransactor: DidContractTransactor{contract: contract}, DidContractFilterer: DidContractFilterer{contract: contract}}, nil
}

// NewDidContractCaller creates a new read-only instance of DidContract, bound to a specific deployed contract.
func NewDidContractCaller(address common.Address, caller bind.ContractCaller) (*DidContractCaller, error) {
	contract, err := bindDidContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DidContractCaller{contract: contract}, nil
}

// NewDidContractTransactor creates a new write-only instance of DidContract, bound to a specific deployed contract.
func NewDidContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DidContractTransactor, error) {
	contract, err := bindDidContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DidContractTransactor{contract: contract}, nil
}

// NewDidContractFilterer creates a new log filterer instance of DidContract, bound to a specific deployed contract.
func NewDidContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DidContractFilterer, error) {
	contract, err := bindDidContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DidContractFilterer{contract: contract}, nil
}

// bindDidContract binds a generic wrapper to an already deployed contract.
func bindDidContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DidContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DidContract *DidContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DidContract.Contract.DidContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DidContract *DidContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidContract.Contract.DidContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DidContract *DidContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DidContract.Contract.DidContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DidContract *DidContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DidContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DidContract *DidContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DidContract *DidContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DidContract.Contract.contract.Transact(opts, method, params...)
}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns(string msgcode, string msg, string didDocument)
func (_DidContract *DidContractCaller) GetDocument(opts *bind.CallOpts, did string) (struct {
	Msgcode     string
	Msg         string
	DidDocument string
}, error) {
	var out []interface{}
	err := _DidContract.contract.Call(opts, &out, "getDocument", did)

	outstruct := new(struct {
		Msgcode     string
		Msg         string
		DidDocument string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Msgcode = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Msg = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.DidDocument = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns(string msgcode, string msg, string didDocument)
func (_DidContract *DidContractSession) GetDocument(did string) (struct {
	Msgcode     string
	Msg         string
	DidDocument string
}, error) {
	return _DidContract.Contract.GetDocument(&_DidContract.CallOpts, did)
}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns(string msgcode, string msg, string didDocument)
func (_DidContract *DidContractCallerSession) GetDocument(did string) (struct {
	Msgcode     string
	Msg         string
	DidDocument string
}, error) {
	return _DidContract.Contract.GetDocument(&_DidContract.CallOpts, did)
}

// IsDidExist is a free data retrieval call binding the contract method 0xb0d8e369.
//
// Solidity: function isDidExist(string did) view returns(bool)
func (_DidContract *DidContractCaller) IsDidExist(opts *bind.CallOpts, did string) (bool, error) {
	var out []interface{}
	err := _DidContract.contract.Call(opts, &out, "isDidExist", did)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDidExist is a free data retrieval call binding the contract method 0xb0d8e369.
//
// Solidity: function isDidExist(string did) view returns(bool)
func (_DidContract *DidContractSession) IsDidExist(did string) (bool, error) {
	return _DidContract.Contract.IsDidExist(&_DidContract.CallOpts, did)
}

// IsDidExist is a free data retrieval call binding the contract method 0xb0d8e369.
//
// Solidity: function isDidExist(string did) view returns(bool)
func (_DidContract *DidContractCallerSession) IsDidExist(did string) (bool, error) {
	return _DidContract.Contract.IsDidExist(&_DidContract.CallOpts, did)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidContract *DidContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DidContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidContract *DidContractSession) Owner() (common.Address, error) {
	return _DidContract.Contract.Owner(&_DidContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DidContract *DidContractCallerSession) Owner() (common.Address, error) {
	return _DidContract.Contract.Owner(&_DidContract.CallOpts)
}

// CreateDid is a paid mutator transaction binding the contract method 0xe51af990.
//
// Solidity: function createDid(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractTransactor) CreateDid(opts *bind.TransactOpts, did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.contract.Transact(opts, "createDid", did, didDocument)
}

// CreateDid is a paid mutator transaction binding the contract method 0xe51af990.
//
// Solidity: function createDid(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractSession) CreateDid(did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.Contract.CreateDid(&_DidContract.TransactOpts, did, didDocument)
}

// CreateDid is a paid mutator transaction binding the contract method 0xe51af990.
//
// Solidity: function createDid(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractTransactorSession) CreateDid(did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.Contract.CreateDid(&_DidContract.TransactOpts, did, didDocument)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidContract *DidContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DidContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidContract *DidContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _DidContract.Contract.RenounceOwnership(&_DidContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DidContract *DidContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DidContract.Contract.RenounceOwnership(&_DidContract.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidContract *DidContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DidContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidContract *DidContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DidContract.Contract.TransferOwnership(&_DidContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DidContract *DidContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DidContract.Contract.TransferOwnership(&_DidContract.TransactOpts, newOwner)
}

// UpdateDocument is a paid mutator transaction binding the contract method 0x4d300e0c.
//
// Solidity: function updateDocument(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractTransactor) UpdateDocument(opts *bind.TransactOpts, did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.contract.Transact(opts, "updateDocument", did, didDocument)
}

// UpdateDocument is a paid mutator transaction binding the contract method 0x4d300e0c.
//
// Solidity: function updateDocument(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractSession) UpdateDocument(did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.Contract.UpdateDocument(&_DidContract.TransactOpts, did, didDocument)
}

// UpdateDocument is a paid mutator transaction binding the contract method 0x4d300e0c.
//
// Solidity: function updateDocument(string did, string didDocument) returns(string msgcode, string msg)
func (_DidContract *DidContractTransactorSession) UpdateDocument(did string, didDocument string) (*types.Transaction, error) {
	return _DidContract.Contract.UpdateDocument(&_DidContract.TransactOpts, did, didDocument)
}

// DidContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DidContract contract.
type DidContractOwnershipTransferredIterator struct {
	Event *DidContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DidContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidContractOwnershipTransferred)
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
		it.Event = new(DidContractOwnershipTransferred)
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
func (it *DidContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidContractOwnershipTransferred represents a OwnershipTransferred event raised by the DidContract contract.
type DidContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DidContract *DidContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DidContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DidContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DidContractOwnershipTransferredIterator{contract: _DidContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DidContract *DidContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DidContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DidContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidContractOwnershipTransferred)
				if err := _DidContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DidContract *DidContractFilterer) ParseOwnershipTransferred(log types.Log) (*DidContractOwnershipTransferred, error) {
	event := new(DidContractOwnershipTransferred)
	if err := _DidContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DidContractCreateDidRetLogIterator is returned from FilterCreateDidRetLog and is used to iterate over the raw logs and unpacked data for CreateDidRetLog events raised by the DidContract contract.
type DidContractCreateDidRetLogIterator struct {
	Event *DidContractCreateDidRetLog // Event containing the contract specifics and raw log

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
func (it *DidContractCreateDidRetLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidContractCreateDidRetLog)
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
		it.Event = new(DidContractCreateDidRetLog)
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
func (it *DidContractCreateDidRetLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidContractCreateDidRetLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidContractCreateDidRetLog represents a CreateDidRetLog event raised by the DidContract contract.
type DidContractCreateDidRetLog struct {
	Msgcode string
	Msg     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateDidRetLog is a free log retrieval operation binding the contract event 0x12bc3c9c8a99ef08acd8f1191e57dc6246466d47e920eba9634c46bf19c9cbc9.
//
// Solidity: event createDidRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) FilterCreateDidRetLog(opts *bind.FilterOpts) (*DidContractCreateDidRetLogIterator, error) {

	logs, sub, err := _DidContract.contract.FilterLogs(opts, "createDidRetLog")
	if err != nil {
		return nil, err
	}
	return &DidContractCreateDidRetLogIterator{contract: _DidContract.contract, event: "createDidRetLog", logs: logs, sub: sub}, nil
}

// WatchCreateDidRetLog is a free log subscription operation binding the contract event 0x12bc3c9c8a99ef08acd8f1191e57dc6246466d47e920eba9634c46bf19c9cbc9.
//
// Solidity: event createDidRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) WatchCreateDidRetLog(opts *bind.WatchOpts, sink chan<- *DidContractCreateDidRetLog) (event.Subscription, error) {

	logs, sub, err := _DidContract.contract.WatchLogs(opts, "createDidRetLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidContractCreateDidRetLog)
				if err := _DidContract.contract.UnpackLog(event, "createDidRetLog", log); err != nil {
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

// ParseCreateDidRetLog is a log parse operation binding the contract event 0x12bc3c9c8a99ef08acd8f1191e57dc6246466d47e920eba9634c46bf19c9cbc9.
//
// Solidity: event createDidRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) ParseCreateDidRetLog(log types.Log) (*DidContractCreateDidRetLog, error) {
	event := new(DidContractCreateDidRetLog)
	if err := _DidContract.contract.UnpackLog(event, "createDidRetLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DidContractGetInitializeDataLogIterator is returned from FilterGetInitializeDataLog and is used to iterate over the raw logs and unpacked data for GetInitializeDataLog events raised by the DidContract contract.
type DidContractGetInitializeDataLogIterator struct {
	Event *DidContractGetInitializeDataLog // Event containing the contract specifics and raw log

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
func (it *DidContractGetInitializeDataLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidContractGetInitializeDataLog)
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
		it.Event = new(DidContractGetInitializeDataLog)
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
func (it *DidContractGetInitializeDataLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidContractGetInitializeDataLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidContractGetInitializeDataLog represents a GetInitializeDataLog event raised by the DidContract contract.
type DidContractGetInitializeDataLog struct {
	Msgcode string
	Msg     []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGetInitializeDataLog is a free log retrieval operation binding the contract event 0xbd7d81e39a13c2876502c5c2520fe0bc2494c940a8183f0fc6e0ff564935afff.
//
// Solidity: event getInitializeDataLog(string msgcode, bytes msg)
func (_DidContract *DidContractFilterer) FilterGetInitializeDataLog(opts *bind.FilterOpts) (*DidContractGetInitializeDataLogIterator, error) {

	logs, sub, err := _DidContract.contract.FilterLogs(opts, "getInitializeDataLog")
	if err != nil {
		return nil, err
	}
	return &DidContractGetInitializeDataLogIterator{contract: _DidContract.contract, event: "getInitializeDataLog", logs: logs, sub: sub}, nil
}

// WatchGetInitializeDataLog is a free log subscription operation binding the contract event 0xbd7d81e39a13c2876502c5c2520fe0bc2494c940a8183f0fc6e0ff564935afff.
//
// Solidity: event getInitializeDataLog(string msgcode, bytes msg)
func (_DidContract *DidContractFilterer) WatchGetInitializeDataLog(opts *bind.WatchOpts, sink chan<- *DidContractGetInitializeDataLog) (event.Subscription, error) {

	logs, sub, err := _DidContract.contract.WatchLogs(opts, "getInitializeDataLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidContractGetInitializeDataLog)
				if err := _DidContract.contract.UnpackLog(event, "getInitializeDataLog", log); err != nil {
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

// ParseGetInitializeDataLog is a log parse operation binding the contract event 0xbd7d81e39a13c2876502c5c2520fe0bc2494c940a8183f0fc6e0ff564935afff.
//
// Solidity: event getInitializeDataLog(string msgcode, bytes msg)
func (_DidContract *DidContractFilterer) ParseGetInitializeDataLog(log types.Log) (*DidContractGetInitializeDataLog, error) {
	event := new(DidContractGetInitializeDataLog)
	if err := _DidContract.contract.UnpackLog(event, "getInitializeDataLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DidContractUpdateDidAuthRetLogIterator is returned from FilterUpdateDidAuthRetLog and is used to iterate over the raw logs and unpacked data for UpdateDidAuthRetLog events raised by the DidContract contract.
type DidContractUpdateDidAuthRetLogIterator struct {
	Event *DidContractUpdateDidAuthRetLog // Event containing the contract specifics and raw log

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
func (it *DidContractUpdateDidAuthRetLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DidContractUpdateDidAuthRetLog)
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
		it.Event = new(DidContractUpdateDidAuthRetLog)
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
func (it *DidContractUpdateDidAuthRetLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DidContractUpdateDidAuthRetLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DidContractUpdateDidAuthRetLog represents a UpdateDidAuthRetLog event raised by the DidContract contract.
type DidContractUpdateDidAuthRetLog struct {
	Msgcode string
	Msg     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateDidAuthRetLog is a free log retrieval operation binding the contract event 0x9af1b91de8cb6c51fb6a8211f6b04fc6c6c3642f77f3e417b9dd5f82f9e860a5.
//
// Solidity: event updateDidAuthRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) FilterUpdateDidAuthRetLog(opts *bind.FilterOpts) (*DidContractUpdateDidAuthRetLogIterator, error) {

	logs, sub, err := _DidContract.contract.FilterLogs(opts, "updateDidAuthRetLog")
	if err != nil {
		return nil, err
	}
	return &DidContractUpdateDidAuthRetLogIterator{contract: _DidContract.contract, event: "updateDidAuthRetLog", logs: logs, sub: sub}, nil
}

// WatchUpdateDidAuthRetLog is a free log subscription operation binding the contract event 0x9af1b91de8cb6c51fb6a8211f6b04fc6c6c3642f77f3e417b9dd5f82f9e860a5.
//
// Solidity: event updateDidAuthRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) WatchUpdateDidAuthRetLog(opts *bind.WatchOpts, sink chan<- *DidContractUpdateDidAuthRetLog) (event.Subscription, error) {

	logs, sub, err := _DidContract.contract.WatchLogs(opts, "updateDidAuthRetLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DidContractUpdateDidAuthRetLog)
				if err := _DidContract.contract.UnpackLog(event, "updateDidAuthRetLog", log); err != nil {
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

// ParseUpdateDidAuthRetLog is a log parse operation binding the contract event 0x9af1b91de8cb6c51fb6a8211f6b04fc6c6c3642f77f3e417b9dd5f82f9e860a5.
//
// Solidity: event updateDidAuthRetLog(string msgcode, string msg)
func (_DidContract *DidContractFilterer) ParseUpdateDidAuthRetLog(log types.Log) (*DidContractUpdateDidAuthRetLog, error) {
	event := new(DidContractUpdateDidAuthRetLog)
	if err := _DidContract.contract.UnpackLog(event, "updateDidAuthRetLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
