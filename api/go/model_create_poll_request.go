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

// CreatePollRequest - This is the object request information needed to add a poll.
type CreatePollRequest struct {

	// Name for Poll entered by the user used as information purposes only.
	PollName string `json:"pollName"`

	MaxNumRankedChoiceCount int32 `json:"maxNumRankedChoiceCount"`

	CandidateList []Candidate `json:"candidateList"`
}