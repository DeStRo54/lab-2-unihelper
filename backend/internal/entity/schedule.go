package entity

import "time"

type OutputClass struct {
	Class    Class
	Homework []Homework
}

type ScheduleDay struct {
	Date                time.Time
	OutputClass         []OutputClass
	IndependentHomework []Homework
}
type HomeworkDay struct {
	Date     time.Time
	Homework []Homework
}
