package repositories

import (
	"go-exercises/final_project/configs"
	"go-exercises/final_project/models"
)

type TimeframesChannel struct {
	Timeframes []models.Timeframe
	Err        error
}

func AddTimeframe(timeframe *models.Timeframe) (*models.Timeframe, error) {
	var newTimeframe models.Timeframe
	q := `
		INSERT INTO timeframes (task_id, "from", "to")
			VALUES ($1, $2, $3)
			RETURNING id, task_id, "from", "to"
	`
	err := configs.Db.QueryRow(
		q, timeframe.TaskId, timeframe.From, timeframe.To,
	).Scan(&newTimeframe.Id, &newTimeframe.TaskId, &newTimeframe.From, &newTimeframe.To)
	if err != nil {
		return nil, err
	}
	return &newTimeframe, nil
}

func DeleteTimeframe(id string) error {
	_, err := configs.Db.Exec("DELETE FROM timeframes WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetTimeframesByTaskId(taskId int, c chan TimeframesChannel) {
	rows, err := configs.Db.Query("SELECT * FROM timeframes WHERE task_id = $1", taskId)
	if err != nil {
		c <- TimeframesChannel{Timeframes: nil, Err: err}
	}
	defer rows.Close()

	timeframes := make([]models.Timeframe, 0)

	for rows.Next() {
		var timeframe models.Timeframe
		if err := rows.Scan(&timeframe.Id, &timeframe.TaskId, &timeframe.From, &timeframe.To); err != nil {
			// Query rows will be closed with defer.
			c <- TimeframesChannel{Timeframes: nil, Err: err}
		}
		timeframes = append(timeframes, timeframe)
	}
	c <- TimeframesChannel{Timeframes: timeframes, Err: nil}
}
