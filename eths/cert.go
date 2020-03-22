// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eths

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CertABI is the input ABI used to generate the binding from.
const CertABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"olhash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"certHash\",\"type\":\"bytes32\"}],\"name\":\"certIssue\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"olHash\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"setAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"}],\"name\":\"getUserCerts\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"uuid\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"olHash\",\"type\":\"bytes32\"}],\"name\":\"verifyCourseHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
func (_Cert *CertRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.CertTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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
func (_Cert *CertTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cert *CertTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cert.Contract.contract.Transact(opts, method, params...)
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

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(bytes32)
func (_Cert *CertCaller) GetHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Cert.contract.Call(opts, out, "getHash")
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(bytes32)
func (_Cert *CertSession) GetHash() ([32]byte, error) {
	return _Cert.Contract.GetHash(&_Cert.CallOpts)
}

// GetHash is a free data retrieval call binding the contract method 0xd13319c4.
//
// Solidity: function getHash() constant returns(bytes32)
func (_Cert *CertCallerSession) GetHash() ([32]byte, error) {
	return _Cert.Contract.GetHash(&_Cert.CallOpts)
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
func (_Cert *CertTransactor) Issue(opts *bind.TransactOpts, uuid string, olHash [32]byte) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "issue", uuid, olHash)
}

// Issue is a paid mutator transaction binding the contract method 0x3a5f4e86.
//
// Solidity: function issue(string uuid, bytes32 olHash) returns()
func (_Cert *CertSession) Issue(uuid string, olHash [32]byte) (*types.Transaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, uuid, olHash)
}

// Issue is a paid mutator transaction binding the contract method 0x3a5f4e86.
//
// Solidity: function issue(string uuid, bytes32 olHash) returns()
func (_Cert *CertTransactorSession) Issue(uuid string, olHash [32]byte) (*types.Transaction, error) {
	return _Cert.Contract.Issue(&_Cert.TransactOpts, uuid, olHash)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertTransactor) SetAdmin(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Cert.contract.Transact(opts, "setAdmin", user)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertSession) SetAdmin(user common.Address) (*types.Transaction, error) {
	return _Cert.Contract.SetAdmin(&_Cert.TransactOpts, user)
}

// SetAdmin is a paid mutator transaction binding the contract method 0x704b6c02.
//
// Solidity: function setAdmin(address user) returns()
func (_Cert *CertTransactorSession) SetAdmin(user common.Address) (*types.Transaction, error) {
	return _Cert.Contract.SetAdmin(&_Cert.TransactOpts, user)
}

// CertCertIssueIterator is returned from FilterCertIssue and is used to iterate over the raw logs and unpacked data for CertIssue events raised by the Cert contract.
type CertCertIssueIterator struct {
	Event *CertCertIssue // Event containing the contract specifics and raw log

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
