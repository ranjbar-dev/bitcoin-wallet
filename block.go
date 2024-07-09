package bitcoinwallet

type BlockchainBlock struct {
	// TODO : implement
}

func FetchCurrentBlockNumber() (int, error) {

	// use config to get explorer type and apikey

	// trezor blockbook api: https://github.com/trezor/blockbook/blob/master/docs/api.md
	// https://btc1.trezor.io

	// TODO : implement

	return 0, nil
}

func FetchCurrentBlockHash() (int, error) {

	// use config to get explorer type and apikey

	// trezor blockbook api: https://github.com/trezor/blockbook/blob/master/docs/api.md
	// https://btc1.trezor.io

	// TODO : implement

	return 0, nil
}

func FetchBlockByNumber() (BlockchainBlock, error) {

	// use config to get explorer type and apikey

	// trezor blockbook api: https://github.com/trezor/blockbook/blob/master/docs/api.md
	// https://btc1.trezor.io

	// TODO : implement

	return BlockchainBlock{}, nil
}
