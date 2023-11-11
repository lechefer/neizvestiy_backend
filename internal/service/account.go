package service

import (
	"context"
	"strings"
)

type AccountService struct {
	accountRepo accountRepo
}

func NewAccountService(accountRepo accountRepo) *AccountService {
	return &AccountService{accountRepo: accountRepo}
}

func (as *AccountService) Create(ctx context.Context, accountId string) error {
	err := as.accountRepo.Create(ctx, accountId)
	switch {
	case err == nil:
	case strings.Contains(err.Error(), "duplicate key"):
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}
