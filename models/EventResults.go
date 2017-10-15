package models

type EventResults struct {
	Recent []Event `bson:"recent" json:"recent"`
	Upcoming []Event `json:"upcoming" json:"upcoming"`
}