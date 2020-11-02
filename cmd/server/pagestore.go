package main

import (
	"flag"
	"pagestore/internal/web"
)

var port = flag.Int("p", 8080, "port")

func main() {
	flag.Parse()
	web.Start(*port)
}
