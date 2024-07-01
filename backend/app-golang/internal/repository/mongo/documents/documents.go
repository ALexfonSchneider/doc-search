package documents

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) GetDocument(ctx context.Context, id string) (*entity.Document, error) {
	res := r.docs.FindOne(ctx, bson.M{"article.article_id": id})
	if err := res.Err(); err != nil {
		return nil, err
	}

	var document entity.Document
	if err := res.Decode(&document); err != nil {
		return nil, err
	}

	return &document, nil
}

func (r *Repository) UpdateDocument(ctx context.Context, document *entity.Document) error {
	_, err := r.docs.UpdateOne(ctx, bson.M{"article.article_id": document.Article.ArticleId}, bson.M{"$set": document})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteDocument(ctx context.Context, id string) error {
	_, err := r.docs.DeleteMany(ctx, bson.M{"article.article_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateDocument(ctx context.Context, document *entity.Document) error {
	_, err := r.docs.InsertOne(ctx, document)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return &entity.DocumentExistsErr{Id: document.Article.ArticleId}
		}
		return err
	}
	return nil
}
