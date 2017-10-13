package models

import "time"

// TODO: In the database, user properties are being saved lowercased (Ignoring json formatting)
type User struct {
	JoinDate time.Time `json:"joinDate"`
	EmailAddress string `json:"emailAddress"`
	Name string `json:"name"`
	Hometown string `json:"hometown"`
	ExperiencePoints int `json:"experiencePoints"`
}