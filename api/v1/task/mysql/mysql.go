package mysql

import (
	"context"
	"database/sql"
	"github.com/sword/api-backend-challenge/api/v1/task"
	"github.com/sword/api-backend-challenge/model"
)

type mysql struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) task.Repository {
	return &mysql{db}
}

func (m *mysql) Create(ctx context.Context, task *model.Task) error {

	insert := `
	INSERT INTO sword.tasks(
			id,
			summary,
			performed_at
		)
	VALUES(?, ?, ?)`

	_, err := m.db.ExecContext(
		ctx,
		insert,
		task.ID,
		task.Summary,
		task.PerformedAt)

	return err
}

func (m *mysql) List(ctx context.Context) ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)

	rows, err := m.db.QueryContext(ctx, `
		SELECT 
			id,
			summary,
			performed_at
		FROM
			sword.tasks
	`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		task := new(model.Task)
		err = rows.Scan(&task.ID, &task.Summary, &task.PerformedAt)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
