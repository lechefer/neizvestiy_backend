package models

import "github.com/google/uuid"

type Riddle struct {
	Id          uuid.UUID `db:"id"`
	QuestStepId uuid.UUID `db:"quest_step_id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	Letter      string    `db:"letter"`
}
