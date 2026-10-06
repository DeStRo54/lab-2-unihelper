package entity

import "time"

type SubjectNote struct {
	NoteID    NoteID
	SubjectID SubjectID
	GroupID   GroupID
	NoteText  string
	CreatedAt time.Time
}

type SubjectWithNotes struct {
	Subject Subject
	Notes   []SubjectNote
}
