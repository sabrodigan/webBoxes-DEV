package repository

import (
	"context"
	"fmt"
	"github.com/sabrodigan/webboxes/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type IMongoRepository interface {
	FindOne(id string, ctx mongo.SessionContext) (interface{}, error)
	FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error)
	Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Create(data interface{}, ctx mongo.SessionContext) (interface{}, error)
	Delete(id string, ctx mongo.SessionContext) (interface{}, error)
	FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error)
	Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error)
}
type MongoRepository struct {
	collection *mongo.Collection
}

func getSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if sessionContext == nil {
		return mongo.NewSessionContext(ctx, mongo.SessionFromContext(ctx))
	}
	return sessionContext
}

func (mr *MongoRepository) FindOne(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objId})
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		log.Fatalf("Error decoding document: %v", err)
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) Update(id string, data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	result, err := mr.collection.UpdateOne(sessionContext, bson.M{"_id": objId}, data)
	if err != nil {
		log.Fatalf("Error updating document: %v", err)
		return nil, err
	}
	return result, err
}

func (mr *MongoRepository) Create(data interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}

func (mr *MongoRepository) Delete(id string, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)
	objId, err := primitive.ObjectIDFromHex(id)
	result, err := mr.collection.DeleteOne(sessionContext, bson.M{"_id": objId})
	if err != nil {
		log.Fatalf("Error updating document: %v", err)
		return nil, err
	}
	return result, err
}

func (mr *MongoRepository) FindAll(filter interface{}, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
	cursor, err := mr.collection.Find(sessionContext, filter)
	if err != nil {
		log.Fatalf("Error updating document: %v", err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing cursor: %v", err)
		}
	}(cursor, sessionContext)

	var results []map[string]interface{}

	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err := cursor.Decode(&document); err != nil {
			log.Fatalf("Error decoding document: %v", err)
			return nil, err
		}
		results = append(results, document)
	}
	return results, cursor.Err()
}

func (mr *MongoRepository) FindOneByKey(key string, value interface{}, ctx mongo.SessionContext) (interface{}, error) {
	sessionContext := getSessionContext(ctx)

	actualValue, err := primitive.ObjectIDFromHex(value.(string))
	var result *mongo.SingleResult
	if err != nil {
		fmt.Printf("%s is not an object id\n", key)
		result = mr.collection.FindOne(sessionContext, bson.M{key: actualValue})
	} else {
		result = mr.collection.FindOne(sessionContext, bson.M{key: value})
	}
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		return nil, err
	}

	return document, nil
}

func (mr *MongoRepository) Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
	cursor, err := mr.collection.Aggregate(sessionContext, pipelines)
	if err != nil {
		log.Fatalf("Error updating document: %v", err)
		return nil, err
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatalf("Error closing cursor: %v", err)
		}
	}(cursor, sessionContext)

	var results []map[string]interface{}

	for cursor.Next(sessionContext) {
		var document map[string]interface{}
		if err := cursor.Decode(&document); err != nil {
			log.Fatalf("Error decoding document: %v", err)
			return nil, err
		}
		results = append(results, document)
	}

	return results, cursor.Err()
}

func GetMongoRepository(dbName string, collectionName string) IMongoRepository {
	collection, _ := config.GetDatabaseCollection(&dbName, collectionName)
	return &MongoRepository{
		collection: collection,
	}

}
