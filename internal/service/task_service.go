package service

import (
	"context"
	"project/internal/model"
	"project/internal/repository"
)

type TaskService interface {
	AddTask(task *model.Task) error
	GetTaskByID(id int) (*model.Task, error)
	GetAllTask() ([]model.Task, error)
	UpdateTask(task *model.Task) error
	Delete(id int) error
}

type taskService struct {
	taskRepository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskService{taskRepository: taskRepository}
}

func (t *taskService) AddTask(task *model.Task) error {
	return t.taskRepository.Create(context.Background(), task)
}

func (t *taskService) Delete(id int) error {
	return t.taskRepository.Delete(context.Background(), id)
}

func (t *taskService) GetAllTask() ([]model.Task, error) {
	return t.taskRepository.GetAll(context.Background())
}

func (t *taskService) GetTaskByID(id int) (*model.Task, error) {
	return t.taskRepository.GetById(context.Background(), id)
}

func (t *taskService) UpdateTask(task *model.Task) error {
	return t.taskRepository.Update(context.Background(), task)
}
