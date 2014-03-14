package main

import (
	"flag"
	"fmt"
	mock "github.com/funkygao/fae/http"
	log "github.com/funkygao/log4go"
	"net/http"
	"os"
)

func init() {
	apis = make(map[string]apiConfig)
}

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

	for _, api := range loadConfig(CONFIG) {
		path := fmt.Sprintf("/%s/%s", api.controller, api.action)
		apis[path] = api
		mock.RegisterHttpApi(path, func(w http.ResponseWriter, req *http.Request,
			params map[string]interface{}) (interface{}, error) {
			return handleHttpQuery(w, req, params)
		}).Methods("GET", "POST")

		log.Debug("registered uri: %s", path)
	}

	select {}
}
