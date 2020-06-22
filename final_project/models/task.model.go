package models

type Task struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func NewTask(id int, title string) Task {
	return Task{Id: id, Title: title}
}

func DeleteTask(id int) int {
	return id
}

func UpdateTask(newId int, newTitle string) *Task {
	return &Task{Id: newId, Title: newTitle}
}
