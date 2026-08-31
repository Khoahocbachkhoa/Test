package repository

import (
	"context"
	"database/sql"
	"errors"
	"project/internal/model"
)

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
		insert into tasks(title, description, status)
		value ($1, $2, $3)
		returning id
	`

	err := t.db.QueryRowContext(ctx, query,
		task.Title,
		task.Description,
		task.Status).
		Scan(&task.ID)

	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Delete(ctx context.Context, id int) error {
	query := `delete from tasks where id = $1`
	res, err := t.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func (t *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	query := `select * from tasks`
	rows, err := t.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if tasks == nil {
		return []model.Task{}, nil
	}

	return tasks, nil
}

func (t *taskRepository) GetById(ctx context.Context, id int) (*model.Task, error) {
	query := `select * from tasks where id = $1`
	var task model.Task

	err := t.db.
		QueryRowContext(ctx, query, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) Update(ctx context.Context, task *model.Task) error {
	query := `
		update tasks
		set title = $1, description = $2, status = $3
		where id = $4
	`

	res, err := t.db.ExecContext(ctx, query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return ErrNotFound
	}

	return nil
}
