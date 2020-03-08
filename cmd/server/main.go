package main

import (
	"github.com/hirakiuc/go-time-table/api"
)

func main() {
	server := api.NewApiServer()
	server.Start()
}
