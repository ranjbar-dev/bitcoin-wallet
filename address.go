package bitcoinwallet

import "crypto/ecdsa"

// P2PKH address
func PrivateKeyToAddress(privateKey *ecdsa.PrivateKey) (string, error) {

	// TODO : implement

	return "", nil
}

// in satoshi
func FetchAddressBalance() (int, error) {

	// use config to get explorer type and apikey

	// TODO : implement

	return 0, nil
}
