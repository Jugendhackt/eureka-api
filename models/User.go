package models

import "time"

type User struct {
	JoinDate time.Time `json:"joinDate"`
	EmailAddress string `json:"emailAddress"`
	Name string `json:"name"`
	Hometown string `json:"hometown"`
	ExperiencePoints int `json:"experiencePoints"`
}