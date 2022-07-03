package main

import (
	"context"
	"fmt"
	"github.com/tonyproenca/stock-service/cmd/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	webPort  = "80"
	mongoURL = "mongodb://mongo:27017"
)

type Config struct {
	Repo data.Repository
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Repo: data.NewMongoDBRepository(mongoClient),
	}

	log.Println("Starting service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic()
	}

}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: os.Getenv("MONGO_USERNAME"),
		Password: os.Getenv("MONGO_PASSWORD"),
	})

	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error Connecting", err)
		return nil, err
	}

	collection := c.Database("stock").Collection("stock")

	indexName, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "product_code", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Panic("error creating indexes", err)
	} else {
		log.Println("Index created: ", indexName)
	}

	return c, nil
}

func (app *Config) setupRepo(conn *mongo.Client) {
	app.Repo = data.NewMongoDBRepository(conn)
}
