package command

import (
	"errors"
	"net"

	"strings"

	"github.com/ralphg6/chat-server/engine"
)

type Quit struct {
	Engine engine.Engine
	conn   net.Conn
}

func (c *Quit) Supports(command string) bool {
	return strings.EqualFold(command, "quit")
}

func (c *Quit) Help() string {
	return "Exit (/quit)"
}

func (c *Quit) Handle(args []string) {
	c.Engine.Quit(c.conn)
}

func (c *Quit) Verify(args []string) error {
	if len(args) != 0 {
		return errors.New("wrong number of arguments, quit command no permit arguments")
	}
	return nil
}

func (c *Quit) SetConn(conn net.Conn) {
	c.conn = conn
}

func (c *Quit) String() string {
	return "quit"
}
