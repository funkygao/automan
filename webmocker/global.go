package main

import (
	"errors"
)

const (
	LISTEN_ADDR = ":9001"
	DEBUG_ADDR  = ":9002"
	CONFIG      = "automan.json"
	VERSION     = "0.1.2.stable"
)

var (
	apis    map[string]apiConfig // key is uri path
	verbose bool

	ErrNotFound = errors.New("uri not found")
)
