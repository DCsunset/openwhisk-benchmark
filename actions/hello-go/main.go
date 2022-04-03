package main

import (
	"log"
)

func Main(obj map[string]interface{}) map[string]interface{} {
	name, ok := obj["name"].(string)
	if !ok {
		name = "world"
	}

	msg := make(map[string]interface{})
	msg["message"] = "hello, " + name + "!"
	log.Printf("name=%s\n", name)
	return msg
}
