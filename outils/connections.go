package nc

import (
	"errors"
	"fmt"
	"net"
	"time"
)

func (s *Server) GetUserInfo(Conn net.Conn) ([]byte, []byte, error) {
	Name := make([]byte, 1024)
	group := make([]byte, 1024)
	res, err := Conn.Read(Name)
	Name = Name[:res-1]
	if err != nil || res == 1 || s.invalidName(Name) {
		Conn.Write([]byte("invalid name"))
		return nil, nil, err
	}
	Conn.Write([]byte("Choose your group chat:\n"))
	i := 1
	for group := range s.Groups {
		Msg := fmt.Sprintf("%d. %s\n", i, group)
		Conn.Write([]byte(Msg))
		i++
	}
	Conn.Write([]byte("Name Of Groupe : "))
	res, err = Conn.Read(group)
	group = group[:res-1]
	if err != nil || res == 1 || s.invalidGrpoup(group) {
		Conn.Write([]byte("invalid group"))
		if err == nil {
			err = errors.New("invalid group")
		}
		return nil, nil, err
	}
	if len(s.Groups[string(group)]) == 4 {
		Conn.Write([]byte("Groupe already have 4 people"))
		return nil, nil, err
	}
	return Name, group, nil
}

func (s *Server) invalidName(name []byte) bool {
	if InvalidMsg(name) {
		return true
	}
	for _, Group := range s.Groups {
		for Name := range Group {
			if Name == string(name) {
				return true
			}
		}
	}
	return false
}

func InvalidMsg(name []byte) bool {
	for _, v := range name {
		if v > 126 || v < 32 {
			return true
		}
	}
	return false
}

func (s *Server) WriteOldMessages(Conn net.Conn, group string) {
	for _, message := range s.Messages[group] {
		Conn.Write([]byte(message))
	}
}

func (s *Server) WriteMessage(Name string, group string, message string) {
	s.mutex.Lock()
	s.Messages[group] = append(s.Messages[group], message)
	s.mutex.Unlock()
	for UserName, UserConn := range s.Groups[group] {
		if UserName != Name {
			UserConn.Write([]byte("\n" + message))
			UserConn.Write([]byte(Format("", string(UserName))))
		}
	}
}

func Format(msg, name string) string {
	time := time.Now().Format("[2020-01-20 16:03:43]")
	return fmt.Sprintf("%s[%s]: %s", time, name, msg)
}

func (s *Server) invalidGrpoup(Grp []byte) bool {
	for group := range s.Groups {
		if string(Grp) == group {
			return false
		}
	}
	return true
}