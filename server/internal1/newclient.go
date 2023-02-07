package internal1

import (
	"net"
)

type User struct {
	Name string
	Addr string
	C    chan string
	Conn net.Conn
}

func NewClient(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		Conn: conn,
	}
	go user.ListenMsg()
	return user
}

func (nc *User) ListenMsg() {
	for {
		msg := <-nc.C
		nc.Conn.Write([]byte(msg))
	}
}

func (nc *User) SendMsg(msg string) {
	nc.Conn.Write([]byte(msg))
}
