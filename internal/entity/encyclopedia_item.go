package entity

import "github.com/google/uuid"

type EncyclopediaItem struct {
	Id           uuid.UUID
	SettlementId uuid.UUID
	Title        string
	Description  string
}
