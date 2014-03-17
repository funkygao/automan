package main

import (
	"flag"
	"fmt"
	mock "github.com/funkygao/fae/http"
	"github.com/funkygao/golib/signal"
	log "github.com/funkygao/log4go"
	"os"
	"syscall"
)

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [ver: %s]\n", os.Args[0], VERSION)
		flag.PrintDefaults()
	}
	flag.Parse()
	if !verbose {
		log.AddFilter("stdout", log.INFO, log.NewConsoleLogWriter())
	}

	if err := mock.LaunchHttpServ(LISTEN_ADDR, DEBUG_ADDR); err != nil {
		panic(err)
	}
	defer mock.StopHttpServ()

	signal.RegisterSignalHandler(syscall.SIGHUP, registerHttpApi)
	signal.Kill(syscall.SIGHUP)

	select {}
}
