package models

import "crypto/ecdsa"

// TODO : change types these are just hints
type TransactionInput struct {
	Address    string
	PrivateKey *ecdsa.PrivateKey // owner of utxo privatekey
	Value      int               // utxo value in satoshi
	Index      int               // utxo index
	TxId       string            // utxo txid
}

func NewTransactionInput(privateKey *ecdsa.PrivateKey, value int, index int, txId string) *TransactionInput {

	return &TransactionInput{
		PrivateKey: privateKey,
		Value:      value,
		Index:      index,
		TxId:       txId,
	}
}
