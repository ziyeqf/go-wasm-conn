package main

import (
	"bytes"
	"fmt"

	"github.com/magodo/go-wasmww"
	wasmconn "github.com/ziyeqf/go-wasm-conn"
)

func main() {
	var stdout, stderr bytes.Buffer
	worker := &wasmww.WasmWebWorkerConn{
		Name:   "server",
		Path:   "/server/server.wasm",
		Stdout: &stdout,
		Stderr: &stderr,
	}

	err := worker.Start()
	if err != nil {
		panic(err)
	}

	conn, err := wasmconn.NewWasmDialer("test", worker).Dial()
	if err != nil {
		panic(err)
	}
	conn.Write([]byte("Ping!"))

	b := make([]byte, 5)
	conn.Read(b)
	fmt.Println(string(b))

	conn.Close()
	worker.Terminate()
}
