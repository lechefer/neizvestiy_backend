package models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Settlement struct {
	Id       uuid.UUID    `db:"id"`
	Name     string       `db:"name"`
	Location pgtype.Point `db:"location"`
}
