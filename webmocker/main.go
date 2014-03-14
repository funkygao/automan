package main

import (
	"fmt"
	mock "github.com/funkygao/fae/http"
	log "github.com/funkygao/log4go"
	"net/http"
)

const (
	LISTEN_ADDR = ":9001"
	DEBUG_ADDR  = ":9002"
	CONFIG      = "automan.json"
)

func main() {
	cf := loadConfig(CONFIG)

	if err := mock.LaunchHttpServ(LISTEN_ADDR, DEBUG_ADDR); err != nil {
		panic(err)
	}
	defer mock.StopHttpServ()

	for _, api := range cf {
		path := fmt.Sprintf("/%s/%s", api.controller, api.action)
		mock.RegisterHttpApi(path, func(w http.ResponseWriter, req *http.Request,
			params map[string]interface{}) (interface{}, error) {
			return handleHttpQuery(w, req, params)
		}).Methods("GET", "POST")

		log.Debug("uri: %s", path)
	}

	select {}
}
