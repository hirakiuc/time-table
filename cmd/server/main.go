package main

import (
	"fmt"
	"os"

	"github.com/hirakiuc/go-time-table/api"
)

const ExitErrCode = 1

func main() {
	server := api.NewServer()

	err := server.Start()
	if err != nil {
		fmt.Println("Failed", err)
		os.Exit(ExitErrCode)
	}
}
