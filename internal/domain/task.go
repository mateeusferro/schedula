package domain

import "time"

type Task struct {
	Id           string                 `json:"id"`
	Name         string                 `json:"name"`
	Payload      map[string]interface{} `json:"payload"`
	Run_at       time.Time              `json:"run_at"`
	Status       string                 `json:"status"`
	Attempts     int                    `json:"attempts"`
	Max_attempts int                    `json:"max_attempts"`
	Created_at   time.Time              `json:"created_at"`
	Updated_at   time.Time              `json:"updated_at"`
}

type TaskInput struct {
	Name         string                 `json:"name"`
	Payload      map[string]interface{} `json:"payload"`
	Run_at       time.Time              `json:"run_at"`
	Status       string                 `json:"status"`
	Attempts     int                    `json:"attempts"`
	Max_attempts int                    `json:"max_attempts"`
}
