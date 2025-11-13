package netcat

import (
	"fmt"
	"net"
)

const (
	Method = "tcp"
	Port   = ":8080"
)

type Servers struct {
	HelloMessage []byte
	Messages     [][]byte
	Users        []net.Conn
	Conns        map[string]net.Conn
}

var Server Servers = Servers{
	HelloMessage :[]byte{},
	Messages     :[][]byte{},
	Users        :[]net.Conn{},
	Conns        :map[string]net.Conn{},
}
func AcceptConnections(listner net.Listener) {
	for {
		// check if the group has a place for the user
		if len(Server.Users) > 10 {
			print_Message_to_con([]byte("we get the limit of the group\ntry later\n"), "")
			continue
		}
		// accept the connexion and ask the user for his name
		con, err := listner.Accept()
		Name := make([]byte, 1024)
		if err != nil {
			fmt.Println(err)
			continue
		}

		con.Write(Server.HelloMessage)
		i, err := con.Read(Name)
		if err != nil {
			continue
		}
		// casse the name until the newline
		Name = Name[:i-1]
		// inform all the participates in the chat that the user entered
		msg := []byte(fmt.Sprintf("%s enter the chat\n", Name))
		Server.Messages = append(Server.Messages, msg)
		// add the user to the users slice
		Server.Users = append(Server.Users, con)
		printMessage(msg, string(Name))
		Server.Conns[string(Name)] = con
		writeOldMessages(string(Name))
		go con_handler(string(Name))
	}
}
