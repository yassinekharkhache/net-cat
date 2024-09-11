package nc

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Users map[string]net.Conn

func CreateNewServer() *Server {
	groups := make(map[string]Users, 3)
	groups["General"] = Users{}
	groups["Random"] = Users{}
	groups["vip"] = Users{}
	return &Server{
		Groups:   groups,
		Messages: make(map[string][]string, 3),
	}
}

type Server struct {
	Groups   map[string]Users
	mutex    sync.Mutex
	Messages map[string][]string
}

func (s *Server) WelcomeToTheServer(Conn net.Conn) {
	defer Conn.Close()
	Conn.Write([]byte("Welcome to our Server Ener your name:\n"))
	Name, group, err := s.GetUserInfo(Conn)
	if err != nil {
		return
	}
	s.WriteOldMessages(Conn, string(group))
	s.mutex.Lock()
	// add new User
	s.Groups[string(group)][string(Name)] = Conn
	s.mutex.Unlock()
	// join Message
	JoinMessage := fmt.Sprintf("%s Joined The groupe chat\n", string(Name))
	s.WriteMessage(string(Name), string(group), JoinMessage)
	s.StartChat(string(Name), string(group))
	DisconnectMsg := fmt.Sprintf("%s Leave The groupe chat\n", string(Name))
	s.WriteMessage(string(Name), string(group), DisconnectMsg)
	s.mutex.Lock()
	delete(s.Groups[string(group)], string(Name))
	s.mutex.Unlock()
}

func (s *Server) StartChat(Name, group string) {
	Conn := s.Groups[group][Name]
	for {
		Conn.Write([]byte(Format("", Name)))
		msg := make([]byte, 2048)
		index, err := Conn.Read(msg)
		if err != nil || index == 1 || InvalidMsg(msg[:index-1]) {
			if !bufio.NewScanner(Conn).Scan() {
				break
			}
			continue
		}
		message := Format(string(msg), Name)
		s.WriteMessage(Name, group, message)
	}
}
