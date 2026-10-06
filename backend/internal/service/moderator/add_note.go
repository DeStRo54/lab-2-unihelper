package moderator

import (
	"context"
	"homewormanager/internal/entity"
)

type AddNote struct {
	SubjectID entity.SubjectID
	GroupID   entity.GroupID
	NoteText  string
}

func (s *Service) AddNote(ctx context.Context, req AddNote) (entity.NoteID, error) {
	note := req.toSubjectNote()

	id, err := s.subjectNoteService.AddNote(ctx, note)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AddNote) toSubjectNote() entity.SubjectNote {
	return entity.SubjectNote{
		SubjectID: r.SubjectID,
		GroupID:   r.GroupID,
		NoteText:  r.NoteText,
	}
}
