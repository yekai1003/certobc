package main

import (
	"certobc/fisco"
	_ "fmt"

	_ "github.com/yekai1003/gobcos/crypto"
)

const adminkey = `{"address":"3f8712acd6ed891ec329fd5ae0a93dd713237e5d","crypto":{"cipher":"aes-128-ctr","ciphertext":"623b85925792e49ac809f474c96a6dc46080d865e5fe1fa89df6c3410fbbfda1","cipherparams":{"iv":"4f0521483a5577b1573f0f63d88b0ede"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":4096,"p":6,"r":8,"salt":"c8ac5e6ee11526b43c2b66a44d0c0bd006fdaff23d22bd64e968406f61e38244"},"mac":"5fd86fc981d37bda5fdab0374db7916244b3dbb3eb71e92b9b6e509e21f9f009"},"id":"2785cb09-649d-4deb-88d2-de152eb78bd5","version":3}`

func main() {
	//1-
	bcos.DeployCert2Fisco("123")
	//0x3380F18596E8824f088177EBd641f50Fb786c1B6

	//2-
	//hash := crypto.Keccak256Hash([]byte("区块链应用工程师"))

	//xh := [32]byte("0x9a2d335fc096b5d9dffbdb0577e00f16d6ccd369a975ee23b7c7f89a9e2cde36")
	//buffer.Write([]byte("0x9a2d335fc096b5d9dffbdb0577e00f16d6ccd369a975ee23b7c7f89a9e2cde36"))
	//fmt.Println(hash)
	// bcos.CertIssue("0x3380F18596E8824f088177EBd641f50Fb786c1B6",
	// 	"123", "123456-67890", hash)
	//3-
	//bcos.Verify("0x3380F18596E8824f088177EBd641f50Fb786c1B6", "123456-67890")
}
