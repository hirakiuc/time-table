package model

import (
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewTimeTable(t *testing.T) {
	assert := assert.New(t)

	testcases := []struct {
		Order    string
		Expected string
	}{
		{
			Order:    OrderInAsc,
			Expected: OrderInAsc,
		},
		{
			Order:    OrderInDesc,
			Expected: OrderInDesc,
		},
		{
			Order:    "unexpected",
			Expected: OrderInAsc,
		},
	}

	for _, testcase := range testcases {
		tt := NewTimeTable(testcase.Order)

		assert.NotNil(tt)
		assert.Equal(testcase.Expected, tt.SortOrder)
	}
}

func TestAddSchedule(t *testing.T) {
	assert := assert.New(t)

	tt := NewTimeTable(OrderInAsc)

	now, err := time.Parse(time.RFC3339, "2020-03-10T10:00:00+09:00")
	assert.Nil(err)

	schedule := Schedule{
		ID:    "12345",
		Cron:  "5 4 * * *",
		Until: now.Add(24 * 10 * time.Hour),
	}

	// nolint:gomnd
	times := []time.Time{
		now.Add(1 * time.Hour),
		now.Add(5 * time.Hour),
		now.Add(10 * time.Hour),
	}

	tt.AddSchedule(schedule, times)
	assert.Len(tt.Events, 3)

	for i, time := range times {
		event := tt.Events[i]

		assert.Equal(schedule, event.Schedule)
		assert.Equal(time, event.At)
	}
}

func TestTimeTableSort(t *testing.T) {
	assert := assert.New(t)

	tt := NewTimeTable(OrderInAsc)

	now, err := time.Parse(time.RFC3339, "2020-03-10T10:00:00+09:00")
	assert.Nil(err)

	schedule := Schedule{
		ID:    "12345",
		Cron:  "5 4 * * *",
		Until: now.Add(24 * 10 * time.Hour),
	}

	times := []time.Time{
		now.Add(1 * time.Hour),
		now.Add(5 * time.Hour),
		now.Add(10 * time.Hour),
	}

	tt.AddSchedule(schedule, times)
	assert.Len(tt.Events, 3)

	sort.Sort(tt)

	// Asc
	for i, event := range tt.Events {
		assert.Equal(schedule, event.Schedule)
		assert.Equal(times[i], event.At)
	}

	// Desc
	tt.SortOrder = OrderInDesc
	sort.Sort(tt)

	sort.Slice(times, func(i, j int) bool {
		return times[i].After(times[j])
	})

	for i, event := range tt.Events {
		assert.Equal(schedule, event.Schedule)
		assert.Equal(times[i], event.At)
	}
}

func TestTimeTableIterator(t *testing.T) {
	assert := assert.New(t)

	tt := NewTimeTable(OrderInAsc)
	now, err := time.Parse(time.RFC3339, "2020-03-10T10:00:00+09:00")
	assert.Nil(err)

	schedule := Schedule{
		ID:    "12345",
		Cron:  "5 4 * * *",
		Until: now.Add(24 * 10 * time.Hour),
	}

	times := []time.Time{
		now.Add(1 * time.Hour),
		now.Add(5 * time.Hour),
		now.Add(10 * time.Hour),
	}

	tt.AddSchedule(schedule, times)
	assert.Len(tt.Events, 3)

	var i int = 0

	for event := range tt.Iterator() {
		assert.Equal(schedule, event.Schedule)
		assert.Equal(times[i], event.At)

		i += 1
	}
}
