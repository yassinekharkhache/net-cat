package netcat

import (
	"bufio"
	"fmt"
	"strconv"
	"time"
)

// this function write the old messages in the chat to the new user
func (Server *Servers) writeOldMessages(name string) {
	for _, msg := range Server.Messages {
		Server.Mutex.Lock()
		Server.Conns[name].Write(msg)
		Server.Mutex.Unlock()
	}
}

// this function take two variables the name of the user and the textmessage
// and format it to be in the right message formula
func formatMessages(name, text string) []byte {
	msg := fmt.Sprintf("[%s]: %s\n", name, text)
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	return []byte(fmt.Sprintf("[%d-%s-%d %d:%d:%d]%s", year, month, day, hour, min, sec, msg))
}

// that function handle the connextion for a single user by his name
func (Server *Servers) con_handler(name string) {
	Conn := Server.Conns[name]
	scanner := bufio.NewScanner(Conn)
	for scanner.Scan() {
		msg := formatMessages(name, scanner.Text())
		Server.Mutex.Lock()
		Server.Messages = append(Server.Messages, msg)
		Server.Mutex.Unlock()
		go Server.printMessage([]byte(msg), name)
	}
	delete(Server.Conns, name)
	Server.printMessage([]byte(fmt.Sprintf("%s left the chat\n", name)), name)
}

// print the the message to the other users
func (Server *Servers) printMessage(msg []byte, name string) {
	for Name, con := range Server.Conns {
		if Name == name {
			continue
		}
		Server.Mutex.Lock()
		con.Write(msg)
		Server.Mutex.Unlock()
	}
}

//check if the port is number plus in the valid range 
func IsValidPort(s string) bool {
	num, err := strconv.Atoi(s)
	if err != nil || num < 1024 || num > 65535 {
		return false
	}
	return true
}