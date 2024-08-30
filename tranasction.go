package bitcoinwallet

import (
	"fmt"

	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type BlockchainTransaction struct {
	// TODO : implment
}

type Transaction struct {
	inputs  []TransactionInput
	outputs []TransactionOutput
}

// returns pointer of tranaction
func NewTransaction(inputs []TransactionInput, outputs []TransactionOutput) *Transaction {

	return &Transaction{
		inputs:  inputs,
		outputs: outputs,
	}
}
func (t *Transaction) Inputs() []TransactionInput {

	return t.inputs
}

func (t *Transaction) Outputs() []TransactionOutput {

	return t.outputs
}

// in bytes
func (t *Transaction) Size() int {

	transactionSize := 10
	eachInputSizeInByte := 150
	eachOutputSizeInByte := 32

	for i := 0; i < len(t.inputs); i++ {
		transactionSize += eachInputSizeInByte
	}

	for i := 0; i < len(t.outputs); i++ {
		transactionSize += eachOutputSizeInByte
	}

	return transactionSize
}

// in satoshi
func (t *Transaction) Fee() int {

	inputsVal := 0
	outputsVal := 0

	for _, input := range t.inputs {
		inputsVal += input.Value
	}

	for _, output := range t.outputs {
		outputsVal += output.Value
	}

	return inputsVal - outputsVal
}

// sign transaction inputs and return transaction hex
func (t *Transaction) SignAndSerialize() ([]byte, error) {

	// TODO : implement

	return nil, nil
}

// broadcast transaction hex in blockchain and returns txID
func (t *Transaction) Broadcast() (string, error) {

	// TODO : implement

	return "", nil
}

func FetchTransactionByTxID(txID string) (models.Transaction, error) {

	tx, err := config.Explorer.GetTransactionByTxID(txID)

	if err != nil {
		fmt.Println(err)
		return models.Transaction{}, err
	}
	return tx, nil
}
