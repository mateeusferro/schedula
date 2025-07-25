package delivery

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateeusferro/schedula/internal/domain"
	"github.com/mateeusferro/schedula/internal/repository"
	"github.com/mateeusferro/schedula/internal/usecase"
)

func Routes(router *gin.Engine, db *sql.DB) {
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)

	router.GET("/task/:id", handleGetTaskInfo(taskUseCase))
	router.GET("/task", handleGetTasksByStatus(taskUseCase))
	router.POST("/task", handleCreateTask(taskUseCase))
	router.DELETE("/task/:id", handleDeleteTask(taskUseCase))
}

func handleGetTaskInfo(taskUseCase *usecase.TaskUseCase) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		task, err := taskUseCase.ExecuteGetTask(id)
		if err != nil {
			fmt.Printf("Error: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}

		context.JSON(http.StatusOK, task)
	}
}

func handleGetTasksByStatus(taskUseCase *usecase.TaskUseCase) gin.HandlerFunc {
	return func(context *gin.Context) {
		status := context.Query("status")

		tasks, err := taskUseCase.ExecuteGetTasksByStatus(status)
		if err != nil {
			fmt.Printf("Error: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
			return
		}

		context.JSON(http.StatusOK, tasks)
	}
}

func handleCreateTask(taskUseCase *usecase.TaskUseCase) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task domain.TaskInput
		err := context.ShouldBindJSON(&task)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Incorret json input"})
			return
		}

		result, err := taskUseCase.ExecuteCreateTask(task)
		if err != nil {
			fmt.Printf("Error: %v", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong while creating the task"})
			return
		}

		context.JSON(http.StatusCreated, result)
	}
}

func handleDeleteTask(taskUseCase *usecase.TaskUseCase) gin.HandlerFunc {
	return func(context *gin.Context) {
		id := context.Param("id")

		result, err := taskUseCase.ExecuteDeleteTask(id)
		if err != nil {
			fmt.Printf("Error: %v", err)
			context.JSON(
				http.StatusInternalServerError,
				gin.H{
					"error": "Something went wrong while deleting the task for id: " + id,
				})
			return
		}

		context.JSON(http.StatusOK, result)
	}
}
