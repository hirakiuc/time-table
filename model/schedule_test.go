package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newSchedule(id string, cron string, until string) (Schedule, error) {
	v, err := time.Parse(time.RFC3339, until)
	if err != nil {
		return Schedule{
			ID:    "",
			Cron:  "",
			Until: time.Now(),
		}, err
	}

	return Schedule{
		ID:    id,
		Cron:  cron,
		Until: v,
	}, nil
}

func TestScheduleHashValue(t *testing.T) {
	assert := assert.New(t)

	schedule, err := newSchedule(
		"12345",
		"5 4 * * *",
		"2020-03-20T10:00:00+09:00",
	)
	assert.Nil(err)

	actual := schedule.HashValue()
	assert.IsType(actual, "string", "Scheddule.HashValue should return a string")
}

// nolint:funlen
func TestScheduleTimesInPeriod(t *testing.T) {
	assert := assert.New(t)

	testcases := []struct {
		Cron     string
		Until    string
		Expected []string
	}{
		{
			Cron:  "5 4 * * *",
			Until: "2020-03-11T10:00:00+09:00",
			Expected: []string{
				"2020-03-01T04:05:00+09:00",
				"2020-03-02T04:05:00+09:00",
				"2020-03-03T04:05:00+09:00",
				"2020-03-04T04:05:00+09:00",
				"2020-03-05T04:05:00+09:00",
				"2020-03-06T04:05:00+09:00",
				"2020-03-07T04:05:00+09:00",
				"2020-03-08T04:05:00+09:00",
				"2020-03-09T04:05:00+09:00",
				"2020-03-10T04:05:00+09:00",
			},
		},
		{
			Cron:  "10 22 */2 * *",
			Until: "2020-03-11T10:00:00+09:00",
			Expected: []string{
				"2020-03-01T22:10:00+09:00",
				"2020-03-03T22:10:00+09:00",
				"2020-03-05T22:10:00+09:00",
				"2020-03-07T22:10:00+09:00",
				"2020-03-09T22:10:00+09:00",
			},
		},
	}

	var from, until time.Time

	var err error

	from, err = time.Parse(time.RFC3339, "2020-03-01T00:00:00+09:00")
	assert.Nil(err)

	until, err = time.Parse(time.RFC3339, "2020-03-10T23:59:59+09:00")
	assert.Nil(err)

	for _, testcase := range testcases {
		var schedule Schedule
		schedule, err = newSchedule(
			"12345",
			testcase.Cron,
			testcase.Until,
		)
		assert.Nil(err)

		var times []time.Time
		times, err = schedule.TimesInPeriod(from, until)
		assert.Nil(err)
		assert.Len(times, len(testcase.Expected))

		strs := []string{}
		for _, v := range times {
			strs = append(strs, v.Format(time.RFC3339))
		}

		assert.ElementsMatch(testcase.Expected, strs)
	}
}

func TestScheduleTimesInPeriodWithInvalidArgument(t *testing.T) {
	assert := assert.New(t)

	testcases := []struct {
		From       string
		Until      string
		Cron       string
		Expected   []time.Time
		ErrMessage string
	}{
		{
			From:       "2020-03-02T00:00:00+09:00",
			Until:      "2019-03-02T00:00:00+09:00",
			Cron:       "5 4 * * *",
			Expected:   []time.Time{},
			ErrMessage: "",
		},
		{
			From:       "2020-03-02T00:00:00+09:00",
			Until:      "2020-03-10T23:59:59+09:00",
			Cron:       "invalid cron syntax",
			Expected:   []time.Time{},
			ErrMessage: "expected exactly 5 fields, found 3: [invalid cron syntax]",
		},
	}

	for _, testcase := range testcases {
		var from, until time.Time

		var err error

		from, err = time.Parse(time.RFC3339, testcase.From)
		assert.Nil(err)

		until, err = time.Parse(time.RFC3339, testcase.Until)
		assert.Nil(err)

		var schedule Schedule
		schedule, err = newSchedule(
			"12345",
			testcase.Cron,
			testcase.Until,
		)
		assert.Nil(err)

		var times []time.Time
		times, err = schedule.TimesInPeriod(from, until)
		assert.Equal(testcase.Expected, times)

		if len(testcase.ErrMessage) == 0 {
			assert.Nil(err)
		} else {
			assert.Equal(testcase.ErrMessage, err.Error())
		}
	}
}
