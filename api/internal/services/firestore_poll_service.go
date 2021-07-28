package services

import (
	"context"
	"reflect"

	"cloud.google.com/go/firestore"
	fs "github.com/core-go/firestore"
	"google.golang.org/api/iterator"

	. "go-service/internal/models"
)

type FirestorePollService struct {
	Collection *firestore.CollectionRef
}

func NewPollService(client *firestore.Client) *FirestorePollService {
	collection := client.Collection("polls")
	return &FirestorePollService{Collection: collection}
}

func (s *FirestorePollService) GetAll(ctx context.Context) (*[]Poll, error) {
	iter := s.Collection.Documents(ctx)
	var polls []Poll
	for {
		doc, er1 := iter.Next()
		if er1 == iterator.Done {
			break
		}
		if er1 != nil {
			return nil, er1
		}
		var poll Poll
		er2 := doc.DataTo(&poll)
		if er2 != nil {
			return &polls, er2
		}

		poll.Id = doc.Ref.ID
		poll.StartDate = &doc.CreateTime
		poll.EndDate = &doc.UpdateTime
		polls = append(polls, poll)
	}
	return &polls, nil
}

func (s *FirestorePollService) Load(ctx context.Context, id string) (*Poll, error) {
	doc, er1 := s.Collection.Doc(id).Get(ctx)
	var poll Poll
	if er1 != nil {
		return nil, er1
	}
	er2 := doc.DataTo(&poll)
	if er2 == nil {
		poll.Id = id
		poll.StartDate = &doc.CreateTime
		poll.EndDate = &doc.UpdateTime
	}
	return &poll, er2
}

func (s *FirestorePollService) Insert(ctx context.Context, poll *Poll) (string, error) {
	_, err := s.Collection.Doc(poll.Id).Set(ctx, poll)
	if err != nil {
		return "Failure", err
	}
	return "Success", nil
}

func (s *FirestorePollService) Update(ctx context.Context, poll *Poll) (string, error) {
	_, err := s.Collection.Doc(poll.Id).Set(ctx, poll)
	if err != nil {
		return "Failure", err
	}
	return "Success", nil
}

func (s *FirestorePollService) Patch(ctx context.Context, json map[string]interface{}) (string, error) {
	pollType := reflect.TypeOf(Poll{})
	maps := fs.MakeFirestoreMap(pollType)

	uid := json["id"]
	id := uid.(string)
	docRef := s.Collection.Doc(id)
	doc, err1 := docRef.Get(ctx)
	if err1 != nil {
		return "Failure", err1
	}
	delete(json, "id")

	dest := fs.MapToFirestore(json, doc, maps)
	_, err := docRef.Set(ctx, dest)
	if err != nil {
		return "Failure", err
	}
	return "Success", nil
}

func (s *FirestorePollService) Delete(ctx context.Context, id string) (string, error) {
	_, err := s.Collection.Doc(id).Delete(ctx)
	if err != nil {
		return "Failure", err
	}
	return "Success", nil
}
