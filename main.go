package main

import (
	"flag"

	"github.com/dinislamdarkhan/simple-wallet/src/app"
)

func main() {
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address only port :8080")
	flag.Parse()

	app.Run(*httpAddr)
}
