package model

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

type Schedule struct {
	ID    string    `json:"id"`
	Cron  string    `json:"cron"`
	Until time.Time `json:"until"`
}

func (s Schedule) HashValue() string {
	str := fmt.Sprintf("%s|%s",
		s.Cron,
		s.Until.Format(time.UnixDate),
	)

	v := sha512.Sum512([]byte(str))

	return hex.EncodeToString(v[:])
}

func (s Schedule) TimesInPeriod(from time.Time, until time.Time) ([]time.Time, error) {
	if from.After(until) {
		return []time.Time{}, nil
	}

	sched, err := cron.ParseStandard(s.Cron)
	if err != nil {
		return []time.Time{}, err
	}

	times := []time.Time{}
	start := from

	for {
		n := sched.Next(start)

		if n.After(until) {
			break
		}

		times = append(times, n)

		start = n
	}

	return times, nil
}
