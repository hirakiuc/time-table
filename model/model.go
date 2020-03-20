package model

import "time"

const (
	OrderInAsc  = "asc"
	OrderInDesc = "desc"
)

type Event struct {
	Schedule Schedule  `json:"schedule"`
	At       time.Time `json:"at"`
}
