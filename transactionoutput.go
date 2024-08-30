package bitcoinwallet

// TODO : change types these are just hints
type TransactionOutput struct {
	Address any // any btc address
	Value   int // in satoshi
}

func NewTransactionOutput(address string, value int) *TransactionOutput {

	return &TransactionOutput{
		Address: address,
		Value:   value,
	}
}
