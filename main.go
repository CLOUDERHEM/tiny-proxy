package main

import (
	"tiny-proxy/acceptor"
	"tiny-proxy/proxy"
)

func main() {
	l := acceptor.Listen(8888)
	acceptor.Accept(l, proxy.HttpProxyHandle)
}
