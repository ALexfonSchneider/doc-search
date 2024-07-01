package actions

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (r *Repository) GetNewActions(ctx context.Context, limit int64, timeout time.Duration) ([]entity.DocAction, error) {
	cursor, err := r.actions.Find(ctx, bson.M{"status": entity.New}, &options.FindOptions{Limit: &limit,
		Sort: bson.M{"created_at": -1}})

	if err != nil {
		return nil, err
	}
	cursor.SetMaxTime(timeout)

	var actions []entity.DocAction
	err = cursor.All(ctx, &actions)
	if err != nil {
		return nil, err
	}

	return actions, err
}

func (r *Repository) AddActions(ctx context.Context, actions []entity.DocAction) error {
	if len(actions) == 0 {
		return nil
	}

	var i []interface{}
	for _, action := range actions {
		i = append(i, action)
	}

	_, err := r.actions.InsertMany(ctx, i)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateStatus(ctx context.Context, entryId string, status entity.Status) error {
	id, err := primitive.ObjectIDFromHex(entryId)
	if err != nil {
		return err
	}

	_, err = r.actions.UpdateOne(ctx, bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}})
	if err != nil {
		return err
	}

	return nil
}
