package main

import (
	"github.com/ghost0036/go-hello/demochain/core"
	"github.com/ghost0036/go-hello/demochain/rpc"
)

func main() {
	bc := core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacky")
	bc.SendData("Send 1 EOS")
	bc.SendData("Why!!!")

	rpc.Run(bc)
}
