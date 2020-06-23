package repositories

import (
	"database/sql"
	"fmt"
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
	"log"
)

func GetAllTasks() *sql.Rows {
	rows, err := configs.Db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func GetGroupByTaskId(taskId int) *models.GroupDb {
	var group models.GroupDb
	fmt.Println(taskId)
	configs.Db.QueryRow("SELECT * FROM groups WHERE id = $1", taskId).Scan(&group.Id, &group.Title)
	return &group
}

func AddTask(task *models.TaskDb) models.TaskDb {
	var newTask models.TaskDb
	configs.Db.QueryRow(
		"INSERT INTO tasks (title, group_id) VALUES ($1, $2) RETURNING id, title, group_id", task.Title, task.GroupId,
	).Scan(&newTask.Id, &newTask.Title, &newTask.GroupId)
	return newTask
}

// TODO add updates for timestamps and groups
func UpdateTask(id string, task *models.Task) models.Task {
	var updatedTask models.Task
	configs.Db.QueryRow(
		"UPDATE tasks SET title = $1 WHERE id = $2 RETURNING id, title", task.Title, id,
	).Scan(&updatedTask.Id, &updatedTask.Title)
	return updatedTask
}

func DeleteTask(id string) {
	_, err := configs.Db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
