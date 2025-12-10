package tasks

import (
	"time"
	model "token-based-auth/internal/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type States int

const (
	Backlog States = iota
	ToDo
	Ready
	WIP
	Finished
)

var statesMapper = map[States]string{
	ToDo:     "To Do",
	Ready:    "Ready",
	WIP:      "Work In Progress",
	Finished: "Finished",
}

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"` // The omitempty tag says: If _id is empty, don't serialize it → Mongo will generate it.
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	State       States             `bson:"state"`
	Priority    int                `bson:"priority"`
	Date        time.Time          `bson:"date"`
}

func (t *Task) ToDomain() *model.Task {
	s, _ := statesMapper[t.State]
	return &model.Task{
		ID:          t.ID.Hex(),
		Title:       t.Title,
		Description: t.Description,
		State:       s,
		Date:        t.Date,
	}
}
