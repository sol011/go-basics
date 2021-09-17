package main

import (
	"sol011/go-basics/my-lib/httpserver"
	"sol011/go-basics/my-lib/interfaces"
)

func main() {
	interfaces.InterfaceTest()

	httpserver.StartServer()
}
