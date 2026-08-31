package service

import (
	"context"
	"project/internal/model"
	"project/internal/repository"
)

type TaskService interface {
	AddTask(ctx context.Context, task *model.Task) error
	GetTaskByID(ctx context.Context, id int) (*model.Task, error)
	GetAllTask(ctx context.Context) ([]model.Task, error)
	UpdateTask(ctx context.Context, task *model.Task) error
	Delete(ctx context.Context, id int) error
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{taskRepository: taskRepository}
}

func (t *taskService) AddTask(ctx context.Context, task *model.Task) error {
	return t.taskRepository.Create(ctx, task)
}

func (t *taskService) Delete(ctx context.Context, id int) error {
	return t.taskRepository.Delete(ctx, id)
}

func (t *taskService) GetAllTask(ctx context.Context) ([]model.Task, error) {
	return t.taskRepository.GetAll(ctx)
}

func (t *taskService) GetTaskByID(ctx context.Context, id int) (*model.Task, error) {
	return t.taskRepository.GetById(ctx, id)
}

func (t *taskService) UpdateTask(ctx context.Context, task *model.Task) error {
	return t.taskRepository.Update(ctx, task)
}
