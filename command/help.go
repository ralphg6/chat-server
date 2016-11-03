package command

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Help struct {
	Commands []Command
	conn     net.Conn
}

func (c *Help) Supports(command string) bool {
	return strings.EqualFold(command, "help")
}

func (c *Help) Help() string {
	return "Help method (/help)"
}

func (c *Help) Handle(args []string) {
	text := "\tMessage Of The Day\n\n"
	text += "Commands:\n"
	for _, cmd := range c.Commands {
		text += fmt.Sprintf(" - %s: %s\n", cmd.String(), cmd.Help())
	}
	c.conn.Write([]byte(text))
}

func (c *Help) Verify(args []string) error {
	if len(args) != 0 {
		return errors.New("wrong number of arguments, get command no suport arguments")
	}
	return nil
}

func (c *Help) SetConn(conn net.Conn) {
	c.conn = conn
}

func (c *Help) String() string {
	return "help"
}
