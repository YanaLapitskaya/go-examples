package models

// it's abstract model
type Group struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Tasks []TaskDb `json:"tasks"`
}

// it's how task is stored in db
type GroupDb struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
