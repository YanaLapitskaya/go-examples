package models

import "time"

// here abstract and db models the same
type Timeframe struct {
	Id     int       `json:"id"`
	TaskId int       `json:"task_id"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}
