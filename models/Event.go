package models

import "time"

type Event struct {
	Place string `json:"place"`
	Activity string `json:"activity"` // Id
	Time time.Time `json:"time"`
	Lang string `json:"lang"`
}

