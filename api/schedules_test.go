package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hirakiuc/time-table/keeper"
	"github.com/hirakiuc/time-table/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestSchedules(t *testing.T) {
	assert := assert.New(t)

	gin.SetMode(gin.TestMode)

	keeper := keeper.NewScheduleKeeper()
	server := NewServer(keeper)

	w := httptest.NewRecorder()
	_, r := gin.CreateTestContext(w)

	r.GET("/schedules", server.GetSchedules)

	req, _ := http.NewRequest("GET", "/schedules", nil)
	r.ServeHTTP(w, req)

	assert.Equal(http.StatusOK, w.Code)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(err)

	res := new(GetSchedulesResponse)
	err = json.Unmarshal(data, res)
	assert.Nil(err)

	assert.Equal(model.OrderInAsc, res.Params.Sort)
	assert.Equal([]model.Event{}, res.Events)
}
