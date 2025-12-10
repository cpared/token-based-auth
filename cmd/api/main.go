package main

import (
	"context"
	"log"
	"os"
	taskshdl "token-based-auth/internal/handlers/tasks"
	tasksrepo "token-based-auth/internal/repositories/tasks"
	tasksserv "token-based-auth/internal/services/tasks"

	loginrepo "token-based-auth/internal/repositories/login"
	loginserv "token-based-auth/internal/services/login"
	loginhdl "token-based-auth/internal/handlers/login"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

const (
	mongoURI   = "mongodb://localhost:27017"
	database   = "test"
	collection = "tasks"
	configPath = "/config/.env"
)

type Config struct {
	JwtSecret string
}

func main() {
	// Init logger
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal("cannot initialize zap logger")
	}

	sugar := logger.Sugar()
	defer logger.Sync()

	// Init router
	r := gin.Default()

	// Connect to mongo
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
	coll := client.Database(database).Collection(collection)

	cfg := ReadConfig()

	// Init repositories
	tasksRepo := tasksrepo.New(coll)
	loginRepo := loginrepo.New()

	// Init servicies
	tasksServ := tasksserv.New(tasksRepo)
	loginServ := loginserv.New(loginRepo)

	// Init handlers
	tasksHdl := taskshdl.New(tasksServ)
	loginHdl := loginhdl.New(loginServ)

	r.POST("login", loginHdl.Login(cfg.JwtSecret))
	r.GET("tasks/:id", tasksHdl.GetByID())
	r.POST("tasks", tasksHdl.Create())

	r.Run()
}


// Reads configuration setting in /config/.env files and returns
// 
// Config object or empty Config file if its fail
func ReadConfig() *Config {
	dir, _ := os.Getwd()
	godotenv.Load(dir + configPath)
	secret := os.Getenv("JWT_SECRET")

	return &Config{
		JwtSecret: secret,
	}
}