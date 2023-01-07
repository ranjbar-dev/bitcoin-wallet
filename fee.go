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

	bd := blockDaemon.NewBlockDaemonService(node.Config)

	res, err := bd.EstimateFee()
	if err != nil {
		return 0, err
	}

	utxoList, _, err := prepareUTXOForTransaction(chain, fromAddress, amount)
	if err != nil {
		return 0, errors.New("vin err " + err.Error())
	}
	if len(utxoList) == 0 {
		return 0, errors.New("insufficient balance")
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

	transactionSizeInByte := transaction.SerializeSize() + 110

	fmt.Println("transactionSizeInByte")
	fmt.Println(transactionSizeInByte)

	return int64(transactionSizeInByte * res.EstimatedFees.Slow), nil
}
