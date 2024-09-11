package bitcoinwallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
)

// P2PKH address
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (string, error) {

	pvBytes, err := PrivateKeyToBytes(privateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	_, pubKey := btcec.PrivKeyFromBytes(pvBytes)

	addr, err := btcutil.NewAddressPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), config.Chaincfg)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return addr.EncodeAddress(), nil
}

// in satoshi
func FetchAddressBalance(address string) (int, error) {

	b, err := config.Explorer.GetAddressBalance(address)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return b, nil
}
