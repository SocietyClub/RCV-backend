package services

import (
	"context"

	. "go-service/internal/models"
)

type PollService interface {
	GetAll(ctx context.Context) (*[]Poll, error)
	Load(ctx context.Context, id string) (*Poll, error)
	Insert(ctx context.Context, poll *Poll) (string, error)
	Update(ctx context.Context, poll *Poll) (string, error)
	Patch(ctx context.Context, poll map[string]interface{}) (string, error)
	Delete(ctx context.Context, id string) (string, error)
}
