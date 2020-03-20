package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/hirakiuc/go-time-table/keeper"

	"github.com/gin-gonic/gin"
)

const OneDay = 24 * time.Hour
const TimeLayout = time.RFC3339

type GetSchedulesParams struct {
	From  time.Time `json:"from"`
	Until time.Time `json:"until"`
	Sort  string    `json:"sort"`
}

type GetSchedulesResponse struct {
	Params GetSchedulesParams `json:"params"`
	Events []keeper.Event     `json:"events"`
}

func (s *Server) parseGetSchedulesParams(ctx *gin.Context) *GetSchedulesParams {
	params := GetSchedulesParams{}

	var err error

	params.From, err = time.Parse(TimeLayout, ctx.Query("from"))
	if err != nil {
		// Set default from
		params.From = time.Now()
	}

	var until time.Time

	until, err = time.Parse(TimeLayout, ctx.Query("until"))
	if err != nil {
		until = (params.From).Add(OneDay)
	}

	if until.Before(params.From) {
		until = (params.From).Add(OneDay)
	}

	params.Until = until

	// Parse sort param
	if strings.EqualFold(ctx.Query("sort"), keeper.OrderInDesc) {
		params.Sort = keeper.OrderInDesc
	} else {
		params.Sort = keeper.OrderInAsc
	}

	return &params
}

func (s *Server) GetSchedules(ctx *gin.Context) {
	params := s.parseGetSchedulesParams(ctx)

	timeTable, err := (s.Keeper).EventsInPeriod(params.From, params.Until, params.Sort)
	if err != nil {
		ResponseError(ctx, http.StatusInternalServerError, err)
		return
	}

	events := []keeper.Event{}
	for event := range timeTable.Iterator() {
		events = append(events, event)
	}

	ctx.JSON(http.StatusOK, events)
}

func ResponseError(ctx *gin.Context, statusCode int, err error) {
	ctx.JSON(statusCode, gin.H{"message": err.Error()})
}
