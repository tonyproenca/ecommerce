package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/tonyproenca/stock-service/repository"
	"github.com/tonyproenca/stock-service/service"
	"github.com/tonyproenca/stock-service/web"
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

type App struct {
	handler web.Handlers
}

func main() {
	// Connect to Mongo
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panic(err)
	}

	// Setup
	repo := repository.NewMongoDBRepository(mongoClient)
	srv := service.NewProductStockService(repo)
	handler := web.NewProductStockHandler(srv)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Create routes
	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Use(middleware.Heartbeat("/health"))

	mux.Post("/product-stock", handler.PostProductStock)
	mux.Get("/product-stock/{productCode}", handler.GetProductStock)
	mux.Delete("/product-stock/{productCode}", handler.DeleteProductStock)
	mux.Put("/product-stock", handler.UpdateProductStock)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: mux,
	}

	// Start Server
	err = server.ListenAndServe()
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
