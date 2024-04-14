package config

import (
	"fmt"
	"os"
	"time"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/pkg/errors"
)

type MongoConfig struct {
	Conn    string       `koanf:"conn"`
	DbName  string       `koanf:"dbName"`
	Docs    MongoDocs    `json:"docs"`
	Actions MongoActions `json:"actions"`
}

type MongoDocs struct {
	CollectionName string `koanf:"collectionName"`
}

type MongoActions struct {
	CollectionName string `koanf:"collectionName"`
}

type ElasticConfig struct {
	Conn            []string               `json:"conn"`
	Docs            ElasticDocs            `json:"docs"`
	SuggestKeywords ElasticSuggestKeywords `json:"suggestKeywords"`
	SuggestQueries  ElasticSuggestQueries  `json:"suggestQueries"`
}

type ElasticDocs struct {
	Index string `koanf:"index"`
}

type ElasticSuggestKeywords struct {
	Index string `koanf:"index"`
}

type ElasticSuggestQueries struct {
	Index string `koanf:"index"`
}

type HTTPConfig struct {
	Port            int `koanf:"port"`
	ReadTimeoutMs   int `koanf:"read_timeout_ms"`
	WriteTimeoutMs  int `koanf:"write_timeout_ms"`
	IdleTimeoutMs   int `koanf:"idle_timeout_ms"`
	MaxHeaderKBytes int `koanf:"max_header_kb"`
}

type Redis struct {
	Conn            string         `koanf:"conn"`
	Password        string         `koanf:"password"`
	CacheExpireTime *time.Duration `json:"cacheExpireTime"`
}

type Config struct {
	Http    HTTPConfig    `koanf:"http"`
	Elastic ElasticConfig `json:"elastic"`
	Mongo   MongoConfig   `json:"mongo"`
	Redis   Redis         `json:"redis"`
}

const (
	appEnvVarName = "APP_ENV"

	AppEnvLocal = "local"
	AppEnvDev   = "dev"
	AppEnvProd  = "prod"
)

func getAppEnv() (string, error) {
	appEnv := os.Getenv(appEnvVarName)

	if appEnv == "" {
		return AppEnvLocal, nil
	}

	switch appEnv {
	case AppEnvDev, AppEnvProd:
		return appEnv, nil
	default:
		return "", errors.New(fmt.Sprintf("unknown %s %s", appEnvVarName, appEnv))
	}
}

func New() (Config, error) {
	k := koanf.New(".")

	appEnv, err := getAppEnv()
	if err != nil {
		return Config{}, errors.Wrap(err, "unable to get app env")
	}

	if err := k.Load(file.Provider(fmt.Sprintf("cfg/%s.yaml", appEnv)), yaml.Parser()); err != nil {
		return Config{}, errors.Wrap(err, "unable to load local env")
	}

	var conf Config
	if err := k.Unmarshal("", &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}
