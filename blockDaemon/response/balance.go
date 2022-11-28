package response

type BalanceResponse []Balance

type Balance struct {
	Currency         Currency `json:"currency"`
	ConfirmedBalance string   `json:"confirmed_balance"`
	ConfirmedBlock   int      `json:"confirmed_block"`
}

type Currency struct {
	AssetPath string `json:"asset_path"`
	Symbol    string `json:"symbol"`
	Name      string `json:"name"`
	Decimals  int    `json:"decimals"`
	Type      string `json:"type"`
}
