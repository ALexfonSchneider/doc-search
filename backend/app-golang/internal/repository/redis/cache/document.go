package cache

import (
	"context"
	"doc-search-app-backend/internal/entity"
)

func (r *Repository) SetDocument(ctx context.Context, document *entity.Document) error {
	if cmd := r.client.Set(ctx, document.Article.ArticleId, *document, r.cacheExpireTime); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}

func (r *Repository) GetDocument(ctx context.Context, id string) (*entity.Document, error) {
	cmd := r.client.Get(ctx, id)

	err := cmd.Err()
	if err != nil {
		return nil, err
	}

	var document entity.Document
	if err := cmd.Scan(&document); err != nil {
		return nil, &entity.NoDocumentInCacheErr{Key: id}
	}

	return &document, nil
}
