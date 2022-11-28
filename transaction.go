package bitcoinWallet

import (
	"bytes"
	"encoding/hex"
	"errors"
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon"
	"github.com/Amirilidan78/bitcoin-wallet/blockDaemon/response"
	"github.com/Amirilidan78/bitcoin-wallet/enums"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
)

func getAddressUTXO(chain *chaincfg.Params, address string) ([]response.UTXO, error) {

	node := enums.MAIN_NODE
	if &chaincfg.TestNet3Params == chain {
		node = enums.TEST_NODE
	}

	bd := blockDaemon.NewBlockDaemonService(node.Config)

	res, err := bd.AddressUTXO(address)
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func prepareUTXOForTransaction(chain *chaincfg.Params, address string, amount int64, fee int64) ([]response.UTXO, int64, error) {

	records, err := getAddressUTXO(chain, address)
	if err != nil {
		return nil, 0, err
	}

	var final []response.UTXO
	var total int64

	for _, record := range records {

		if total >= (amount + fee) {
			break
		}

		if record.Mined.Confirmations > 2 {

			final = append(final, record)

			total += int64(record.Value)
		}
	}

	return final, total, nil
}

func createTransactionAndSignTransaction(chain *chaincfg.Params, fromAddress string, privateKey *btcec.PrivateKey, toAddress string, amount int64, fee int64) (*wire.MsgTx, error) {

	fromAddr, err := btcutil.DecodeAddress(fromAddress, chain)
	if err != nil {
		return nil, errors.New("DecodeAddress fromAddr err " + err.Error())
	}

	fromAddrScriptByte, err := txscript.PayToAddrScript(fromAddr)

	if err != nil {
		return nil, errors.New("fromAddr PayToAddrScript err " + err.Error())
	}

	toAddr, err := btcutil.DecodeAddress(toAddress, chain)
	if err != nil {
		return nil, errors.New("DecodeAddress destAddrStr err " + err.Error())
	}

	toAddrByte, err := txscript.PayToAddrScript(toAddr)
	if err != nil {
		return nil, errors.New("toAddr PayToAddrScript err " + err.Error())
	}

	fromAddrByte, err := txscript.PayToAddrScript(fromAddr)
	if err != nil {
		return nil, errors.New("fromAddrByte PayToAddrScript err " + err.Error())
	}

	utxoList, totalAmount, err := prepareUTXOForTransaction(chain, fromAddress, amount, fee)
	if err != nil {
		return nil, errors.New("vin err " + err.Error())
	}
	if totalAmount < amount || len(utxoList) == 0 {
		return nil, errors.New("insufficient balance")
	}

	t, err := createTransactionInputsAndSign(privateKey, utxoList, fromAddrByte, fromAddrScriptByte, toAddrByte, totalAmount, amount, fee)
	if err != nil {
		return nil, errors.New("vin err " + err.Error())
	}

	return t, nil
}

func createTransactionInputsAndSign(privateKey *btcec.PrivateKey, utxos []response.UTXO, fromAddressByte []byte, fromAddressScriptByte []byte, toAddressByte []byte, totalAmount int64, amount int64, fee int64) (*wire.MsgTx, error) {

	transaction := wire.NewMsgTx(2)

	// vin
	for _, utxo := range utxos {

		hash, err := chainhash.NewHashFromStr(utxo.Mined.TxId)
		if err != nil {
			return nil, err
		}

		txIn := wire.NewTxIn(wire.NewOutPoint(hash, uint32(utxo.Mined.Index)), nil, [][]byte{})
		txIn.Sequence = txIn.Sequence - 2
		transaction.AddTxIn(txIn)
	}

	// vout
	changeAmount := totalAmount - amount - fee
	transaction.AddTxOut(wire.NewTxOut(amount, toAddressByte))
	if changeAmount > 0 {
		transaction.AddTxOut(wire.NewTxOut(changeAmount, fromAddressByte))
	}

	transaction.LockTime = 0

	signerMap := make(map[wire.OutPoint]*wire.TxOut)
	for _, in := range transaction.TxIn {
		signerMap[in.PreviousOutPoint] = &wire.TxOut{}
	}
	sigHashes := txscript.NewTxSigHashes(transaction, txscript.NewMultiPrevOutFetcher(signerMap))

	// sign
	for index, utxo := range utxos {

		signature, err := txscript.WitnessSignature(transaction, sigHashes, index, int64(utxo.Value), fromAddressScriptByte, txscript.SigHashAll, privateKey, true)
		if err != nil {
			return nil, errors.New("WitnessSignature err " + err.Error())
		}

		transaction.TxIn[index].Witness = signature
	}

	return transaction, nil
}

func getRawTransaction(tx *wire.MsgTx) (string, error) {

	var signedTx bytes.Buffer

	err := tx.Serialize(&signedTx)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(signedTx.Bytes()), nil
}

func broadcastHex(chain *chaincfg.Params, hex string) (string, error) {

	node := enums.MAIN_NODE
	if &chaincfg.TestNet3Params == chain {
		node = enums.TEST_NODE
	}

	res, err := blockDaemon.NewBlockDaemonService(node.Config).Broadcast(hex)
	if err != nil {
		return "", err
	}

	return res.Id, nil
}

func createSignAndBroadcastTransaction(chain *chaincfg.Params, privateKey *btcec.PrivateKey, fromAddress string, toAddress string, amount int64, fee int64) (string, error) {

	// signed tx
	tx, err := createTransactionAndSignTransaction(chain, fromAddress, privateKey, toAddress, amount, fee)
	if err != nil {
		return "", err
	}

	// raw
	raw, err := getRawTransaction(tx)
	if err != nil {
		return "", err
	}

	return broadcastHex(chain, raw)
}
