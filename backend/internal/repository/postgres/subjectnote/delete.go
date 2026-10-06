package subjectnote

import (
	"context"
	"homewormanager/internal/entity"
)

func (r *Repo) Delete(ctx context.Context, id entity.NoteID) error {
	q := "DELETE FROM subjectnotes WHERE note_id = $1"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, id)

	if err != nil {
		return err
	}

	return nil
}
