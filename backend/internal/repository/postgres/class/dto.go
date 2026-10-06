package class

import (
	"time"

	"homewormanager/internal/entity"
)

type class struct {
	ClassID        entity.ClassID       `db:"class_id"`
	GroupID        entity.GroupID       `db:"group_id"`
	SubjectID      entity.SubjectID     `db:"subject_id"`
	StartTime      time.Time            `db:"start_time"`
	EndTime        time.Time            `db:"end_time"`
	Summary        string               `db:"summary"`
	Description    string               `db:"description"`
	SemClassNumber int64                `db:"class_sem_number"`
	Location       string               `db:"location"`
	Category       entity.ClassCategory `db:"category"`
}

func (c class) toClass() entity.Class {
	return entity.Class{
		ClassID:        c.ClassID,
		GroupID:        c.GroupID,
		SubjectID:      c.SubjectID,
		StartTime:      c.StartTime.Local(),
		EndTime:        c.EndTime.Local(),
		Summary:        c.Summary,
		Description:    c.Description,
		SemClassNumber: c.SemClassNumber,
		Location:       c.Location,
		Category:       c.Category,
	}
}
