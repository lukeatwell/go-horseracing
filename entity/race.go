package entity

type Race struct {
	RaceNum   int
	Length    int
	Class     Class
	Rating    Rating
	Age       Age
	Sex       Sex
	Horses    map[string]*Horse
}