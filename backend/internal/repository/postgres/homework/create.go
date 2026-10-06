package homework

import (
	"context"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, homework entity.Homework) (entity.HomeworkID, error) {

	q := "INSERT INTO homeworks(class_sem_number, group_id, subject_id, homework_text, category, due_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING homework_id"

	t := r.manager.GetTxOrDefault(ctx)

	var id entity.HomeworkID

	err := t.QueryRowContext(ctx, q, homework.SemClassNumber, homework.GroupID, homework.SubjectID, homework.HomeworkText, homework.Category, homework.DueDate).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil

}
