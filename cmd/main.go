package main

import (
	"fmt"
	bitcoinWallet "github.com/ranjbar-dev/bitcoin-wallet"
	"github.com/ranjbar-dev/bitcoin-wallet/enums"
)

func main() {

	w, _ := bitcoinWallet.CreateBitcoinWallet(enums.TEST_NODE, "88414dbb373a211bc157265a267f3de6a4cec210f3a5da12e89630f2c447ad27")

	fmt.Println(w.Transfer("tb1q0r23g66m9rhhak8aahsg53wfp5egt2huuc4tnu", 1000))
	fmt.Println(w.EstimateTransferFee("tb1q0r23g66m9rhhak8aahsg53wfp5egt2huuc4tnu", 1000))
	//service := blockDaemon.NewBlockDaemonService(blockDaemon.TestNet)
	//fmt.Println(service.Tx("ba67fc134b884b3409d16bb7dd5538bf37b9def6327ce104002faa75e20e1c88"))
	//fmt.Println(service.AddressTxs("tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0"))
	//fmt.Println(service.AddressTxs("tb1qppv790u4dz48ctnk3p7ss7fmspckagp3wrfyp0"))
	//fmt.Println(service.CurrentBlockNumber())
	//fmt.Println(service.CurrentBlockHash())
	//fmt.Println(service.BlockByNumber(2408786))
	//fmt.Println(service.BlockByHash("000000000000000d1ef6490743e15a4b6d63f19e7f8d3899f817579ebc12d7ae"))
	//fmt.Println(service.Broadcast("000000000000000d1ef6490743e15a4b6d63f19e7f8d3899f817579ebc12d7ae"))
}
