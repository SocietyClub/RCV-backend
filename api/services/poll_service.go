package services

import (
	"context"

	. "go-service/models"
)

type PollService interface {
	GetAll(ctx context.Context) (*[]Poll, error)
	Load(ctx context.Context, pollid string) (*Poll, error)
	Insert(ctx context.Context, poll *Poll) (int64, error)
	Update(ctx context.Context, poll *Poll) (int64, error)
	Patch(ctx context.Context, poll map[string]interface{}) (int64, error)
	Delete(ctx context.Context, pollid string) (int64, error)
}
