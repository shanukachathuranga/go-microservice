package account

import (
	"context"
	"github.com/segmentio/ksuid"
)

type Service interface {
	PostAccount(ctx context.Context, name string) (*Account, error)
	GetAccount(ctx context.Context, id string) (*Account, error)
	GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error)
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type accountService struct {
	repository Repository
}

// constructor of the accountService
func NewService(r Repository) Service {
	return &accountService{r}
}

func (s *accountService) PostAccount(ctx context.Context, name string) (*Account, error) {
	a := &Account{
		Name: name,
		ID:   ksuid.New().String(),
	}
	err := s.repository.PutAccount(ctx, *a)
	if err != nil {
		return nil, err
	}
	return a, err

}

func (s *accountService) GetAccount(ctx context.Context, id string) (*Account, error) {
	return s.repository.GetAccountById(ctx, id)
}

func (s *accountService) GetAccounts(ctx context.Context, skip uint64, take uint64) ([]Account, error) {
	return s.repository.ListAccounts(ctx, skip, take)

}
