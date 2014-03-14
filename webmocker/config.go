package main

import (
	conf "github.com/funkygao/jsconf"
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

}
