package tasks

import (
	"context"
	"time"
	model "token-based-auth/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	coll *mongo.Collection
}

func New(c *mongo.Collection) *Repository {
	return &Repository{
		coll: c,
	}
}

func (r *Repository) Get(ctx context.Context, tID string) *model.Task {
	var tsk Task
	id, err := primitive.ObjectIDFromHex(tID)
	if err != nil {
		return &model.Task{}
	}

	err = r.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&tsk)
	if err != nil {
		return &model.Task{}
	}
	return tsk.ToDomain()
}

func (r *Repository) Create(ctx context.Context, i model.Info) *model.Task {
	t := Task{
		Title:       i.Title,
		Description: i.Description,
		State:       ToDo,
		Priority:    1,
		Date:        time.Now(),
	}

	res, err := r.coll.InsertOne(ctx, t)
	if err != nil {
		return &model.Task{}
	}
	
	t.ID = res.InsertedID.(primitive.ObjectID)
	return t.ToDomain()
}
