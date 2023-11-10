package entity

import "github.com/google/uuid"

type Achievement struct {
	Id          uuid.UUID
	Name        string
	Icon        string
	Steps       int
	Description string
	Passed      int
	IsCompleted bool
}

type ListAchievementsOptions struct {
	AccountId string
}
