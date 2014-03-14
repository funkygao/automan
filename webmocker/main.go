package main

import (
	"github.com/funkygao/fae/http"
)

const (
	LISTEN_ADDR = ":9001"
	DEBUG_ADDR  = ":9002"
	CONFIG      = "automan.cf"
)

func main() {
	loadConfig(CONFIG)

	if err := http.LaunchHttpServ(LISTEN_ADDR, DEBUG_ADDR); err != nil {
		panic(err)
	}
	defer http.StopHttpServ()

	select {}
}
