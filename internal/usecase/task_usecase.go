package usecase

import (
	"github.com/mateeusferro/schedula/internal/repository"
)

type TaskUseCase struct {
	repo *repository.TaskRepository
}

func NewTaskUseCase(repo *repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{repo: repo}
}
