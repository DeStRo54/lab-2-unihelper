package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, note entity.SubjectNote) (entity.NoteID, error) {

	q := "INSERT INTO subjectnotes(subject_id, group_id, note_text) VALUES ($1, $2, $3)  RETURNING note_id"

	t := r.manager.GetTxOrDefault(ctx)

	var id entity.NoteID

	err := t.QueryRowContext(ctx, q, note.SubjectID, note.GroupID, note.NoteText).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil

}
