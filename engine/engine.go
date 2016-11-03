package engine

import (
	"net"
	"strings"
)

//"fmt"
//"strings"

//"golang.org/x/net/context"

type Engine struct {
}

func (e *Engine) Echo(conn net.Conn, message string) (err error) {
	newmessage := strings.ToUpper(message)

	_, err = conn.Write([]byte(newmessage + "\n"))
	if err != nil {
		return
	}

	return
}

func (e *Engine) Quit(conn net.Conn) (err error) {

	err = conn.Close()
	if err != nil {
		return
	}

	return
}
