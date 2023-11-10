package entity

import "github.com/google/uuid"

type Settlement struct {
	Id        uuid.UUID
	Name      string
	Latitude  float64
	Longitude float64
}

type ListSettlementsOptions struct {
	Query string
	Page  int
	Count int
}
