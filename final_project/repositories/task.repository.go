package repositories

import (
	"database/sql"
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

func AddTask(task *models.Task) models.Task {
	var newTask models.Task
	configs.Db.QueryRow(
		"INSERT INTO tasks (title) VALUES ($1) RETURNING id, title", task.Title,
	).Scan(&newTask.Id, &newTask.Title)
	return newTask
}

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
