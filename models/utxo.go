package models

type UTXO struct {
	Amount int64
	TxID   string
	Index  uint32
	Confirmations int
}
