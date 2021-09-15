package mongodb

import (
	"WorkerPlace/app/domain/entity"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccessMongoDbRepository struct {
	collection *mongo.Collection
}

func NewAccessMongoDbRepository(mongoDb *MongoDbConnection) *AccessMongoDbRepository {
	collection :=mongoDb.client.Database("workerplace").Collection("records")
	return &AccessMongoDbRepository{collection: collection}
}

func (c *AccessMongoDbRepository)GetAccessRecordDocument(idDocument string) (*entity.AccessRecordDocument, error) {
	filter := bson.D{{"id", idDocument}}
	var result *entity.AccessRecordDocument

	err := c.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Found a single document: %+v\n", result)

	return result, nil
}
func (c *AccessMongoDbRepository)SaveAccessRecordDocument(record *entity.AccessRecordDocument) error {
	insertResult, err := c.collection.InsertOne(context.TODO(),record)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saved a single document: ", insertResult.InsertedID)

	return nil
}

func (c *AccessMongoDbRepository)UpdateAccessRecordDocument(record *entity.AccessRecordDocument) error {
	filter := bson.M{"id": record.Id}
	insertResult, err := c.collection.ReplaceOne(context.TODO(), filter,record)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Saved a single document: ", insertResult.UpsertedID)

	return nil
}

