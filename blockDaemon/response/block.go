package response

type CurrentBlockNumberResponse int

type CurrentBlockHashResponse string

type BlockResponse struct {
	Number   int           `json:"number"`
	Id       string        `json:"id"`
	ParentId string        `json:"parent_id"`
	Date     int           `json:"date"`
	NumTxs   int           `json:"num_txs"`
	Meta     interface{}   `json:"meta"`
	Txs      []Transaction `json:"txs"`
}
