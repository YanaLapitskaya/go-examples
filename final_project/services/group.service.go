package services

import (
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
)

func GetGroups() ([]models.Group, error) {
	groupsDb, err := repositories.GetAllGroups()
	if err != nil {
		return nil, err
	}
	groups := make([]models.Group, 0)

	for _, groupDb := range groupsDb {
		tasks, err := repositories.GetTasksByGroupId(groupDb.Id)
		if err != nil {
			return nil, err
		}
		group := models.Group{Id: groupDb.Id, Title: groupDb.Title, Tasks: tasks}
		groups = append(groups, group)
	}
	return groups, nil
}
