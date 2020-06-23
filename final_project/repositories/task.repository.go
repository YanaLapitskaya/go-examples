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

func GetTasksByGroupId(groupId int) []models.TaskDb {
	rows, err := configs.Db.Query("SELECT * FROM tasks where group_id = $1", groupId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	tasks := make([]models.TaskDb, 0)

	for rows.Next() {
		var taskDb models.TaskDb
		if err := rows.Scan(&taskDb.Id, &taskDb.Title, &taskDb.GroupId); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		tasks = append(tasks, taskDb)
	}
	return tasks
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
