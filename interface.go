package wasmconn

import "github.com/hack-pad/safejs"

type PostFunc func(data safejs.Value, transfers []safejs.Value) error
