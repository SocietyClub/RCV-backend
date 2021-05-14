/*
 * Ranked Choice Voting API
 *
 * This is the API for Creating, Managing and Fetching Polls and Votes for the RCV Project.
 *
 * API version: 1.0.0
 * Contact: test@test.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type VoteCount struct {

	ChoicePosition int32 `json:"choicePosition,omitempty"`

	VoteCount int32 `json:"voteCount,omitempty"`
}
