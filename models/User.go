package models

import (
	"time"
)

// TODO: In the database, user properties are being saved lowercased (Ignoring json formatting)
type User struct {
	Id               string    `bson:"_id" json:"id"`
	JoinDate         time.Time `bson:"joinDate" json:"joinDate"`
	EmailAddress     string    `bson:"emailAddress" json:"emailAddress"`
	Password         string    `bson:"password" json:"password"`
	Name             string    `bson:"name" json:"name"`
	Hometown         string    `bson:"hometown" json:"hometown"`
	ExperiencePoints int       `bson:"experiencePoints" json:"experiencePoints"'`
}
