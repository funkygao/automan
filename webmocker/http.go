package main

import (
	"encoding/json"
	log "github.com/funkygao/log4go"
	"io"
	"net/http"
	"net/url"
)

func handleHttpQuery(w http.ResponseWriter, req *http.Request,
	params map[string]interface{}) (interface{}, error) {
	body := make(map[string]interface{})
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&body); err != nil && err != io.EOF {
		log.Error("%v", err)
		return nil, err
	}

	uri, _ := url.Parse(req.RequestURI)
	log.Info("request: %s {url:%s, body:%+v}", req.Method, uri.Path, body)
	api, ok := apis[uri.Path]
	if !ok {
		return nil, ErrNotFound
	}

	log.Info("response: %+v", api.output)

	return api.output, nil
}
