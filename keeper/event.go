package keeper

import "time"

const (
	OrderInAsc  = "asc"
	OrderInDesc = "desc"
)

type Event struct {
	Schedule Schedule
	At       time.Time
}

type TimeTable struct {
	Events    []Event
	SortOrder string
}

func NewTimeTable(order string) *TimeTable {
	sortOrder := ""

	switch order {
	case OrderInAsc:
	case OrderInDesc:
		sortOrder = order
	default:
		sortOrder = OrderInAsc
	}

	return &TimeTable{
		Events:    []Event{},
		SortOrder: sortOrder,
	}
}

func (t *TimeTable) Len() int {
	return len(t.Events)
}

func (t *TimeTable) Swap(i int, j int) {
	t.Events[i], t.Events[j] = t.Events[j], t.Events[i]
}

func (t *TimeTable) Less(i int, j int) bool {
	if t.SortOrder == OrderInAsc {
		return (t.Events[i].At).Before(t.Events[j].At)
	}

	return (t.Events[i].At).After(t.Events[j].At)
}

func (t *TimeTable) AddSchedule(schedule Schedule, times []time.Time) {
	events := make([]Event, len(times))

	for idx, time := range times {
		e := Event{
			Schedule: schedule,
			At:       time,
		}

		events[idx] = e
	}

	t.Events = append(t.Events, events...)
}

func (t *TimeTable) Iterator() <-chan Event {
	ch := make(chan Event)

	go func() {
		for _, event := range t.Events {
			ch <- event
		}
	}()

	return ch
}
