package explorer

type Explorer interface {
	GetAddressBalance(address string) (int, error)
}

func NewTrezorExplorer() TrezorExplorer {

	return TrezorExplorer{
		baseURL: "https://btc1.trezor.io/api/v2",
	}
}

func NewBlockdaemonExplorer(protocol string, network string, apiKey string) BlockdaemonExplorer {

	return BlockdaemonExplorer{
		Protocol: protocol,
		Network:  network,
		baseURL:  "https://svc.blockdaemon.com/universal/v1",
		ApiKey:   apiKey,
	}
}
