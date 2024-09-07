package netcat

import (
	"fmt"
	"net"
	"sync"
)

const (
	Method = "tcp"
	Usage = "[USAGE]: ./TCPChat $port"
)

type Servers struct {
	HelloMessage []byte
	Messages     [][]byte
	Conns map[string]net.Conn
	Mutex sync.Mutex
}

func Startserver() *Servers {
	return &Servers{
		HelloMessage: []byte{},
		Messages:     [][]byte{},
		Conns:        map[string]net.Conn{},
	}
}

// Accept the connection to the socket and redirect them.
func (Server *Servers) AcceptConnections(listner net.Listener) {
	for {
		// check if the group has a place for the user
		if len(Server.Conns) > 10 {
			continue
		}
		// accept the connexion and ask the user for his name
		con, err := listner.Accept()
		Name := make([]byte, 1024)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Server.Mutex.Lock()
		con.Write(Server.HelloMessage)
		Server.Mutex.Unlock()
		i, err := con.Read(Name)
		if err != nil {
			continue
		}
		// casse the name until the newline
		Name = Name[:i-1]
		// add the user to the  connection map
		Server.Conns[string(Name)] = con
		go Server.writeOldMessages(string(Name))
		// inform all the participates in the chat that the user entered
		msg := []byte(fmt.Sprintf("%s enter the chat\n", Name))
		Server.Mutex.Lock()
		Server.Messages = append(Server.Messages, msg)
		Server.Mutex.Unlock()
		go Server.printMessage(msg, string(Name))
		go Server.con_handler(string(Name))
	}
}
