package consumer

import (
	"context"
	"doc-search-app-backend/internal/entity"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"time"
)

type Service struct {
	Actions    Actions
	Index      Index
	Repository Repository

	BatchSize      int64
	Timeout        time.Duration
	OnErrorTimeout time.Duration
	Delay          time.Duration

	Log *zerolog.Logger
}

func (s *Service) onErrorDelay() {
	time.Sleep(time.Duration(s.OnErrorTimeout))
}

func (s *Service) delay() {
	time.Sleep(time.Duration(s.Delay))
}

func (s *Service) handleNewDocument(ctx context.Context, document *entity.Document) error {
	if err := s.Repository.CreateDocument(ctx, document); err != nil {
		if errors.Is(err, &entity.DocumentExistsErr{}) {
			//pass
		} else {
			return errors.WithMessage(err, "Could not add document to storage")
		}
	}

	if err := s.Index.IndexDocument(ctx, document); err != nil {
		return errors.WithMessage(err, "Could not index document")
	}

	if len(document.Article.Keywords) > 0 {
		createAt := time.Now()
		var newActions []entity.DocAction
		for _, keyword := range document.Article.Keywords {
			newActions = append(newActions, entity.DocAction{
				Keyword:   keyword,
				ArticleId: document.Article.ArticleId,
				Status:    entity.New,
				Action:    entity.AddKeyword,
				CreatedAt: createAt,
				UpdatedAt: nil,
			})
		}

		if err := s.Actions.AddActions(ctx, newActions); err != nil {
			return errors.WithMessage(err, "Critical. Could not index document")
		}
	}

	return nil
}

func (s *Service) handleDeleteDocument(ctx context.Context, document *entity.Document) error {
	if err := s.Index.DeleteDocument(ctx, document.Article.ArticleId); err != nil {
		return errors.WithMessage(err, "Could not delete document")
	}

	createAt := time.Now()
	var newActions []entity.DocAction
	for _, keyword := range document.Article.Keywords {
		newActions = append(newActions, entity.DocAction{
			Keyword:   keyword,
			ArticleId: document.Article.ArticleId,
			Status:    entity.New,
			Action:    entity.DeleteKeyword,
			CreatedAt: createAt,
			UpdatedAt: nil,
		})
	}

	if err := s.Actions.AddActions(ctx, newActions); err != nil {
		return errors.WithMessage(err, "Could not add action")
	}

	if err := s.Repository.DeleteDocument(ctx, document.Article.ArticleId); err != nil {
		return errors.WithMessage(err, "Could not delete delete document")
	}

	return nil
}

func (s *Service) handleModifyDocument(ctx context.Context, document *entity.Document) error {
	//TODO: доделать

	//storedDocument, err := s.Repository.GetDocument(ctx, document.Article.ArticleId)
	//if err != nil {
	//	return err
	//}
	//
	//createAt := time.Now()
	//var newActions []entity.DocAction

	//for _, keyword := range document.Article.Keywords {
	//	if slices.Contains(storedDocument.Article.Keywords, keyword) {
	//		newActions = append(newActions, entity.DocAction{
	//			Keyword:   keyword,
	//			ArticleId: document.Article.ArticleId,
	//			Status:    entity.New,
	//			Action:    entity.Delete,
	//			CreatedAt: createAt,
	//			UpdatedAt: nil,
	//		})
	//	}
	//}
	//
	//if err := s.Actions.AddActions(ctx, newActions); err != nil {
	//	return errors.Wrap(err, "Could not index document")
	//}
	//
	//if err := s.Repository.DeleteDocument(ctx, document.Article.ArticleId); err != nil {
	//	return errors.Wrap(err, "Could not delete document from store")
	//}

	return nil
}

func (s *Service) handleAddKeyword(ctx context.Context, keyword string) error {
	if err := s.Index.IndexKeyword(ctx, keyword); err != nil {
		return errors.WithMessage(err, "Could not add keyword to index")
	}
	return nil
}

func (s *Service) handleDeleteKeyword(ctx context.Context, keyword string) error {
	if err := s.Index.UnIndexKeyword(ctx, keyword); err != nil {
		return errors.WithMessage(err, "Could not delete keyword from index")
	}
	return nil
}

func (s *Service) Start(ctx context.Context) {
	s.Log.Info().Msg("Start service")

	for {
		actions, err := s.Actions.GetNewActions(ctx, s.BatchSize, s.Timeout)
		if err != nil {
			s.Log.Error().Timestamp().Err(err).Msg("Could not get actions")
			s.onErrorDelay()
			continue
		}

		for _, action := range actions {
			s.Log.Info().Timestamp().Any("action", action).Msg("Start processing action")

			if action.Action == entity.Add {
				if err := s.handleNewDocument(ctx, action.Document); err != nil {
					s.Log.Error().Timestamp().Msg(err.Error())
					s.onErrorDelay()
					continue
				}
			} else if action.Action == entity.Delete {
				if err := s.handleDeleteDocument(ctx, action.Document); err != nil {
					s.Log.Error().Timestamp().Msg(err.Error())
					s.onErrorDelay()
					continue
				}
			} else if action.Action == entity.Modify {
				if err := s.handleModifyDocument(ctx, action.Document); err != nil {
					s.Log.Error().Timestamp().Msg(err.Error())
					s.onErrorDelay()
					continue
				}
			} else if action.Action == entity.AddKeyword {
				if err := s.handleAddKeyword(ctx, action.Keyword); err != nil {
					s.Log.Error().Timestamp().Msg(err.Error())
					s.onErrorDelay()
					continue
				}
			} else if action.Action == entity.DeleteKeyword {
				if err := s.handleDeleteKeyword(ctx, action.Keyword); err != nil {
					s.Log.Error().Timestamp().Msg(err.Error())
					s.onErrorDelay()
					continue
				}
			} else {
				s.Log.Panic().Timestamp().Any("action", action).Err(err).Msg("Unknown action")
				continue
			}

			if err := s.Actions.UpdateStatus(ctx, action.Id, entity.Processed); err != nil {
				s.Log.Error().Timestamp().Msg(errors.WithMessage(err, "Could update status for action").Error())
				continue
			}
			s.Log.Info().Timestamp().Any("action", action).Msg("Successful processed document")
		}
		s.delay()
	}
}
