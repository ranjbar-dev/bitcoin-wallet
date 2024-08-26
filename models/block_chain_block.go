package models

type BlockChainBlock struct {
	Hash              string     `json:"hash"`
	Page              int        `json:"page"`
	TotalPages        int        `json:"totalPages"`
	PreviousBlockHash string     `json:"previousBlockHash"`
	NextBlockHash     string     `json:"nextBlockHash"`
	Height            int        `json:"height"`
	Confirmations     int        `json:"confirmations"`
	Size              int        `json:"size"`
	Time              int        `json:"time"`
	Version           int        `json:"version"`
	MerkleRoot        string     `json:"merkleRoot"`
	Nonce             string     `json:"nonce"`
	Bits              string     `json:"bits"`
	Difficulty        string     `json:"difficulty"`
	TxCount           int        `json:"txCount"`
	Txs               []BlockTxs `json:"txs"`
}
type BlockTxs struct {
	TxID string `json:"txid"`
	Vin  []struct {
		N     int    `json:"n"`
		Value string `json:"value"`
	}
	Vout []struct {
		N         int      `json:"n"`
		Value     string   `json:"value"`
		Addresses []string `json:"addresses"`
		IsAddress bool     `json:"isAddress"`
	}
	BlockHash     string `json:"blockHash"`
	BlockHeight   int    `json:"blockHeight"`
	Confirmations int    `json:"confirmations"`
	BlockTime     int    `json:"blockTime"`
	Value         string `json:"value"`
	ValueIn       string `json:"valueIn"`
	Fees          string `json:"fees"`
}
