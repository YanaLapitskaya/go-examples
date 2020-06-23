package models

// it's abstract model
type Task struct {
	Id         int         `json:"id"`
	Title      string      `json:"title"`
	Group      *GroupDb    `json:"group"`
	Timeframes []Timeframe `json:"timeframes"`
}

// it's how task is stored in db
type TaskDb struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	GroupId int    `json:"group_id"`
}
