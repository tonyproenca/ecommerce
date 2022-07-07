package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type MongoDBRepository struct {
	Conn *mongo.Client
}

func NewMongoDBRepository(db *mongo.Client) *MongoDBRepository {
	return &MongoDBRepository{
		Conn: db,
	}
}

type ProductStockEntity struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	ProductCode string    `bson:"product_code" json:"productCode"`
	ProductName string    `bson:"product_name" json:"productName"`
	Quantity    int       `bson:"quantity" json:"quantity"`
	CreatedAt   time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updatedAt"`
}

func (m *MongoDBRepository) Insert(entity ProductStockEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := m.Conn.Database("stock").Collection("stock")

	_, err := collection.InsertOne(ctx, ProductStockEntity{
		ProductCode: entity.ProductCode,
		ProductName: entity.ProductName,
		Quantity:    entity.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into stock", err)
		return err
	}
	return nil
}

func (m *MongoDBRepository) GetOne(productCode string) (*ProductStockEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := m.Conn.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", productCode}}
	var entry ProductStockEntity
	err := collection.FindOne(ctx, filter).Decode(&entry)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (m *MongoDBRepository) Update(productStock ProductStockEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := m.Conn.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", productStock.ProductCode}}
	var entry ProductStockEntity
	err := collection.FindOne(ctx, filter).Decode(&entry)
	if err != nil {
		return err
	}

	idPrimitive, err := primitive.ObjectIDFromHex(entry.ID)

	_, err = collection.UpdateOne(
		ctx,
		bson.M{"_id": idPrimitive},
		bson.D{
			{"$set", bson.D{
				{"product_name", productStock.ProductName},
				{"quantity", productStock.Quantity},
				{"updated_at", time.Now()},
			}},
		},
	)

	if err != nil {
		return err
	}

	return nil
}

func (m *MongoDBRepository) Delete(productCode string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := m.Conn.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", productCode}}
	var entry ProductStockEntity
	err := collection.FindOne(ctx, filter).Decode(&entry)
	if err != nil {
		return err
	}

	idPrimitive, err := primitive.ObjectIDFromHex(entry.ID)

	_, err = collection.DeleteOne(
		ctx,
		bson.M{"_id": idPrimitive},
	)

	if err != nil {
		return err
	}
	return nil
}
