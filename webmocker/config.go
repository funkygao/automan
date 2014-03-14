package main

import (
	"fmt"
	conf "github.com/funkygao/jsconf"
	log "github.com/funkygao/log4go"
)

var (
	cf *conf.Conf
)

type apiConfig struct {
	controller string
	action     string
	input      map[string]interface{}
	output     map[string]interface{}
}

func (this *apiConfig) loadConfig(section *conf.Conf) {
	this.controller = section.String("controller", "")
	this.action = section.String("action", "")
	io, err := section.Section("io")
	if err != nil {
		panic(err)
	}
	this.input = io.Object("input", nil)
	this.output = io.Object("output", nil)
}

func loadConfig(fn string) (apis []*apiConfig) {
	var err error
	cf, err = conf.Load(fn)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(cf.List("api", nil)); i++ {
		section, err := cf.Section(fmt.Sprintf("api[%d]", i))
		if err != nil {
			panic(err)
		}

		apiInfo := new(apiConfig)
		apiInfo.loadConfig(section)
		apis = append(apis, apiInfo)

		log.Debug("api: %+v", *apiInfo)
	}

	return
}
