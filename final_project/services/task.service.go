package services

import (
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"log"
)

func GetTasks() []models.Task {
	rows := repositories.GetAllTasks()
	defer rows.Close()

	tasks := make([]models.Task, 0)

	for rows.Next() {
		var taskDb models.TaskDb
		if err := rows.Scan(&taskDb.Id, &taskDb.Title, &taskDb.GroupId); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		groupChan := make(chan *models.GroupDb)
		timeframesChan := make(chan []models.Timeframe)
		go repositories.GetGroupByGroupId(taskDb.GroupId, groupChan)
		go repositories.GetTimeframesByTaskId(taskDb.Id, timeframesChan)

		group, timeframes := <-groupChan, <-timeframesChan

		task := models.Task{Id: taskDb.Id, Title: taskDb.Title, Group: group, Timeframes: timeframes}
		tasks = append(tasks, task)
	}
	return tasks
}
