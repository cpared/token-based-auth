package model

import (
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	State       string
	Priority    int
	Date        time.Time
}

type Info struct {
	Title       string
	Description string
}
