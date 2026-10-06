package subjectnote

import (
	"homewormanager/internal/entity"
)

type subject struct {
	SubjectID   entity.SubjectID `json:"subject_id"`
	SubjectName string           `json:"subject_name"`
}

type subjectNote struct {
	NoteID    entity.NoteID    `json:"note_id"`
	SubjectID entity.SubjectID `json:"subject_id"`
	NoteText  string           `json:"note_text"`
}

type subjectWithNotes struct {
	Subject subject       `json:"subject"`
	Notes   []subjectNote `json:"notes"`
}

func toSubjectNotes(n []entity.SubjectWithNotes) []subjectWithNotes {
	res := make([]subjectWithNotes, len(n))
	for i := range n {
		res[i] = subjectWithNotes{
			Subject: toSubject(n[i].Subject),
			Notes:   toSubjectNote(n[i].Notes),
		}
	}
	return res
}

func toSubject(g entity.Subject) subject {
	return subject{
		SubjectID:   g.SubjectID,
		SubjectName: g.SubjectName,
	}
}

func toSubjectNote(g []entity.SubjectNote) []subjectNote {
	res := make([]subjectNote, len(g))
	for i := range g {
		res[i] = subjectNote{
			NoteID:    g[i].NoteID,
			SubjectID: g[i].SubjectID,
			NoteText:  g[i].NoteText,
		}
	}
	return res
}
