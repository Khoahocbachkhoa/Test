package repository

import (
	"context"
	"errors"
	"project/internal/model"
)

var (
	ErrNotFound      = errors.New("resource not found")
	ErrAlreadyExists = errors.New("resource already exists")
)

type TaskRepository interface {
	Create(ctx context.Context, task *model.Task) error
	GetById(ctx context.Context, id int) (*model.Task, error)
	GetAll(ctx context.Context) ([]model.Task, error)
	Update(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int) error
}
