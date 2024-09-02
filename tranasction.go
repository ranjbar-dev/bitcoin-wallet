package bitcoinwallet

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ranjbar-dev/bitcoin-wallet/models"
)

type Transaction struct {
	hex     string
	sweep   bool
	inputs  []models.TransactionInput
	outputs []models.TransactionOutput
}

// returns pointer of tranaction
func NewTransaction(inputs []models.TransactionInput, outputs []models.TransactionOutput) *Transaction {

	return &Transaction{
		inputs:  inputs,
		outputs: outputs,
	}
}
func (t *Transaction) Inputs() []models.TransactionInput {

	return t.inputs
}

func (t *Transaction) Outputs() []models.TransactionOutput {

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
		outputsVal += int(output.Value)
	}

	return inputsVal - outputsVal
}

// sign transaction inputs and return transaction hex
func (t *Transaction) SignAndSerialize() error {

	inputsVal := 0
	outputsVal := 0

	for _, input := range t.inputs {
		inputsVal += input.Value
	}

	for _, output := range t.outputs {
		outputsVal += int(output.Value)
	}

	// fee := inputsVal - outputsVal

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
		fmt.Println(output)
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

// broadcast transaction hex in blockchain and returns txID
func (t *Transaction) Broadcast() (string, error) {

	res, err := config.Explorer.BroadcastTransaction(t.hex)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return res, nil
}

func FetchTransactionByTxID(txID string) (models.Transaction, error) {

	tx, err := config.Explorer.GetTransactionByTxID(txID)

	if err != nil {
		fmt.Println(err)
		return models.Transaction{}, err
	}
	return tx, nil
}
