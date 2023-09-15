package wasmconn

import (
	"syscall/js"

	"github.com/hack-pad/safejs"
)

type wasmConnRequest struct {
	ConnectStr string
	ConnId     string
}

type WasmConnResponse struct {
	ConnId string
}

type wasmConnClose struct {
	ConnId string
}

type WasmConnMessage struct {
	ConnId string
	Bytes  []byte
}
