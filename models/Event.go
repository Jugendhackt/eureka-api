package models

import "time"

type Event struct {
	Id			string		`bson:"_id" json:"id"`
	CreationDate time.Time `bson:"creationDate" json:"creationDate"`
	CreatorId    string    `bson:"creatorId" json:"creatorId"`
	Place        string    `bson:"place" json:"place"`
	ActivityId   string    `bson:"activityId" json:"activityId"`
	Time         time.Time `bson:"time" json:"time"`
	Lang         string    `bson:"lang" json:"lang"`
}
