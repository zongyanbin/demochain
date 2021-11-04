package main

import "demochain/core"

func main()  {
	bc :=core.NewBlockchain()
	bc.SendData("Send 1 BTC to Jacy")
	bc.SendData("Send 1 EOS to Jack")
	bc.Print()
}
