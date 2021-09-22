package openapi

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func GetFirestoreClient(ctx context.Context) *firestore.Client {
	// Todo: Investigate the impact of new connections on every API call.
	// Sets your Google Cloud Platform project ID.
	projectID := "societyclub-rcv-backend"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}
