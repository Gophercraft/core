package models

import "time"

type SchedTrigger struct {
}

type SchedEvent struct {
	ID       string
	Trigger  string
	Duration time.Duration
}
