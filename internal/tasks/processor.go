package tasks

import (
	"fmt"
	"log"

	"github.com/mateeusferro/schedula/internal/repository"
)

func ProcessPendingTasks(repository *repository.TaskRepository) {
	data, err := repository.GetPendingTask()
	if err != nil {
		log.Fatalf("Error while retrieving pending tasks: %v", err)
	}

	for _, r := range data {
		fmt.Printf("%v", r)
	}
}
