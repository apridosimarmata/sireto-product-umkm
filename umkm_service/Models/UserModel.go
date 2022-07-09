package Models

import "time"

type User struct {
	Id int `json:"id"`

	Username          string `json:"username" gorm:"unique;not null;size:20"`
	Name              string `json:"name" gorm:"not null;size:50"`
	ProfilePictureUrl string `json:"profile_picture_url" gorm:"size:30"`
	Email             string `json:"email" gorm:"unique;size:50"`
	Phone             string `json:"phone" gorm:"unique;size:15"`
	Password          string `json:"password"`

	TokenFamily string `json:"token_family"`

	DateOfBirth time.Time `json:"date_of_birth"`
	CreatedAt   time.Time `json:"created_at"`
}
