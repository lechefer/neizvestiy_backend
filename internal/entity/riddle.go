package entity

import "github.com/google/uuid"

type Riddle struct {
	Id          uuid.UUID
	QuestStepId uuid.UUID
	Name        string
	Description string
	Status      string
	Letters     string
}

type ListRiddlesOptions struct {
	QuestStepId uuid.UUID
}
