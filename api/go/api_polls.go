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
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A PollsApiController binds http requests to an api service and writes the service results to the http response
type PollsApiController struct {
	service PollsApiServicer
}

// NewPollsApiController creates a default api controller
func NewPollsApiController(s PollsApiServicer) Router {
	return &PollsApiController{service: s}
}

// Routes returns all of the api route for the PollsApiController
func (c *PollsApiController) Routes() Routes {
	return Routes{ 
		{
			"CreatePoll",
			strings.ToUpper("Post"),
			"/ranked-choice-vote/v1/poll",
			c.CreatePoll,
		},
		{
			"DeletePoll",
			strings.ToUpper("Delete"),
			"/ranked-choice-vote/v1/poll/{PollID}",
			c.DeletePoll,
		},
		{
			"GetPoll",
			strings.ToUpper("Get"),
			"/ranked-choice-vote/v1/poll/{PollID}",
			c.GetPoll,
		},
		{
			"GetPollResults",
			strings.ToUpper("Get"),
			"/ranked-choice-vote/v1/poll/{PollID}/results",
			c.GetPollResults,
		},
		{
			"UpdatePoll",
			strings.ToUpper("Patch"),
			"/ranked-choice-vote/v1/poll/{PollID}",
			c.UpdatePoll,
		},
	}
}


// CreatePoll - Creates a new Poll
func (c *PollsApiController) CreatePoll(w http.ResponseWriter, r *http.Request) {
	xUSERID := r.Header.Get("X-USER-ID")
	createPollRequest := &CreatePollRequest{}
	if err := json.NewDecoder(r.Body).Decode(&createPollRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.CreatePoll(r.Context(), xUSERID, *createPollRequest)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// DeletePoll - Deletes an existing Poll
func (c *PollsApiController) DeletePoll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	xUSERID := r.Header.Get("X-USER-ID")
	pollID := params["PollID"]
	
	result, err := c.service.DeletePoll(r.Context(), xUSERID, pollID)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPoll - Gets a specific Poll by its ID.
func (c *PollsApiController) GetPoll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	xUSERID := r.Header.Get("X-USER-ID")
	pollID := params["PollID"]
	
	result, err := c.service.GetPoll(r.Context(), xUSERID, pollID)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetPollResults - Gets the Results of a specific Poll by its ID
func (c *PollsApiController) GetPollResults(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	xUSERID := r.Header.Get("X-USER-ID")
	pollID := params["PollID"]
	
	result, err := c.service.GetPollResults(r.Context(), xUSERID, pollID)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// UpdatePoll - Updates an existing Poll
func (c *PollsApiController) UpdatePoll(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	xUSERID := r.Header.Get("X-USER-ID")
	pollID := params["PollID"]
	
	updatePollRequest := &UpdatePollRequest{}
	if err := json.NewDecoder(r.Body).Decode(&updatePollRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result, err := c.service.UpdatePoll(r.Context(), xUSERID, pollID, *updatePollRequest)
	// If an error occurred, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
