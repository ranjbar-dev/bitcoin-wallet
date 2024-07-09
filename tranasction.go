package bitcoinwallet

type BlockchainTransaction struct {
	// TODO : implment
}

type Transaction struct {
	inputs  []TransactionInput
	outputs []TransactionOutput
}

// returns pointer of tranaction
func NewTransaction(inputs []TransactionInput, outputs []TransactionOutput) *Transaction {

	// TODO : implement

	return nil
}

func (t *Transaction) Inputs() []TransactionInput {

	return t.inputs
}

func (t *Transaction) Outputs() []TransactionOutput {

	return t.outputs
}

// in bytes
func (t *Transaction) Size() int {

	// TODO : implement

	return 0
}

// in satoshi
func (t *Transaction) Fee() int {

	// TODO : implement

	// sum inputs value - sum outputs value

	return 0
}

// sign transaction inputs and return transaction hex
func (t *Transaction) SignAndSerialize() ([]byte, error) {

	// TODO : implement

	return nil, nil
}

// broadcast transaction hex in blockchain and returns txID
func (t *Transaction) Broadcase() (string, error) {

	// TODO : implement

	return "", nil
}

func FetchTransactionByTxID(txID string) (BlockchainTransaction, error) {

	// TODO : implement

	return BlockchainTransaction{}, nil
}
