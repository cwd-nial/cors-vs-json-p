package main

import "github.com/cwd-nial/cors-vs-json-p/pkg/shutdown"

func main() {
	Server().RunWithCancelCtx(shutdown.SignalContext())
}
