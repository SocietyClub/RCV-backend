package app

import (
	"context"

	firebase "firebase.google.com/go"
	"github.com/core-go/health"
	"github.com/core-go/health/firestore"
	"google.golang.org/api/option"

	"go-service/internal/handlers"
	"go-service/internal/services"
)

type ApplicationContext struct {
	HealthHandler *health.Handler
	PollHandler   *handlers.PollHandler
}

func NewApp(ctx context.Context, root Root) (*ApplicationContext, error) {
	opts := option.WithCredentialsJSON([]byte(root.Credentials))
	app, er1 := firebase.NewApp(ctx, nil, opts)
	if er1 != nil {
		return nil, er1
	}

	client, er2 := app.Firestore(ctx)
	if er2 != nil {
		return nil, er2
	}

	pollService := services.NewPollService(client)
	pollHandler := handlers.NewPollHandler(pollService)

	firestoreChecker := firestore.NewHealthChecker(ctx, []byte(root.Credentials))
	healthHandler := health.NewHandler(firestoreChecker)

	return &ApplicationContext{
		HealthHandler: healthHandler,
		PollHandler:   pollHandler,
	}, nil
}
