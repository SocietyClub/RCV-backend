package models

type Poll struct {
	PollId            string `json:"pollId gorm:"column:pollId;primary_key" bson:"_pollId" dynamodbav:"pollId" firestore:"-" validate:"required,max=40"`
	RankedChoiceCount string `json:"rankedChoiceCount gorm:"column:rankedChoiceCount" bson:"rankedChoiceCount" dynamodbav:"rankedChoiceCount" firestore:"rankedChoiceCount" validate:"required,rankedChoiceCount,max=100"`
}
