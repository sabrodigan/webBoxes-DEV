package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func setUpSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	cont, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if sessionContext == nil {
		return mongo.NewSessionContext(cont, mongo.SessionFromContext(cont))
	}
	return sessionContext
}

func (mr *MongoRepository) FindOne(id string, cont mongo.SessionContext) (interface{}, error) {
	sessionContext := setUpSessionContext(cont)
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := mr.collection.FindOne(sessionContext, bson.M{"_id": objId})
	var document map[string]interface{}
	if err := result.Decode(&document); err != nil {
		return nil, err
	}
	return document, nil
}

func (mr *MongoRepository) Create(data interface{}, cont mongo.SessionContext) (interface{}, error) {
	sessionContext := setUpSessionContext(cont)
	result, err := mr.collection.InsertOne(sessionContext, data)
	return result, err
}
