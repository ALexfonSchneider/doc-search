package main

import (
	"context"
	"doc-search-app-backend/internal/config"
	"doc-search-app-backend/internal/handlers/metrics"
	"doc-search-app-backend/internal/handlers/search"
	"doc-search-app-backend/internal/handlers/suggest_keywords"
	"doc-search-app-backend/internal/handlers/suggest_queries"
	logging "doc-search-app-backend/internal/logger"
	"doc-search-app-backend/internal/middlewares"
	elasticrepo "doc-search-app-backend/internal/repository/elastic"
	"doc-search-app-backend/internal/repository/mongo"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	//ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer cancel()

	log, err := logging.New()
	if err != nil {
		log.Error(err.Error())
		return
	}
	logging.Log = log

	cfg, err := config.New()
	if err != nil {
		log.Error(err.Error())
		return
	}

	e := echo.New()
	e.Use(middlewares.ErrorHandler)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	elasticRepo, err := elasticrepo.NewRepository("states", "keywords-suggest", "queries-suggest",
		elasticsearch.Config{
			Addresses: []string{cfg.Elastic.Conn},
		})

	mongoRepo, err := mongo.NewRepository(context.TODO(), cfg.Mongo.Conn, "doc-search", "documents")

	if err != nil {
		panic(err)
	}

	SearchHandler := search.NewHadler(elasticRepo)
	SuggestKeywordsHandler := suggest_keywords.NewHadler(elasticRepo)
	SuggestQueriesHandler := suggest_queries.NewHadler(elasticRepo)
	MetricsHandler := metrics.NewHandler(mongoRepo)

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
		log.Error(err.Error())
	}
}
