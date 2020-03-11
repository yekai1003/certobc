package bcos

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/satori/go.uuid"
	"github.com/yekai1003/gobcos/accounts/abi/bind"
	"github.com/yekai1003/gobcos/client"
	"github.com/yekai1003/gobcos/common"
	"github.com/yekai1003/gobcos/crypto"
)

const FISCO_GROUP = 1

var FISCO_NETWORK = "http://localhost:8545"
var ContractAddr = "0x016431d5333CFcad5A2acc2274EE53F3B5343334"

const adminkey = `{"address":"3f8712acd6ed891ec329fd5ae0a93dd713237e5d","crypto":{"cipher":"aes-128-ctr","ciphertext":"623b85925792e49ac809f474c96a6dc46080d865e5fe1fa89df6c3410fbbfda1","cipherparams":{"iv":"4f0521483a5577b1573f0f63d88b0ede"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"c8ac5e6ee11526b43c2b66a44d0c0bd006fdaff23d22bd64e968406f61e38244"},"mac":"5fd86fc981d37bda5fdab0374db7916244b3dbb3eb71e92b9b6e509e21f9f009"},"id":"2785cb09-649d-4deb-88d2-de152eb78bd5","version":3}`

func init() {
	url := os.Getenv("FISCO_NETWORK")
	if url != "" {
		FISCO_NETWORK = url
	}
	crt := os.Getenv("CERT_CONTRACT_ADDR")
	if crt != "" {
		ContractAddr = crt
	} else {
		log.Panic("CERT_CONTRACT_ADDR must be set")
	}
}

func DeployCert2Fisco(pass string) {
	//1. 连接到fisco节点
	cli, err := client.Dial(FISCO_NETWORK, FISCO_GROUP)
	if err != nil {
		log.Panic("Failed to Dial", err)
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 合约部署
	//common.Address, *types.RawTransaction, *Cert, error
	address, tx, _, err := DeployCert(auth, cli)
	if err != nil {
		log.Panic("Failed to DeployCert", err)
	}
	fmt.Println(tx.Hash().Hex(), address.Hex())
}

func CertIssue(pass, uuid string, olHash [32]byte) error {
	//1. 连接到fisco节点
	cli, err := client.Dial(FISCO_NETWORK, FISCO_GROUP)
	if err != nil {
		fmt.Println("Failed to Dial", err)
		return err
	}
	//延迟关闭连接
	defer cli.Close()

	//2.签名
	keyin := strings.NewReader(adminkey)
	auth, err := bind.NewTransactor(keyin, pass)
	//3. 调用函数
	cert, err := NewCert(common.HexToAddress(ContractAddr), cli)
	if err != nil {
		fmt.Println("Failed to NewCert", err)
		return err
	}
	fmt.Printf("CertIssue:uuid = [%s], olhash=[%x]\n", uuid, olHash)
	tx, err := cert.Issue(auth, uuid, olHash)
	if err != nil {
		fmt.Println("Failed to Issue", err)
		return err
	}
	fmt.Println(tx.Hash().Hex())
	return nil
}

func Verify(uuid string, olHash [32]byte) ([32]byte, error) {
	//1. 连接到fisco节点
	cli, err := client.Dial(FISCO_NETWORK, FISCO_GROUP)
	if err != nil {
		fmt.Println("Failed to Dial fisco ", err)
		return [32]byte{}, err
	}
	//延迟关闭连接
	defer cli.Close()

	//2. 构造合约实例
	cert, err := NewCert(common.HexToAddress(ContractAddr), cli)
	if err != nil {
		fmt.Println("Failed to NewCert", err)
		return [32]byte{}, err
	}
	//3. 验证用户证书
	opts := bind.CallOpts{
		From: common.HexToAddress("0x3f8712aCD6Ed891ec329fd5AE0a93dd713237E5d"),
	}
	fmt.Printf("Verify:uuid = [%s], olhash=[%x]\n", uuid, olHash)
	hash, err := cert.VerifyCourseHash(&opts, uuid, olHash)
	if err != nil {
		fmt.Println("Failed to VerifyCourseHash", err)
		return [32]byte{}, err
	}
	fmt.Printf("%s\n", hash)
	return hash, nil
}

func GetUUID() string {
	return fmt.Sprint("%s", uuid.NewV4())
}

func GetHash(name string) [32]byte {
	return crypto.Keccak256Hash([]byte(name))
}
