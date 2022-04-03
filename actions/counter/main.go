package main

import (
	"net/http"
)

func Main(obj map[string]interface{}) map[string]interface{} {
	// process input params
	result := make(map[string]interface{})

	url := obj["url"].(string)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		result["error"] = err.Error()
		return result
	}
	req.Header.Set("Connection", "close")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		result["error"] = err.Error()
		return result
	}
	defer resp.Body.Close()
	result["resp"] = "OK"

	return result
}
