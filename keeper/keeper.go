package keeper

import (
	"sort"
	"time"
)

type Keeper interface {
	AddSchedules(entries *[]Schedule) error
	EventsInPeriod(from time.Time, until time.Time, order string) (*TimeTable, error)
}

type ScheduleKeeper struct {
	ScheduleMap map[string]Schedule
}

func NewScheduleKeeper() ScheduleKeeper {
	return ScheduleKeeper{
		ScheduleMap: map[string]Schedule{},
	}
}

func (k *ScheduleKeeper) AddSchedules(schedules []Schedule) error {
	for _, schedule := range schedules {
		k.ScheduleMap[schedule.ID] = schedule
	}

	return nil
}

func (k *ScheduleKeeper) EventsInPeriod(from time.Time, until time.Time, order string) (*TimeTable, error) {
	tt := NewTimeTable(order)

	for _, schedule := range k.ScheduleMap {
		times, err := schedule.TimesInPeriod(from, until)
		if err != nil {
			return NewTimeTable(order), err
		}

		tt.AddSchedule(schedule, times)
	}

	sort.Sort(tt)

	return tt, nil
}
