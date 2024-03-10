package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

import (
	"context"
)

type Repository struct {
	client *mongo.Client
	db     *mongo.Database
	docs   *mongo.Collection

	dbName         string
	collectionName string
}

func NewRepository(ctx context.Context, conn string, dbName string, collectionName string) (*Repository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conn).SetConnectTimeout(10*time.Second))
	if err != nil {
		return nil, err
	}

	db := client.Database(dbName)

	return &Repository{
		dbName:         dbName,
		collectionName: collectionName,
		client:         client,
		db:             db,
		docs:           db.Collection(collectionName),
	}, nil
}

func (r *Repository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
