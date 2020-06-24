package services

import (
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
)

func GetTasks() ([]models.Task, error) {
	tasksDb, err := repositories.GetAllTasks()
	if err != nil {
		return nil, err
	}

	tasks := make([]models.Task, 0)

	for _, taskDb := range tasksDb {
		groupChan := make(chan repositories.GroupChannel)
		timeframesChan := make(chan repositories.TimeframesChannel)
		go repositories.GetGroupByGroupId(taskDb.GroupId, groupChan)
		go repositories.GetTimeframesByTaskId(taskDb.Id, timeframesChan)

		groupResult, timeframesResult := <-groupChan, <-timeframesChan

		if groupResult.Err != nil {
			return nil, groupResult.Err
		}

		if timeframesResult.Err != nil {
			return nil, timeframesResult.Err
		}

		task := models.Task{Id: taskDb.Id, Title: taskDb.Title, Group: groupResult.Group, Timeframes: timeframesResult.Timeframes}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
