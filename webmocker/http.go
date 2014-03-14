package main

import (
	"net/http"
)

func handleHttpQuery(w http.ResponseWriter, req *http.Request,
	params map[string]interface{}) (interface{}, error) {
	output := make(map[string]interface{})
	return output, nil
}
