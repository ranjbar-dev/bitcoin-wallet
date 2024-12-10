package models

type BlockInfo struct {
}

type Block struct {
	Hash              string
	PreviousBlockHash string
	NextBlockHash     string
	Height            int
	Confirmations     int
	TxCount           int
	Txs               []Transaction
}

type Transaction struct {
	TxID              string
	Inputs            []Input
	Outputs           []Output
	Confirmations     int
	TotalInputsValue  int
	TotalOutputsValue int
	Size              int
	VBytes            int
	Fee               int
	BlockNumber       int
	BlockHash         string
	Timestamp         int
}

type Input struct {
	Address string
	TxID    string
	Index   int
	Value   int
}

type Output struct {
	Address string
	Index   int
	Value   int
}
