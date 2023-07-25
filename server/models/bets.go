package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ThreeOddsBet struct {
	gorm.Model
	title string
	home string
	draw string
	away string
	homeOdds float64
	drawOdds float64
	awayOdds float64
	gameType string
	league string
	profit float64
	bookmarkerRegion string
	gameTime time.Time
}

type TwoOddsBet struct {
	gorm.Model
	title string
	home string
	away string
	homeOdds float64
	awayOdds float64
	gameType string
	league string
	profit float64
	bookmarkerRegion string
	gameTime time.Time
}