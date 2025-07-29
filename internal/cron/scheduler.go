package cron

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New()

	c.AddFunc("@every 1m10s", task)

	go c.Start()
}

func task() {
	fmt.Println(time.Now().String() + " - Start Task")
	time.Sleep(1 * time.Second)
	fmt.Println(time.Now().String() + " - End Task")
}
