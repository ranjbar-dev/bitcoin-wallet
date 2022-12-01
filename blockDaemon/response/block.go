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

type EstimateFee struct {
	MostRecentBlock int `json:"most_recent_block"`
	EstimatedFees   struct {
		Fast   int `json:"fast"`
		Medium int `json:"medium"`
		Slow   int `json:"slow"`
	} `json:"estimated_fees"`
}
