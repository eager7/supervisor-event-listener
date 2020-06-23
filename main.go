package main

import (
	"github.com/eager7/supervisor-event-listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
