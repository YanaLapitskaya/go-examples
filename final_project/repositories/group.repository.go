package repositories

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
)

type GroupChannel struct {
	Group *models.GroupDb
	Err   error
}

func GetAllGroups() ([]models.GroupDb, error) {
	rows, err := configs.Db.Query("SELECT * FROM groups")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	groups := make([]models.GroupDb, 0)

	for rows.Next() {
		var groupDb models.GroupDb
		if err := rows.Scan(&groupDb.Id, &groupDb.Title); err != nil {
			// Query rows will be closed with defer.
			return nil, err
		}

		groups = append(groups, groupDb)
	}
	return groups, nil
}

func GetGroupByGroupId(groupId int, c chan GroupChannel) {
	var group models.GroupDb
	err := configs.Db.QueryRow("SELECT * FROM groups WHERE id = $1", groupId).Scan(&group.Id, &group.Title)
	if err != nil {
		c <- GroupChannel{Group: nil, Err: err}
	}
	c <- GroupChannel{Group: &group, Err: nil}
}

func AddGroup(group *models.GroupDb) (*models.GroupDb, error) {
	var newGroup models.GroupDb
	err := configs.Db.QueryRow(
		"INSERT INTO groups (title) VALUES ($1) RETURNING id, title", group.Title,
	).Scan(&newGroup.Id, &newGroup.Title)
	if err != nil {
		return nil, err
	}
	return &newGroup, nil
}

func UpdateGroup(id string, group *models.GroupDb) (*models.GroupDb, error) {
	var updatedGroup models.GroupDb
	err := configs.Db.QueryRow(
		"UPDATE groups SET title = $1 WHERE id = $2 RETURNING id, title", &group.Title, id,
	).Scan(&updatedGroup.Id, &updatedGroup.Title)
	if err != nil {
		return nil, err
	}
	return &updatedGroup, nil
}

func DeleteGroup(id string) error {
	_, err := configs.Db.Exec("DELETE FROM groups WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
