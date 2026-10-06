package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (s *Service) GetByGroup(ctx context.Context, groupID entity.GroupID) ([]entity.SubjectWithNotes, error) {
	notes, err := s.subjectNoteRepo.GetByGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}

	subjects, err := s.subjectService.GetByGroupId(ctx, groupID)
	if err != nil {
		return nil, err
	}

	notesMap := make(map[entity.SubjectID][]entity.SubjectNote)

	for _, note := range notes {
		notesMap[note.SubjectID] = append(notesMap[note.SubjectID], note)
	}

	subjectsWithNotes := make([]entity.SubjectWithNotes, len(subjects))

	for i, subject := range subjects {
		subjectsWithNotes[i] = entity.SubjectWithNotes{
			Subject: subject,
			Notes:   notesMap[subject.SubjectID],
		}
	}

	return subjectsWithNotes, nil
}
