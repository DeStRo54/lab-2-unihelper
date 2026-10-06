package homework

import "context"

func (r *Repo) Clear(ctx context.Context) error {

	q := "DELETE FROM homeworks"

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q)
	if err != nil {
		return err
	}

	return nil
}
