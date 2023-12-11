package models

import "github.com/jinzhu/gorm"

// User object that represents current structure of our user
type User struct {
	gorm.Model
	Name string	`gorm:"not null;default:null"`
	Email string `gorm:"unique;not null;type:varchar(100);default:null"`
	Password string `gorm:"not null;default:null"`
	BookmarkerRegion string
	SubscriptionPaid bool
}