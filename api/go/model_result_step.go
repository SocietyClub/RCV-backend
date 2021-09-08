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

type ResultStep struct {

	StepId int64 `json:"stepId,omitempty"`

	CandidateList []CandidateResults `json:"candidateList,omitempty"`
}