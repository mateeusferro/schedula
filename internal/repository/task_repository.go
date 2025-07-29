package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/mateeusferro/schedula/internal/domain"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (repository *TaskRepository) CreateTask(task domain.TaskToSave) (bool, error) {
	_, err := repository.DB.Exec(`
			INSERT INTO SCHEDULED_TASKS
				(NAME, PAYLOAD, RUN_AT, STATUS, ATTEMPTS, MAX_ATTEMPTS)
			VALUES
				($1, $2, $3, $4, $5, $6)
		`, task.Name, task.Payload, task.Run_at,
		task.Status, task.Attempts, task.Max_attempts)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repository *TaskRepository) GetTaskInfo(id string) (*domain.Task, error) {
	row := repository.DB.QueryRow(`
			SELECT * FROM SCHEDULED_TASKS 
			WHERE ID = $1
		`, id)

	var task domain.Task
	var rawPayload []byte

	err := row.Scan(
		&task.Id,
		&task.Name,
		&rawPayload,
		&task.Run_at,
		&task.Status,
		&task.Attempts,
		&task.Max_attempts,
		&task.Created_at,
		&task.Updated_at,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(rawPayload, &task.Payload)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (repository *TaskRepository) GetTasksByStatus(status string) ([]domain.Task, error) {
	rows, err := repository.DB.Query(`
		SELECT * FROM SCHEDULED_TASKS
		WHERE STATUS = $1
	`, status)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		var rawPayload []byte

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&rawPayload,
			&task.Run_at,
			&task.Status,
			&task.Attempts,
			&task.Max_attempts,
			&task.Created_at,
			&task.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(rawPayload, &task.Payload)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *TaskRepository) DeleteTask(id string) (bool, error) {
	_, err := repository.DB.Exec("DELETE FROM SCHEDULED_TASKS WHERE id = $1", id)
	if err != nil {
		return false, err
	}

	return true, err
}

func (repository *TaskRepository) GetPendingTask() ([]domain.Task, error) {
	rows, err := repository.DB.Query(`
		SELECT * FROM SCHEDULED_TASKS
		WHERE RUN_AT <= NOW() AND STATUS = 'pending'
	`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		var rawPayload []byte

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&rawPayload,
			&task.Run_at,
			&task.Status,
			&task.Attempts,
			&task.Max_attempts,
			&task.Created_at,
			&task.Updated_at,
		)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(rawPayload, &task.Payload)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (repository *TaskRepository) UpdateTaskStatus(id string, status string) (bool, error) {
	_, err := repository.DB.Exec(`
		UPDATE SCHEDULED_TASKS SET STATUS = $1 WHERE id = $2
	`, status, id)
	if err != nil {
		return false, err
	}

	return true, err
}
