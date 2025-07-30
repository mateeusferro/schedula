package tasks

import (
	"encoding/json"
	"log"
	"time"

	"github.com/mateeusferro/schedula/internal/repository"
)

func ProcessPendingTasks(repository *repository.TaskRepository) {
	data, err := repository.GetPendingTask()
	if err != nil {
		log.Fatalf("Error while retrieving pending tasks: %v", err)
	}

	for _, task := range data {
		payload, _ := json.MarshalIndent(task.Payload, "", "  ")
		log.Printf("Processing task: %v at %v", task.Id, time.Now())

		log.Printf("---- At this moment the task action is mocked ----")
		log.Printf("Task name: %v", task.Name)
		log.Printf("Task payload: %v", string(payload))

		log.Printf("Task processed: %v at %v", task.Id, time.Now())

		_, err := repository.UpdateTaskStatus(task.Id, "completed")
		if err != nil {
			log.Fatalf("Error while updating task status: %v", err)
		}
	}
}
