package models

type Group struct {
	Id    int      `json:"id"`
	Title string   `json:"title"`
	Tasks []TaskDb `json:"tasks"`
}

type GroupDb struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
