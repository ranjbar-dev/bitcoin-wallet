package bitcoinwallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// GeneratePrivateKey generates a new ECDSA private key using the P-256 curve.
// It returns a pointer to the generated private key and an error if the key generation fails.
//
// Returns:
//   - *ecdsa.PrivateKey: A pointer to the generated ECDSA private key.
//   - error: An error if the key generation fails, otherwise nil.
func GeneratePrivateKey() (*ecdsa.PrivateKey, error) {

	pk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	return pk, err
}

// HexToPrivateKey converts a hexadecimal string to an ECDSA private key.
//
// Parameters:
// - hex: A string representing the private key in hexadecimal format.
//
// Returns:
// - *ecdsa.PrivateKey: A pointer to the ECDSA private key.
// - error: An error if the conversion fails.
func HexToPrivateKey(hex string) (*ecdsa.PrivateKey, error) {

	pk, err := crypto.HexToECDSA(hex)

	return pk, err
}

// BytesToPrivateKey converts a byte slice to an ECDSA private key.
// It takes a byte slice as input and returns a pointer to an ECDSA private key and an error.
// If the conversion is successful, the error will be nil. Otherwise, the error will contain
// information about what went wrong.
//
// Parameters:
// - bytes: A byte slice representing the private key.
//
// Returns:
// - *ecdsa.PrivateKey: A pointer to the ECDSA private key.
// - error: An error if the conversion fails, otherwise nil.
func BytesToPrivateKey(bytes []byte) (*ecdsa.PrivateKey, error) {

	pk, err := crypto.ToECDSA(bytes)

	return pk, err
}

// PrivateKeyToHex converts an ECDSA private key to its hexadecimal string representation.
//
// Parameters:
// - privateKey: A pointer to an ECDSA private key.
//
// Returns:
// - A string containing the hexadecimal representation of the private key.
// - An error, if any occurs during the conversion process.
func PrivateKeyToHex(privateKey *ecdsa.PrivateKey) (string, error) {

	pkBytes := crypto.FromECDSA(privateKey)

	pk := hexutil.Encode(pkBytes)[2:]

	return pk, nil
}

// PrivateKeyToBytes converts an ECDSA private key to a byte slice.
//
// Parameters:
// - privateKey: A pointer to an ECDSA private key.
//
// Returns:
// - A byte slice representing the private key.
// - An error, which is always nil in this implementation.
func PrivateKeyToBytes(privateKey *ecdsa.PrivateKey) ([]byte, error) {

	bytes := crypto.FromECDSA(privateKey)

	return bytes, nil
}
