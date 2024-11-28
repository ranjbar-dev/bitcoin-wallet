package bitcoinwallet

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

// PrivateKeyToAddress generates a P2WPKH address from a given ECDSA private key.
// It returns the generated address as a string and an error if any occurs during the process.
// The function first converts the private key to bytes, then derives the public key from it,
// and finally generates the address using the public key hash.
//
// Parameters:
// - privateKey: A pointer to an ECDSA private key.
//
// Returns:
// - A string representing the generated P2WPKH address.
// - An error if any occurs during the address generation process.
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (string, error) {

	pvBytes, err := PrivateKeyToBytes(privateKey)
	if err != nil {

		return "", err
	}

	_, pubKey := btcec.PrivKeyFromBytes(pvBytes)

	addr, err := btcutil.NewAddressWitnessPubKeyHash(btcutil.Hash160(pubKey.SerializeCompressed()), config.Chaincfg)
	if err != nil {

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

		return -1, err
	}

	return b, nil
}

// FetchAddressUTXOs retrieves the unspent transaction outputs (UTXOs) for a given Bitcoin address.
// It uses the configured blockchain explorer to fetch the UTXOs within a specified timeout period.
//
// Parameters:
//   - address: A string representing the Bitcoin address for which to fetch UTXOs.
//
// Returns:
//   - []models.UTXO: A slice of UTXO objects associated with the given address.
//   - error: An error object if there was an issue fetching the UTXOs, otherwise nil.
func FetchAddressUTXOs(address string) ([]models.UTXO, error) {

	v, err := config.Explorer.GetAddressUTXOs(address, config.Timeout)

	if err != nil {

		return nil, err
	}

	return v, nil
}
