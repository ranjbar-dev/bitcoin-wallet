package response

type UTXOResponse struct {
	Total int    `json:"total"`
	Data  []UTXO `json:"data"`
}

type UTXO struct {
	Status  string `json:"status"`
	IsSpent bool   `json:"is_spent"`
	Value   int    `json:"value"`
	Mined   struct {
		Index         int    `json:"index"`
		TxId          string `json:"tx_id"`
		Date          int    `json:"date"`
		BlockId       string `json:"block_id"`
		BlockNumber   int    `json:"block_number"`
		Confirmations int    `json:"confirmations"`
		Meta          struct {
			Index      int      `json:"index"`
			Script     string   `json:"script"`
			Addresses  []string `json:"addresses"`
			ScriptType string   `json:"script_type"`
		} `json:"meta"`
	} `json:"mined"`
	Spent struct {
		TxId          string `json:"tx_id"`
		Date          int    `json:"date"`
		BlockId       string `json:"block_id"`
		BlockNumber   int    `json:"block_number"`
		Confirmations int    `json:"confirmations"`
	} `json:"spent,omitempty"`
}
