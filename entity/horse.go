package entity

type Horse struct {
	Name         string
	Jockey       Jockey
	Trainer      Trainer
	Rating       int
	Weight       float64
	Barrier      int
	Blinkers     bool
	Sectionals   []Sectional
}