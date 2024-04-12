package entity

import "time"

type Sectional struct {
	Date       time.Time
	Track      Track
	RaceNum    int
	Time       time.Time
	Speed600   float64
	Speed400   float64
	Speed200   float64
	Speed000   float64
}