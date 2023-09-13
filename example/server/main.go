package main

import (
	"log"
	"syscall/js"

	"github.com/magodo/go-wasmww"
	wasmconn "github.com/ziyeqf/go-wasm-conn"
)

func main() {
	self, err := wasmww.NewSelfConn()
	if err != nil {
		log.Fatal(err)
	}
	ch, err := self.SetupConn()
	if err != nil {
		log.Fatal(err)
	}

	listener := wasmconn.NewListener("test", self.PostMessage, ch, self.Close)

	doneCh := make(chan interface{}, 0)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			go func() {
				b := make([]byte, 5)
				conn.Read(b)
				js.Global().Get("console").Call("log", js.ValueOf(string(b)))
				conn.Write([]byte("Pong!"))
				conn.Close()
			}()
		}
	}()

	<-doneCh
}
