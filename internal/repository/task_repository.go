package repository

import (
	"context"
	"errors"
	"project/internal/model"

	"github.com/jackc/pgx/v5"
)

var (
	ErrTaskNotFound = errors.New("Task not found")
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetById(ctx context.Context, id int) (*model.Task, error)
	GetAll(ctx context.Context) ([]model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int) error
}

type taskRepository struct {
	db *pgx.Conn
}

func NewTaskRepository(db *pgx.Conn) TaskRepository {
	return &taskRepository{db: db}
}

func (t *taskRepository) Create(ctx context.Context, task *model.Task) error {
	query := `
		insert into tasks(title, description, status)
		values ($1, $2, $3)
		returning id
	`

	err := t.db.QueryRow(ctx, query,
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
	res, err := t.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (t *taskRepository) GetAll(ctx context.Context) ([]model.Task, error) {
	query := `select * from tasks`
	rows, err := t.db.Query(ctx, query)
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
		QueryRow(ctx, query, id).
		Scan(&task.ID, &task.Title, &task.Description, &task.Status)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrTaskNotFound
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

	res, err := t.db.Exec(ctx, query, task.Title, task.Description, task.Status, task.ID)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}
