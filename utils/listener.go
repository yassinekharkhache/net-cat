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
		if err != nil {
			fmt.Println(err)
			continue
		}
		Name := make([]byte, 1024)
		Server.Mutex.Lock()
		con.Write(Server.HelloMessage)
		Server.Mutex.Unlock()
		i, err := con.Read(Name)
		if err != nil || i == 1 {
			if err != nil {
				fmt.Println("Error reading name:", err)
			}
			if len(Name[:i]) == 1 {
				errorMsg := "Invalid name. Connection will be closed.\n"
				con.Write([]byte(errorMsg))
			}
			con.Close()
			continue
		}
		// casse the name until the newline
		Name = Name[:i-1]
		// add the user to the  connection map
		Server.writeOldMessages(string(Name))
		Server.Conns[string(Name)] = con
		// inform all the participates in the chat that the user entered
		msg := []byte(fmt.Sprintf("%s enter the chat\n", Name))
		Server.Mutex.Lock()
		Server.Messages = append(Server.Messages, msg)
		Server.Mutex.Unlock()
		go Server.printMessage(msg, string(Name))
		go Server.con_handler(string(Name))
	}
}
