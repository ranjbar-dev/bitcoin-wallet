package bitcoinwallet

import "crypto/ecdsa"

// TODO : change types these are just hints
type TransactionInput struct {
	PrivateKey any // owner of utxo privatekey
	Value      any // utxo value in satoshi
	Index      any // utxo index
	TxId       any // utxo txid
}

func NewTransactionInput(privateKey *ecdsa.PrivateKey, value int, index int, txId string) *TransactionInput {

	return &TransactionInput{
		PrivateKey: privateKey,
		Value:      value,
		Index:      index,
		TxId:       txId,
	}
}
