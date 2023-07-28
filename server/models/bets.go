package models

import (
	"github.com/jinzhu/gorm"
)

type ThreeOddsBet struct {
	gorm.Model
	Title string `gorm:"not null;default:null"`
	Home string `gorm:"not null;default:null"`
	Draw string `gorm:"not null;default:null"`
	Away string `gorm:"not null;default:null"`
	HomeOdds float64 `gorm:"not null;default:null"`
	DrawOdds float64 `gorm:"not null;default:null"`
	AwayOdds float64 `gorm:"not null;default:null"`
	GameType string `gorm:"not null;default:null"`
	League string `gorm:"not null;default:null"`
	Profit float64 `gorm:"not null;default:null"`
	BookmarkerRegion string `gorm:"not null;default:null"`
	GameTime string `gorm:"not null;default:null"`
}

type TwoOddsBet struct {
	gorm.Model
	Title string `gorm:"not null;default:null"`
	Home string `gorm:"not null;default:null"`
	Away string `gorm:"not null;default:null"`
	HomeOdds float64 `gorm:"not null;default:null"`
	AwayOdds float64 `gorm:"not null;default:null"`
	GameType string `gorm:"not null;default:null"`
	League string `gorm:"not null;default:null"`
	Profit float64 `gorm:"not null;default:null"`
	BookmarkerRegion string `gorm:"not null;default:null"`
	GameTime string `gorm:"not null;default:null"`
}

type Sport struct {
	Key     	  string `json:"key"`
	Group     	  string `json:"group"`
	Title     	  string `json:"title"`
	Description   string `json:"description"`
	Active        bool   `json:"active"`
	HasOutrights  bool `json:"has_outrights"`
}

type Outcome struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Market struct {
	Key         string    `json:"key"`
	LastUpdate  string    `json:"last_update"`
	Outcomes    []Outcome `json:"outcomes"`
}

type Bookmaker struct {
	Key        string   `json:"key"`
	Title      string   `json:"title"`
	LastUpdate string   `json:"last_update"`
	Markets    []Market `json:"markets"`
}

type Odds struct {
	ID           string      `json:"id"`
	SportKey     string      `json:"sport_key"`
	SportTitle   string      `json:"sport_title"`
	CommenceTime string      `json:"commence_time"`
	HomeTeam     string      `json:"home_team"`
	AwayTeam     string      `json:"away_team"`
	Bookmakers   []Bookmaker `json:"bookmakers"`
}
