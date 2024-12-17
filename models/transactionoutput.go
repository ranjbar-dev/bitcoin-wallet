package models

// TODO : change types these are just hints
type TransactionOutput struct {
	Address string // any btc address
	Value   int64  // in satoshi
}

func NewTransactionOutput(address string, value int64) *TransactionOutput {

	return &TransactionOutput{
		Address: address,
		Value:   value,
	}
}
