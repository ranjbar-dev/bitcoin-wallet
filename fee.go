package bitcoinWallet

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/ranjbar-dev/bitcoin-wallet/blockDaemon"
	"github.com/ranjbar-dev/bitcoin-wallet/enums"
)

func estimateTransactionFee(chain *chaincfg.Params, fromAddress string, toAddress string, amount int64) (int64, error) {

	node := enums.MAIN_NODE
	if &chaincfg.TestNet3Params == chain {
		node = enums.TEST_NODE
	}

	toAddr, err := btcutil.DecodeAddress(toAddress, chain)
	if err != nil {
		return 0, errors.New("DecodeAddress destAddrStr err " + err.Error())
	}

	toAddressByte, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return 0, errors.New("toAddr PayToAddrScript err " + err.Error())
	}

	feePerByte, err := estimateFeePerByte(node)
	if err != nil {
		return 0, err
	}

	utxoList, totalAmount, err := prepareUTXOForTransaction(chain, fromAddress, amount)
	if err != nil {
		return 0, errors.New("vin err " + err.Error())
	}
	if len(utxoList) == 0 {
		return 0, errors.New("insufficient balance")
	}

	for _, utxo := range utxoList {
		fmt.Println(utxo.Mined.TxId)
		fmt.Println(float64(utxo.Value) / 100000000)
	}

	transaction := wire.NewMsgTx(2)

	for _, utxo := range utxoList {

		hash, err := chainhash.NewHashFromStr(utxo.Mined.TxId)
		if err != nil {
			return 0, err
		}

		txIn := wire.NewTxIn(wire.NewOutPoint(hash, uint32(utxo.Mined.Index)), nil, [][]byte{})
		txIn.Sequence = txIn.Sequence - 2
		transaction.AddTxIn(txIn)
	}

	transaction.AddTxOut(wire.NewTxOut(amount, toAddressByte))

	transactionSizeInByte := calculateTransactionSize(transaction)

	fee := int64(transactionSizeInByte * feePerByte)

	changeAmount := totalAmount - fee
	// to avoid dust
	if changeAmount > 500 {
		transaction.AddTxOut(wire.NewTxOut(changeAmount, toAddressByte))
	}

	transactionSizeInByte = calculateTransactionSize(transaction)

	fee = int64(transactionSizeInByte * feePerByte)

	return fee, nil
}

func estimateFeePerByte(node enums.Node) (int, error) {

	bd := blockDaemon.NewBlockDaemonService(node.Config)

	res, err := bd.EstimateFee()
	if err != nil {
		return 0, err
	}

	return res.EstimatedFees.Slow, nil
}

func calculateTransactionSize(transaction *wire.MsgTx) int {

	return calculateTransactionSizeWithInputAndOutput(len(transaction.TxIn), len(transaction.TxOut))
}

func calculateTransactionSizeWithInputAndOutput(inputs int, outputs int) int {

	transactionSize := 10
	eachInputSizeInByte := 150
	eachOutputSizeInByte := 32

	for i := 0; i < inputs; i++ {
		transactionSize += eachInputSizeInByte
	}

	for i := 0; i < outputs; i++ {
		transactionSize += eachOutputSizeInByte
	}

	return transactionSize
}
