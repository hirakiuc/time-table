package api

import (
	"github.com/hirakiuc/time-table/keeper"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Keeper keeper.Keeper
}

func NewServer(keeper keeper.Keeper) *Server {
	return &Server{
		Keeper: keeper,
	}
}

func (s *Server) Start() error {
	r := gin.Default()
	r.GET("/schedules", s.GetSchedules)

	// listen and serve on 0.0.0.0:8080
	return r.Run()
}
