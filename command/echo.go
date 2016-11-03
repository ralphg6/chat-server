package command

import (
	"errors"
	"net"

	"strings"

	"github.com/ralphg6/chat-server/engine"
)

type Echo struct {
	Engine engine.Engine
	conn   net.Conn
}

func (c *Echo) Supports(command string) bool {
	return strings.EqualFold(command, "echo")
}

func (c *Echo) Help() string {
	return "Echo method (/echo arg1)"
}

func (c *Echo) Handle(args []string) {
	c.Engine.Echo(c.conn, args[0])
}

func (c *Echo) Verify(args []string) error {
	if len(args) != 1 {
		return errors.New("wrong number of arguments, get command requires one argument")
	}
	return nil
}

func (c *Echo) SetConn(conn net.Conn) {
	c.conn = conn
}

func (c *Echo) String() string {
	return "echo"
}
