package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

//AppConfig Application Configuration
type AppConfig struct {
	Cacheuri       string `envconfig:"CACHE_URI"`       // Cacheuri
	RedisPassword  string `envconfig:"REDIS_PASSWORD"`  // Redis Password
	Dburi          string `envconfig:"DB_URI"`          //MongDB URI
	Env            string `envconfig:"ENV"`             //PROD,DEV
	MongoDatabase  string `envconfig:"MONGO_DATABASE"`  // Defines mongo database
	RedisDB        string `envconfig:"REDIS_DB"`        // Defines logical redis database
}

//GetCacheURI returs cache uri for redis
func (a *AppConfig) GetCacheURI() string {
	return a.Cacheuri
}

//GetDBURI returns mongodb URI
func (a *AppConfig) GetDBURI() string {
	return a.GetDBURI()
}

//GetEnv get the run time environment
func (a *AppConfig) GetEnv() string {
	return a.Env
}

// GetConfig get the current configuration from the environment
func GetConfig() (AppConfig, error) {
	var cfg AppConfig
	err := envconfig.Process("", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	return cfg, nil
}
