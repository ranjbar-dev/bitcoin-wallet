package bitcoinWallet

import (
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ranjbar-dev/bitcoin-wallet/blockDaemon"
	"github.com/ranjbar-dev/bitcoin-wallet/blockDaemon/response"
	"github.com/ranjbar-dev/bitcoin-wallet/enums"
	"sort"
)

func sortUTXOsDESC(records []response.UTXO) []response.UTXO {

	sort.Slice(records, func(i, j int) bool {
		return records[i].Value > records[j].Value
	})

	return records
}

func sortUTXOsASC(records []response.UTXO) []response.UTXO {

	sort.Slice(records, func(i, j int) bool {
		return records[i].Value < records[j].Value
	})

	return records
}

func getAddressUTXO(chain *chaincfg.Params, address string) ([]response.UTXO, error) {

	node := enums.MAIN_NODE
	if &chaincfg.TestNet3Params == chain {
		node = enums.TEST_NODE
	}

	bd := blockDaemon.NewBlockDaemonService(node.Config)

	res, err := bd.AddressUTXO(address, "")
	if err != nil {
		return nil, err
	}

	return res.Data, nil
}

func prepareUTXOForTransaction(chain *chaincfg.Params, address string, amount int64) ([]response.UTXO, int64, error) {

	outputs := 2

	node := enums.MAIN_NODE
	if &chaincfg.TestNet3Params == chain {
		node = enums.TEST_NODE
	}

	feePerByte, err := estimateFeePerByte(node)
	if err != nil {
		return nil, 0, err
	}

	records, err := getAddressUTXO(chain, address)
	if err != nil {
		return nil, 0, err
	}

	remainingAmount := amount

	var final []response.UTXO
	var total int64

	if len(records) == 0 {
		return nil, 0, errors.New("not enough confirmed UTXOs, please try again later")
	}

	for remainingAmount > 0 {

		usedUTXO := len(final)

		for _, record := range sortUTXOsASC(records) {

			if usedUTXO == len(records) {
				return nil, 0, errors.New("not enough confirmed UTXOs, please try again later")
			}

			usedUTXO++

			if record.IsSpent {
				continue
			}

			recordValue := int64(record.Value)

			if recordValue > remainingAmount || usedUTXO+1 == len(records) {
				total += recordValue
				remainingAmount = remainingAmount - recordValue
				final = append(final, record)
				if remainingAmount < -int64(calculateTransactionSizeWithInputAndOutput(1, outputs)*feePerByte) {
					return final, total, nil
				} else {
					break
				}
			}
		}
	}

	return final, total, nil
}

func prepareUTXOForSweepTransaction(chain *chaincfg.Params, address string, usedUTXOs []string) ([]response.UTXO, int64, error) {

	records, err := getAddressUTXO(chain, address)
	if err != nil {
		return nil, 0, err
	}

	var final []response.UTXO
	var total int64

	for _, record := range sortUTXOsDESC(records) {

		// we used this txID but blockchain has not updated yet
		utxoAlreadyUsed := false
		for _, used := range usedUTXOs {
			if record.Mined.TxId == used {
				utxoAlreadyUsed = true
			}
		}
		if utxoAlreadyUsed {
			continue
		}

		if record.IsSpent {
			continue
		}

		if record.Mined.Confirmations >= 1 {

			final = append(final, record)

			total += int64(record.Value)
		}
	}

	return final, total, nil
}
