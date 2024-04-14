package elastic

import (
	"context"
	"crypto"
	"doc-search-app-backend/internal/entity"
	"fmt"
	"strings"
)

func (r *Repository) IndexQuery(ctx context.Context, query string) error {
	indexQuery := r.BuildIndexQuery(query)

	h := crypto.SHA256.New()
	h.Write([]byte(query))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	_, err := r.client.Update(r.querySuggestIndex, hash, strings.NewReader(indexQuery), r.client.Update.WithContext(ctx))
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) IndexKeyword(ctx context.Context, keyword string) error {
	indexKeyword := r.BuildIndexKeyword(keyword)

	h := crypto.SHA256.New()
	h.Write([]byte(keyword))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	_, err := r.client.Update(r.keywordSuggestIndex, hash, strings.NewReader(indexKeyword), r.client.Update.WithContext(ctx))
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UnIndexKeyword(ctx context.Context, keyword string) error {
	indexKeyword := r.BuildUnIndexKeyword(keyword)

	h := crypto.SHA256.New()
	h.Write([]byte(keyword))
	hash := fmt.Sprintf("%x", h.Sum(nil))

	_, err := r.client.Update(r.keywordSuggestIndex, hash, strings.NewReader(indexKeyword), r.client.Update.WithContext(ctx))
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) IndexDocument(ctx context.Context, document *entity.Document) error {
	if _, err := r.TypedClient.Index(r.searchIndex).Id(document.Article.ArticleId).Document(*document).Do(ctx); err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteDocument(ctx context.Context, id string) error {
	_, err := r.TypedClient.Delete(r.searchIndex, id).Do(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateDocument(ctx context.Context, document *entity.Document) error {
	if _, err := r.TypedClient.Update(r.searchIndex, document.Article.ArticleId).Doc(*document).Do(ctx); err != nil {
		return err
	}

	return nil
}
