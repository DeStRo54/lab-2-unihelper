package class

import (
	"context"
	"fmt"
	"reflect"

	"homewormanager/internal/entity"
)

func (r *Repo) Create(ctx context.Context, classes []entity.Class) error {

	//TODO: заюзать здесь sql билдер
	q := "INSERT INTO classes(group_id, subject_id, start_time, end_time, summary, description, class_sem_number, location, category) VALUES "

	numField := reflect.TypeOf(entity.Class{}).NumField() - 1

	var args []interface{}

	for i, class := range classes {
		if i > 0 {
			q += ", "
		}
		q += fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", i*numField+1, i*numField+2, i*numField+3, i*numField+4, i*numField+5, i*numField+6, i*numField+7, i*numField+8, i*numField+9)
		args = append(args, class.GroupID, class.SubjectID, class.StartTime, class.EndTime, class.Summary, class.Description, class.SemClassNumber, class.Location, class.Category)
	}

	t := r.manager.GetTxOrDefault(ctx)

	_, err := t.ExecContext(ctx, q, args...)
	if err != nil {
		return err
	}

	return nil

}
