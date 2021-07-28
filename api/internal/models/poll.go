package models

import "time"

type Poll struct {
	Id                      string     `json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"-" validate:"required,max=40"`
	StartDate               *time.Time `json:"startDate" gorm:"column:start_date" bson:"startDate" dynamodbav:"startDate" firestore:"startDate"`
	EndDate                 *time.Time `json:"endDate" gorm:"column:end_date" bson:"endDate" dynamodbav:"endDate" firestore:"-"`
	MaxNumRankedChoiceCount string     `json:"maxNumRankedChoiceCount" gorm:"column:maxNumRankedChoiceCount" bson:"maxNumRankedChoiceCount" dynamodbav:"maxNumRankedChoiceCount" firestore:"maxNumRankedChoiceCount"`
	Candidate               string     `json:"candidate" gorm:"column:candidate" bson:"candidate" dynamodbav:"candidate" firestore:"candidate"`
}
