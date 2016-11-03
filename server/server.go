package server

import (
	"bufio"
	"flag"
	"fmt"
	"net"

	"strings" // only needed below for sample processing

	"github.com/ralphg6/chat-server/command"
	"github.com/ralphg6/chat-server/engine"
)

type Server struct {
	Engine   engine.Engine
	Commands []command.Command
}

/*var ch = make(chan ChannelStruct, 100)

type Server struct {
	Users  map[string]User
	Groups map[string]Group
}*/

func (s *Server) handleConnection(conn net.Conn) {
	for {
		found := false
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		if message == "" {
			break
		}

		if len(message) > 2 {
			// remove "\r\n" end of message
			message = message[1 : len(message)-2]

			tokens := strings.Split(message, " ")
			if len(tokens) == 0 {
				continue
			}

			//console.AppendHistory(line)

			command := tokens[0]
			args := tokens[1:]

			for _, commandHandler := range s.Commands {
				commandHandler.SetConn(conn)
				if commandHandler.Supports(command) {
					found = true
					err := commandHandler.Verify(args)
					if err != nil {
						conn.Write([]byte(fmt.Sprintln(err)))
					} else {
						commandHandler.Handle(args)
					}
					break
				}
			}
		}

		if !found {
			conn.Write([]byte("invalid command\n"))
		}

	}
}

func (s *Server) Serve() (err error) {
	var port = flag.String("port", "8081", "server port")

	flag.Parse()

	fmt.Printf("Launching server in...\n")

	ifaces, err := net.Interfaces()
	if err != nil {
		return
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return err
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			// process IP address
			if ip.To4() != nil {
				fmt.Printf(" - tcp://%s:%s\n", ip, *port)
			}
		}
	}

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":"+*port)

	// run loop forever (or until ctrl-c)
	for {
		// accept connection on port
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		defer conn.Close()

		go s.handleConnection(conn)
	}

	return
}
