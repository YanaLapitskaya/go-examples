package services

import (
	"go-exercises/final_project/models"
	"go-exercises/final_project/repositories"
	"log"
)

func GetGroups() []models.Group {
	groupsRows := repositories.GetAllGroups()
	defer groupsRows.Close()

	groups := make([]models.Group, 0)

	for groupsRows.Next() {
		var groupDb models.GroupDb
		if err := groupsRows.Scan(&groupDb.Id, &groupDb.Title); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}

		tasksRows := repositories.GetTasksByGroupId(groupDb.Id)
		tasks := ReadTasksFromSql(tasksRows)

		group := models.Group{Id: groupDb.Id, Title: groupDb.Title, Tasks: tasks}
		groups = append(groups, group)
	}
	return groups
}
