package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Filename         string
	Status           string
	Filesize         int
	Channel          string
	SenderAdd        string
	SubscriberAmount int
}
