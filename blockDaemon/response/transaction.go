package response

type AddressTransactionResponse struct {
	Total int           `json:"total"`
	Data  []Transaction `json:"data"`
}

type BroadcastTransactionResponse struct {
	Id string `json:"id"`
}

type Transaction struct {
	Id            string             `json:"id"`
	BlockId       string             `json:"block_id"`
	Date          int                `json:"date"`
	Status        string             `json:"status"`
	NumEvents     int                `json:"num_events"`
	Meta          TransactionMeta    `json:"meta"`
	BlockNumber   int                `json:"block_number"`
	Confirmations int                `json:"confirmations"`
	Events        []TransactionEvent `json:"events"`
}

type TransactionMeta struct {
	Vsize int `json:"vsize"`
}

type TransactionEvent struct {
	Id            string                `json:"id"`
	TransactionId string                `json:"transaction_id"`
	Type          string                `json:"type"`
	Denomination  string                `json:"denomination"`
	Source        string                `json:"source,omitempty"`
	Meta          *TransactionEventMeta `json:"meta"`
	Date          int                   `json:"date"`
	Amount        int                   `json:"amount"`
	Decimals      int                   `json:"decimals"`
	Destination   string                `json:"destination,omitempty"`
}

type TransactionEventMeta struct {
	Index      int      `json:"index"`
	Script     string   `json:"script"`
	ScriptType string   `json:"script_type"`
	Addresses  []string `json:"addresses,omitempty"`
}
