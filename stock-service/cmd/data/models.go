package data

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

var client *mongo.Client

type MongoDBRepository struct {
	Conn *mongo.Client
}

func NewMongoDBRepository(db *mongo.Client) *MongoDBRepository {
	client = db
	return &MongoDBRepository{
		Conn: db,
	}
}

//type Models struct {
//	StockProduct StockProduct
//}

type StockProduct struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	ProductCode string    `bson:"product_code" json:"productCode"`
	ProductName string    `bson:"product_name" json:"productName"`
	Quantity    int       `bson:"quantity" json:"quantity"`
	CreatedAt   time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updatedAt"`
}

//func New(mongo *mongo.Client) Models {
//	client = mongo
//	return Models{
//		StockProduct: StockProduct{},
//	}
//}

func (m *MongoDBRepository) Insert(stockProduct StockProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("stock").Collection("stock")

	_, err := collection.InsertOne(ctx, StockProduct{
		ProductCode: stockProduct.ProductCode,
		ProductName: stockProduct.ProductName,
		Quantity:    stockProduct.Quantity,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		log.Println("Error inserting into stock", err)
		return err
	}
	return nil
}

func (m *MongoDBRepository) GetOne(productCode string) (*StockProduct, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", productCode}}
	var entry StockProduct
	err := collection.FindOne(ctx, filter).Decode(&entry)

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (m *MongoDBRepository) Update(stockProduct StockProduct) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := client.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", stockProduct.ProductCode}}
	var entry StockProduct
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
				{"product_name", stockProduct.ProductName},
				{"quantity", stockProduct.Quantity},
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

	collection := client.Database("stock").Collection("stock")

	filter := bson.D{{"product_code", productCode}}
	var entry StockProduct
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
