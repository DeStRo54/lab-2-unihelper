package entity

import "time"

type Homework struct {
	HomeworkID     HomeworkID
	SemClassNumber *int64
	GroupID        GroupID
	SubjectID      SubjectID
	SubjectName    string
	HomeworkText   string
	Category       *ClassCategory
	IsCompleted    bool
	Files          []HomeworkFile
	DueDate        time.Time
	CreatedAt      time.Time
}
