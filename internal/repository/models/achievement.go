package models

import "github.com/google/uuid"

type Achievement struct {
	Id          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Icon        string    `db:"icon"`
	Steps       int       `db:"steps"`
	Description string    `db:"description"`
	Passed      int       `db:"passed"`
	IsCompleted bool      `db:"is_completed"`
}
