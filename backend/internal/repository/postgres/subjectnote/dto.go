package subjectnote

import (
	"homewormanager/internal/entity"
	"time"
)

type subjectNote struct {
	NoteID    entity.NoteID    `db:"note_id"`
	SubjectID entity.SubjectID `db:"subject_id"`
	GroupID   entity.GroupID   `db:"group_id"`
	NoteText  string           `db:"note_text"`
	CreatedAt time.Time        `db:"created_at"`
}

func (h subjectNote) toSubjectNote() entity.SubjectNote {
	return entity.SubjectNote{
		NoteID:    h.NoteID,
		SubjectID: h.SubjectID,
		GroupID:   h.GroupID,
		NoteText:  h.NoteText,
		CreatedAt: h.CreatedAt,
	}
}
