package bitcoinwallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {

	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return pk, err
}

func HexToPrivateKey(hex string) (*ecdsa.PrivateKey, error) {

	pk, err := crypto.HexToECDSA(hex)

	return pk, err
}

func BytesToPrivateKey(bytes []byte) (*ecdsa.PrivateKey, error) {

	pk, err := crypto.ToECDSA(bytes)

	return pk, err
}

func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) (string, error) {

	pkBytes := crypto.FromECDSA(privateKey)

	pk := hexutil.Encode(pkBytes)[2:]

	return pk, nil
}

func PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) ([]byte, error) {

	bytes := crypto.FromECDSA(privateKey)

	return bytes, nil
}
