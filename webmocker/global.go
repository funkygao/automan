package main

import (
	conf "github.com/funkygao/jsconf"
)

const (
	LISTEN_ADDR = ":9001"
	DEBUG_ADDR  = ":9002"
	CONFIG      = "automan.json"
)

var (
	cf      *conf.Conf
	apis    map[string]apiConfig // key is uri path
	verbose bool
)
