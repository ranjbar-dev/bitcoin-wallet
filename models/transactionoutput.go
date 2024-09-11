package models

// TODO : change types these are just hints
type TransactionOutput struct {
	Address []byte // any btc address
	Value   int64  // in satoshi
}

func NewTransactionOutput(address []byte, value int64) *TransactionOutput {

	return &TransactionOutput{
		Address: address,
		Value:   value,
	}
}
