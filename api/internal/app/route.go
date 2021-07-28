package app

import (
	"context"

	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

func Route(r *mux.Router, ctx context.Context, root Root) error {
	app, err := NewApp(ctx, root)
	if err != nil {
		return err
	}

	r.HandleFunc("/health", app.HealthHandler.Check).Methods(GET)

	pollPath := "/poll"
	r.HandleFunc(pollPath, app.PollHandler.GetAll).Methods(GET)
	r.HandleFunc(pollPath+"/{id}", app.PollHandler.Load).Methods(GET)
	r.HandleFunc(pollPath, app.PollHandler.Insert).Methods(POST)
	r.HandleFunc(pollPath+"/{id}", app.PollHandler.Update).Methods(PUT)
	r.HandleFunc(pollPath+"/{id}", app.PollHandler.Patch).Methods(PATCH)
	r.HandleFunc(pollPath+"/{id}", app.PollHandler.Delete).Methods(DELETE)

	return nil
}
