package models

import "time"

type Event struct {
	CreationTime time.Time `json:"creationTime"`
	CreatorId string `json:"creatorId"`
	Place string `json:"place"`
	ActivityId string `json:"activityId"` // Id
	Time time.Time `json:"time"`
	Lang string `json:"lang"`
}

