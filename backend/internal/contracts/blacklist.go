// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// BlacklistUserOperation is an auto generated low-level Go binding around an user-defined struct.
type BlacklistUserOperation struct {
	Sender          common.Address
	ReportedAddress common.Address
	Category        *big.Int
	Comments        string
	Source          string
	Date            *big.Int
	Nonce           *big.Int
	Signature       []byte
}

// BlacklistMetaData contains all meta data concerning the Blacklist contract.
var BlacklistMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"MAX_BATCH_SIZE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"executeBatchOperations\",\"inputs\":[{\"name\":\"userOps\",\"type\":\"tuple[]\",\"internalType\":\"structBlacklist.UserOperation[]\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reportedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"comments\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"date\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOperationHash\",\"inputs\":[{\"name\":\"userOp\",\"type\":\"tuple\",\"internalType\":\"structBlacklist.UserOperation\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reportedAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"category\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"comments\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"source\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"date\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getReportByAddress\",\"inputs\":[{\"name\":\"reportedAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Blacklisted\",\"inputs\":[{\"name\":\"reportedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"category\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"date\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"comments\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"source\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidNonce\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoReportFound\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610bd88061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c806322aa20b4146100595780635fe7f452146100865780637ecebe001461009b5780639299a502146100c8578063cfdbf254146100db575b5f5ffd5b61006c6100673660046107ea565b6100e3565b604080519283526020830191909152015b60405180910390f35b610099610094366004610817565b610134565b005b6100ba6100a93660046107ea565b60016020525f908152604090205481565b60405190815260200161007d565b6100ba6100d6366004610888565b61043c565b6100ba606481565b6001600160a01b0381165f908152602081905260408120600581015482919060ff1661012257604051630142c73560e01b815260040160405180910390fd5b80546001909101549094909350915050565b606481111561017c5760405162461bcd60e51b815260206004820152600f60248201526e426174636820746f6f206c6172676560881b60448201526064015b60405180910390fd5b5f82825f81811061018f5761018f6108c0565b90506020028101906101a191906108d4565b6101af9060208101906107ea565b6001600160a01b0381165f908152600160205260408120549192505b838110156103ca57826001600160a01b03168585838181106101ef576101ef6108c0565b905060200281019061020191906108d4565b61020f9060208101906107ea565b6001600160a01b0316146102555760405162461bcd60e51b815260206004820152600d60248201526c4d697865642073656e6465727360981b6044820152606401610173565b61025f81836108f2565b858583818110610271576102716108c0565b905060200281019061028391906108d4565b60c00135146102cd5760405162461bcd60e51b8152602060048201526016602482015275496e76616c6964206e6f6e63652073657175656e636560501b6044820152606401610173565b5f6102f58686848181106102e3576102e36108c0565b90506020028101906100d691906108d4565b90505f61036987878581811061030d5761030d6108c0565b905060200281019061031f91906108d4565b61032d9060e0810190610911565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525086939250506104ba9050565b9050846001600160a01b0316816001600160a01b0316146103c05760405162461bcd60e51b8152602060048201526011602482015270496e76616c6964207369676e617475726560781b6044820152606401610173565b50506001016101cb565b505f5b83811015610409576104018585838181106103ea576103ea6108c0565b90506020028101906103fc91906108d4565b6104e4565b6001016103cd565b506001600160a01b0382165f90815260016020526040812080548592906104319084906108f2565b909155505050505050565b5f3061044b60208401846107ea565b61045b60408501602086016107ea565b604085013561046d6060870187610911565b61047a6080890189610911565b8960a001358a60c0013560405160200161049d9a9998979695949392919061095b565b604051602081830303815290604052805190602001209050919050565b5f5f5f5f6104c8868661061d565b9250925092506104d88282610666565b50909150505b92915050565b5f80806104f760408501602086016107ea565b6001600160a01b0316815260208101919091526040015f20600581015490915060ff166105865760018082556040830135908201556105396060830183610911565b6003830191610549919083610a66565b506105576080830183610911565b6004830191610567919083610a66565b5060a0820135600282015560058101805460ff19166001179055610599565b8054600190810182556040830135908201555b6105a960408301602084016107ea565b6001600160a01b03167f2dd6af485f4185bb41c0d8d27334ab623a3093debe674d29ee6a4ba7c87c1eb7825f015484604001358560a001358680606001906105f19190610911565b6105fe60808a018a610911565b6040516106119796959493929190610b48565b60405180910390a25050565b5f5f5f8351604103610654576020840151604085015160608601515f1a61064688828585610722565b95509550955050505061065f565b505081515f91506002905b9250925092565b5f82600381111561067957610679610b8e565b03610682575050565b600182600381111561069657610696610b8e565b036106b45760405163f645eedf60e01b815260040160405180910390fd5b60028260038111156106c8576106c8610b8e565b036106e95760405163fce698f760e01b815260048101829052602401610173565b60038260038111156106fd576106fd610b8e565b0361071e576040516335e2f38360e21b815260048101829052602401610173565b5050565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a084111561075b57505f915060039050826107e0565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa1580156107ac573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b0381166107d757505f9250600191508290506107e0565b92505f91508190505b9450945094915050565b5f602082840312156107fa575f5ffd5b81356001600160a01b0381168114610810575f5ffd5b9392505050565b5f5f60208385031215610828575f5ffd5b823567ffffffffffffffff81111561083e575f5ffd5b8301601f8101851361084e575f5ffd5b803567ffffffffffffffff811115610864575f5ffd5b8560208260051b8401011115610878575f5ffd5b6020919091019590945092505050565b5f60208284031215610898575f5ffd5b813567ffffffffffffffff8111156108ae575f5ffd5b82016101008185031215610810575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f823560fe198336030181126108e8575f5ffd5b9190910192915050565b808201808211156104de57634e487b7160e01b5f52601160045260245ffd5b5f5f8335601e19843603018112610926575f5ffd5b83018035915067ffffffffffffffff821115610940575f5ffd5b602001915036819003821315610954575f5ffd5b9250929050565b6001600160601b03198b60601b1681526001600160601b03198a60601b1660148201526001600160601b03198960601b16602882015287603c8201528587605c8301375f868201605c81015f81528688823750909401605c81019390935250607c820152609c0198975050505050505050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806109f657607f821691505b602082108103610a1457634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610a6157805f5260205f20601f840160051c81016020851015610a3f5750805b601f840160051c820191505b81811015610a5e575f8155600101610a4b565b50505b505050565b67ffffffffffffffff831115610a7e57610a7e6109ce565b610a9283610a8c83546109e2565b83610a1a565b5f601f841160018114610ac3575f8515610aac5750838201355b5f19600387901b1c1916600186901b178355610a5e565b5f83815260208120601f198716915b82811015610af25786850135825560209485019460019092019101610ad2565b5086821015610b0e575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b87815286602082015285604082015260a060608201525f610b6d60a083018688610b20565b8281036080840152610b80818587610b20565b9a9950505050505050505050565b634e487b7160e01b5f52602160045260245ffdfea2646970667358221220767d17010cbe576d05e6f8eb8193d7a10d4daa9892e3da98ae56e9f8d6bed30b64736f6c634300081c0033",
}

