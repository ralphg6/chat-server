package main

import (
	"github.com/ralphg6/chat-server/command"
	"github.com/ralphg6/chat-server/engine"
	"github.com/ralphg6/chat-server/server"
)

func main() {

	s := &server.Server{
		Engine: engine.Engine{},
	}

	s.Commands = []command.Command{
		&command.Echo{Engine: s.Engine},
		&command.Quit{Engine: s.Engine},
	}

	s.Commands = append(s.Commands, &command.Help{Commands: s.Commands})

	s.Serve()
}
