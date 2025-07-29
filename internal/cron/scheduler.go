package cron

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mateeusferro/schedula/internal/repository"
	"github.com/mateeusferro/schedula/internal/tasks"
	"github.com/robfig/cron/v3"
)

func InitCron(db *sql.DB) {
	taskRepository := repository.NewTaskRepository(db)
	c := cron.New()

	c.AddFunc("@every 1m", func() {
		task(taskRepository)
	})

	go c.Start()
}

func task(repository *repository.TaskRepository) {
	fmt.Println(time.Now().String() + " - Start Task")
	tasks.ProcessPendingTasks(repository)
	fmt.Println(time.Now().String() + " - End Task")
}
