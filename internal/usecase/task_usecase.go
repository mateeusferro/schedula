package usecase

import (
	"github.com/mateeusferro/schedula/internal/domain"
	"github.com/mateeusferro/schedula/internal/repository"
)

type TaskUseCase struct {
	repo *repository.TaskRepository
}

func NewTaskUseCase(repo *repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (usecase *TaskUseCase) ExecuteGetTask(id string) (*domain.Task, error) {
	result, err := usecase.repo.GetTaskInfo(id)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (usecase *TaskUseCase) ExecuteGetTasksByStatus(status string) ([]domain.Task, error) {
	result, err := usecase.repo.GetTasksByStatus(status)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (usecase *TaskUseCase) ExecuteCreateTask(task domain.TaskInput) (bool, error) {
	_, err := usecase.repo.CreateTask(task)
	if err != nil {
		return false, err
	}

	return true, err
}

func (usecase *TaskUseCase) ExecuteDeleteTask(id string) (bool, error) {
	_, err := usecase.repo.DeleteTask(id)
	if err != nil {
		return false, err
	}

	return true, err
}
