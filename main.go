package main

import (
	"net"
	"os"

	netcat "nc/outils"
)

func main() {
	if len(os.Args) > 2 {
		return
	}
	port := "8989"
	if len(os.Args) == 2 {
		port = os.Args[1]
	}
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	defer l.Close()
	Server := netcat.CreateNewServer()
	for {
		Conn, err := l.Accept()
		if err != nil {
			continue
		}
		go Server.WelcomeToTheServer(Conn)
	}
}
