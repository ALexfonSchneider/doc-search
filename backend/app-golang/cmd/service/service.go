package main

import (
	"context"
	worker_actions_consumer "doc-search-app-backend/cmd/service/workers/actions/consumer"
	"doc-search-app-backend/internal/config"
	"doc-search-app-backend/internal/handlers/metrics"
	"doc-search-app-backend/internal/handlers/search"
	"doc-search-app-backend/internal/handlers/suggest_keywords"
	"doc-search-app-backend/internal/handlers/suggest_queries"
	"doc-search-app-backend/internal/middlewares"
	elasticrepo "doc-search-app-backend/internal/repository/elastic"
	"doc-search-app-backend/internal/repository/mongo/documents"
	"doc-search-app-backend/internal/repository/redis/cache"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"sync"
	"time"
)

func NewLogger() (*zerolog.Logger, error) {
	file, err := os.OpenFile("logs/log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	log := zerolog.New(file)

	return &log, err
}

func main() {
	//ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer cancel()

	ctx := context.Background()

	log, err := NewLogger()
	if err != nil {
		log.Error().Timestamp().Err(err)
		return
	}

	cfg, err := config.New()
	if err != nil {
		log.Error().Timestamp().Err(err)
		return
	}

	fmt.Println(cfg)

	elasticRepo, err := elasticrepo.NewRepository(cfg.Elastic.Docs.Index, cfg.Elastic.SuggestKeywords.Index, cfg.Elastic.SuggestQueries.Index,
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

	redisRepo, err := cache.NewRepository(ctx, cfg.Redis.Conn, cfg.Redis.Password, (*cfg.Redis.CacheExpireTime)*time.Minute)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer func() {
			//recover()
		}()
		worker_actions_consumer.Start(ctx, cfg)
	}()

	e := echo.New()
	e.Use(middlewares.ErrorHandler)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	SearchHandler := search.NewHadler(elasticRepo)
	SuggestKeywordsHandler := suggest_keywords.NewHadler(elasticRepo)
	SuggestQueriesHandler := suggest_queries.NewHadler(elasticRepo)
	MetricsHandler := metrics.NewHandler(mongoRepo, redisRepo)

	e.GET("/search", SearchHandler.Handle)

	suggestGroup := e.Group("/suggest")
	{
		suggestGroup.GET("/keywords", SuggestKeywordsHandler.Handle)
		suggestGroup.GET("/queries", SuggestQueriesHandler.Handle)
	}

	e.GET("/metrics", MetricsHandler.Handle)

	err = e.StartServer(&http.Server{
		Addr:                         fmt.Sprintf(":%d", cfg.Http.Port),
		DisableGeneralOptionsHandler: true,
		Handler:                      e,
		//BaseContext:                  func(_ net.Listener) context.Context { return ctx },
	})

	if err != nil {
		log.Error().Timestamp().Err(err)
	}
}
