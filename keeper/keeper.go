package keeper

import (
	"sort"
	"time"

	"github.com/hirakiuc/time-table/model"
)

type Keeper interface {
	AddSchedules(schedules []model.Schedule) error
	EventsInPeriod(from time.Time, until time.Time, order string) (*model.TimeTable, error)
}

type ScheduleKeeper struct {
	ScheduleMap map[string]model.Schedule
}

func NewScheduleKeeper() *ScheduleKeeper {
	return &ScheduleKeeper{
		ScheduleMap: map[string]model.Schedule{},
	}
}

func (k *ScheduleKeeper) AddSchedules(schedules []model.Schedule) error {
	for _, schedule := range schedules {
		k.ScheduleMap[schedule.ID] = schedule
	}

	return nil
}

func (k *ScheduleKeeper) EventsInPeriod(from time.Time, until time.Time, order string) (*model.TimeTable, error) {
	tt := model.NewTimeTable(order)

	for _, schedule := range k.ScheduleMap {
		times, err := schedule.TimesInPeriod(from, until)
		if err != nil {
			return model.NewTimeTable(order), err
		}

		tt.AddSchedule(schedule, times)
	}

	sort.Sort(tt)

	return tt, nil
}
