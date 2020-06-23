package repositories

import (
	"database/sql"
	"fmt"
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
	"log"
)

func GetAllGroups() *sql.Rows {
	rows, err := configs.Db.Query("SELECT * FROM groups")
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

func AddGroup(group *models.GroupDb) models.GroupDb {
	var newGroup models.GroupDb
	configs.Db.QueryRow(
		"INSERT INTO groups (title) VALUES ($1) RETURNING id, title", group.Title,
	).Scan(&newGroup.Id, &newGroup.Title)
	return newGroup
}

func UpdateGroup(id string, group *models.GroupDb) models.GroupDb {
	var updatedGroup models.GroupDb
	configs.Db.QueryRow(
		"UPDATE groups SET title = $1 WHERE id = $2 RETURNING id, title", &group.Title, id,
	).Scan(&updatedGroup.Id, &updatedGroup.Title)
	return updatedGroup
}

func DeleteGroup(id string) {
	_, err := configs.Db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}
