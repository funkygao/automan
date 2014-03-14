package main

import (
	"flag"
	"fmt"
	mock "github.com/funkygao/fae/http"
	log "github.com/funkygao/log4go"
	"net/http"
)

func main() {
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.Parse()
	if !verbose {
		log.AddFilter("stdout", log.INFO, log.NewConsoleLogWriter())
	}

	cf := loadConfig(CONFIG)

	if err := mock.LaunchHttpServ(LISTEN_ADDR, DEBUG_ADDR); err != nil {
		panic(err)
	}
	defer mock.StopHttpServ()

	for _, api := range cf {
		path := fmt.Sprintf("/%s/%s", api.controller, api.action)
		mock.RegisterHttpApi(path, func(w http.ResponseWriter, req *http.Request,
			params map[string]interface{}) (interface{}, error) {
			return handleHttpQuery(w, req, params, api)
		}).Methods("GET", "POST")

		log.Debug("registered uri: %s", path)
	}

	select {}
}
