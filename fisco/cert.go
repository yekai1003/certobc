// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bcos

import (
	"math/big"
	"strings"

	"github.com/yekai1003/gobcos/accounts/abi"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/common"
	"github.com/yekai1003/gobcos/core/types"
	"github.com/yekai1003/gobcos/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = common.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CertABI is the input ABI used to generate the binding from.
const CertABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"uuid\",\"type\":\"string\"},{\"name\":\"olHash\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"getCourseByHash\",\"outputs\":[{\"components\":[{\"name\":\"uuid\",\"type\":\"string\"},{\"name\":\"ok\",\"type\":\"bool\"},{\"name\":\"olHash\",\"type\":\"bytes32\"}],\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"string\"},{\"name\":\"olHash\",\"type\":\"bytes32\"}],\"name\":\"verifyCourseHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getUserCerts\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"olhash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"certIssue\",\"type\":\"event\"}]"

// CertBin is the compiled bytecode used for deploying new contracts.
var CertBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110df806100606000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680633a5f4e861461007d57806362e4e6ac146100a6578063704b6c02146100e35780637b8677a81461010c578063d849070414610149578063f851a44014610186575b600080fd5b34801561008957600080fd5b506100a4600480360361009f9190810190610b2b565b6101b1565b005b3480156100b257600080fd5b506100cd60048036036100c89190810190610ac1565b6104ad565b6040516100da9190610f3c565b60405180910390f35b3480156100ef57600080fd5b5061010a60048036036101059190810190610a98565b6105b2565b005b34801561011857600080fd5b50610133600480360361012e9190810190610b2b565b6106f8565b6040516101409190610e13565b60405180910390f35b34801561015557600080fd5b50610170600480360361016b9190810190610aea565b610862565b60405161017d9190610df1565b60405180910390f35b34801561019257600080fd5b5061019b610929565b6040516101a89190610dd6565b60405180910390f35b60006101bb61094e565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561024c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024390610f1c565b60405180910390fd5b838360405160200161025f929190610e2e565b6040516020818303038152906040526040518082805190602001908083835b6020831015156102a3578051825260208201915060208101905060208303925061027e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020915060016000836000191660001916815260200190815260200160002060010160009054906101000a900460ff16151515610340576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033790610ebc565b60405180910390fd5b6060604051908101604052808581526020016001151581526020018460001916815250905080600160008460001916600019168152602001908152602001600020600082015181600001908051906020019061039d929190610975565b5060208201518160010160006101000a81548160ff021916908315150217905550604082015181600201906000191690559050506002846040518082805190602001908083835b60208310151561040957805182526020820191506020810190506020830392506103e4565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390208290806001815401808255809150509060018203906000526020600020016000909192909190915090600019169055507f99247928a0ab5edc06080d1fe47eefec81d542f647c177fc5f1a4bee4b4d8d2484848460405161049f93929190610e5e565b60405180910390a150505050565b6104b561094e565b60016000836000191660001916815260200190815260200160002060606040519081016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105755780601f1061054a57610100808354040283529160200191610575565b820191906000526020600020905b81548152906001019060200180831161055857829003601f168201915b505050505081526020016001820160009054906101000a900460ff1615151515815260200160028201546000191660001916815250509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610643576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063a90610f1c565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff16141515156106b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ac90610e9c565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080838360405160200161070e929190610e2e565b6040516020818303038152906040526040518082805190602001908083835b602083101515610752578051825260208201915060208101905060208303925061072d565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020905060016000826000191660001916815260200190815260200160002060010160009054906101000a900460ff1615156107ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107e590610efc565b60405180910390fd5b82600019166001600083600019166000191681526020019081526020016000206002015460001916141515610858576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084f90610edc565b60405180910390fd5b8091505092915050565b60606002826040518082805190602001908083835b60208310151561089c5780518252602082019150602081019050602083039250610877565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902080548060200260200160405190810160405280929190818152602001828054801561091d57602002820191906000526020600020905b81546000191681526020019060010190808311610905575b50505050509050919050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606040519081016040528060608152602001600015158152602001600080191681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106109b657805160ff19168380011785556109e4565b828001600101855582156109e4579182015b828111156109e35782518255916020019190600101906109c8565b5b5090506109f191906109f5565b5090565b610a1791905b80821115610a135760008160009055506001016109fb565b5090565b90565b6000610a268235611028565b905092915050565b6000610a3a8235611048565b905092915050565b600082601f8301121515610a5557600080fd5b8135610a68610a6382610f8b565b610f5e565b91508082526020830160208301858383011115610a8457600080fd5b610a8f838284611052565b50505092915050565b600060208284031215610aaa57600080fd5b6000610ab884828501610a1a565b91505092915050565b600060208284031215610ad357600080fd5b6000610ae184828501610a2e565b91505092915050565b600060208284031215610afc57600080fd5b600082013567ffffffffffffffff811115610b1657600080fd5b610b2284828501610a42565b91505092915050565b60008060408385031215610b3e57600080fd5b600083013567ffffffffffffffff811115610b5857600080fd5b610b6485828601610a42565b9250506020610b7585828601610a2e565b9150509250929050565b610b8881610ff2565b82525050565b6000610b9982610fc4565b808452602084019350610bab83610fb7565b60005b82811015610bdd57610bc1868351610bf8565b610bca82610fe5565b9150602086019550600181019050610bae565b50849250505092915050565b610bf281611012565b82525050565b610c018161101e565b82525050565b6000610c1282610fda565b808452610c26816020860160208601611061565b610c2f81611094565b602085010191505092915050565b6000610c4882610fcf565b808452610c5c816020860160208601611061565b610c6581611094565b602085010191505092915050565b6000601082527f75736572206d75737420657869737473000000000000000000000000000000006020830152604082019050919050565b6000600f82527f6d757374206e6f742065786973747300000000000000000000000000000000006020830152604082019050919050565b6000601082527f68617368206d75737420726967687421000000000000000000000000000000006020830152604082019050919050565b6000600e82527f6d757374206265206578697374730000000000000000000000000000000000006020830152604082019050919050565b6000601182527f6f6e6c792061646d696e2063616e20646f0000000000000000000000000000006020830152604082019050919050565b60006060830160008301518482036000860152610da38282610c3d565b9150506020830151610db86020860182610be9565b506040830151610dcb6040860182610bf8565b508091505092915050565b6000602082019050610deb6000830184610b7f565b92915050565b60006020820190508181036000830152610e0b8184610b8e565b905092915050565b6000602082019050610e286000830184610bf8565b92915050565b60006040820190508181036000830152610e488185610c07565b9050610e576020830184610bf8565b9392505050565b60006060820190508181036000830152610e788186610c07565b9050610e876020830185610bf8565b610e946040830184610bf8565b949350505050565b60006020820190508181036000830152610eb581610c73565b9050919050565b60006020820190508181036000830152610ed581610caa565b9050919050565b60006020820190508181036000830152610ef581610ce1565b9050919050565b60006020820190508181036000830152610f1581610d18565b9050919050565b60006020820190508181036000830152610f3581610d4f565b9050919050565b60006020820190508181036000830152610f568184610d86565b905092915050565b6000604051905081810181811067ffffffffffffffff82111715610f8157600080fd5b8060405250919050565b600067ffffffffffffffff821115610fa257600080fd5b601f19601f8301169050602081019050919050565b6000602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561107f578082015181840152602081019050611064565b8381111561108e576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a723058205beaa7d73f6024cbd302c9128b9130c416c2bfa69bdd74395dfa6172cdd624336c6578706572696d656e74616cf50037"

// DeployCert deploys a new Ethereum contract, binding an instance of Cert to it.
func DeployCert(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.RawTransaction, *Cert, error) {
	parsed, err := abi.JSON(strings.NewReader(CertABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CertBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// Cert is an auto generated Go binding around an Ethereum contract.
type Cert struct {
	CertCaller     // Read-only binding to the contract
	CertTransactor // Write-only binding to the contract
	CertFilterer   // Log filterer for contract events
}

// CertCaller is an auto generated read-only Go binding around an Ethereum contract.
type CertCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CertTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CertFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CertSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CertSession struct {
	Contract     *Cert             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CertCallerSession struct {
	Contract *CertCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CertTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CertTransactorSession struct {
	Contract     *CertTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CertRaw is an auto generated low-level Go binding around an Ethereum contract.
type CertRaw struct {
	Contract *Cert // Generic contract binding to access the raw methods on
}

// CertCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CertCallerRaw struct {
	Contract *CertCaller // Generic read-only contract binding to access the raw methods on
}

// CertTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CertTransactorRaw struct {
	Contract *CertTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCert creates a new instance of Cert, bound to a specific deployed contract.
func NewCert(address common.Address, backend bind.ContractBackend) (*Cert, error) {
	contract, err := bindCert(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cert{CertCaller: CertCaller{contract: contract}, CertTransactor: CertTransactor{contract: contract}, CertFilterer: CertFilterer{contract: contract}}, nil
}

// NewCertCaller creates a new read-only instance of Cert, bound to a specific deployed contract.
func NewCertCaller(address common.Address, caller bind.ContractCaller) (*CertCaller, error) {
	contract, err := bindCert(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CertCaller{contract: contract}, nil
}

// NewCertTransactor creates a new write-only instance of Cert, bound to a specific deployed contract.
func NewCertTransactor(address common.Address, transactor bind.ContractTransactor) (*CertTransactor, error) {
	contract, err := bindCert(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CertTransactor{contract: contract}, nil
}

// NewCertFilterer creates a new log filterer instance of Cert, bound to a specific deployed contract.
func NewCertFilterer(address common.Address, filterer bind.ContractFilterer) (*CertFilterer, error) {
	contract, err := bindCert(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CertFilterer{contract: contract}, nil
}

// bindCert binds a generic wrapper to an already deployed contract.
func bindCert(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CertABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.CertCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cert *CertCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Cert.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.RawTransaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.RawTransaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	Uuid   string
	Ok     bool
	OlHash [32]byte
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Cert *CertCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Cert.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Cert *CertSession) Admin() (common.Address, error) {
	return _Cert.Contract.Admin(&_Cert.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Cert *CertCallerSession) Admin() (common.Address, error) {
	return _Cert.Contract.Admin(&_Cert.CallOpts)
}

// GetCourseByHash is a free data retrieval call binding the contract method 0x62e4e6ac.
//
// Solidity: function getCourseByHash(bytes32 certHash) constant returns(Struct0)
func (_Cert *CertCaller) GetCourseByHash(opts *bind.CallOpts, certHash [32]byte) (Struct0, error) {
	var (
		ret0 = new(Struct0)
	)
	out := ret0
	err := _Cert.contract.Call(opts, out, "getCourseByHash", certHash)
	return *ret0, err
}

// GetCourseByHash is a free data retrieval call binding the contract method 0x62e4e6ac.
//
// Solidity: function getCourseByHash(bytes32 certHash) constant returns(Struct0)
func (_Cert *CertSession) GetCourseByHash(certHash [32]byte) (Struct0, error) {
	return _Cert.Contract.GetCourseByHash(&_Cert.CallOpts, certHash)
}

// GetCourseByHash is a free data retrieval call binding the contract method 0x62e4e6ac.
//
// Solidity: function getCourseByHash(bytes32 certHash) constant returns(Struct0)
func (_Cert *CertCallerSession) GetCourseByHash(certHash [32]byte) (Struct0, error) {
	return _Cert.Contract.GetCourseByHash(&_Cert.CallOpts, certHash)
}

// GetUserCerts is a free data retrieval call binding the contract method 0xd8490704.
//
// Solidity: function getUserCerts(string uuid) constant returns(bytes32[])
func (_Cert *CertCaller) GetUserCerts(opts *bind.CallOpts, uuid string) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Cert.contract.Call(opts, out, "getUserCerts", uuid)
	return *ret0, err
}

// GetUserCerts is a free data retrieval call binding the contract method 0xd8490704.
//
// Solidity: function getUserCerts(string uuid) constant returns(bytes32[])
func (_Cert *CertSession) GetUserCerts(uuid string) ([][32]byte, error) {
	return _Cert.Contract.GetUserCerts(&_Cert.CallOpts, uuid)
}

// GetUserCerts is a free data retrieval call binding the contract method 0xd8490704.
//
// Solidity: function getUserCerts(string uuid) constant returns(bytes32[])
func (_Cert *CertCallerSession) GetUserCerts(uuid string) ([][32]byte, error) {
	return _Cert.Contract.GetUserCerts(&_Cert.CallOpts, uuid)
}

// VerifyCourseHash is a free data retrieval call binding the contract method 0x7b8677a8.
//
// Solidity: function verifyCourseHash(string uuid, bytes32 olHash) constant returns(bytes32)
func (_Cert *CertCaller) VerifyCourseHash(opts *bind.CallOpts, uuid string, olHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Cert.contract.Call(opts, out, "verifyCourseHash", uuid, olHash)
	return *ret0, err
}

// VerifyCourseHash is a free data retrieval call binding the contract method 0x7b8677a8.
//
// Solidity: function verifyCourseHash(string uuid, bytes32 olHash) constant returns(bytes32)
func (_Cert *CertSession) VerifyCourseHash(uuid string, olHash [32]byte) ([32]byte, error) {
	return _Cert.Contract.VerifyCourseHash(&_Cert.CallOpts, uuid, olHash)
}

// VerifyCourseHash is a free data retrieval call binding the contract method 0x7b8677a8.
//
// Solidity: function verifyCourseHash(string uuid, bytes32 olHash) constant returns(bytes32)
func (_Cert *CertCallerSession) VerifyCourseHash(uuid string, olHash [32]byte) ([32]byte, error) {
	return _Cert.Contract.VerifyCourseHash(&_Cert.CallOpts, uuid, olHash)
}

// Issue is a paid mutator transaction binding the contract method 0x3a5f4e86.
//
// Solidity: function issue(string uuid, bytes32 olHash) returns()
func (_Cert *CertTransactor) Issue(opts *bind.TransactOpts, uuid string, olHash [32]byte) (*types.RawTransaction, error) {
	return _Cert.contract.Transact(opts, "issue", uuid, olHash)
}

// Issue is a paid mutator transaction binding the contract method 0x3a5f4e86.
//
// Solidity: function issue(string uuid, bytes32 olHash) returns()
func (_Cert *CertSession) Issue(uuid string, olHash [32]byte) (*types.RawTransaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, uuid, olHash)
}

// Issue is a paid mutator transaction binding the contract method 0x3a5f4e86.
//
// Solidity: function issue(string uuid, bytes32 olHash) returns()
func (_Cert *CertTransactorSession) Issue(uuid string, olHash [32]byte) (*types.RawTransaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, uuid, olHash)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertTransactor) SetAdmin(opts *bind.TransactOpts, user common.Address) (*types.RawTransaction, error) {
	return _Cert.contract.Transact(opts, "setAdmin", user)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertSession) SetAdmin(user common.Address) (*types.RawTransaction, error) {
	return _Cert.Contract.SetAdmin(&_Cert.TransactOpts, user)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertTransactorSession) SetAdmin(user common.Address) (*types.RawTransaction, error) {
	return _Cert.Contract.SetAdmin(&_Cert.TransactOpts, user)
}

// CertCertIssueIterator is returned from FilterCertIssue and is used to iterate over the raw logs and unpacked data for CertIssue events raised by the Cert contract.
type CertCertIssueIterator struct {
	Event *CertCertIssue // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  common.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CertCertIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CertCertIssue)
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
		it.Event = new(CertCertIssue)
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
func (it *CertCertIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CertCertIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CertCertIssue represents a CertIssue event raised by the Cert contract.
type CertCertIssue struct {
	Uuid     string
	Olhash   [32]byte
	CertHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCertIssue is a free log retrieval operation binding the contract event 0x99247928a0ab5edc06080d1fe47eefec81d542f647c177fc5f1a4bee4b4d8d24.
//
// Solidity: event certIssue(string uuid, bytes32 olhash, bytes32 certHash)
func (_Cert *CertFilterer) FilterCertIssue(opts *bind.FilterOpts) (*CertCertIssueIterator, error) {

	logs, sub, err := _Cert.contract.FilterLogs(opts, "certIssue")
	if err != nil {
		return nil, err
	}
	return &CertCertIssueIterator{contract: _Cert.contract, event: "certIssue", logs: logs, sub: sub}, nil
}

// WatchCertIssue is a free log subscription operation binding the contract event 0x99247928a0ab5edc06080d1fe47eefec81d542f647c177fc5f1a4bee4b4d8d24.
//
// Solidity: event certIssue(string uuid, bytes32 olhash, bytes32 certHash)
func (_Cert *CertFilterer) WatchCertIssue(opts *bind.WatchOpts, sink chan<- *CertCertIssue) (event.Subscription, error) {

	logs, sub, err := _Cert.contract.WatchLogs(opts, "certIssue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CertCertIssue)
				if err := _Cert.contract.UnpackLog(event, "certIssue", log); err != nil {
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

// ParseCertIssue is a log parse operation binding the contract event 0x99247928a0ab5edc06080d1fe47eefec81d542f647c177fc5f1a4bee4b4d8d24.
//
// Solidity: event certIssue(string uuid, bytes32 olhash, bytes32 certHash)
func (_Cert *CertFilterer) ParseCertIssue(log types.Log) (*CertCertIssue, error) {
	event := new(CertCertIssue)
	if err := _Cert.contract.UnpackLog(event, "certIssue", log); err != nil {
		return nil, err
	}
	return event, nil
}
