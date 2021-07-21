package services

import (
	"context"
	"reflect"

	"cloud.google.com/go/firestore"
	fs "github.com/core-go/firestore"
	"google.golang.org/api/iterator"

	. "go-service/models"
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

		poll.PollId = doc.Ref.ID
		polls = append(polls, poll)
	}
	return &polls, nil
}

func (s *FirestorePollService) Load(ctx context.Context, pollid string) (*Poll, error) {
	doc, er1 := s.Collection.Doc(pollid).Get(ctx)
	var poll Poll
	if er1 != nil {
		return nil, er1
	}
	er2 := doc.DataTo(&poll)
	if er2 == nil {
		poll.PollId = pollid
	}
	return &poll, er2
}

func (s *FirestorePollService) Insert(ctx context.Context, poll *Poll) (int64, error) {
	_, err := s.Collection.Doc(poll.PollId).Set(ctx, poll)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func (s *FirestorePollService) Update(ctx context.Context, poll *Poll) (int64, error) {
	_, err := s.Collection.Doc(poll.PollId).Set(ctx, poll)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func (s *FirestorePollService) Patch(ctx context.Context, json map[string]interface{}) (int64, error) {
	pollType := reflect.TypeOf(Poll{})
	maps := fs.MakeFirestoreMap(pollType)

	uid := json["pollid"]
	pollid := uid.(string)
	docRef := s.Collection.Doc(pollid)
	doc, err1 := docRef.Get(ctx)
	if err1 != nil {
		return -1, err1
	}
	delete(json, "pollid")

	dest := fs.MapToFirestore(json, doc, maps)
	_, err := docRef.Set(ctx, dest)
	if err != nil {
		return -1, err
	}
	return 1, nil
}

func (s *FirestorePollService) Delete(ctx context.Context, pollid string) (int64, error) {
	_, err := s.Collection.Doc(pollid).Delete(ctx)
	if err != nil {
		return -1, err
	}
	return 1, nil
}
