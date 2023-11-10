package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type Quest struct {
	Id           uuid.UUID       `db:"id"`
	SettlementId uuid.UUID       `db:"settlement_id"`
	Name         string          `db:"name"`
	Description  string          `db:"description"`
	Type         string          `db:"type"`
	AvgDuration  time.Duration   `db:"avg_duration"`
	Reward       decimal.Decimal `db:"reward"`

	Steps []QuestStep
}

type QuestStep struct {
	Id      uuid.UUID `db:"id"`
	QuestId uuid.UUID `db:"quest_id"`
	Order   int       `db:"order"`

	Name      string `db:"name"`
	PlaceType string `db:"place_type"`
	Address   string `db:"address"`
	Phone     string `db:"phone"`
	Email     string `db:"email"`
	Website   string `db:"website"`
	Schedule  string `db:"schedule"`

	Location pgtype.Point `db:"location"`
}
