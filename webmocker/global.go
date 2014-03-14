package main

import (
	"errors"
	conf "github.com/funkygao/jsconf"
)

const (
	LISTEN_ADDR = ":9001"
	DEBUG_ADDR  = ":9002"
	CONFIG      = "automan.json"
	VERSION     = "0.1.2.stable"
)

var (
	cf      *conf.Conf
	apis    map[string]apiConfig // key is uri path
	verbose bool

	ErrNotFound = errors.New("uri not found")
)
