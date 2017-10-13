package models

type User struct {
	EmailAddress string `json:"emailAddress"`
	Name string `json:"name"`
	Hometown string `json:"hometown"`
	ExperiencePoints int `json:"experiencePoints"`
}