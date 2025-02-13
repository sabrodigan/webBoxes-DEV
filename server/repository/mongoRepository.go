package repository

import (
	"context"
	"github.com/sabrodigan/webboxes/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type IMongoRepository interface {
	ReadOne(id string, ctx mongo.SessionContext) (interface{}, error)
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

func (mr *MongoRepository) ReadOne(id string, ctx mongo.SessionContext) (interface{}, error) {
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
	defer cursor.Close(sessionContext)

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

func (mr *MongoRepository) Aggregate(pipelines mongo.Pipeline, ctx mongo.SessionContext) ([]map[string]interface{}, error) {
	sessionContext := getSessionContext(ctx)
	cursor, err := mr.collection.Aggregate(sessionContext, pipelines)
	if err != nil {
		log.Fatalf("Error updating document: %v", err)
		return nil, err
	}
	defer cursor.Close(sessionContext)

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