// BlacklistABI is the input ABI used to generate the binding from.
// Deprecated: Use BlacklistMetaData.ABI instead.
var BlacklistABI = BlacklistMetaData.ABI

// BlacklistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlacklistMetaData.Bin instead.
var BlacklistBin = BlacklistMetaData.Bin

// DeployBlacklist deploys a new Ethereum contract, binding an instance of Blacklist to it.
func DeployBlacklist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blacklist, error) {
	parsed, err := BlacklistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlacklistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blacklist{BlacklistCaller: BlacklistCaller{contract: contract}, BlacklistTransactor: BlacklistTransactor{contract: contract}, BlacklistFilterer: BlacklistFilterer{contract: contract}}, nil
}

// Blacklist is an auto generated Go binding around an Ethereum contract.
type Blacklist struct {
	BlacklistCaller     // Read-only binding to the contract
	BlacklistTransactor // Write-only binding to the contract
	BlacklistFilterer   // Log filterer for contract events
}

// BlacklistCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlacklistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlacklistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlacklistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlacklistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlacklistSession struct {
	Contract     *Blacklist        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlacklistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlacklistCallerSession struct {
	Contract *BlacklistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlacklistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlacklistTransactorSession struct {
	Contract     *BlacklistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlacklistRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlacklistRaw struct {
	Contract *Blacklist // Generic contract binding to access the raw methods on
}

// BlacklistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlacklistCallerRaw struct {
	Contract *BlacklistCaller // Generic read-only contract binding to access the raw methods on
}

// BlacklistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlacklistTransactorRaw struct {
	Contract *BlacklistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlacklist creates a new instance of Blacklist, bound to a specific deployed contract.
func NewBlacklist(address common.Address, backend bind.ContractBackend) (*Blacklist, error) {
	contract, err := bindBlacklist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blacklist{BlacklistCaller: BlacklistCaller{contract: contract}, BlacklistTransactor: BlacklistTransactor{contract: contract}, BlacklistFilterer: BlacklistFilterer{contract: contract}}, nil
}

// NewBlacklistCaller creates a new read-only instance of Blacklist, bound to a specific deployed contract.
func NewBlacklistCaller(address common.Address, caller bind.ContractCaller) (*BlacklistCaller, error) {
	contract, err := bindBlacklist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlacklistCaller{contract: contract}, nil
}

// NewBlacklistTransactor creates a new write-only instance of Blacklist, bound to a specific deployed contract.
func NewBlacklistTransactor(address common.Address, transactor bind.ContractTransactor) (*BlacklistTransactor, error) {
	contract, err := bindBlacklist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlacklistTransactor{contract: contract}, nil
}

// NewBlacklistFilterer creates a new log filterer instance of Blacklist, bound to a specific deployed contract.
func NewBlacklistFilterer(address common.Address, filterer bind.ContractFilterer) (*BlacklistFilterer, error) {
	contract, err := bindBlacklist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlacklistFilterer{contract: contract}, nil
}

// bindBlacklist binds a generic wrapper to an already deployed contract.
func bindBlacklist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlacklistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blacklist *BlacklistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blacklist.Contract.BlacklistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blacklist *BlacklistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blacklist.Contract.BlacklistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blacklist *BlacklistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blacklist.Contract.BlacklistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blacklist *BlacklistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blacklist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blacklist *BlacklistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blacklist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blacklist *BlacklistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blacklist.Contract.contract.Transact(opts, method, params...)
}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_Blacklist *BlacklistCaller) MAXBATCHSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blacklist.contract.Call(opts, &out, "MAX_BATCH_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_Blacklist *BlacklistSession) MAXBATCHSIZE() (*big.Int, error) {
	return _Blacklist.Contract.MAXBATCHSIZE(&_Blacklist.CallOpts)
}

