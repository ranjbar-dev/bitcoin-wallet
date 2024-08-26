package explorer

type Explorer interface {
	GetAddressBalance(address string) (int, error)
}

func NewTrezorExplorer() TrezorExplorer {

	return TrezorExplorer{}
}

func NewBlockdaemonExplorer(protocol string, network string) BlockdaemonExplorer {

	return BlockdaemonExplorer{
		Protocol: protocol,
		Network:  network,
		ApiKey:   "im4YrpAa9tjvFcwlZDci22aQGzp4JtAqnQtdzcMXAIdj-Aoi",
	}
}
