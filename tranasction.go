package bitcoinwallet

import (
	"bytes"
	"encoding/hex"
	"errors"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type Transaction struct {
	hex     string
	inputs  []models.TransactionInput
	outputs []models.TransactionOutput
}

// NewTransaction creates a new transaction with the given inputs and outputs.
func NewTransaction(inputs []models.TransactionInput, outputs []models.TransactionOutput) *Transaction {

	return &Transaction{
		inputs:  inputs,
		outputs: outputs,
	}
}

// Inputs returns copy of transaction inputs
func (t *Transaction) Inputs() []models.TransactionInput {

	return t.inputs
}

// Outputs returns copy of transaction outputs
func (t *Transaction) Outputs() []models.TransactionOutput {

	return t.outputs
}

// Hex
func (t *Transaction) Hex() string {

	return t.hex
}

// Size returns the size of the transaction in bytes.
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

// Fee returns the fee of the transaction in satoshi.
func (t *Transaction) Fee() int64 {

	var inputsVal int64
	var outputsVal int64

	for _, input := range t.inputs {
		inputsVal += input.Value
	}

	for _, output := range t.outputs {
		outputsVal += output.Value
	}

	return inputsVal - outputsVal
}

// SignAndSerialize signs the transaction inputs and returns the transaction hex.
func (t *Transaction) SignAndSerialize() error {

	var inputsVal int64
	var outputsVal int64

	for _, input := range t.inputs {

		inputsVal += input.Value
	}

	for _, output := range t.outputs {

		outputsVal += output.Value
	}

	tx := wire.NewMsgTx(2)

	for _, utxo := range t.inputs {

		hash, err := chainhash.NewHashFromStr(utxo.TxId)
		if err != nil {

			return err
		}

		txIn := wire.NewTxIn(wire.NewOutPoint(hash, uint32(utxo.Index)), nil, [][]byte{})
		txIn.Sequence = txIn.Sequence - 2
		tx.AddTxIn(txIn)
	}

	for _, output := range t.outputs {

		tx.AddTxOut(wire.NewTxOut(output.Value, output.Address))
	}

	tx.LockTime = 0

	signerMap := make(map[wire.OutPoint]*wire.TxOut)

	for _, in := range tx.TxIn {
		signerMap[in.PreviousOutPoint] = &wire.TxOut{}
	}
	sigHashes := txscript.NewTxSigHashes(tx, txscript.NewMultiPrevOutFetcher(signerMap))

	// sign
	for index, utxo := range t.inputs {

		fromAddr, err := btcutil.DecodeAddress(utxo.Address, config.Chaincfg)
		if err != nil {
			return errors.New("DecodeAddress fromAddr err " + err.Error())
		}

		fromAddrByte, err := txscript.PayToAddrScript(fromAddr)
		if err != nil {
			return errors.New("fromAddrByte PayToAddrScript err " + err.Error())
		}

		secpPrivKey := secp256k1.PrivKeyFromBytes(utxo.PrivateKey.D.Bytes())

		signature, err := txscript.WitnessSignature(tx, sigHashes, index, int64(utxo.Value), fromAddrByte, txscript.SigHashAll, secpPrivKey, true)
		if err != nil {
			return errors.New("WitnessSignature err " + err.Error())
		}

		tx.TxIn[index].Witness = signature
	}

	var signedTx bytes.Buffer

	err := tx.Serialize(&signedTx)
	if err != nil {
		return err
	}

	t.hex = hex.EncodeToString(signedTx.Bytes())

	return nil
}

// Broadcast broadcasts the transaction hex to the blockchain and returns the transaction ID.
func (t *Transaction) Broadcast() (string, error) {

	res, err := config.Explorer.BroadcastTransaction(t.hex)

	if err != nil {

		return "", err
	}

	return res, nil
}

// FetchTransactionByTxID fetches a transaction by its transaction ID.
func FetchTransactionByTxID(txID string) (models.Transaction, error) {

	tx, err := config.Explorer.GetTransactionByTxID(txID)

	if err != nil {

		return models.Transaction{}, err
	}
	return tx, nil
}
