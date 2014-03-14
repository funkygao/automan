package main

import (
	"encoding/json"
	log "github.com/funkygao/log4go"
	"io"
	"net/http"
	"net/url"
)

func handleHttpQuery(w http.ResponseWriter, req *http.Request,
	params map[string]interface{}, api apiConfig) (interface{}, error) {
	body := make(map[string]interface{})
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&body); err != nil && err != io.EOF {
		log.Error("%v", err)
		return nil, err
	}

	uri, _ := url.Parse(req.RequestURI)
	log.Info("request: {url:%s, body:%+v}", uri.Path, body)
	log.Info("response: %+v", api.output)

	return api.output, nil
}
