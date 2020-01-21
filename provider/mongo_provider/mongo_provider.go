package mongo_provider

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoProvider(ctx context.Context, connectString string, dbName string) (*mongo.Database, error){
	clientOpts := options.Client().ApplyURI(connectString)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil{
		logrus.Fatal(err)
		return nil, err
	}
	db := client.Database(dbName)
	return db, nil
}