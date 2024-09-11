package main

import (
	"net"
	"os"

	netcat "nc/outils"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	port := os.Args[1]
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	Server := netcat.CreateNewServer()
	for {
		Conn, err := l.Accept()
		if err != nil {
			continue
		}
		go Server.WelcomeToTheServer(Conn)
	}
}
