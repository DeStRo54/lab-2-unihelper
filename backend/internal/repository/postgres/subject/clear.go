package subject

import "context"

func (r *Repo) Clear(ctx context.Context) error {

	q := "DELETE FROM subjects"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
