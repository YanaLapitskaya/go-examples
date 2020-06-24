package repositories

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
)

func GetAllTasks() ([]models.TaskDb, error) {
	rows, err := configs.Db.Query("SELECT * FROM tasks")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	tasksDb := make([]models.TaskDb, 0)

	for rows.Next() {
		var taskDb models.TaskDb
		if err := rows.Scan(&taskDb.Id, &taskDb.Title, &taskDb.GroupId); err != nil {
			// Query rows will be closed with defer.
			return nil, err
		}
		tasksDb = append(tasksDb, taskDb)
	}
	return tasksDb, nil
}

func GetTasksByGroupId(groupId int) ([]models.TaskDb, error) {
	rows, err := configs.Db.Query("SELECT * FROM tasks where group_id = $1", groupId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	tasksDb := make([]models.TaskDb, 0)

	for rows.Next() {
		var taskDb models.TaskDb
		if err := rows.Scan(&taskDb.Id, &taskDb.Title, &taskDb.GroupId); err != nil {
			// Query rows will be closed with defer.
			return nil, err
		}
		tasksDb = append(tasksDb, taskDb)
	}
	return tasksDb, nil
}

func AddTask(task *models.TaskDb) (*models.TaskDb, error) {
	var newTask models.TaskDb
	err := configs.Db.QueryRow(
		"INSERT INTO tasks (title, group_id) VALUES ($1, $2) RETURNING id, title, group_id", task.Title, task.GroupId,
	).Scan(&newTask.Id, &newTask.Title, &newTask.GroupId)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

// TODO add updates for timestamps and groups
func UpdateTask(id string, task *models.Task) (*models.Task, error) {
	var updatedTask models.Task
	err := configs.Db.QueryRow(
		"UPDATE tasks SET title = $1 WHERE id = $2 RETURNING id, title", task.Title, id,
	).Scan(&updatedTask.Id, &updatedTask.Title)
	if err != nil {
		return nil, err
	}
	return &updatedTask, nil
}

func DeleteTask(id string) error {
	_, err := configs.Db.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
