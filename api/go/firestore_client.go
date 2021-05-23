package openapi

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func CreateFirestoreClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "societyclub_rcb_backend"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
