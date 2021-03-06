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

type Vote struct {

	ChoicePosition int32 `json:"choicePosition,omitempty"`

	Candidate CandidateVote `json:"candidate,omitempty"`
}
