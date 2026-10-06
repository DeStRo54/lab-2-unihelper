package subjectnote

import "context"

func (r *Repo) Clear(ctx context.Context) error {

	q := "DELETE FROM subjectnotes"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
