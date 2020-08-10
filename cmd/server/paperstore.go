package main

import (
	"flag"
	"paperstore/internal/web"
)

var port = flag.Int("p", 8080, "port")

func main() {
	flag.Parse()
	web.Start(*port)
}
