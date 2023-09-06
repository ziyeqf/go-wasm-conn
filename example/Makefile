serve: client server
	ln -sf "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .
	go run ./http_server.go

build:
	GOOS=js GOARCH=wasm go build -C ./server -o ./server.wasm
	GOOS=js GOARCH=wasm go build -C ./client -o ./client.wasm