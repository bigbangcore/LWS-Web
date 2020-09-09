package libs

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/jamesruan/sodium"
)

//KeyPair ,keypair struct
type KeyPair struct {
	PublicKey  string
	PrivateKey string
}

//MakeKeyPairJSON is used to make key pair,type is json
func MakeKeyPairJSON() string {
	keyPair := MakeKeyPair()
	js, err := json.Marshal(keyPair)
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	return string(js)
}

//MakeKeyPair is used to make key pair,type is KeyPair
func MakeKeyPair() (keypair KeyPair) {
	skp := new(sodium.SignKP)
	MakeKeyPairBytes(skp)

	keypair.PublicKey = hex.EncodeToString(skp.PublicKey.Bytes)
	keypair.PrivateKey = hex.EncodeToString(skp.SecretKey.Bytes)[0:64]
	return keypair
}

//MakeKeyPairBytes is used to make key pair,type is Bytes
func MakeKeyPairBytes(skp *sodium.SignKP) (pk sodium.Bytes) {
	for {
		tmp := sodium.MakeSignKP()
		//skp = &tmp
		*skp = tmp
		pk = tmp.PublicKey.Bytes
		count := 0
		p := skp.SecretKey.Bytes

		for i := 0; i < 31; i++ {
			if p[i] != 0x00 && p[i] != 0xFF {
				count++
				if count >= 4 {
					return pk
				}
			}
		}
	}
}
