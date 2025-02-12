package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoRepository struct {
	collection *mongo.Collection
}

func setUpSessionContext(sessionContext mongo.SessionContext) mongo.SessionContext {
	cont, cancel := context.WithTimeout(context.Background(), 10*time.second)
	defer cancel()

	if sessionContext == nil  {
		return mongo.NewSessionContext(cont, mongo.SessionFromContext(cont))
}
	return sessionContext

func (mr *MongoRepository) Read(id string cont mongo.SessionContext) (interface{}, error) {
	sessionContext := setUpSessionContext(sessionContext)
}
