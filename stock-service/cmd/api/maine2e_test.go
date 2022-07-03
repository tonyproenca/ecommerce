package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"github.com/tonyproenca/stock-service/cmd/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var db *mongo.Client
var testApp Config

func TestMain(m *testing.M) {
	// Setup
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	environmentVariables := []string{
		"MONGO_INITDB_ROOT_USERNAME=admin",
		"MONGO_INITDB_ROOT_PASSWORD=password",
		"MONGO_INITDB_DATABASE=stock",
	}

	log.Println("Starting pool")
	resource, err := pool.Run("mongo", "5.0", environmentVariables)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}
	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err = pool.Retry(func() error {
		var err error
		log.Println("Trying to connect...")
		db, err = mongo.Connect(
			context.TODO(),
			options.Client().ApplyURI(
				fmt.Sprintf("mongodb://admin:password@localhost:%s", resource.GetPort("27017/tcp")),
			),
		)
		if err != nil {
			return err
		}
		log.Println("Connected!")
		return db.Ping(context.TODO(), nil)
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// seed data
	collection := db.Database("stock").Collection("stock")

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

	_, err = collection.InsertOne(context.TODO(), data.StockProduct{
		ProductCode: "123",
		ProductName: "test",
		Quantity:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	result := collection.FindOne(context.TODO(), "62c210b26d113efcbb322074")
	log.Println("Seed result:", result)
	testApp.Repo = data.NewMongoDBRepository(db)

	if err != nil {
		log.Fatalf("Could not seed the data %s", err)
	}

	filter := bson.D{{"product_code", "123"}}
	var entry data.StockProduct
	_ = collection.FindOne(context.TODO(), filter).Decode(&entry)
	log.Println(entry)
	// Run tests
	exitCode := m.Run()

	// Teardown
	// When you're done, kill and remove the container
	if err = pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	// Exit
	os.Exit(exitCode)
}

func TestPostStockProductE2E(t *testing.T) {
	postBody := map[string]interface{}{
		"ProductName": "test",
		"ProductCode": "123",
		"Quantity":    "1",
	}

	body, _ := json.Marshal(postBody)
	req, _ := http.NewRequest("POST", "/stock-product", bytes.NewReader(body))
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.StoreNewStockProduct)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
}

func TestGetStockProductE2E(t *testing.T) {
	productCode := "123"
	req, _ := http.NewRequest("GET", "/stock-product/"+productCode, nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testApp.RetrieveStockProduct)

	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
