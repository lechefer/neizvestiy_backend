package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"time"
)

type Quest struct {
	Id           uuid.UUID
	SettlementId uuid.UUID
	Preview      Image
	Name         string
	Description  string
	Type         QuestType
	AvgDuration  time.Duration
	Reward       decimal.Decimal

	Steps []QuestStep
}

type QuestType string

const (
	RouteQuestType QuestType = "route"
)

type QuestStep struct {
	Id      uuid.UUID
	QuestId uuid.UUID
	Order   int

	Images []Image

	Name      string
	PlaceType string
	Address   string
	Phone     string
	Email     string
	Website   string
	Schedule  []Schedule

	Latitude  float64
	Longitude float64
}

type Schedule struct {
	// WeekDay
	WeekDay WeekDay
	// FromTo format `10:00 - 18:00`
	FromTo string
}

type WeekDay string

const (
	MondayWeekDay    = "Mon"
	TuesdayWeekDay   = "Tue"
	WednesdayWeekDay = "Wed"
	ThursdayWeekDay  = "Thu"
	FridayWeekDay    = "Fri"
	SaturdayWeekDay  = "Sat"
	SundayWeekDay    = "Sun"
)

type ListQuestsOptions struct {
	SettlementId uuid.UUID
}
