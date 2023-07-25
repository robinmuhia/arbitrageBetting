package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string
	Email string `gorm:"unique"`
	Password string
	Image string
	bookmarkerRegion string
	subscriptionPaid bool
}