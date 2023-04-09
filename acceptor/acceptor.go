package acceptor

import (
	"log"
	"net"
	"strconv"
)

func Listen(port int) net.Listener {
	l, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Panic(err)
	}
	return l
}

func Accept(listener net.Listener, handler func(conn net.Conn)) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error | %v", err)
		} else {
			go handler(conn)
		}
	}
}
