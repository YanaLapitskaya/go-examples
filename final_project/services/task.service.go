package services

import (
	"database/sql"
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"log"
)

func GetTasks() []models.Task {
	rows := repositories.GetAllTasks()
	return ReadTasksFromSqlWithGroups(rows)
}

// more like helpers methods, not services methods... anyway
func ReadTasksFromSql(rows *sql.Rows) []models.TaskDb {
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

func ReadTasksFromSqlWithGroups(rows *sql.Rows) []models.Task {
	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var taskDb models.TaskDb
		if err := rows.Scan(&taskDb.Id, &taskDb.Title, &taskDb.GroupId); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		group := repositories.GetGroupByTaskId(taskDb.GroupId)
		task := models.Task{Id: taskDb.Id, Title: taskDb.Title, Group: group}
		tasks = append(tasks, task)
	}
	return tasks
}
