package main

import (
	"encoding/json"
	"fmt"
	mock "github.com/funkygao/fae/http"
	log "github.com/funkygao/log4go"
	"io"
	"net/http"
	"net/url"
	"os"
)

func registerHttpApi(sig os.Signal) {
	log.Info("got signal %s, reloading...", sig)

	// initialize
	apis = make(map[string]apiConfig)
	mock.UnregisterAllHttpApi()

	for _, api := range loadConfig(CONFIG) {
		path := fmt.Sprintf("/%s/%s", api.controller, api.action)
		apis[path] = api // register into all api's
		mock.RegisterHttpApi(path, func(w http.ResponseWriter, req *http.Request,
			params map[string]interface{}) (interface{}, error) {
			return handleHttpQuery(w, req, params)
		}).Methods("GET", "POST")

		log.Debug("registered uri: %s", path)
	}
}

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
