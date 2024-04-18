package worker_actions_consumer

import (
	"context"
	"doc-search-app-backend/internal/config"
	elasticrepo "doc-search-app-backend/internal/repository/elastic"
	"doc-search-app-backend/internal/repository/mongo/actions"
	"doc-search-app-backend/internal/repository/mongo/documents"
	"doc-search-app-backend/internal/services/actions/consumer"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func newLogger(consumer string) (*zerolog.Logger, error) {
	file, err := os.OpenFile(fmt.Sprintf("logs/actions/consumers/%s", consumer), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log := zerolog.New(file)

	return &log, err
}

func Start(ctx context.Context, cfg config.Config) {
	elasticRepo, err := elasticrepo.NewRepository(cfg.Elastic.Docs.Index, cfg.Elastic.SuggestKeywords.Index, cfg.Elastic.SuggestQueries.Index, "udk",
		elasticsearch.Config{
			Addresses: cfg.Elastic.Conn,
		})
	if err != nil {
		panic(err)
	}

	mongoRepo, err := documents.NewRepository(ctx, cfg.Mongo.Conn, cfg.Mongo.DbName, cfg.Mongo.Docs.CollectionName)
	if err != nil {
		panic(err)
	}

	actionsrepo, err := actions.NewRepository(ctx, cfg.Mongo.Conn, cfg.Mongo.DbName, cfg.Mongo.Actions.CollectionName)
	if err != nil {
		panic(err)
	}

	log, err := newLogger("consumers.txt")
	if err != nil {
		panic(err)
	}

	service := consumer.Service{
		Actions:        actionsrepo,
		Index:          elasticRepo,
		Repository:     mongoRepo,
		BatchSize:      10,
		Timeout:        100 * time.Second,
		OnErrorTimeout: 10000,
		Delay:          0,
		Log:            log,
	}

	service.Start(ctx)
}
