package main

import (
	"fmt"

	"incolore.net/stint/core"
)

func init() {
	fmt.Printf("%v\n", "AuthX init")
}

func OnEvent(eventName string, data interface{}) {

}

func About() *core.PluginInfo {
	info := &core.PluginInfo{
		Name:    "AuthX",
		Author:  "Pluveto",
		Version: "0.0.1",
	}
	return info
}
