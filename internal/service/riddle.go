package service

import (
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"smolathon/internal/entity"
	"strings"
)

type RiddleService struct {
	riddleRepo riddleRepo
}

func NewRiddleService(riddleRepo riddleRepo) *RiddleService {
	return &RiddleService{
		riddleRepo: riddleRepo,
	}
}

func (rs *RiddleService) List(ctx context.Context, options entity.ListRiddlesOptions) ([]entity.Riddle, error) {
	riddles, err := rs.riddleRepo.ListRiddles(ctx, options)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	default:
		return nil, err
	}

	return riddles, nil
}

func (rs *RiddleService) Update(ctx context.Context, accountId string, riddleId uuid.UUID) (entity.Riddle, error) {
	err := rs.riddleRepo.UpdateRiddle(ctx, accountId, riddleId)
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "duplicate key"):
		return entity.Riddle{}, ErrAlreadyExists
	default:
		return entity.Riddle{}, err
	}

	riddle, err := rs.riddleRepo.GetRiddle(ctx, accountId, riddleId)
	switch {
	case err == nil:
	case errors.Is(err, sql.ErrNoRows):
		return entity.Riddle{}, ErrNotFound
	default:
		return entity.Riddle{}, err
	}

	return riddle, nil
}
