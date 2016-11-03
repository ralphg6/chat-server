package command

import "net"

type Command interface {
	Supports(string) bool
	Help() string
	Handle([]string)
	Verify([]string) error
	String() string
	SetConn(net.Conn)
}
