package bitcoinwallet

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
)

// PrivateKeyToAddress generates a P2PKH address from a given ECDSA private key.
// It returns the generated address as a string and an error if any occurs during the process.
// The function first converts the private key to bytes, then derives the public key from it,
// and finally generates the address using the public key hash.
//
// Parameters:
// - privateKey: A pointer to an ECDSA private key.
//
// Returns:
// - A string representing the generated P2PKH address.
// - An error if any occurs during the address generation process.
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

// FetchAddressBalance retrieves the balance of a given Bitcoin address in satoshis
// from the configured blockchain explorer. It returns the balance as an integer
// and an error if any occurs during the process.
//
// Parameters:
// - address: A string representing the Bitcoin address whose balance is to be fetched.
//
// Returns:
// - An integer representing the balance of the address in satoshis.
// - An error if any occurs during the balance retrieval process.
func FetchAddressBalance(address string) (int, error) {

	b, err := config.Explorer.GetAddressBalance(address)

	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return b, nil
}
