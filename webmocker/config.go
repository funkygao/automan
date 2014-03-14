package main

import (
	conf "github.com/funkygao/jsconf"
	log "github.com/funkygao/log4go"
)

var (
	cf *conf.Conf
)

func loadConfig(fn string) {
	var err error
	cf, err = conf.Load(fn)
	if err != nil {
		panic(err)
	}

	api, _ := cf.Section("api")
	log.Debug("%+v", api)
}
