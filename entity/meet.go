package entity

import (
	"time"
)

type Meet struct {
	Date    time.Time
	Track   Track
	Races   map[int]*Race
}