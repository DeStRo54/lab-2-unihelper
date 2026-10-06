package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) Update(ctx context.Context, id entity.NoteID, noteText string) error {
	q := "UPDATE subjectnotes SET note_text = $1 WHERE note_id = $2"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, noteText, id)

	if err != nil {
		return err
	}

	return nil
}
