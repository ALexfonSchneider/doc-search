package consumer

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"time"
)

type Repository interface {
	GetDocument(ctx context.Context, id string) (*entity.Document, error)
	UpdateDocument(ctx context.Context, document *entity.Document) error
	CreateDocument(ctx context.Context, document *entity.Document) error
	DeleteDocument(ctx context.Context, id string) error
}

type Actions interface {
	AddActions(ctx context.Context, actions []entity.DocAction) error
	GetNewActions(ctx context.Context, limit int64, timeout time.Duration) ([]entity.DocAction, error)
	UpdateStatus(ctx context.Context, entryId string, status entity.Status) error
}

type Index interface {
	IndexDocument(ctx context.Context, document *entity.Document) error
	UpdateDocument(ctx context.Context, document *entity.Document) error
	DeleteDocument(ctx context.Context, id string) error
	IndexKeyword(ctx context.Context, keyword string) error
	UnIndexKeyword(ctx context.Context, keyword string) error
}
