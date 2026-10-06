package entity

import "time"

type Class struct {
	ClassID        ClassID
	GroupID        GroupID
	SubjectID      SubjectID
	StartTime      time.Time
	EndTime        time.Time
	Summary        string
	Description    string
	SemClassNumber int64
	Location       string
	Category       ClassCategory
}
