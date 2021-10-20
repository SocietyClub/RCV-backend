/*
 * Ranked Choice Voting API
 *
 * This is the API for Creating, Managing and Fetching Polls and Votes for the RCV Project.
 *
 * API version: 1.0.0-dev
 * Contact: teamsocietyclub@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// PollsApiService is a service that implents the logic for the PollsApiServicer
// This service should implement the business logic for every endpoint for the PollsApi API.
// Include any external packages or services that will be required by this service.
type PollsApiService struct {
}

// NewPollsApiService creates a default api service
func NewPollsApiService() PollsApiServicer {
	return &PollsApiService{}
}

// CreatePoll - Creates a new Poll
func (s *PollsApiService) CreatePoll(ctx context.Context, xUSERID string, createPollRequest CreatePollRequest) (ImplResponse, error) {
	// TODO - update CreatePoll with the required logic for this service method.
	// Add api_polls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, AddPollResponse{}) or use other options such as http.Ok ...
	//return Response(200, AddPollResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, Messages{}) or use other options such as http.Ok ...
	//return Response(400, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(401, Messages{}) or use other options such as http.Ok ...
	//return Response(401, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(403, Messages{}) or use other options such as http.Ok ...
	//return Response(403, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(404, Messages{}) or use other options such as http.Ok ...
	//return Response(404, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(422, Messages{}) or use other options such as http.Ok ...
	//return Response(422, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(500, Messages{}) or use other options such as http.Ok ...
	//return Response(500, Messages{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreatePoll method not implemented")
}

// DeletePoll - Deletes an existing Poll
func (s *PollsApiService) DeletePoll(ctx context.Context, xUSERID string, pollID string) (ImplResponse, error) {
	// TODO - update DeletePoll with the required logic for this service method.
	// Add api_polls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(204, Messages{}) or use other options such as http.Ok ...
	//return Response(204, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(400, Messages{}) or use other options such as http.Ok ...
	//return Response(400, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(401, Messages{}) or use other options such as http.Ok ...
	//return Response(401, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(403, Messages{}) or use other options such as http.Ok ...
	//return Response(403, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(404, Messages{}) or use other options such as http.Ok ...
	//return Response(404, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(422, Messages{}) or use other options such as http.Ok ...
	//return Response(422, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(500, Messages{}) or use other options such as http.Ok ...
	//return Response(500, Messages{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePoll method not implemented")
}

// GetPoll - Gets a specific Poll by its ID.
func (s *PollsApiService) GetPoll(ctx context.Context, xUSERID string, pollID string) (ImplResponse, error) {
	var messages Messages
	var poll_model GetPollResponse

	context_background := context.Background()
	firestore_client := GetFirestoreClient(context_background)

	defer firestore_client.Close()

	poll, err := firestore_client.Collection("polls").Doc(pollID).Get(ctx)

	if err != nil {
		AddMessage(&messages, Severity(ERROR), "GetPoll-0", fmt.Sprintf("GetPoll could not find the given pollID(%s): %s", pollID, err))
		poll_model.Messages = messages
		return Response(http.StatusNotFound, poll_model), err
	}

	err = poll.DataTo(&poll_model.GetPollData)

	if err != nil {
		AddMessage(&messages, Severity(ERROR), "GetPoll-1", fmt.Sprintf("GetPoll could not load the given pollID(%s): %s", pollID, err))
		poll_model.Messages = messages
		return Response(http.StatusNotAcceptable, poll_model), err
	}

	AddMessage(&messages, Severity(INFO), "GetPoll-OK", "Poll successfully retreived")
	poll_model.Messages = messages
	return Response(http.StatusOK, poll_model), nil
}

// GetPollResults - Gets the Results of a specific Poll by its ID
func (s *PollsApiService) GetPollResults(ctx context.Context, xUSERID string, pollID string) (ImplResponse, error) {
	// TODO - update GetPollResults with the required logic for this service method.
	// Add api_polls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, PollResults{}) or use other options such as http.Ok ...
	//return Response(200, PollResults{}), nil

	//TODO: Uncomment the next line to return response Response(400, Messages{}) or use other options such as http.Ok ...
	//return Response(400, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(401, Messages{}) or use other options such as http.Ok ...
	//return Response(401, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(403, Messages{}) or use other options such as http.Ok ...
	//return Response(403, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(404, Messages{}) or use other options such as http.Ok ...
	//return Response(404, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(422, Messages{}) or use other options such as http.Ok ...
	//return Response(422, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(500, Messages{}) or use other options such as http.Ok ...
	//return Response(500, Messages{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPollResults method not implemented")
}

// UpdatePoll - Updates an existing Poll
func (s *PollsApiService) UpdatePoll(ctx context.Context, xUSERID string, pollID string, updatePollRequest UpdatePollRequest) (ImplResponse, error) {
	// TODO - update UpdatePoll with the required logic for this service method.
	// Add api_polls_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, GetPollResponse{}) or use other options such as http.Ok ...
	//return Response(200, GetPollResponse{}), nil

	//TODO: Uncomment the next line to return response Response(400, Messages{}) or use other options such as http.Ok ...
	//return Response(400, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(401, Messages{}) or use other options such as http.Ok ...
	//return Response(401, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(403, Messages{}) or use other options such as http.Ok ...
	//return Response(403, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(404, Messages{}) or use other options such as http.Ok ...
	//return Response(404, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(422, Messages{}) or use other options such as http.Ok ...
	//return Response(422, Messages{}), nil

	//TODO: Uncomment the next line to return response Response(500, Messages{}) or use other options such as http.Ok ...
	//return Response(500, Messages{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePoll method not implemented")
}
