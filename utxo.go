package bitcoinWallet

import (
	"github.com/ranjbar-dev/bitcoin-wallet/blockDaemon/response"
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
