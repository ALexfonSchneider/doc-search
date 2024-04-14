package actions

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Repository struct {
	client  *mongo.Client
	db      *mongo.Database
	actions *mongo.Collection

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
		client:  client,
		db:      db,
		actions: db.Collection(collectionName),

		dbName:         dbName,
		collectionName: collectionName,
	}, nil
}

func (r *Repository) Close(ctx context.Context) error {
	return r.client.Disconnect(ctx)
}
