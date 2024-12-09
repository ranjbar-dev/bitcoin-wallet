package models

import "crypto/ecdsa"

// TODO : change types these are just hints
type TransactionInput struct {
	Address    string
	PrivateKey *ecdsa.PrivateKey // owner of utxo privatekey
	Value      int64             // utxo value in satoshi
	Index      uint32            // utxo index
	TxId       string            // utxo txid
}

func NewTransactionInput(privateKey *ecdsa.PrivateKey, value int64, index uint32, txId string) *TransactionInput {

	return &TransactionInput{
		PrivateKey: privateKey,
		Value:      value,
		Index:      index,
		TxId:       txId,
	}
}
