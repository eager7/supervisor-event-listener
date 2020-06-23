package main

import (
	"flag"
	"github.com/eager7/supervisor-event-listener/listener"
)

func main() {
	var key = flag.String("key", "", "wechat robot key")
	flag.Parse()
	for {
		listener.Start(*key)
	}
}