// MAXBATCHSIZE is a free data retrieval call binding the contract method 0xcfdbf254.
//
// Solidity: function MAX_BATCH_SIZE() view returns(uint256)
func (_Blacklist *BlacklistCallerSession) MAXBATCHSIZE() (*big.Int, error) {
	return _Blacklist.Contract.MAXBATCHSIZE(&_Blacklist.CallOpts)
}

// GetOperationHash is a free data retrieval call binding the contract method 0x9299a502.
//
// Solidity: function getOperationHash((address,address,uint256,string,string,uint256,uint256,bytes) userOp) view returns(bytes32)
func (_Blacklist *BlacklistCaller) GetOperationHash(opts *bind.CallOpts, userOp BlacklistUserOperation) ([32]byte, error) {
	var out []interface{}
	err := _Blacklist.contract.Call(opts, &out, "getOperationHash", userOp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOperationHash is a free data retrieval call binding the contract method 0x9299a502.
//
// Solidity: function getOperationHash((address,address,uint256,string,string,uint256,uint256,bytes) userOp) view returns(bytes32)
func (_Blacklist *BlacklistSession) GetOperationHash(userOp BlacklistUserOperation) ([32]byte, error) {
	return _Blacklist.Contract.GetOperationHash(&_Blacklist.CallOpts, userOp)
}

// GetOperationHash is a free data retrieval call binding the contract method 0x9299a502.
//
// Solidity: function getOperationHash((address,address,uint256,string,string,uint256,uint256,bytes) userOp) view returns(bytes32)
func (_Blacklist *BlacklistCallerSession) GetOperationHash(userOp BlacklistUserOperation) ([32]byte, error) {
	return _Blacklist.Contract.GetOperationHash(&_Blacklist.CallOpts, userOp)
}

// GetReportByAddress is a free data retrieval call binding the contract method 0x22aa20b4.
//
// Solidity: function getReportByAddress(address reportedAddress) view returns(uint256 count, uint256 category)
func (_Blacklist *BlacklistCaller) GetReportByAddress(opts *bind.CallOpts, reportedAddress common.Address) (struct {
	Count    *big.Int
	Category *big.Int
}, error) {
	var out []interface{}
	err := _Blacklist.contract.Call(opts, &out, "getReportByAddress", reportedAddress)

	outstruct := new(struct {
		Count    *big.Int
		Category *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Category = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReportByAddress is a free data retrieval call binding the contract method 0x22aa20b4.
//
// Solidity: function getReportByAddress(address reportedAddress) view returns(uint256 count, uint256 category)
func (_Blacklist *BlacklistSession) GetReportByAddress(reportedAddress common.Address) (struct {
	Count    *big.Int
	Category *big.Int
}, error) {
	return _Blacklist.Contract.GetReportByAddress(&_Blacklist.CallOpts, reportedAddress)
}

// GetReportByAddress is a free data retrieval call binding the contract method 0x22aa20b4.
//
// Solidity: function getReportByAddress(address reportedAddress) view returns(uint256 count, uint256 category)
func (_Blacklist *BlacklistCallerSession) GetReportByAddress(reportedAddress common.Address) (struct {
	Count    *big.Int
	Category *big.Int
}, error) {
	return _Blacklist.Contract.GetReportByAddress(&_Blacklist.CallOpts, reportedAddress)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blacklist *BlacklistCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blacklist.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blacklist *BlacklistSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blacklist.Contract.Nonces(&_Blacklist.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blacklist *BlacklistCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blacklist.Contract.Nonces(&_Blacklist.CallOpts, arg0)
}

// ExecuteBatchOperations is a paid mutator transaction binding the contract method 0x5fe7f452.
//
// Solidity: function executeBatchOperations((address,address,uint256,string,string,uint256,uint256,bytes)[] userOps) returns()
func (_Blacklist *BlacklistTransactor) ExecuteBatchOperations(opts *bind.TransactOpts, userOps []BlacklistUserOperation) (*types.Transaction, error) {
	return _Blacklist.contract.Transact(opts, "executeBatchOperations", userOps)
}

// ExecuteBatchOperations is a paid mutator transaction binding the contract method 0x5fe7f452.
//
// Solidity: function executeBatchOperations((address,address,uint256,string,string,uint256,uint256,bytes)[] userOps) returns()
func (_Blacklist *BlacklistSession) ExecuteBatchOperations(userOps []BlacklistUserOperation) (*types.Transaction, error) {
	return _Blacklist.Contract.ExecuteBatchOperations(&_Blacklist.TransactOpts, userOps)
}

// ExecuteBatchOperations is a paid mutator transaction binding the contract method 0x5fe7f452.
//
// Solidity: function executeBatchOperations((address,address,uint256,string,string,uint256,uint256,bytes)[] userOps) returns()
func (_Blacklist *BlacklistTransactorSession) ExecuteBatchOperations(userOps []BlacklistUserOperation) (*types.Transaction, error) {
	return _Blacklist.Contract.ExecuteBatchOperations(&_Blacklist.TransactOpts, userOps)
}

// BlacklistBlacklistedIterator is returned from FilterBlacklisted and is used to iterate over the raw logs and unpacked data for Blacklisted events raised by the Blacklist contract.
type BlacklistBlacklistedIterator struct {
	Event *BlacklistBlacklisted // Event containing the contract specifics and raw log

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
func (it *BlacklistBlacklistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlacklistBlacklisted)
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
		it.Event = new(BlacklistBlacklisted)
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
func (it *BlacklistBlacklistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlacklistBlacklistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlacklistBlacklisted represents a Blacklisted event raised by the Blacklist contract.
type BlacklistBlacklisted struct {
	ReportedAddress common.Address
	Count           *big.Int
	Category        *big.Int
	Date            *big.Int
	Comments        string
	Source          string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlacklisted is a free log retrieval operation binding the contract event 0x2dd6af485f4185bb41c0d8d27334ab623a3093debe674d29ee6a4ba7c87c1eb7.
//
// Solidity: event Blacklisted(address indexed reportedAddress, uint256 count, uint256 category, uint256 date, string comments, string source)
func (_Blacklist *BlacklistFilterer) FilterBlacklisted(opts *bind.FilterOpts, reportedAddress []common.Address) (*BlacklistBlacklistedIterator, error) {

	var reportedAddressRule []interface{}
	for _, reportedAddressItem := range reportedAddress {
		reportedAddressRule = append(reportedAddressRule, reportedAddressItem)
	}

	logs, sub, err := _Blacklist.contract.FilterLogs(opts, "Blacklisted", reportedAddressRule)
	if err != nil {
		return nil, err
	}
	return &BlacklistBlacklistedIterator{contract: _Blacklist.contract, event: "Blacklisted", logs: logs, sub: sub}, nil
}

// WatchBlacklisted is a free log subscription operation binding the contract event 0x2dd6af485f4185bb41c0d8d27334ab623a3093debe674d29ee6a4ba7c87c1eb7.
//
// Solidity: event Blacklisted(address indexed reportedAddress, uint256 count, uint256 category, uint256 date, string comments, string source)
func (_Blacklist *BlacklistFilterer) WatchBlacklisted(opts *bind.WatchOpts, sink chan<- *BlacklistBlacklisted, reportedAddress []common.Address) (event.Subscription, error) {

	var reportedAddressRule []interface{}
	for _, reportedAddressItem := range reportedAddress {
		reportedAddressRule = append(reportedAddressRule, reportedAddressItem)
	}

	logs, sub, err := _Blacklist.contract.WatchLogs(opts, "Blacklisted", reportedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlacklistBlacklisted)
				if err := _Blacklist.contract.UnpackLog(event, "Blacklisted", log); err != nil {
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

// ParseBlacklisted is a log parse operation binding the contract event 0x2dd6af485f4185bb41c0d8d27334ab623a3093debe674d29ee6a4ba7c87c1eb7.
//
// Solidity: event Blacklisted(address indexed reportedAddress, uint256 count, uint256 category, uint256 date, string comments, string source)
func (_Blacklist *BlacklistFilterer) ParseBlacklisted(log types.Log) (*BlacklistBlacklisted, error) {
	event := new(BlacklistBlacklisted)
	if err := _Blacklist.contract.UnpackLog(event, "Blacklisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
