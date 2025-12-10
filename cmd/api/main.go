package main

import (
	"context"
	"log"
	taskshdl "token-based-auth/internal/handlers/tasks"
	tasksrepo "token-based-auth/internal/repositories/tasks"
	tasksserv "token-based-auth/internal/services/tasks"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	mongoURI   = "mongodb://localhost:27017"
	Database   = "test"
	Collection = "tasks"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("cannot initialize zap logger")
	}

	r := gin.Default()
	sugar := logger.Sugar()
	defer logger.Sync()

	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic("cannot connect to mongo cluster")
	}

	if err := client.Ping(ctx, nil); err != nil {
    	sugar.Fatal("Could not connect to Mongo:", err)
		panic("Could not connect to Mongo")
	}

	sugar.Info("Connected to mongo cluster succesfully!")

	defer client.Disconnect(ctx)
	coll := client.Database(Database).Collection(Collection)

	// Init repositories
	tasksRepo := tasksrepo.New(coll)

	// Init servicies
	tasksServ := tasksserv.New(tasksRepo)

	// Init handlers
	taskshdl := taskshdl.New(tasksServ)

	r.GET("tasks/:id", taskshdl.GetByID())
	r.POST("tasks", taskshdl.Create())

	r.Run()
}
