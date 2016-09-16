package main

import (
	"fmt"
	"net/http"
	"time"
)

var ch = make(chan ChannelStruct, 100)

type Server struct {
	Users  map[string]User
	Groups map[string]Group
}

func handlerOK(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func handlerPooling(w http.ResponseWriter, r *http.Request) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(30 * time.Second)
		timeout <- true
	}()
	select {
	case msg := <-ch:
		fmt.Fprintf(w, "%s", msg.Teste)
	case <-timeout:
		fmt.Fprintf(w, "TIMEOUT")
	}
}

func handlerTest(w http.ResponseWriter, r *http.Request) {
	ch <- ChannelStruct{Teste: "mensagem de teste"}
}

func Serve() {
	http.HandleFunc("/", handlerOK)
	http.HandleFunc("/pooling", handlerPooling)
	http.HandleFunc("/test", handlerTest)
	http.ListenAndServe(":8080", nil)
}
