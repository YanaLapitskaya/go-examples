package repositories

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
	"log"
)

func AddTimeframe(timeframe *models.Timeframe) models.Timeframe {
	var newTimeframe models.Timeframe
	q := `
		INSERT INTO timeframes (task_id, "from", "to")
			VALUES ($1, $2, $3)
			RETURNING id, task_id, "from", "to"
	`
	configs.Db.QueryRow(
		q, timeframe.TaskId, timeframe.From, timeframe.To,
	).Scan(&newTimeframe.Id, &newTimeframe.TaskId, &newTimeframe.From, &newTimeframe.To)
	return newTimeframe
}

func DeleteTimeframe(id string) {
	_, err := configs.Db.Exec("DELETE FROM timeframes WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

func GetTimeframesByTaskId(taskId int, c chan []models.Timeframe) {
	rows, err := configs.Db.Query("SELECT * FROM timeframes WHERE task_id = $1", taskId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	timeframes := make([]models.Timeframe, 0)

	for rows.Next() {
		var timeframe models.Timeframe
		if err := rows.Scan(&timeframe.Id, &timeframe.TaskId, &timeframe.From, &timeframe.To); err != nil {
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		timeframes = append(timeframes, timeframe)
	}
	c <- timeframes
}
