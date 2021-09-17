package main

import (
	"sol011/go-basics/my-lib/concurrency"
	"sol011/go-basics/my-lib/httpserver"
	"sol011/go-basics/my-lib/interfaces"
)

func main() {
	interfaces.InterfaceTest()

	concurrency.SimpleWait()
	// concurrency.DeadlockReadingFromForeverEmptyChan()
	// concurrency.UnbufferedChansBlockSendUntilReceive()

	httpserver.StartServer()
}
