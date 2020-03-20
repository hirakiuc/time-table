package main

import (
	"fmt"
	"os"

	"github.com/hirakiuc/time-table/api"
	"github.com/hirakiuc/time-table/keeper"
)

const ExitErrCode = 1

func main() {
	keeper := keeper.NewScheduleKeeper()
	server := api.NewServer(keeper)

	err := server.Start()
	if err != nil {
		fmt.Println("Failed", err)
		os.Exit(ExitErrCode)
	}
}
